package main

import "encoding/xml"

type A struct {
	XMLName xml.Name `xml:"a"`
	Text    string   `xml:",chardata"`
	AttrB   string   `xml:"b,attr"`
	B       struct {
		Text string `xml:",chardata"`
	} `xml:"b"`
}
