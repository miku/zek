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

// truncateString after n chars and append ellipsis.
func truncateString(s string, n int, ellipsis string) string {
	if len(s) < n {
		return s
	}
	return fmt.Sprintf("%s%s", s[:n], ellipsis)
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
	WithComments      bool                // Annotate struct with comments and examples.
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
	if node == nil || reflect.DeepEqual(node, new(Node)) {
		return nil
	}
	return sw.writeNode(node, true)
}

// writeField writes a field with a simple xml struct tag to writer.
func (sw *StructWriter) writeNameField(w io.Writer, node *Node) (int, error) {
	return fmt.Fprintf(w, "XMLName xml.Name `xml:\"%s\"`\n", node.Name.Local)
}

func (sw *StructWriter) writeStructTag(w io.Writer, node *Node) (int, error) {
	return fmt.Fprintf(w, "`xml:\"%s\"`", node.Name.Local)
}

// writeChardataField writes a chardata field. Might add a comment as well.
func (sw *StructWriter) writeChardataField(w io.Writer, node *Node) (int, error) {
	s := fmt.Sprintf("%s string `xml:\",chardata\"`", sw.TextFieldName)
	if sw.WithComments && len(node.Examples) > 0 {
		s = fmt.Sprintf("%s // %s", s, truncateString(strings.Join(node.Examples, ", "), 25, "..."))
	}
	return fmt.Fprintf(w, "%s\n", s)
}

// writeAttrField writes an attribute field.
func (sw *StructWriter) writeAttrField(w io.Writer, name, typeName, tag string) (int, error) {
	return fmt.Fprintf(w, "%s %s `xml:\"%s,attr\"`\n", name, typeName, tag)
}

// writeStructIntro writes the nodes current field name name and struct introduction.
func (sw *StructWriter) writeStructIntro(w io.Writer, node *Node) {
	io.WriteString(w, sw.NameFunc(node.Name.Local))
	io.WriteString(w, " ")
	if node.IsMultivalued() {
		io.WriteString(w, "[]")
	}
	io.WriteString(w, "struct {\n")
}

// writeNode writes out the node as a struct. Output is not formatted.
func (sw *StructWriter) writeNode(node *Node, top bool) (err error) {
	sew := stickyErrWriter{w: sw.w, err: &err}
	if top {
		io.WriteString(sew, "type ")
	}
	sw.writeStructIntro(sew, node)
	if top {
		sw.writeNameField(sew, node)
	}
	sw.writeChardataField(sew, node)

	// Helper to check for name clash with any generated field name.
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
			return fmt.Errorf("name clash: %s", attr.Name.Local)
		}
		sw.writeAttrField(sew, name, "string", attr.Name.Local)
	}

	// Write children.
	for _, child := range node.Children {
		sw.writeNode(child, false)
	}

	// Write outro.
	io.WriteString(sew, "} ")
	if !top {
		sw.writeStructTag(sew, node)
	}
	io.WriteString(sew, "\n")
	return err
}
