zek
===

Zek is a **prototype** for turning an XML directly to a Go struct.

Upsides:

* it works fine for non-recursive structures,
* it is relatively convenient to access attributes, children and text,
* will generate a single struct, which make for a quite compact representation,
* simple user interface,
* comments with examples,
* O(1) space complexity.

Downsides:

* no support for recursive types (similar to *Russian Doll-ish* strategy, [[1](https://medbiq.org/std_specs/techguidelines/xmldesignguidelines.pdf#page=7)])
* no type inference, everything is accessible as string.

Usage
-----

```shell
$ zek -h
Usage of zek:
  -d    debug output
  -p    write out an example program
```

Examples:

```xml
$ cat fixtures/a.xml
<a></a>

$ zek < fixtures/a.xml
type A struct {
    XMLName xml.Name `xml:"a"`
    Text    string   `xml:",chardata"`
}
```

Debug output dumps the internal tree as JSON to stdout.

```json
$ zek -d < fixtures/a.xml
{"name":{"Space":"","Local":"a"}}
```

Example program:

```go
$ zek -p < fixtures/a.xml
package main

import "encoding/xml"
import "os"
import "encoding/json"
import "log"
import "fmt"

type A struct {
	XMLName xml.Name `xml:"a"`
	Text    string   `xml:",chardata"`
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var doc A
	if err := dec.Decode(&doc); err != nil {
		log.Fatal(err)
	}
	b, err := json.Marshal(doc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}

$ zek -p < fixtures/a.xml > sample.go && go run sample.go < fixtures/a.xml | jq . && rm sample.go
{
  "XMLName": {
    "Space": "",
    "Local": "a"
  },
  "Text": ""
}
```

More complex example:

```go
$ zek < fixtures/d.xml
type Root struct {
	XMLName xml.Name `xml:"root"`
	Text    string   `xml:",chardata"`
	A       []struct {
		Text string `xml:",chardata"`
		B    []struct {
			Text string `xml:",chardata"`
			C    struct {
				Text string `xml:",chardata"`
			} `xml:"c"`
			D struct {
				Text string `xml:",chardata"`
			} `xml:"d"`
		} `xml:"b"`
	} `xml:"a"`
}

$ zek -p < fixtures/d.xml > sample.go && go run sample.go < fixtures/d.xml | jq . && rm sample.go
{
  "XMLName": {
    "Space": "",
    "Local": "root"
  },
  "Text": "\n\n\n\n",
  "A": [
    {
      "Text": "\n  \n  \n",
      "B": [
        {
          "Text": "\n    \n  ",
          "C": {
            "Text": "Hi"
          },
          "D": {
            "Text": ""
          }
        },
        {
          "Text": "\n    \n    \n  ",
          "C": {
            "Text": "World"
          },
          "D": {
            "Text": ""
          }
        }
      ]
    },
    {
      "Text": "\n  \n",
      "B": [
        {
          "Text": "\n    \n  ",
          "C": {
            "Text": "Hello"
          },
          "D": {
            "Text": ""
          }
        }
      ]
    },
    {
      "Text": "\n  \n",
      "B": [
        {
          "Text": "\n    \n  ",
          "C": {
            "Text": ""
          },
          "D": {
            "Text": "World"
          }
        }
      ]
    }
  ]
}
```
