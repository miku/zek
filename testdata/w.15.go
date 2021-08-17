import "encoding/xml"

type Yo struct {
	XMLName xml.Name `xml:"Yo"`
	Text    string   `xml:",chardata"`
	V30day  string   `xml:"_30day,attr"`
}

