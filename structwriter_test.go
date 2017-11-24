package zek

import (
	"bytes"
	"testing"
)

func TestWriteNode(t *testing.T) {
	var cases = []struct {
		node   *Node
		result string
		err    error
	}{}

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
