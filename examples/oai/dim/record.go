package dim

import "encoding/xml"

type Record struct {
	XMLName xml.Name `xml:"Record"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text       string `xml:",chardata"`
		Status     string `xml:"status,attr"`
		Identifier struct {
			Text string `xml:",chardata"` // oai:idus.us.es:11441/1161...
		} `xml:"identifier"`
		Datestamp struct {
			Text string `xml:",chardata"` // 2016-11-30T10:33:37Z, 201...
		} `xml:"datestamp"`
		SetSpec []struct {
			Text string `xml:",chardata"` // com_11441_2824, com_11441...
		} `xml:"setSpec"`
	} `xml:"header"`
	Metadata struct {
		Text string `xml:",chardata"`
		Dim  struct {
			Text           string `xml:",chardata"` // 15, 21, 15, 11, 25, 15, 1...
			Dim            string `xml:"dim,attr"`
			Doc            string `xml:"doc,attr"`
			Xsi            string `xml:"xsi,attr"`
			SchemaLocation string `xml:"schemaLocation,attr"`
			Field          []struct {
				Text       string `xml:",chardata"` // Talego Vázquez, Félix, ...
				Mdschema   string `xml:"mdschema,attr"`
				Element    string `xml:"element,attr"`
				Lang       string `xml:"lang,attr"`
				Qualifier  string `xml:"qualifier,attr"`
				Authority  string `xml:"authority,attr"`
				Confidence string `xml:"confidence,attr"`
			} `xml:"field"`
		} `xml:"dim"`
	} `xml:"metadata"`
	About struct {
		Text string `xml:",chardata"`
	} `xml:"about"`
}
