package oaimarc

import "encoding/xml"

type Record struct {
	XMLName xml.Name `xml:"Record"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text       string `xml:",chardata"`
		Status     string `xml:"status,attr"`
		Identifier struct {
			Text string `xml:",chardata"` // oai:ojs.www.ledonline.it:...
		} `xml:"identifier"`
		Datestamp struct {
			Text string `xml:",chardata"` // 2013-12-20T09:21:34Z, 201...
		} `xml:"datestamp"`
		SetSpec struct {
			Text string `xml:",chardata"` // ECPS-Journal:SRC, ECPS-Jo...
		} `xml:"setSpec"`
	} `xml:"header"`
	Metadata struct {
		Text    string `xml:",chardata"`
		OaiMarc struct {
			Text           string `xml:",chardata"`
			Status         string `xml:"status,attr"`
			Type           string `xml:"type,attr"`
			Level          string `xml:"level,attr"`
			EncLvl         string `xml:"encLvl,attr"`
			CatForm        string `xml:"catForm,attr"`
			Xmlns          string `xml:"xmlns,attr"`
			Xsi            string `xml:"xsi,attr"`
			SchemaLocation string `xml:"schemaLocation,attr"`
			Fixfield       struct {
				Text string `xml:",chardata"` // "100622 2010             ...
				ID   string `xml:"id,attr"`
			} `xml:"fixfield"`
			Varfield []struct {
				Text     string `xml:",chardata"`
				Tag      string `xml:"tag,attr"`
				Ind1     string `xml:"ind1,attr"`
				Ind2     string `xml:"ind2,attr"`
				Subfield []struct {
					Text  string `xml:",chardata"` // 2037-7924, 2037-7932, dc,...
					Label string `xml:"label,attr"`
				} `xml:"subfield"`
			} `xml:"varfield"`
			DataField struct {
				Text     string `xml:",chardata"`
				Tag      string `xml:"tag,attr"`
				Ind1     string `xml:"ind1,attr"`
				Ind2     string `xml:"ind2,attr"`
				Subfield struct {
					Text  string `xml:",chardata"` // 2013-08-22 00:00:00, 2013...
					Label string `xml:"label,attr"`
				} `xml:"subfield"`
			} `xml:"dataField"`
		} `xml:"oai_marc"`
	} `xml:"metadata"`
	About struct {
		Text string `xml:",chardata"`
	} `xml:"about"`
}
