package main

import "encoding/xml"

type List struct {
	XMLName xml.Name `xml:"list"`
	Text    string   `xml:",chardata"`
	Value   struct {
		Text string `xml:",chardata"` // a
	} `xml:"value"`
	List struct {
		Text  string `xml:",chardata"`
		Value struct {
			Text string `xml:",chardata"` // b
		} `xml:"value"`
		List struct {
			Text  string `xml:",chardata"`
			Value struct {
				Text string `xml:",chardata"` // c
			} `xml:"value"`
		} `xml:"list"`
	} `xml:"list"`
}
