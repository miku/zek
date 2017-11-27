package main

import "encoding/xml"

type X struct {
	XMLName xml.Name `xml:"x"`
	Text    string   `xml:",chardata"`
	ID      string   `xml:"id,attr"`
	Y       []struct {
		Text string `xml:",chardata"` // L'Rell, Michael, Voq, Wis...
	} `xml:"y"`
}
