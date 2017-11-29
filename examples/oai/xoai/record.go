package xoai

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
		Text     string `xml:",chardata"`
		Metadata struct {
			Text           string `xml:",chardata"`
			Xmlns          string `xml:"xmlns,attr"`
			Xsi            string `xml:"xsi,attr"`
			SchemaLocation string `xml:"schemaLocation,attr"`
			Element        []struct {
				Text    string `xml:",chardata"`
				Name    string `xml:"name,attr"`
				Element []struct {
					Text    string `xml:",chardata"`
					Name    string `xml:"name,attr"`
					Element []struct {
						Text  string `xml:",chardata"`
						Name  string `xml:"name,attr"`
						Field []struct {
							Text string `xml:",chardata"` // Talego Vázquez, Félix, ...
							Name string `xml:"name,attr"`
						} `xml:"field"`
						Element []struct {
							Text  string `xml:",chardata"`
							Name  string `xml:"name,attr"`
							Field []struct {
								Text string `xml:",chardata"` // 2014-11-26T13:23:23Z, 201...
								Name string `xml:"name,attr"`
							} `xml:"field"`
						} `xml:"element"`
					} `xml:"element"`
					Field struct {
						Text string `xml:",chardata"` // ORIGINAL, ORIGINAL, ORIGI...
						Name string `xml:"name,attr"`
					} `xml:"field"`
				} `xml:"element"`
				Field []struct {
					Text string `xml:",chardata"` // 11441/11610, oai:idus.us....
					Name string `xml:"name,attr"`
				} `xml:"field"`
			} `xml:"element"`
		} `xml:"metadata"`
	} `xml:"metadata"`
	About struct {
		Text string `xml:",chardata"`
	} `xml:"about"`
}
