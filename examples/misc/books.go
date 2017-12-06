package misc

import "encoding/xml"

// Books was generated 2017-12-06 01:09:22 by tir on apollo.
type Books struct {
	XMLName xml.Name `xml:"books"`
	Text    string   `xml:",chardata"`
	Book    []struct {
		Text  string `xml:",chardata"`
		Title struct {
			Text string `xml:",chardata"` // Alice in Wonderland, Arou...
		} `xml:"title"`
		Translation []struct {
			Text    string `xml:",chardata"` // Spanish, Chinese, Japanes...
			Edition string `xml:"edition,attr"`
		} `xml:"translation"`
	} `xml:"book"`
}
