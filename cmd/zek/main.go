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

var createExampleProgram = flag.Bool("p", false, "write out an example program")

func main() {
	flag.Parse()

	root := new(zek.Node)
	if _, err := root.ReadFrom(os.Stdin); err != nil {
		log.Fatal(err)
	}

	if *createExampleProgram {
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
		os.Exit(0)
	}

	b, err := json.Marshal(root)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
