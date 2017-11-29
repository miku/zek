package marc21

import "encoding/xml"

type Record struct {
	XMLName xml.Name `xml:"Record"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text       string `xml:",chardata"`
		Status     string `xml:"status,attr"`
		Identifier struct {
			Text string `xml:",chardata"` // oai:idsluzern.ch:ILU01-00...
		} `xml:"identifier"`
		Datestamp struct {
			Text string `xml:",chardata"` // 2012-07-15T20:38:46Z, 201...
		} `xml:"datestamp"`
		SetSpec struct {
			Text string `xml:",chardata"` // OAI_IDS_LUZERN, OAI_IDS_L...
		} `xml:"setSpec"`
	} `xml:"header"`
	Metadata struct {
		Text   string `xml:",chardata"`
		Record struct {
			Text           string `xml:",chardata"`
			Marc           string `xml:"marc,attr"`
			Xsi            string `xml:"xsi,attr"`
			SchemaLocation string `xml:"schemaLocation,attr"`
			Leader         struct {
				Text string `xml:",chardata"` // 00763nam  22001934c 4500,...
			} `xml:"leader"`
			Controlfield []struct {
				Text string `xml:",chardata"` // cr, 930517s1925    sz    ...
				Tag  string `xml:"tag,attr"`
			} `xml:"controlfield"`
			Datafield []struct {
				Text     string `xml:",chardata"`
				Tag      string `xml:"tag,attr"`
				Ind1     string `xml:"ind1,attr"`
				Ind2     string `xml:"ind2,attr"`
				Subfield []struct {
					Text string `xml:",chardata"` // SzZuIDS LU LUZHB, kids, 5...
					Code string `xml:"code,attr"`
				} `xml:"subfield"`
			} `xml:"datafield"`
		} `xml:"record"`
	} `xml:"metadata"`
	About struct {
		Text string `xml:",chardata"`
	} `xml:"about"`
}
