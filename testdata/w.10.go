package main

import "encoding/xml"

type A struct {
	XMLName xml.Name `xml:"a"`
	Text    string   `xml:",chardata"`
	B       struct {
		Text string `xml:",chardata"` // Europe
	} `xml:"b"`
	A struct {
		Text string `xml:",chardata"`
		B    struct {
			Text string `xml:",chardata"` // Germany
		} `xml:"b"`
		A struct {
			Text string `xml:",chardata"`
			B    struct {
				Text string `xml:",chardata"` // Leipzig
			} `xml:"b"`
		} `xml:"a"`
	} `xml:"a"`
}
