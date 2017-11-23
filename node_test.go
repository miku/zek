package zek

import (
	"encoding/xml"
	"reflect"
	"testing"
)

func TestAttrListContains(t *testing.T) {
	var cases = []struct {
		attrs  []xml.Attr
		attr   xml.Attr
		result bool
	}{
		{
			attrs:  []xml.Attr{xml.Attr{}},
			attr:   xml.Attr{},
			result: true,
		},
		{
			attrs:  []xml.Attr{xml.Attr{}},
			attr:   xml.Attr{Name: xml.Name{Local: "id"}},
			result: false,
		},
		{
			attrs:  []xml.Attr{xml.Attr{Name: xml.Name{Local: "id"}}},
			attr:   xml.Attr{Name: xml.Name{Local: "id"}},
			result: true,
		},
		{
			attrs:  []xml.Attr{xml.Attr{Name: xml.Name{Local: "id"}, Value: "yyy"}},
			attr:   xml.Attr{Name: xml.Name{Local: "id"}, Value: "xxx"},
			result: true,
		},
	}

	for _, c := range cases {
		got := attrListContains(c.attrs, c.attr)
		if c.result != got {
			t.Errorf("got %v, want %v", got, c.result)
		}
	}
}

func TestNodeMergeAttr(t *testing.T) {
	var cases = []struct {
		node   *Node
		attr   []xml.Attr
		result []xml.Attr
	}{
		{
			node:   &Node{Attr: []xml.Attr{}},
			attr:   []xml.Attr{},
			result: []xml.Attr{},
		},
		{
			node:   &Node{Attr: []xml.Attr{}},
			attr:   []xml.Attr{xml.Attr{Name: xml.Name{Local: "id"}}},
			result: []xml.Attr{xml.Attr{Name: xml.Name{Local: "id"}}},
		},
		{
			node:   &Node{Attr: []xml.Attr{xml.Attr{Name: xml.Name{Local: "id"}}}},
			attr:   []xml.Attr{xml.Attr{Name: xml.Name{Local: "id"}}},
			result: []xml.Attr{xml.Attr{Name: xml.Name{Local: "id"}}},
		},
		{
			node: &Node{Attr: []xml.Attr{xml.Attr{Name: xml.Name{Local: "id"}}}},
			attr: []xml.Attr{xml.Attr{Name: xml.Name{Local: "xx"}}},
			result: []xml.Attr{
				xml.Attr{Name: xml.Name{Local: "id"}},
				xml.Attr{Name: xml.Name{Local: "xx"}},
			},
		},
	}
	for _, c := range cases {
		c.node.mergeAttr(c.attr)
		if !reflect.DeepEqual(c.node.Attr, c.result) {
			t.Errorf("got %v, want %v", c.node.Attr, c.result)
		}
	}
}
