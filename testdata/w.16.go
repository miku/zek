package main

import "encoding/xml"

type A struct {
	XMLName xml.Name `xml:"a"`
	Text    string   `xml:",chardata"`
	B       B        `xml:"b"`
	A       A        `xml:"a"`
}
