package misc

import "encoding/xml"

// Data was generated 2017-12-05 23:29:18 by tir on apollo.
type Data struct {
	XMLName xml.Name `xml:"data"`
	Text    string   `xml:",chardata"`
	Type    string   `xml:"type,attr"`
	Date    string   `xml:"date,attr"`
	Head    struct {
		Text string `xml:",chardata"`
		Db   []struct {
			Text    string `xml:",chardata"`
			Name    string `xml:"name,attr"`
			Code    string `xml:"code,attr"`
			Fields  string `xml:"fields,attr"`
			Server  string `xml:"server,attr"`
			Query   string `xml:"query,attr"`
			Link    string `xml:"link,attr"`
			Accept  string `xml:"accept,attr"`
			Message string `xml:"message,attr"`
		} `xml:"db"`
	} `xml:"head"`
	Body struct {
		Text    string `xml:",chardata"`
		Journal []struct {
			Text       string `xml:",chardata"`
			Title      string `xml:"title,attr"`
			Abbr       string `xml:"abbr,attr"`
			ISSN       string `xml:"issn,attr"`
			Years      string `xml:"years,attr"`
			Volumes    string `xml:"volumes,attr"`
			Prefix     string `xml:"prefix,attr"`
			RealTitle  string `xml:"real_title,attr"`
			Misspelled string `xml:"misspelled,attr"`
		} `xml:"journal"`
	} `xml:"body"`
	Tail struct {
		Text    string `xml:",chardata"`
		Summary struct {
			Text     string `xml:",chardata"`
			Journals string `xml:"journals,attr"`
			Prefixes string `xml:"prefixes,attr"`
			Records  string `xml:"records,attr"`
			Years    string `xml:"years,attr"`
			Volumes  string `xml:"volumes,attr"`
		} `xml:"summary"`
	} `xml:"tail"`
}
