package main

import "encoding/xml"

type List struct {
	XMLName xml.Name `xml:"list"`
	Text    string   `xml:",chardata"`
	Value   struct {
		XMLName xml.Name `xml:"value"`
		Text    string   `xml:",chardata"` // a
	}
	List struct {
		XMLName xml.Name `xml:"list"`
		Text    string   `xml:",chardata"`
		Value   struct {
			XMLName xml.Name `xml:"value"`
			Text    string   `xml:",chardata"` // b
		}
		List struct {
			XMLName xml.Name `xml:"list"`
			Text    string   `xml:",chardata"`
			Value   struct {
				XMLName xml.Name `xml:"value"`
				Text    string   `xml:",chardata"` // c
			}
		}
	}
}
