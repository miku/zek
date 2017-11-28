package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"io"
	"log"
	"os"

	"github.com/miku/zek"
)

var (
	withComments         = flag.Bool("e", false, "add comments with example")
	maxExamples          = flag.Int("max-examples", 10, "limit number of examples")
	debug                = flag.Bool("d", false, "debug output")
	createExampleProgram = flag.Bool("p", false, "write out an example program")
	tagName              = flag.String("t", "", "emit struct for tag matching this name")
)

func main() {
	flag.Parse()

	root := new(zek.Node)
	root.MaxExamples = *maxExamples
	if _, err := root.ReadFrom(os.Stdin); err != nil {
		log.Fatal(err)
	}

	// Move root, if we have a tagName. Ignore unknown names.
	if *tagName != "" {
		if n := root.ByName(*tagName); n != nil {
			root = n
		}
	}

	switch {
	default:
		var buf bytes.Buffer
		sw := zek.NewStructWriter(&buf)
		sw.WithComments = *withComments
		if err := sw.WriteNode(root); err != nil {
			log.Fatal(err)
		}
		b, err := format.Source(buf.Bytes())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
	case *debug:
		b, err := json.Marshal(root)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
	case *createExampleProgram:
		var buf bytes.Buffer
		io.WriteString(&buf, `
			package main
			import "encoding/xml"
			import "os"
			import "encoding/json"
			import "log"
			import "fmt"
		`)

		sw := zek.NewStructWriter(&buf)
		sw.WithComments = *withComments
		if err := sw.WriteNode(root); err != nil {
			log.Fatal(err)
		}

		io.WriteString(&buf, fmt.Sprintf(`
			func main() {
				dec := xml.NewDecoder(os.Stdin)
				var doc %s
				if err := dec.Decode(&doc); err != nil {
					log.Fatal(err)
				}
				b, err := json.Marshal(doc)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(string(b))
			}
		`, sw.NameFunc(root.Name.Local)))

		b, err := format.Source(buf.Bytes())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
	}
}
