package marcxml

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
		Record struct {
			Text           string `xml:",chardata"`
			Xmlns          string `xml:"xmlns,attr"`
			Xsi            string `xml:"xsi,attr"`
			SchemaLocation string `xml:"schemaLocation,attr"`
			Leader         struct {
				Text string `xml:",chardata"` // nmb a2200000Iu 4500, nmb ...
			} `xml:"leader"`
			Controlfield struct {
				Text string `xml:",chardata"` // "150130 2015             ...
				Tag  string `xml:"tag,attr"`
			} `xml:"controlfield"`
			Datafield []struct {
				Text     string `xml:",chardata"`
				Tag      string `xml:"tag,attr"`
				Ind1     string `xml:"ind1,attr"`
				Ind2     string `xml:"ind2,attr"`
				Subfield []struct {
					Text string `xml:",chardata"` // 2297-3249, dc, 7. Wildaue...
					Code string `xml:"code,attr"`
				} `xml:"subfield"`
			} `xml:"datafield"`
			DataField struct {
				Text     string `xml:",chardata"`
				Tag      string `xml:"tag,attr"`
				Ind1     string `xml:"ind1,attr"`
				Ind2     string `xml:"ind2,attr"`
				Subfield struct {
					Text string `xml:",chardata"` // 2015-02-04 11:14:27, 2016...
					Code string `xml:"code,attr"`
				} `xml:"subfield"`
			} `xml:"dataField"`
		} `xml:"record"`
	} `xml:"metadata"`
	About struct {
		Text string `xml:",chardata"`
	} `xml:"about"`
}
