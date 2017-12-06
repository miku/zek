package misc

import "encoding/xml"

// Ruleset was generated 2017-12-06 01:11:21 by tir on apollo.
type Ruleset struct {
	XMLName xml.Name `xml:"ruleset"`
	Text    string   `xml:",chardata"`
	Name    string   `xml:"name,attr"`
	Target  []struct {
		Text string `xml:",chardata"`
		Host string `xml:"host,attr"`
	} `xml:"target"`
	Rule []struct {
		Text string `xml:",chardata"`
		From string `xml:"from,attr"`
		To   string `xml:"to,attr"`
	} `xml:"rule"`
}
