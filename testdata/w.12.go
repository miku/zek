import "encoding/xml"

// List was generated 2018-11-23 14:50:36 by tir on nexus.
type List struct {
	XMLName xml.Name `xml:"list"`
	Text    string   `xml:",chardata"`
	Value   []struct {
		Text string `xml:",chardata"` // a, a
	} `xml:"value"`
}

