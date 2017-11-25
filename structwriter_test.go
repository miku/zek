package zek

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"os"
	"reflect"
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
		input  string // Input XML filename.
		result string // Generated struct filename.
		err    error
	}{
		{
			input:  "testdata/w.1.xml",
			result: "testdata/w.1.go",
			err:    nil,
		},
		{
			input:  "testdata/w.2.xml",
			result: "testdata/w.2.go",
			err:    nil,
		},
		{
			input:  "testdata/w.3.xml",
			result: "testdata/w.3.go",
			err:    nil,
		},
	}

	for _, c := range cases {
		f, err := os.Open(c.input)
		if err != nil {
			t.Errorf("cannot open test input file: %s", err)
		}
		defer f.Close()

		node := new(Node)
		if _, err := node.ReadFrom(f); err != nil {
			t.Errorf("failed to read XML input: %s", err)
		}

		var buf bytes.Buffer
		sw := NewStructWriter(&buf)
		if err := sw.WriteNode(node); err != c.err {
			t.Errorf("WriteNode failed: got %v, want %v", err, c.err)
		}

		b, err := ioutil.ReadFile(c.result)
		if err != nil {
			t.Errorf("cannot open test output file: %s", err)
		}

		if ok, err := codeEquals(buf.Bytes(), b); !ok || err != nil {
			if err != nil {
				t.Errorf("failed to compare code snippets: %s", err)
			}
			t.Errorf("got %v, want %v", buf.String(), string(b))
		}
	}
}
