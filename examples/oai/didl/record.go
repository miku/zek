package didl

import "encoding/xml"

type Record struct {
	XMLName xml.Name `xml:"Record"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text       string `xml:",chardata"`
		Status     string `xml:"status,attr"`
		Identifier struct {
			Text string `xml:",chardata"`
		} `xml:"identifier"`
		Datestamp struct {
			Text string `xml:",chardata"`
		} `xml:"datestamp"`
		SetSpec []struct {
			Text string `xml:",chardata"`
		} `xml:"setSpec"`
	} `xml:"header"`
	Metadata struct {
		Text string `xml:",chardata"`
		DIDL struct {
			Text           string `xml:",chardata"`
			D              string `xml:"d,attr"`
			Doc            string `xml:"doc,attr"`
			Xsi            string `xml:"xsi,attr"`
			SchemaLocation string `xml:"schemaLocation,attr"`
			Item           struct {
				Text       string `xml:",chardata"`
				ID         string `xml:"id,attr"`
				Descriptor []struct {
					Text      string `xml:",chardata"`
					Statement struct {
						Text       string `xml:",chardata"`
						MimeType   string `xml:"mimeType,attr"`
						Identifier struct {
							Text           string `xml:",chardata"`
							Dii            string `xml:"dii,attr"`
							SchemaLocation string `xml:"schemaLocation,attr"`
						} `xml:"Identifier"`
						Dc struct {
							Text           string `xml:",chardata"`
							OaiDc          string `xml:"oai_dc,attr"`
							Dc             string `xml:"dc,attr"`
							SchemaLocation string `xml:"schemaLocation,attr"`
							Title          []struct {
								Text string `xml:",chardata"`
							} `xml:"title"`
							Creator []struct {
								Text string `xml:",chardata"`
							} `xml:"creator"`
							Description []struct {
								Text string `xml:",chardata"`
							} `xml:"description"`
							Date struct {
								Text string `xml:",chardata"`
							} `xml:"date"`
							Type struct {
								Text string `xml:",chardata"`
							} `xml:"type"`
							Identifier []struct {
								Text string `xml:",chardata"`
							} `xml:"identifier"`
							Language []struct {
								Text string `xml:",chardata"`
							} `xml:"language"`
							Relation []struct {
								Text string `xml:",chardata"`
							} `xml:"relation"`
							Rights []struct {
								Text string `xml:",chardata"`
							} `xml:"rights"`
							Publisher struct {
								Text string `xml:",chardata"`
							} `xml:"publisher"`
							Subject []struct {
								Text string `xml:",chardata"`
							} `xml:"subject"`
							Contributor []struct {
								Text string `xml:",chardata"`
							} `xml:"contributor"`
						} `xml:"dc"`
					} `xml:"Statement"`
				} `xml:"Descriptor"`
				Component []struct {
					Text     string `xml:",chardata"`
					ID       string `xml:"id,attr"`
					Resource struct {
						Text     string `xml:",chardata"`
						Ref      string `xml:"ref,attr"`
						MimeType string `xml:"mimeType,attr"`
					} `xml:"Resource"`
				} `xml:"Component"`
			} `xml:"Item"`
		} `xml:"DIDL"`
	} `xml:"metadata"`
	About struct {
		Text string `xml:",chardata"`
	} `xml:"about"`
}
