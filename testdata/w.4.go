package main

import "encoding/xml"

type Records struct {
	XMLName xml.Name `xml:"Records"`
	Text    string   `xml:",chardata"`
	Xsi     string   `xml:"xsi,attr"`
	Record  []struct {
		Text   string `xml:",chardata"`
		Header struct {
			Text       string `xml:",chardata"`
			Status     string `xml:"status,attr"`
			Identifier struct {
				Text string `xml:",chardata"`
			} `xml:"identifier"`
			Datestamp struct {
				Text string `xml:",chardata"`
			} `xml:"datestamp"`
			SetSpec struct {
				Text string `xml:",chardata"`
			} `xml:"setSpec"`
		} `xml:"header"`
		Metadata struct {
			Text    string `xml:",chardata"`
			Rfc1807 struct {
				Text           string `xml:",chardata"`
				Xmlns          string `xml:"xmlns,attr"`
				Xsi            string `xml:"xsi,attr"`
				SchemaLocation string `xml:"schemaLocation,attr"`
				BibVersion     struct {
					Text string `xml:",chardata"`
				} `xml:"bib-version"`
				ID struct {
					Text string `xml:",chardata"`
				} `xml:"id"`
				Entry struct {
					Text string `xml:",chardata"`
				} `xml:"entry"`
				Organization []struct {
					Text string `xml:",chardata"`
				} `xml:"organization"`
				Title struct {
					Text string `xml:",chardata"`
				} `xml:"title"`
				Type struct {
					Text string `xml:",chardata"`
				} `xml:"type"`
				Author []struct {
					Text string `xml:",chardata"`
				} `xml:"author"`
				Copyright struct {
					Text string `xml:",chardata"`
				} `xml:"copyright"`
				OtherAccess struct {
					Text string `xml:",chardata"`
				} `xml:"other_access"`
				Keyword struct {
					Text string `xml:",chardata"`
				} `xml:"keyword"`
				Period []struct {
					Text string `xml:",chardata"`
				} `xml:"period"`
				Monitoring struct {
					Text string `xml:",chardata"`
				} `xml:"monitoring"`
				Language struct {
					Text string `xml:",chardata"`
				} `xml:"language"`
				Abstract struct {
					Text string `xml:",chardata"`
				} `xml:"abstract"`
				Date struct {
					Text string `xml:",chardata"`
				} `xml:"date"`
			} `xml:"rfc1807"`
		} `xml:"metadata"`
		About struct {
			Text string `xml:",chardata"`
		} `xml:"about"`
	} `xml:"Record"`
}
