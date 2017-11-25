package zek

import (
	"fmt"
	"io"
	"reflect"
	"strings"
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

// stringSliceContains returns true, if a string is found in a slice.
func stringSliceContains(ss []string, s string) bool {
	for _, v := range ss {
		if v == s {
			return true
		}
	}
	return false
}

// createNameFunc returns a function that converts a tag into a canonical Go
// name. Given list of strings will be uppercased.
func createNameFunc(upper []string) func(string) string {
	f := func(name string) string {
		var capped []string
		splitter := func(c rune) bool {
			return c == '_' || c == '-'
		}
		for _, s := range strings.FieldsFunc(name, splitter) {
			switch {
			case stringSliceContains(upper, strings.ToLower(s)):
				capped = append(capped, strings.ToUpper(s))
			default:
				capped = append(capped, strings.Title(s))
			}
		}
		return strings.Join(capped, "")
	}
	return f
}

// StructWriter can turn a node into a struct and can be configured.
type StructWriter struct {
	w        io.Writer
	NameFunc func(string) string // Turns xml tag names into Go names.
}

// NewStructWriter can write a node to a given writer. Default list of abbreviations to wholly uppercase.
func NewStructWriter(w io.Writer) *StructWriter {
	exceptions := []string{"id", "isbn", "ismn", "eissn", "issn", "lccn", "rsn", "url", "urn", "zdb"}
	return &StructWriter{w: w, NameFunc: createNameFunc(exceptions)}
}

// WriteNode writes a node to a writer. XXX: Implement.
func (sw *StructWriter) WriteNode(node *Node) (err error) {
	if sw.w == nil {
		return nil
	}
	if node == nil || reflect.DeepEqual(node, new(Node)) {
		return nil
	}
	return sw.writeNode(node, true)
}

// writeNode writes a node.
func (sw *StructWriter) writeNode(node *Node, top bool) (err error) {
	sew := stickyErrWriter{w: sw.w, err: &err}
	if top {
		io.WriteString(sew, "type ")
	}
	io.WriteString(sew, sw.NameFunc(node.Name.Local))
	io.WriteString(sew, " ")
	if node.IsMultivalued() {
		io.WriteString(sew, "[]")
	}
	io.WriteString(sew, "struct { \n")
	io.WriteString(sew, fmt.Sprintf("XMLName xml.Name `xml:\"%s\"`\n", node.Name.Local))
	io.WriteString(sew, fmt.Sprintf("Text string `xml:\",chardata\"`\n"))
	// XXX: check, if there is a direct child with this name.
	for _, attr := range node.Attr {
		for _, child := range node.Children {
			if sw.NameFunc(child.Name.Local) == sw.NameFunc(attr.Name.Local) {
				// XXX: resolve.
				return fmt.Errorf("name clash between attribute and child node")
			}
		}
		io.WriteString(sew, fmt.Sprintf("%s string `xml:\"%s,attr\"`\n", sw.NameFunc(attr.Name.Local), attr.Name.Local))
	}
	// XXX: check, if there is a child named "text" and rename it.
	for _, child := range node.Children {
		sw.writeNode(child, false)
	}
	io.WriteString(sew, "}\n")
	return err
}
