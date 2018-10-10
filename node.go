package zek

import (
	"encoding/xml"
	"io"
	"reflect"
	"strings"

	"golang.org/x/net/html/charset"
)

// Node represents an element in the XML tree. It keeps track of its name,
// attributes, childnodes and example chardata and basic statistics, e.g. how
// often a node has been seen within its parent node.
type Node struct {
	Name        xml.Name   `json:"name,omitempty"`
	Attr        []xml.Attr `json:"attr,omitempty"`
	Examples    []string   `json:"examples,omitempty"`
	Children    []*Node    `json:"children,omitempty"`
	Freqs       []int      `json:"-"` // Collect number of occurences of this node within parent.
	MaxExamples int        `json:"-"` // Maximum number of examples to keep, gets passed to children.

	childFreqs map[xml.Name]int // Count child tag occurences, used temporarily.
}

// readNode reads XML from a reader and returns a parsed node. If node is
// given, it is reused, allowing for multiple passes (e.g. from multiple
// files). XXX: maxExamples should be factored out into options.
func readNode(r io.Reader, root *Node, maxExamples int) (node *Node, n int64, err error) {
	cw := countwriter{}
	rr := io.TeeReader(r, &cw)

	dec := xml.NewDecoder(rr)
	dec.CharsetReader = charset.NewReaderLabel
	dec.Strict = false

	if root == nil {
		root = &Node{}
	}

	stack := Stack{}
	stack.Put(root)

	for {
		token, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return root, cw.n, err
		}
		switch t := token.(type) {
		case xml.StartElement:
			parent := stack.Peek().(*Node)
			n := parent.CreateOrGetChild(t.Name, t.Attr)
			stack.Put(n)
		case xml.EndElement:
			n := stack.Pop().(*Node)
			n.End()
		case xml.CharData:
			v := strings.TrimSpace(string(t))
			if v == "" {
				break
			}
			n := stack.Peek().(*Node)
			if len(n.Examples) < maxExamples {
				// XXX: sample better, e.g. reservoir dictionary.
				n.Examples = append(n.Examples, v)
			}
		}
	}

	stack.Pop()
	return root, cw.n, nil
}

// ReadFromAll builds a single node from all readers.
func (node *Node) ReadFromAll(readers []io.Reader) (n int64, err error) {
	root := &Node{}
	var nr int64
	for _, r := range readers {
		if root, nr, err = readNode(r, root, node.MaxExamples); err != nil {
			return n, err
		}
		n = n + nr
	}
	if len(root.Children) > 0 {
		// Decapitate node.
		*node = *root.Children[0]
	}
	return n, nil
}

// ReadFrom reads XML from a reader.
func (node *Node) ReadFrom(r io.Reader) (int64, error) {
	nn, n, err := readNode(r, nil, node.MaxExamples)
	if err != nil {
		return n, err
	}
	if len(nn.Children) > 0 {
		// Decapitate node.
		*node = *nn.Children[0]
	}
	return n, nil
}

// attrListContains determines containment only on attribute name, not value.
func attrListContains(attrs []xml.Attr, attr xml.Attr) bool {
	for _, a := range attrs {
		if a.Name == attr.Name {
			return true
		}
	}
	return false
}

// mergeAttr adds attributes to node, which are not already there.
func (node *Node) mergeAttr(attr []xml.Attr) {
	for _, a := range attr {
		if attrListContains(node.Attr, a) {
			continue
		}
		node.Attr = append(node.Attr, a)
	}
}

// IsMultivalued returns true, if this node appeared more than once.
func (node *Node) IsMultivalued() bool {
	for _, f := range node.Freqs {
		if f > 1 {
			return true
		}
	}
	return false
}

// CreateOrGetChild creates a child if no child with the same tag name exists,
// otherwise returns the existing node with that name. We want to collect node
// and attribute information for a node and not replicate the XML tree.
func (node *Node) CreateOrGetChild(name xml.Name, attr []xml.Attr) *Node {
	for _, child := range node.Children {
		if child.Name.Local != name.Local {
			continue
		}
		child.mergeAttr(attr)
		node.childFreqs[name]++
		return child
	}

	if node.childFreqs == nil {
		node.childFreqs = make(map[xml.Name]int)
	}
	node.childFreqs[name]++

	n := &Node{Name: name, Attr: attr, MaxExamples: node.MaxExamples}
	node.Children = append(node.Children, n)
	return n
}

// End signals end of an element.
func (node *Node) End() {
	// Take note of frequencies, collect them inside child for later stats.
	for name, freq := range node.childFreqs {
		for _, c := range node.Children {
			if c.Name != name {
				continue
			}
			c.Freqs = append(c.Freqs, freq)
		}
	}
	// Reset counter.
	node.childFreqs = make(map[xml.Name]int)
}

// Height returns the height of the tree. A tree with zero nodes has height
// zero, a single node tree has height 1.
func (node *Node) Height() int {
	if node == nil || reflect.DeepEqual(node, new(Node)) {
		return 0
	}
	max := 0
	for _, c := range node.Children {
		h := c.Height()
		if h > max {
			max = h
		}
	}
	return 1 + max
}

// ByName finds a node in the tree by name. Comparisons start at the current
// node. First match is returned. If nothing matches, nil is returned.
func (node *Node) ByName(name string) *Node {
	if node == nil {
		return nil
	}
	if name == "" || node.Name.Local == name {
		return node
	}
	for _, c := range node.Children {
		if n := c.ByName(name); n != nil {
			return n
		}
	}
	return nil
}
