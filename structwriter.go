package zek

import (
	"fmt"
	"io"
	"reflect"
	"strings"
	"sync"
)

type stickyErrWriter struct {
	w   io.Writer
	err *error
}

// Write returns early, if write has failed in the past.
func (sew stickyErrWriter) Write(p []byte) (n int, err error) {
	if *sew.err != nil {
		return 0, *sew.err
	}
	n, err = sew.w.Write(p)
	*sew.err = err
	return
}

// GoName generates a camel case name for a given string. Allows exceptions for
// certain fields, e.g. ISBN.
func GoName(name string) string {
	var capped []string
	splitter := func(c rune) bool {
		return c == '_' || c == '-'
	}
	for _, s := range strings.FieldsFunc(name, splitter) {
		switch strings.ToLower(s) {
		case "id", "isbn", "ismn", "eissn", "issn", "lccn", "rsn", "url", "urn", "zdb":
			capped = append(capped, strings.ToUpper(s))
		default:
			capped = append(capped, strings.Title(s))
		}
	}
	return strings.Join(capped, "")
}

// StructWriter can turn a node into a struct and can be configured.
type StructWriter struct {
	w        io.Writer
	NameFunc func(string) string
	onStart  sync.Once
	onEnd    sync.Once
}

// NewStructWriter can write a node to a given writer.
func NewStructWriter(w io.Writer) *StructWriter {
	return &StructWriter{w: w, NameFunc: GoName}
}

// WriteNode writes a node to a writer. XXX: Implement.
func (sw *StructWriter) WriteNode(node *Node) (err error) {
	if sw.w == nil {
		return nil
	}
	if node == nil || reflect.DeepEqual(node, new(Node)) {
		return nil
	}

	sew := stickyErrWriter{w: sw.w, err: &err}

	sw.onStart.Do(func() {
		io.WriteString(sew, "type Document struct {\n")
	})

	name := sw.NameFunc(node.Name.Local)

	io.WriteString(sew, name)
	io.WriteString(sew, " ")

	if node.IsMultivalued() {
		io.WriteString(sew, "[]")
	}

	if len(node.Children) == 0 {
		s := fmt.Sprintf("string `xml:\"%s\"`", node.Name.Local)
		io.WriteString(sew, s)
	} else {
		io.WriteString(sew, "struct {\n")
		for _, child := range node.Children {
			sw.WriteNode(child)
		}
		io.WriteString(sew, "}")
	}

	sw.onEnd.Do(func() {
		io.WriteString(sew, "}\n")
	})

	return *sew.err
}
