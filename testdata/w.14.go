import "encoding/xml"

// List was generated 2018-11-23 14:50:36 by tir on nexus.
type List struct {
	XMLName xml.Name `xml:"list"`
	Value   []struct {
		Text string `xml:",chardata"` // a
	} `xml:"value"`
}

