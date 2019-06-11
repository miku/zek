package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
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
	withJSONTags         = flag.Bool("j", false, "add JSON tags")
	maxExamples          = flag.Int("max-examples", 10, "limit number of examples")
	debug                = flag.Bool("d", false, "debug output")
	createExampleProgram = flag.Bool("p", false, "write out an example program")
	tagName              = flag.String("t", "", "emit struct for tag matching this name")
	skipFormatting       = flag.Bool("F", false, "skip formatting")
	strict               = flag.Bool("s", false, "strict parsing and writing")
	exampleMaxChars      = flag.Int("x", 25, "max chars for example")
	version              = flag.Bool("version", false, "show version")
	structName           = flag.String("n", "", "use a different name for the top-level struct")
	compact              = flag.Bool("c", false, "emit more compact struct (noop, as this is the default since 0.1.7)")
	nonCompact           = flag.Bool("C", false, "emit less compact struct")
	uniqueExamples       = flag.Bool("u", false, "filter out duplicated examples")
)

func main() {
	flag.Parse()

	root := new(zek.Node)
	root.MaxExamples = *maxExamples

	if *version {
		fmt.Println(zek.Version)
		os.Exit(0)
	}

	// Read one or more XML files given as arguments.
	if flag.NArg() > 0 {
		var readers []io.Reader
		for _, fn := range flag.Args() {
			f, err := os.Open(fn)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			readers = append(readers, f)
		}
		if _, err := root.ReadFromAll(readers); err != nil {
			log.Fatal(err)
		}
	} else {
		if _, err := root.ReadFrom(os.Stdin); err != nil {
			log.Fatal(err)
		}
	}

	// Move root, if we have a tagName. Ignore unknown names.
	if *tagName != "" {
		if n := root.ByName(*tagName); n != nil {
			root = n
		}
	}

	if *structName != "" {
		root.Name = xml.Name{Space: "", Local: *structName}
	}

	switch {
	default:
		var buf bytes.Buffer
		sw := zek.NewStructWriter(&buf)
		sw.WithComments = *withComments
		sw.WithJSONTags = *withJSONTags
		sw.Strict = *strict
		sw.ExampleMaxChars = *exampleMaxChars
		sw.Compact = !*nonCompact
		sw.UniqueExamples = *uniqueExamples

		if err := sw.WriteNode(root); err != nil {
			log.Fatal(err)
		}
		if !*skipFormatting {
			b, err := format.Source(buf.Bytes())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(b))
		} else {
			fmt.Println(buf.String())
		}
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

			import (
				"encoding/json"
				"encoding/xml"
				"fmt"
				"log"
				"os"

				"golang.org/x/net/html/charset"
			)
		`)

		sw := zek.NewStructWriter(&buf)
		sw.WithComments = *withComments
		sw.Strict = *strict
		sw.ExampleMaxChars = *exampleMaxChars
		sw.Compact = *compact

		if err := sw.WriteNode(root); err != nil {
			log.Fatal(err)
		}

		io.WriteString(&buf, fmt.Sprintf(`
			func main() {
				dec := xml.NewDecoder(os.Stdin)
				dec.CharsetReader = charset.NewReaderLabel
				dec.Strict = false

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

		if !*skipFormatting {
			b, err := format.Source(buf.Bytes())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(b))
		} else {
			fmt.Println(buf.String())
		}
	}
}
