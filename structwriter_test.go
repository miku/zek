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
	"bytes"
	"go/format"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
	"testing"
)

// startsWithType pattern to match where a type declaration starts.
var startsWithType = regexp.MustCompile(`(?m)^type`)

// removeComments removes simple comments.
func removeComments(b []byte) []byte {
	var buf bytes.Buffer
	re := regexp.MustCompile(`^//.*`)
	for _, line := range bytes.Split(b, []byte("\n")) {
		buf.Write(re.ReplaceAll(line, nil))
		io.WriteString(&buf, "\n")
	}
	return buf.Bytes()
}

// codeEquals returns true, if two code snippets match.
func codeEquals(a, b []byte) (bool, error) {
	a = bytes.TrimSpace(removeComments(a))
	b = bytes.TrimSpace(removeComments(b))

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

// skipUntilType remove any line before the first type declaration. Used for
// testing, so snippets may be valid go code (and formatable).
func skipUntilType(b []byte) []byte {
	if loc := startsWithType.FindIndex(b); loc != nil {
		return b[loc[0]:]
	}
	return b
}

func TestWriteNode(t *testing.T) {
	var cases = []struct {
		input          string // Input XML filename.
		result         string // Generated struct filename.
		withComments   bool
		uniqueExamples bool
		err            error
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
		{
			input:  "testdata/w.4.xml",
			result: "testdata/w.4.go",
			err:    nil,
		},
		{
			input:  "testdata/w.5.xml",
			result: "testdata/w.5.go",
			err:    nil,
		},
		{
			input:  "testdata/w.6.xml",
			result: "testdata/w.6.go",
			err:    nil,
		},
		{
			input:  "testdata/w.7.xml",
			result: "testdata/w.7.go",
			err:    nil,
		},
		{
			input:        "testdata/w.8.xml",
			result:       "testdata/w.8.go",
			withComments: true,
			err:          nil,
		},
		{
			input:        "testdata/w.9.xml",
			result:       "testdata/w.9.go",
			withComments: true,
			err:          nil,
		},
		{
			input:        "testdata/w.10.xml",
			result:       "testdata/w.10.go",
			withComments: true,
			err:          nil,
		},
		{
			input:        "testdata/w.11.xml",
			result:       "testdata/w.11.go",
			withComments: true,
			err:          nil,
		},
		{
			input:          "testdata/w.12.xml",
			result:         "testdata/w.12.go",
			withComments:   true,
			uniqueExamples: false,
			err:            nil,
		},
		{
			input:          "testdata/w.13.xml",
			result:         "testdata/w.13.go",
			withComments:   true,
			uniqueExamples: true,
			err:            nil,
		},
	}

	for _, c := range cases {
		f, err := os.Open(c.input)
		if err != nil {
			t.Errorf("cannot open test input file: %s", err)
		}
		defer f.Close()

		node := new(Node)
		node.MaxExamples = 10

		if _, err := node.ReadFrom(f); err != nil {
			t.Errorf("failed to read XML input: %s", err)
		}

		var buf bytes.Buffer
		sw := NewStructWriter(&buf)
		sw.WithComments = c.withComments
		sw.UniqueExamples = c.uniqueExamples

		if err := sw.WriteNode(node); err != c.err {
			t.Errorf("WriteNode failed: got %v, want %v", err, c.err)
		}

		b, err := ioutil.ReadFile(c.result)
		if err != nil {
			t.Errorf("cannot open test output file: %s", err)
		}
		b = skipUntilType(b)

		if ok, err := codeEquals(buf.Bytes(), b); !ok || err != nil {
			if err != nil {
				t.Errorf("failed to compare code snippets: %s", err)
			}
			t.Errorf("[%s] got %v, want %v", c.input, buf.String(), string(b))
		}
	}
}
