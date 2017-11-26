package main

import "encoding/xml"

type A struct {
	XMLName xml.Name `xml:"a"`
	Text    string   `xml:",chardata"`
	B       struct {
		XMLName xml.Name `xml:"b"`
		Text    string   `xml:",chardata"` // Berlin
	}
}
