type A struct {
	XMLName xml.Name `xml:"a"`
	Text    string   `xml:",chardata"`
	ID      string   `xml:"id,attr"`
}