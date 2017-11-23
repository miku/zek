package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/miku/zek"
)

var maxExamples = flag.Int("x", 10, "max number of examples")

func main() {
	flag.Parse()

	dec := xml.NewDecoder(os.Stdin)
	dec.Strict = false

	stack, root := zek.Stack{}, &zek.Node{}
	stack.Put(root)

	for {
		token, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		switch t := token.(type) {
		case xml.StartElement:
			parent := stack.Peek().(*zek.Node)
			n := parent.CreateOrGetChild(t.Name, t.Attr)
			stack.Put(n)
		case xml.EndElement:
			n := stack.Pop().(*zek.Node)
			n.End()
		case xml.CharData:
			v := strings.TrimSpace(string(t))
			if v == "" {
				break
			}
			node := stack.Peek().(*zek.Node)
			if len(node.Examples) < *maxExamples {
				node.Examples = append(node.Examples, v)
			}
		}
	}
	b, err := json.Marshal(root.Children[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
