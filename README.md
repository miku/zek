zek
===

Zek is a **prototype** for creating a Go struct from an XML document.

[![Build Status](https://travis-ci.org/miku/zek.svg?branch=master)](https://travis-ci.org/miku/zek)

Upsides:

* it works fine for non-recursive structures,
* does not need XSD or DTD,
* it is relatively convenient to access attributes, children and text,
* will generate a single struct, which make for a quite compact representation,
* simple user interface,
* comments with examples,
* schema inference across multiple files.

Downsides:

* experimental, early, buggy, unstable prototype,
* no support for recursive types (similar to *Russian Doll* strategy, [[1](https://medbiq.org/std_specs/techguidelines/xmldesignguidelines.pdf#page=7)])
* no type inference, everything is accessible as string.

Bugs:

> Mapping between XML elements and data structures is inherently flawed: an XML
> element is an order-dependent collection of anonymous values, while a data
> structure is an order-independent collection of named values.

https://golang.org/pkg/encoding/xml/#pkg-note-BUG

Related projects:

* https://github.com/bemasher/JSONGen
* https://github.com/dutchcoders/XMLGen
* https://github.com/gnewton/chidley

Install
-------

```
$ go get github.com/miku/zek/cmd/...
```

Debian and RPM packages:

* https://github.com/miku/zek/releases


![](https://github.com/miku/zek/blob/master/docs/114391.png)

Usage
-----

```shell
$ zek -h
Usage of zek:
  -F    skip formatting
  -c    emit more compact struct
  -d    debug output
  -e    add comments with example
  -j    add JSON tags
  -max-examples int
        limit number of examples (default 10)
  -n string
        use a different name for the top-level struct
  -p    write out an example program
  -s    strict parsing and writing
  -t string
        emit struct for tag matching this name
  -version
        show version
  -x int
        max chars for example (default 25)
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
package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

// A was generated 2017-12-05 17:35:21 by tir on apollo.
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

The above struct can be made a bit more compact - use the `-c` flag to see how:

```
$ zek -c -e < fixtures/l.xml
// Records was generated 2018-08-09 14:00:07 by tir on sol.
type Records struct {
        XMLName xml.Name `xml:"Records"`
        Text    string   `xml:",chardata"` // \n
        Xsi     string   `xml:"xsi,attr"`
        Record  []struct {
                Text   string `xml:",chardata"`
                Header struct {
                        Text       string `xml:",chardata"`
                        Status     string `xml:"status,attr"`
                        Identifier string `xml:"identifier"`
                        Datestamp  string `xml:"datestamp"`
                        SetSpec    string `xml:"setSpec"`
                } `xml:"header"`
                Metadata struct {
                        Text    string `xml:",chardata"`
                        Rfc1807 struct {
                                Text           string   `xml:",chardata"`
                                Xmlns          string   `xml:"xmlns,attr"`
                                Xsi            string   `xml:"xsi,attr"`
                                SchemaLocation string   `xml:"schemaLocation,attr"`
                                BibVersion     string   `xml:"bib-version"`
                                ID             string   `xml:"id"`
                                Entry          string   `xml:"entry"`
                                Organization   []string `xml:"organization"`
                                Title          string   `xml:"title"`
                                Type           string   `xml:"type"`
                                Author         []string `xml:"author"`
                                Copyright      string   `xml:"copyright"`
                                OtherAccess    string   `xml:"other_access"`
                                Keyword        string   `xml:"keyword"`
                                Period         []string `xml:"period"`
                                Monitoring     string   `xml:"monitoring"`
                                Language       string   `xml:"language"`
                                Abstract       string   `xml:"abstract"`
                                Date           string   `xml:"date"`
                        } `xml:"rfc1807"`
                } `xml:"metadata"`
                About string `xml:"about"`
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

Inference across files
----------------------

```go
$ zek fixtures/a.xml fixtures/b.xml fixtures/c.xml
// A was generated 2017-12-05 17:40:14 by tir on apollo.
type A struct {
	XMLName xml.Name `xml:"a"`
	Text    string   `xml:",chardata"`
	B       []struct {
		Text string `xml:",chardata"`
	} `xml:"b"`
}
```

This is also useful, if you deal with archives containing XML files:

```shell
$ unzip -p 4082359.zip '*.xml' | zek -e
```

Given a directory full of zip files, you can combined find, unzip and zek:

```shell
$ for i in $(find ftp/b571 -type f -name "*zip"); do unzip -p $i '*xml'; done | zek -e
```

Misc
----

As a side effect, zek seems to be a useful for debugging. Example:

* https://git.io/vbUNa

This record is emitted from a typical [OAI](https://www.openarchives.org/)
server ([OJS](https://pkp.sfu.ca/ojs/), not even uncommon), yet one can quickly
spot the flaw in the structure.

Over 30 different struct generated manually in the course of a few hours
(around five minutes per source): https://git.io/vbTDo.

-- Current extent leader: [1532](https://github.com/miku/zek/blob/master/fixtures/r.txt) lines struct

