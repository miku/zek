package zek

import (
	"encoding/xml"
	"io"
	"strings"
)

// countwriter counts bytes written to it.
type countwriter struct {
	n int64
}

// Write increments the number of bytes read by len(p).
func (w *countwriter) Write(p []byte) (n int, err error) {
	w.n += w.n + int64(len(p))
	return len(p), nil
}

// Node represents an element in the XML tree.
type Node struct {
	Name     xml.Name   `json:"name,omitempty"`
	Attr     []xml.Attr `json:"attr,omitempty"`
	Examples []string   `json:"examples,omitempty"`
	Children []*Node    `json:"children,omitempty"`

	maxExamples      int   // Maximum number of examples for this node.
	freqs            []int // Counts the number of occurences of this node within parent.
	childFrequencies map[xml.Name]int
}

// ReadFrom reads XML from a reader.
func (node *Node) ReadFrom(r io.Reader) (n int64, err error) {
	cw := countwriter{}
	rr := io.TeeReader(r, &cw)

	dec := xml.NewDecoder(rr)
	dec.Strict = false

	stack, root := Stack{}, &Node{}
	stack.Put(root)

	for {
		token, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return cw.n, err
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
			n.Examples = append(n.Examples, v)
		}
	}
	if len(root.Children) > 0 {
		*node = *root.Children[0]
	}
	return cw.n, nil
}

// attrListContains returns true, if the attribute name is contained in the
// given attribute list.
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
	for _, f := range node.freqs {
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
		node.mergeAttr(attr)
		node.childFrequencies[name]++
		return child
	}

	if node.childFrequencies == nil {
		node.childFrequencies = make(map[xml.Name]int)
	}
	node.childFrequencies[name]++

	n := &Node{Name: name, Attr: attr}
	node.Children = append(node.Children, n)
	return n
}

// End signals end of an element.
func (node *Node) End() {
	// Take note of frequencies, collect them inside node for later stats.
	for name, freq := range node.childFrequencies {
		for _, c := range node.Children {
			if c.Name != name {
				continue
			}
			node.freqs = append(node.freqs, freq)
		}
	}
	// Reset counter.
	node.childFrequencies = make(map[xml.Name]int)
}
