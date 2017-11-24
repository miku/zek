package zek

import (
	"bytes"
	"go/format"
	"reflect"
	"strings"
	"testing"
)

// codeEquals returns true, if two code snippets match.
func codeEquals(a, b []byte) (bool, error) {
	fa, err := format.Source(bytes.TrimSpace(a))
	if err != nil {
		return false, err
	}
	fb, err := format.Source(bytes.TrimSpace(b))
	if err != nil {
		return false, err
	}
	return reflect.DeepEqual(fa, fb), nil
}

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
			input:  "<a></a>",
			result: `type Document struct { A string ` + "`xml:" + `"a"` + "`\n}",
			err:    nil,
		},
		{
			input:  "<a><b></b></a>",
			result: `type Document struct { A string ` + "`xml:" + `"a"` + "`\n}",
			err:    nil,
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
		if ok, err := codeEquals(buf.Bytes(), []byte(c.result)); !ok || err != nil {
			if err != nil {
				t.Errorf("failed to compare code snippets: %s", err)
			}
			t.Errorf("got %v, want %v", buf.String(), c.result)
		}
	}
}
