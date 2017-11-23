package zek

import (
	"fmt"
	"io"
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

// StructWriter can turn a tree of nodes into a struct.
type StructWriter struct {
	w              io.Writer
	UppercaseNames []string
}

func NewStructWriter(w io.Writer) *StructWriter {
	return &StructWriter{w: w}
}

// WriteNode writes a node to a writer. XXX: Implement.
func (sw *StructWriter) WriteNode(node *Node) (err error) {
	if sw.w == nil {
		return nil
	}
	if node == nil {
		return nil
	}

	sew := stickyErrWriter{w: sw.w, err: &err}

	// Leaf element.
	if len(node.Children) == 0 {
		s := fmt.Sprintf("string `xml:\"%s\"`", node.Name.Local)
		io.WriteString(sew, s)
	}
	return nil
}
