import "encoding/xml"

type Root struct {
	XMLName xml.Name `xml:"root"`
	Text    string   `xml:",chardata"`
	A       []struct {
		XMLName xml.Name `xml:"a"`
		Text    string   `xml:",chardata"`
		B       []struct {
			XMLName xml.Name `xml:"b"`
			Text    string   `xml:",chardata"`
			C       struct {
				XMLName xml.Name `xml:"c"`
				Text    string   `xml:",chardata"`
			}
			D struct {
				XMLName xml.Name `xml:"d"`
				Text    string   `xml:",chardata"`
			}
		}
	}
}
