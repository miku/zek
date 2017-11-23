package zek

import (
	"encoding/xml"
)

// Node represents an element in the XML tree.
type Node struct {
	Name     xml.Name   `json:"name"`
	Attr     []xml.Attr `json:"attr"`
	Examples []string   `json:"examples"`
	Children []*Node

	freqs            []int // freqs counts the number of occurences of this node within parent.
	childFrequencies map[xml.Name]int
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
