package main

import "encoding/xml"

type A struct {
	XMLName xml.Name `xml:"a"`
	Text    string   `xml:",chardata"`
	AttrB   string   `xml:"b,attr"`
	B       struct {
		XMLName xml.Name `xml:"b"`
		Text    string   `xml:",chardata"`
	}
}
