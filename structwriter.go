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
// name. Given list of strings will be wholly upper cased.
func createNameFunc(upper []string) func(string) string {
	f := func(name string) string {
		var capped []string
		splitter := func(c rune) bool {
			return c == '_' || c == '-' || c == '.'
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
	w                 io.Writer
	NameFunc          func(string) string // Turns xml tag names into Go names.
	TextFieldName     string              // Field name for chardata.
	AttributePrefixes []string            // In case of a name clash, try these prefixes.
}

// NewStructWriter can write a node to a given writer. Default list of
// abbreviations to wholly uppercase.
func NewStructWriter(w io.Writer) *StructWriter {
	exceptions := []string{"id", "isbn", "ismn", "eissn", "issn", "lccn", "rfc", "rsn", "url", "urn", "zdb"}
	return &StructWriter{
		w:                 w,
		NameFunc:          createNameFunc(exceptions),
		TextFieldName:     "Text",
		AttributePrefixes: []string{"Attr", "Attribute"},
	}
}

// WriteNode writes a node to a writer.
func (sw *StructWriter) WriteNode(node *Node) (err error) {
	if sw.w == nil {
		return nil
	}
	empty := new(Node)
	if node == nil || reflect.DeepEqual(node, empty) {
		return nil
	}
	return sw.writeNode(node, true)
}

// writeField writes a field with a simple xml struct tag to writer.
func writeField(w io.Writer, name, typeName, tag string) (int, error) {
	return fmt.Fprintf(w, "%s %s `xml:\"%s\"`\n", name, typeName, tag)
}

// writeChardataField writes a chardata field.
func writeChardataField(w io.Writer, name, typeName string) (int, error) {
	return fmt.Fprintf(w, "%s %s `xml:\",chardata\"`\n", name, typeName)
}

// writeAttrField writes an attribute field.
func writeAttrField(w io.Writer, name, typeName, tag string) (int, error) {
	return fmt.Fprintf(w, "%s %s `xml:\"%s,attr\"`\n", name, typeName, tag)
}

// writeNode writes out the node as a struct. Output is not formatted.
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
	io.WriteString(sew, "struct {\n")

	writeField(sew, "XMLName", "xml.Name", node.Name.Local)
	writeChardataField(sew, sw.TextFieldName, "string")

	// Returns false if the given name would clash with any of the generated field names.
	isValidName := func(name string) bool {
		if name == sw.TextFieldName {
			return false
		}
		for _, child := range node.Children {
			if name == sw.NameFunc(child.Name.Local) {
				return false
			}
		}
		return true
	}

	// Write attributes.
	for _, attr := range node.Attr {
		name := sw.NameFunc(attr.Name.Local)
		for _, prefix := range sw.AttributePrefixes {
			if isValidName(name) {
				break
			}
			name = fmt.Sprintf("%s%s", prefix, name)
		}
		if !isValidName(name) {
			return fmt.Errorf("name clash (adjust AttributePrefix list): %s",
				attr.Name.Local)
		}
		writeAttrField(sew, name, "string", attr.Name.Local)
	}

	// Write children.
	for _, child := range node.Children {
		sw.writeNode(child, false)
	}

	io.WriteString(sew, "}\n")
	return err
}
