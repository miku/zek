import "encoding/xml"

type P struct {
	XMLName xml.Name `xml:"p"`
	Text    string   `xml:",chardata"`
	Br      string   `xml:"br"`
}
