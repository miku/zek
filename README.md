# zek

Zek is a **prototype** for creating a Go [struct](https://golang.org/ref/spec#Struct_types) from an XML document.

> Skip the fluff, just the code.

Given some [XML](https://raw.githubusercontent.com/miku/zek/master/fixtures/e.xml), run:

```go
$ curl -s https://raw.githubusercontent.com/miku/zek/master/fixtures/e.xml | zek -e -c
// Rss was generated 2018-08-30 20:24:14 by tir on sol.
type Rss struct {
    XMLName xml.Name `xml:"rss"`
    Text    string   `xml:",chardata"`
    Rdf     string   `xml:"rdf,attr"`
    Dc      string   `xml:"dc,attr"`
    Geoscan string   `xml:"geoscan,attr"`
    Media   string   `xml:"media,attr"`
    Gml     string   `xml:"gml,attr"`
    Taxo    string   `xml:"taxo,attr"`
    Georss  string   `xml:"georss,attr"`
    Content string   `xml:"content,attr"`
    Geo     string   `xml:"geo,attr"`
    Version string   `xml:"version,attr"`
    Channel struct {
        Text          string `xml:",chardata"`
        Title         string `xml:"title"`         // ESS New Releases (Display...
        Link          string `xml:"link"`          // http://tinyurl.com/ESSNew...
        Description   string `xml:"description"`   // New releases from the Ear...
        LastBuildDate string `xml:"lastBuildDate"` // Mon, 27 Nov 2017 00:06:35...
        Item          []struct {
            Text        string `xml:",chardata"`
            Title       string `xml:"title"`       // Surficial geology, Aberde...
            Link        string `xml:"link"`        // https://geoscan.nrcan.gc....
            Description string `xml:"description"` // Geological Survey of Cana...
            Guid        struct {
                Text        string `xml:",chardata"` // 304279, 306212, 306175, 3...
                IsPermaLink string `xml:"isPermaLink,attr"`
            } `xml:"guid"`
            PubDate       string   `xml:"pubDate"`      // Fri, 24 Nov 2017 00:00:00...
            Polygon       []string `xml:"polygon"`      // 64.0000 -98.0000 64.0000 ...
            Download      string   `xml:"download"`     // https://geoscan.nrcan.gc....
            License       string   `xml:"license"`      // http://data.gc.ca/eng/ope...
            Author        string   `xml:"author"`       // Geological Survey of Cana...
            Source        string   `xml:"source"`       // Geological Survey of Cana...
            SndSeries     string   `xml:"SndSeries"`    // Bedford Institute of Ocea...
            Publisher     string   `xml:"publisher"`    // Natural Resources Canada,...
            Edition       string   `xml:"edition"`      // prelim., surficial data m...
            Meeting       string   `xml:"meeting"`      // Geological Association of...
            Documenttype  string   `xml:"documenttype"` // serial, open file, serial...
            Language      string   `xml:"language"`     // English, English, English...
            Maps          string   `xml:"maps"`         // 1 map, 5 maps, Publicatio...
            Mapinfo       string   `xml:"mapinfo"`      // surficial geology, surfic...
            Medium        string   `xml:"medium"`       // on-line; digital, digital...
            Province      string   `xml:"province"`     // Nunavut, Northwest Territ...
            Nts           string   `xml:"nts"`          // 066B, 095J; 095N; 095O; 0...
            Area          string   `xml:"area"`         // Aberdeen Lake, Mackenzie ...
            Subjects      string   `xml:"subjects"`
            Program       string   `xml:"program"`       // GEM2: Geo-mapping for Ene...
            Project       string   `xml:"project"`       // Rae Province Project Mana...
            Projectnumber string   `xml:"projectnumber"` // 340521, 343202, 340557, 3...
            Abstract      string   `xml:"abstract"`      // This new surficial geolog...
            Links         string   `xml:"links"`         // Online - En ligne (PDF, 9...
            Readme        string   `xml:"readme"`        // readme | https://geoscan....
            PPIid         string   `xml:"PPIid"`         // 34532, 35096, 35438, 2563...
        } `xml:"item"`
    } `xml:"channel"`
}
```

## Online

Try it online at [https://www.onlinetool.io/xmltogo/](https://www.onlinetool.io/xmltogo/).

## About

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

## Install

```
$ go get github.com/miku/zek/cmd/zek/...
```

Debian and RPM packages:

* https://github.com/miku/zek/releases


![](https://github.com/miku/zek/blob/master/docs/114391.png)

## Usage

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
  -u    filter out duplicated examples
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

The above struct can be made a bit more compact - use the `-c` flag (since 0.1.4) to see how:

```go
$ zek -c -e < fixtures/l.xml
// Records was generated 2018-08-09 14:10:25 by tir on sol.
type Records struct {
    XMLName xml.Name `xml:"Records"`
    Text    string   `xml:",chardata"` // \n
    Xsi     string   `xml:"xsi,attr"`
    Record  []struct {
        Text   string `xml:",chardata"`
        Header struct {
            Text       string `xml:",chardata"`
            Status     string `xml:"status,attr"`
            Identifier string `xml:"identifier"` // oai:ojs.localhost:article...
            Datestamp  string `xml:"datestamp"`  // 2009-06-24T14:48:23Z, 200...
            SetSpec    string `xml:"setSpec"`    // eppp:ART, eppp:ART, eppp:...
        } `xml:"header"`
        Metadata struct {
            Text    string `xml:",chardata"`
            Rfc1807 struct {
                Text           string   `xml:",chardata"`
                Xmlns          string   `xml:"xmlns,attr"`
                Xsi            string   `xml:"xsi,attr"`
                SchemaLocation string   `xml:"schemaLocation,attr"`
                BibVersion     string   `xml:"bib-version"`  // v2, v2, v2, v2, v2, v2, v...
                ID             string   `xml:"id"`           // http://journals.zpid.de/i...
                Entry          string   `xml:"entry"`        // 2009-06-24T14:48:23Z, 200...
                Organization   []string `xml:"organization"` // Proceedings of the Worksh...
                Title          string   `xml:"title"`        // Introduction and some Ide...
                Type           string   `xml:"type"`
                Author         []string `xml:"author"`       // KRAMPEN, Günter, CARBON,...
                Copyright      string   `xml:"copyright"`    // Das Urheberrecht liegt be...
                OtherAccess    string   `xml:"other_access"` // url:http://journals.zpid....
                Keyword        string   `xml:"keyword"`
                Period         []string `xml:"period"`
                Monitoring     string   `xml:"monitoring"`
                Language       string   `xml:"language"` // en, en, en, en, en, en, e...
                Abstract       string   `xml:"abstract"` // After a short description...
                Date           string   `xml:"date"`     // 2009-06-22 12:12:00, 2009...
            } `xml:"rfc1807"`
        } `xml:"metadata"`
        About string `xml:"about"`
    } `xml:"Record"`
}
```

## Only consider a nested element

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

## Inference across files

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

Another example (tarball with thousands of XML files, seemingly MARC):

```shell
$ tar -xOzf /tmp/20180725.125255.tar.gz | zek -e
// OAIPMH was generated 2018-09-26 15:03:29 by tir on sol.
type OAIPMH struct {
        XMLName        xml.Name `xml:"OAI-PMH"`
        Text           string   `xml:",chardata"`
        Xmlns          string   `xml:"xmlns,attr"`
        Xsi            string   `xml:"xsi,attr"`
        SchemaLocation string   `xml:"schemaLocation,attr"`
        ListRecords    struct {
                Text   string `xml:",chardata"`
                Record struct {
                        Text   string `xml:",chardata"`
                        Header struct {
                                Text       string `xml:",chardata"`
                                Identifier struct {
                                        Text string `xml:",chardata"` // aleph-publish:000000001, ...
                                } `xml:"identifier"`
                        } `xml:"header"`
                        Metadata struct {
                                Text   string `xml:",chardata"`
                                Record struct {
                                        Text           string `xml:",chardata"`
                                        Xmlns          string `xml:"xmlns,attr"`
                                        Xsi            string `xml:"xsi,attr"`
                                        SchemaLocation string `xml:"schemaLocation,attr"`
                                        Leader         struct
                                                Text string `xml:",chardata"` // 00001nM2.01200024
                                        } `xml:"leader"`
                                        Controlfield []struct {
                                                Text string `xml:",chardata"` // 00001nM2.01200024
                                                Tag  string `xml:"tag,attr"`
                                        } `xml:"controlfield"`
                                        Datafield []struct {
                                                Text     string `xml:",chardata"`
                                                Tag      string `xml:"tag,attr"`
                                                Ind1     string `xml:"ind1,attr"`
                                                Ind2     string `xml:"ind2,attr"`
                                                Subfield []struct {
                                                        Text string `xml:",chardata"` // KM0000002
                                                        Code string `xml:"code,attr"`
                                                } `xml:"subfield"`
                                        } `xml:"datafield"`
                                } `xml:"record"`
                        } `xml:"metadata"`
                } `xml:"record"`
        } `xml:"ListRecords"`
}
```

## Misc

As a side effect, zek seems to be a useful for debugging. Example:

* https://git.io/vbUNa

This record is emitted from a typical [OAI](https://www.openarchives.org/)
server ([OJS](https://pkp.sfu.ca/ojs/), not even uncommon), yet one can quickly
spot the flaw in the structure.

Over 30 different struct generated manually in the course of a few hours
(around five minutes per source): https://git.io/vbTDo.

-- Current extent leader: [1532](https://github.com/miku/zek/blob/master/fixtures/r.txt) lines struct

