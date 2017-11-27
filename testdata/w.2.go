package main

import "encoding/xml"

type A struct {
	XMLName xml.Name `xml:"a"`
	Text    string   `xml:",chardata"`
	B       struct {
		Text string `xml:",chardata"`
	} `xml:"b"`
}
