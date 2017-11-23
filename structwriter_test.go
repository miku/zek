package zek

import (
	"bytes"
	"encoding/xml"
	"testing"
)

func TestWriteNode(t *testing.T) {
	var cases = []struct {
		node   *Node
		result string
		err    error
	}{
		{
			node:   &Node{},
			result: "",
			err:    nil,
		},
		{
			node: &Node{
				Name: xml.Name{Local: "a"},
			},
			result: "type A string",
			err:    nil,
		},
	}

	for _, c := range cases {
		var buf bytes.Buffer
		sw := NewStructWriter(&buf)
		err := sw.WriteNode(c.node)

		if err != c.err {
			t.Errorf("got %v, want %v", err, c.err)
		}
		if buf.String() != c.result {
			t.Errorf("got %v, want %v", buf.String(), c.result)
		}
	}
}
