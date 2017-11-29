package qdc

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
		Text        string `xml:",chardata"`
		Qualifieddc struct {
			Text           string `xml:",chardata"`
			Qdc            string `xml:"qdc,attr"`
			Doc            string `xml:"doc,attr"`
			Xsi            string `xml:"xsi,attr"`
			Dcterms        string `xml:"dcterms,attr"`
			Dc             string `xml:"dc,attr"`
			SchemaLocation string `xml:"schemaLocation,attr"`
			Title          []struct {
				Text string `xml:",chardata"` // Democracia asamblearia y ...
			} `xml:"title"`
			Creator []struct {
				Text string `xml:",chardata"` // Talego Vázquez, Félix, ...
			} `xml:"creator"`
			Contributor []struct {
				Text string `xml:",chardata"` // Universidad de Sevilla. D...
			} `xml:"contributor"`
			DateAccepted struct {
				Text string `xml:",chardata"` // 2014-11-26T13:23:23Z, 201...
			} `xml:"dateAccepted"`
			Available struct {
				Text string `xml:",chardata"` // 2014-11-26T13:23:23Z, 201...
			} `xml:"available"`
			Created struct {
				Text string `xml:",chardata"` // 2014-11-26T13:23:23Z, 201...
			} `xml:"created"`
			Issued struct {
				Text string `xml:",chardata"` // 1996, 2009, 2002, 2004, 2...
			} `xml:"issued"`
			Type []struct {
				Text string `xml:",chardata"` // info:eu-repo/semantics/ar...
			} `xml:"type"`
			Identifier []struct {
				Text string `xml:",chardata"` // Talego Vázquez, F. (1996...
			} `xml:"identifier"`
			Language []struct {
				Text string `xml:",chardata"` // spa, spa, spa, spa, spa, ...
			} `xml:"language"`
			Relation []struct {
				Text string `xml:",chardata"` // Revista de estudios andal...
			} `xml:"relation"`
			Rights []struct {
				Text string `xml:",chardata"` // info:eu-repo/semantics/op...
			} `xml:"rights"`
			Publisher []struct {
				Text string `xml:",chardata"` // Universidad de Sevilla, U...
			} `xml:"publisher"`
			Subject []struct {
				Text string `xml:",chardata"` // MMOG, production, develop...
			} `xml:"subject"`
			Abstract []struct {
				Text string `xml:",chardata"` // El MMOG o ‘Massively Mu...
			} `xml:"abstract"`
			Source []struct {
				Text string `xml:",chardata"` // Tres grandes cuestiones d...
			} `xml:"source"`
			Coverage struct {
				Text string `xml:",chardata"` // Laayoune, Sahara, Malabo,...
			} `xml:"coverage"`
		} `xml:"qualifieddc"`
	} `xml:"metadata"`
	About struct {
		Text string `xml:",chardata"`
	} `xml:"about"`
}
