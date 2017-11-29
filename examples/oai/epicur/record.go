package epicur

import "encoding/xml"

type Record struct {
	XMLName xml.Name `xml:"Record"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text       string `xml:",chardata"`
		Status     string `xml:"status,attr"`
		Identifier struct {
			Text string `xml:",chardata"` // oai:journals.ub.uni-heide...
		} `xml:"identifier"`
		Datestamp struct {
			Text string `xml:",chardata"` // 2015-02-08T19:48:44Z, 201...
		} `xml:"datestamp"`
		SetSpec struct {
			Text string `xml:",chardata"` // hzfi:ART, test:ART, teste...
		} `xml:"setSpec"`
	} `xml:"header"`
	Metadata struct {
		Text   string `xml:",chardata"`
		Epicur struct {
			Text               string `xml:",chardata"`
			Xmlns              string `xml:"xmlns,attr"`
			Xsi                string `xml:"xsi,attr"`
			SchemaLocation     string `xml:"schemaLocation,attr"`
			AdministrativeData struct {
				Text     string `xml:",chardata"`
				Delivery struct {
					Text         string `xml:",chardata"`
					UpdateStatus struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
					} `xml:"update_status"`
				} `xml:"delivery"`
			} `xml:"administrative_data"`
			Record []struct {
				Text       string `xml:",chardata"`
				Identifier struct {
					Text   string `xml:",chardata"` // urn:nbn:de:bsz:16-ip-1649...
					Scheme string `xml:"scheme,attr"`
				} `xml:"identifier"`
				Resource []struct {
					Text       string `xml:",chardata"`
					Identifier struct {
						Text   string `xml:",chardata"` // http://journals.ub.uni-he...
						Scheme string `xml:"scheme,attr"`
						Type   string `xml:"type,attr"`
						Role   string `xml:"role,attr"`
					} `xml:"identifier"`
					Format struct {
						Text   string `xml:",chardata"` // text/html, application/pd...
						Scheme string `xml:"scheme,attr"`
					} `xml:"format"`
				} `xml:"resource"`
			} `xml:"record"`
		} `xml:"epicur"`
	} `xml:"metadata"`
	About struct {
		Text string `xml:",chardata"`
	} `xml:"about"`
}
