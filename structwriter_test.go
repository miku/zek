package zek

import (
	"bytes"
	"go/format"
	"strings"
	"testing"
)

func TestWriteNode(t *testing.T) {
	var cases = []struct {
		input  string // Input XML.
		result string // Generated struct.
		err    error
	}{
		{
			input:  "",
			result: "",
			err:    nil,
		},
		{
			input: "<a></a>",
			result: `
type Document struct {
	A string ` + "`xml:" + `"a"` + "`" + "\n}",
			err: nil,
		},
	}

	for _, c := range cases {
		node := new(Node)
		if _, err := node.ReadFrom(strings.NewReader(c.input)); err != nil {
			t.Errorf("failed to read XML input: %s", err)
		}

		var buf bytes.Buffer
		sw := NewStructWriter(&buf)

		if err := sw.WriteNode(node); err != c.err {
			t.Errorf("got %v, want %v", err, c.err)
		}
		formatted, err := format.Source(buf.Bytes())
		if err != nil {
			t.Errorf("format failed: %s", err)
		}
		if string(formatted) != c.result {
			t.Errorf("got %v, want %v", string(formatted), c.result)
		}
	}
}
