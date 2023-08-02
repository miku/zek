// Copyright 2021 by Leipzig University Library, http://ub.uni-leipzig.de
//                   The Finc Authors, http://finc.info
//                   Martin Czygan, <martin.czygan@uni-leipzig.de>
//
// This file is part of some open source application.
//
// Some open source application is free software: you can redistribute
// it and/or modify it under the terms of the GNU General Public
// License as published by the Free Software Foundation, either
// version 3 of the License, or (at your option) any later version.
//
// Some open source application is distributed in the hope that it will
// be useful, but WITHOUT ANY WARRANTY; without even the implied warranty
// of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Foobar.  If not, see <http://www.gnu.org/licenses/>.
//
// @license GPL-3.0+ <http://spdx.org/licenses/GPL-3.0+>

package zek

import (
	"encoding/xml"
	"io"
	"reflect"
	"strings"

	"golang.org/x/net/html/charset"
)

var emptyNode = new(Node)

// Node represents an element in the XML tree. It keeps track of its name,
// attributes, childnodes and example chardata and basic statistics, e.g. how
// often a node has been seen within its parent node.
type Node struct {
	Name        xml.Name         `json:"name,omitempty"`
	Attr        []xml.Attr       `json:"attr,omitempty"`
	Examples    []string         `json:"examples,omitempty"`
	Children    []*Node          `json:"children,omitempty"`
	Freqs       []int            `json:"-"` // Collect number of occurrences of this node within parent.
	MaxExamples int              `json:"-"` // Maximum number of examples to keep, gets passed to children.
	childFreqs  map[xml.Name]int // Count child tag occurrences, used temporarily.
}

// readNode reads XML from a reader and returns a parsed node. If node is
// given, it is reused, allowing for multiple passes (e.g. from multiple
// files). XXX: maxExamples should be factored out into options.
func readNode(r io.Reader, root *Node, maxExamples int) (node *Node, n int64, err error) {
	var (
		cw  = countwriter{}
		dec = xml.NewDecoder(io.TeeReader(r, &cw))
	)
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

// ReadFrom reads XML from a reader. TODO: pass read options.
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

// mergeAttr adds attributes to node, which are not already there.
func (node *Node) mergeAttr(attr []xml.Attr) {
LOOP:
	for _, a := range attr {
		for _, na := range node.Attr {
			if a.Name == na.Name {
				continue LOOP
			}
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
	n := &Node{
		Name:        name,
		Attr:        attr,
		MaxExamples: node.MaxExamples,
	}
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
	if node == nil || reflect.DeepEqual(node, emptyNode) {
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

// ByName finds a node in the tree (dfs) by name. Comparisons start at the
// current node. First match is returned. If nothing matches, nil is returned.
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
