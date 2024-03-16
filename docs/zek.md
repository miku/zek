% ZEK(1)
% Martin Czygan (https://github.com/miku)
% November, 23 2018

# NAME

**zek** - generate Go struct definitions from XML

# SYNOPSIS

zek [*options*] < [*file*]

# DESCRIPTION

The Go programming language supports XML serialization through struct tags. The
`zek` command line tool will convert a stream of XML documents into a suitable
struct definition. The XML data is read from a file or, if none given, from
stdin.

    zek < file.xml

The idea of code generation is not new and has been layed out in the blog post
[*Generating code*](https://blog.golang.org/generate) from 2014-12-22 by Rob
Pike. Go has many source code generators, notably
[*JSONGen*](https://github.com/bemasher/JSONGen) and also a few for XML. The
`zek` tool is a bit different, as it will *always* emit a valid struct - as
opposed to the otherwise great
[*XMLGen*](https://github.com/dutchcoders/XMLGen) - and it will work on raw
data, that has no DTD or XSD.

Online versions have been created by Krzysztof Kowalczyk, available at
https://www.onlinetool.io/xmltogo/ and Yaroslav Podorvanov, available at
https://xml-to-go.github.io/.

# OPTIONS

-B
:  Use a fixed banner string (e.g. for CI).

-C
:  Emit less compact representation. This mimicks the former default setting.

-F
:  Skip formatting. The formatter might choke on funny characters.

-I
:  Use verbatim innerxml instead of chardata.

-P *NAME*
:  If set, write out struct within a package with the given name.

-S int
: Read at most this many tags, approximately (0=unlimited).

-c
:  Emit more compact representation. This is the default as of 0.1.8.

-d
:  Emit debug output.

-e
:  Add comments with examples. Can be used in conjuction with `-u` to make examples unique.

-j
:  Add tags for JSON as well.

-m
:  Omit empty Text fields.

-max-examples *INT*
:  Limit the number of examples (default 10).

-n *NAME*
:  Use a different name for the struct.

-o *FILE*
:  If set, write to output file, not stdout.

-p
:  Write an example program to stdout.

-s
:  Strict parsing.

-t
:  Emit struct for tag matching this name.

-u
:  Unique examples only.

-version
:  Show program version.

-x
:  Max length for a single example (default 25).

# BUGS

Mapping between XML elements and data structures is inherently flawed: an XML
element is an order-dependent collection of anonymous values, while a data
structure is an order-independent collection of named values.

From: https://golang.org/pkg/encoding/xml/#pkg-note-BUG

# EXAMPLE

There are numerous examples of XML files on the web (e.g. try:
https://www.google.com/search?q=filetype:xml+inurl:rss). Below is an example
from a publishers publication feed.

	$ curl -sL http://epubs.siam.org/rss/SJOPE8.xml | zek -e
	// RDF was generated 2018-11-23 14:21:21 by tir on nexus.
	type RDF struct {
		XMLName xml.Name `xml:"RDF"`
		Text    string   `xml:",chardata"`
		Rdf     string   `xml:"rdf,attr"`
		Cc      string   `xml:"cc,attr"`
		Rss     string   `xml:"rss,attr"`
		Content string   `xml:"content,attr"`
		Prism   string   `xml:"prism,attr"`
		Dc      string   `xml:"dc,attr"`
		Channel struct {
			Text            string `xml:",chardata"`
			About           string `xml:"about,attr"`
			Title           string `xml:"title"`           // Society for Industrial an...
			Description     string `xml:"description"`     // Table of Contents for SIA...
			Link            string `xml:"link"`            // https://epubs.siam.org/lo...
			Publisher       string `xml:"publisher"`       // Society for Industrial an...
			Language        string `xml:"language"`        // en-US
			PublicationName string `xml:"publicationName"` // SIAM Journal on Optimizat...
			Items           struct {
				Text string `xml:",chardata"`
				Seq  struct {
					Text string `xml:",chardata"`
					Li   []struct {
						Text     string `xml:",chardata"`
						Resource string `xml:"resource,attr"`
					} `xml:"li"`
				} `xml:"Seq"`
			} `xml:"items"`
		} `xml:"channel"`
		Image struct {
			Text  string `xml:",chardata"`
			About string `xml:"about,attr"`
			Title string `xml:"title"` // SIAM Journal on Optimizat...
			URL   string `xml:"url"`   // https://epubs.siam.org/na...
			Link  string `xml:"link"`  // https://epubs.siam.org/lo...
		} `xml:"image"`
		Item []struct {
			Text             string   `xml:",chardata"`
			About            string   `xml:"about,attr"`
			Title            string   `xml:"title"`            // Random Gradient Extrapola...
			Link             string   `xml:"link"`             // https://epubs.siam.org/do...
			Encoded          string   `xml:"encoded"`          // SIAM Journal on Optimizat...
			Description      string   `xml:"description"`      // SIAM Journal on Optimizat...
			Identifier       string   `xml:"identifier"`       // doi:10.1137/17M1157891, d...
			Source           string   `xml:"source"`           // SIAM Journal on Optimizat...
			Date             string   `xml:"date"`             // 2018-10-02T07:00:00Z, 201...
			Rights           string   `xml:"rights"`           // © 2018, Society for Indu...
			Creator          []string `xml:"creator"`          // Guanghui Lan, Yi Zhou, An...
			PublicationName  string   `xml:"publicationName"`  // Random Gradient Extrapola...
			Volume           string   `xml:"volume"`           // 28, 28, 28, 28, 28, 28, 2...
			Number           string   `xml:"number"`           // 4, 4, 4, 4, 4, 4, 4, 4, 4...
			StartingPage     string   `xml:"startingPage"`     // 2753, 2783, 2809, 2839, 2...
			EndingPage       string   `xml:"endingPage"`       // 2782, 2808, 2838, 2871, 2...
			CoverDate        string   `xml:"coverDate"`        // 2018-01-01T08:00:00Z, 201...
			CoverDisplayDate string   `xml:"coverDisplayDate"` // 2018-01-01T08:00:00Z, 201...
			Doi              string   `xml:"doi"`              // 10.1137/17M1157891, 10.11...
			URL              string   `xml:"url"`              // https://epubs.siam.org/do...
			Copyright        string   `xml:"copyright"`        // © 2018, Society for Indu...
		} `xml:"item"`
	}

# CONTRIBUTORS

* [Mateusz Burniak](https://github.com/matbur)
* [Christian G. Warden](https://github.com/cwarden)
