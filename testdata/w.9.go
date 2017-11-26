package main

import "encoding/xml"

type X struct {
	XMLName xml.Name `xml:"x"`
	Text    string   `xml:",chardata"`
	ID      string   `xml:"id,attr"`
	Y       []struct {
		XMLName xml.Name `xml:"y"`
		Text    string   `xml:",chardata"` // L'Rell, Michael, Voq, Wis...
	}
}
