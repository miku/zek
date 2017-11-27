package main

import "encoding/xml"

type Root struct {
	XMLName xml.Name `xml:"root"`
	Text    string   `xml:",chardata"`
	A       []struct {
		Text string `xml:",chardata"`
		B    []struct {
			Text string `xml:",chardata"`
			C    struct {
				Text string `xml:",chardata"`
			} `xml:"c"`
			D struct {
				Text string `xml:",chardata"`
			} `xml:"d"`
		} `xml:"b"`
	} `xml:"a"`
}
