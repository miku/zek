zek
===

Zek is a **prototype** for creating Go struct from an XML document.

Upsides:

* it works fine for non-recursive structures,
* does not need XSD or DTD,
* it is relatively convenient to access attributes, children and text,
* will generate a single struct, which make for a quite compact representation,
* simple user interface,
* comments with examples,
* O(1) space complexity.

Downsides:

* no support for recursive types (similar to *Russian Doll-ish* strategy, [[1](https://medbiq.org/std_specs/techguidelines/xmldesignguidelines.pdf#page=7)])
* no type inference, everything is accessible as string.

Install
-------

```
$ go get github.com/miku/zek/cmd/...
```

Debian and RPM packages:

* https://github.com/miku/zek/releases

Usage
-----

```shell
$ zek -h
Usage of zek:
  -d    debug output
  -e    add comments with example
  -max-examples int
        limit number of examples (default 10)
  -p    write out an example program
  -t string
        emit struct for tag matching this name
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

Annotate with comments:

```go
$ zek -e < fixtures/l.xml
type Records struct {
	XMLName xml.Name `xml:"Records"`
	Text    string   `xml:",chardata"` // \n
	Xsi     string   `xml:"xsi,attr"`
	Record  []struct {
		Text   string `xml:",chardata"`
		Header struct {
			Text       string `xml:",chardata"`
			Status     string `xml:"status,attr"`
			Identifier struct {
				Text string `xml:",chardata"` // oai:ojs.localhost:article...
			} `xml:"identifier"`
			Datestamp struct {
				Text string `xml:",chardata"` // 2009-06-24T14:48:23Z, 200...
			} `xml:"datestamp"`
			SetSpec struct {
				Text string `xml:",chardata"` // eppp:ART, eppp:ART, eppp:...
			} `xml:"setSpec"`
		} `xml:"header"`
		Metadata struct {
			Text    string `xml:",chardata"`
			Rfc1807 struct {
				Text           string `xml:",chardata"`
				Xmlns          string `xml:"xmlns,attr"`
				Xsi            string `xml:"xsi,attr"`
				SchemaLocation string `xml:"schemaLocation,attr"`
				BibVersion     struct {
					Text string `xml:",chardata"` // v2, v2, v2, v2, v2, v2, v...
				} `xml:"bib-version"`
				ID struct {
					Text string `xml:",chardata"` // http://journals.zpid.de/i...
				} `xml:"id"`
				Entry struct {
					Text string `xml:",chardata"` // 2009-06-24T14:48:23Z, 200...
				} `xml:"entry"`
				Organization []struct {
					Text string `xml:",chardata"` // Proceedings of the Worksh...
				} `xml:"organization"`
				Title struct {
					Text string `xml:",chardata"` // Introduction and some Ide...
				} `xml:"title"`
				Type struct {
					Text string `xml:",chardata"`
				} `xml:"type"`
				Author []struct {
					Text string `xml:",chardata"` // KRAMPEN, Günter, CARBON,...
				} `xml:"author"`
				Copyright struct {
					Text string `xml:",chardata"` // Das Urheberrecht liegt be...
				} `xml:"copyright"`
				OtherAccess struct {
					Text string `xml:",chardata"` // url:http://journals.zpid....
				} `xml:"other_access"`
				Keyword struct {
					Text string `xml:",chardata"`
				} `xml:"keyword"`
				Period []struct {
					Text string `xml:",chardata"`
				} `xml:"period"`
				Monitoring struct {
					Text string `xml:",chardata"`
				} `xml:"monitoring"`
				Language struct {
					Text string `xml:",chardata"` // en, en, en, en, en, en, e...
				} `xml:"language"`
				Abstract struct {
					Text string `xml:",chardata"` // After a short description...
				} `xml:"abstract"`
				Date struct {
					Text string `xml:",chardata"` // 2009-06-22 12:12:00, 2009...
				} `xml:"date"`
			} `xml:"rfc1807"`
		} `xml:"metadata"`
		About struct {
			Text string `xml:",chardata"`
		} `xml:"about"`
	} `xml:"Record"`
}
```

Only consider a nested element
------------------------------

```go
$ zek -t thesis < fixtures/z.xml
type Thesis struct {
	XMLName        xml.Name `xml:"thesis"`
	Text           string   `xml:",chardata"`
	Xmlns          string   `xml:"xmlns,attr"`
	Doc            string   `xml:"doc,attr"`
	Xsi            string   `xml:"xsi,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Title          []struct {
		Text string `xml:",chardata"`
	} `xml:"title"`
	Creator []struct {
		Text string `xml:",chardata"`
	} `xml:"creator"`
	Date []struct {
		Text string `xml:",chardata"`
	} `xml:"date"`
	Identifier []struct {
		Text string `xml:",chardata"`
	} `xml:"identifier"`
	Language []struct {
		Text string `xml:",chardata"`
	} `xml:"language"`
	Rights []struct {
		Text string `xml:",chardata"`
	} `xml:"rights"`
	Coverage []struct {
		Text string `xml:",chardata"`
	} `xml:"coverage"`
	Publisher []struct {
		Text string `xml:",chardata"`
	} `xml:"publisher"`
	Contributor []struct {
		Text string `xml:",chardata"`
	} `xml:"contributor"`
	Subject []struct {
		Text string `xml:",chardata"`
	} `xml:"subject"`
	Description []struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	Source struct {
		Text string `xml:",chardata"`
	} `xml:"source"`
	Type struct {
		Text string `xml:",chardata"`
	} `xml:"type"`
	Relation []struct {
		Text string `xml:",chardata"`
	} `xml:"relation"`
}
```
