# zek

Zek is a **prototype** for creating a Go
[struct](https://golang.org/ref/spec#Struct_types) from an XML document. The
resulting struct works best for *reading* XML (see also
[#14](https://github.com/miku/zek/issues/14)), to create XML, you might want to
use [something else](https://github.com/avelino/awesome-go#xml).

It was developed at [ubleipzig](https://github.com/ubleipzig) to shorten the
time to go from raw XML to a struct that allows to access XML data in Go
programs.

> Skip the fluff, just the code.

Given some [XML](https://raw.githubusercontent.com/miku/zek/master/fixtures/e.xml), run:

```go
$ curl -s https://raw.githubusercontent.com/miku/zek/master/fixtures/e.xml | zek -e
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

The zek tools has been developed at [Leipzig University Library](https://ub.uni-leipzig.de).

## Online

Try it online at [https://www.onlinetool.io/xmltogo/](https://www.onlinetool.io/xmltogo/).

## About

[![Build Status](https://travis-ci.org/miku/zek.svg?branch=master)](https://travis-ci.org/miku/zek) [![Project Status: Active – The project has reached a stable, usable state and is being actively developed.](https://www.repostatus.org/badges/latest/active.svg)](https://www.repostatus.org/#active)

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

Presentations:

* [Lightning talk about struct generators](https://gist.github.com/miku/39e4273d15abfd7e4297071338da3349) at [GoLab 2018](https://www.golab.io/)

## Install

```
$ go get github.com/miku/zek/cmd/...
```

Debian and RPM packages:

* https://github.com/miku/zek/releases

It's in [AUR](https://aur.archlinux.org/packages/zek-bin/), too.

![](https://github.com/miku/zek/blob/master/docs/114391.png)

## Usage

```shell
$ zek -h
Usage of zek:
  -C    emit less compact struct
  -F    skip formatting
  -c    emit more compact struct (noop, as this is the default since 0.1.7)
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

$ zek -C < fixtures/a.xml
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

$ zek -C -p < fixtures/a.xml > sample.go && go run sample.go < fixtures/a.xml | jq . && rm sample.go
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
// Root was generated 2019-06-11 16:27:04 by tir on hayiti.
type Root struct {
        XMLName xml.Name `xml:"root"`
        Text    string   `xml:",chardata"`
        A       []struct {
                Text string `xml:",chardata"`
                B    []struct {
                        Text string `xml:",chardata"`
                        C    string `xml:"c"`
                        D    string `xml:"d"`
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
          "C": "Hi",
          "D": ""
        },
        {
          "Text": "\n    \n    \n  ",
          "C": "World",
          "D": ""
        }
      ]
    },
    {
      "Text": "\n  \n",
      "B": [
        {
          "Text": "\n    \n  ",
          "C": "Hello",
          "D": ""
        }
      ]
    },
    {
      "Text": "\n  \n",
      "B": [
        {
          "Text": "\n    \n  ",
          "C": "",
          "D": "World"
        }
      ]
    }
  ]
}
```

Annotate with comments:

```go
$ zek -e < fixtures/l.xml
// Records was generated 2019-06-11 16:29:35 by tir on hayiti.
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
                                BibVersion     string   `xml:"bib-version"`  // v2, v2, v2...
                                ID             string   `xml:"id"`           // http://jou...
                                Entry          string   `xml:"entry"`        // 2009-06-24...
                                Organization   []string `xml:"organization"` // Proceeding...
                                Title          string   `xml:"title"`        // Introducti...
                                Type           string   `xml:"type"`
                                Author         []string `xml:"author"`       // KRAMPEN, G..
                                Copyright      string   `xml:"copyright"`    // Das Urhebe...
                                OtherAccess    string   `xml:"other_access"` // url:http:/...
                                Keyword        string   `xml:"keyword"`
                                Period         []string `xml:"period"`
                                Monitoring     string   `xml:"monitoring"`
                                Language       string   `xml:"language"` // en, en, en, e...
                                Abstract       string   `xml:"abstract"` // After a short...
                                Date           string   `xml:"date"`     // 2009-06-22 12...
                        } `xml:"rfc1807"`
                } `xml:"metadata"`
                About string `xml:"about"`
        } `xml:"Record"`
}
```

## Only consider a nested element

```go
$ zek -t metadata fixtures/z.xml
// Metadata was generated 2019-06-11 16:33:26 by tir on hayiti.
type Metadata struct {
        XMLName xml.Name `xml:"metadata"`
        Text    string   `xml:",chardata"`
        Dc      struct {
                Text  string `xml:",chardata"`
                Xmlns string `xml:"xmlns,attr"`
                Title struct {
                        Text  string `xml:",chardata"`
                        Xmlns string `xml:"xmlns,attr"`
                } `xml:"title"`
                Identifier struct {
                        Text  string `xml:",chardata"`
                        Xmlns string `xml:"xmlns,attr"`
                } `xml:"identifier"`
                Rights struct {
                        Text  string `xml:",chardata"`
                        Xmlns string `xml:"xmlns,attr"`
                        Lang  string `xml:"lang,attr"`
                } `xml:"rights"`
                AccessRights struct {
                        Text  string `xml:",chardata"`
                        Xmlns string `xml:"xmlns,attr"`
                } `xml:"accessRights"`
        } `xml:"dc"`
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

