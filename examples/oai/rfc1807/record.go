package rfc1807

import "encoding/xml"

type Record struct {
	XMLName xml.Name `xml:"Record"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text       string `xml:",chardata"`
		Status     string `xml:"status,attr"`
		Identifier struct {
			Text string `xml:",chardata"` // oai:ojs.ub-ojs.ub.unibas....
		} `xml:"identifier"`
		Datestamp struct {
			Text string `xml:",chardata"` // 2013-03-21T09:05:39Z, 201...
		} `xml:"datestamp"`
		SetSpec struct {
			Text string `xml:",chardata"` // cdrs_0277:ART, cdrs_0277:...
		} `xml:"setSpec"`
	} `xml:"header"`
	Metadata struct {
		Text    string `xml:",chardata"`
		Rfc1807 struct {
			Text           string `xml:",chardata"`
			Xmlns          string `xml:"xmlns,attr"`
			Xsi            string `xml:"xsi,attr"`
			SchemaLocation string `xml:"schemaLocation,attr"`
			BibVersion     struct {
				Text string `xml:",chardata"` // v2, v2, v2, v2, v2, v2, v...
			} `xml:"bib-version"`
			ID struct {
				Text string `xml:",chardata"` // http://0277.ch/ojs/index....
			} `xml:"id"`
			Entry struct {
				Text string `xml:",chardata"` // 2013-03-21T09:05:39Z, 201...
			} `xml:"entry"`
			Organization []struct {
				Text string `xml:",chardata"` // 027.7 Zeitschrift für Bi...
			} `xml:"organization"`
			Title struct {
				Text string `xml:",chardata"` // Die Bibliothek 2.0 ist to...
			} `xml:"title"`
			Type struct {
				Text string `xml:",chardata"`
			} `xml:"type"`
			Author []struct {
				Text string `xml:",chardata"` // Tréfás, David; Universi...
			} `xml:"author"`
			Date struct {
				Text string `xml:",chardata"` // 2013-03-12 08:25:58, 2013...
			} `xml:"date"`
			Copyright struct {
				Text string `xml:",chardata"`
			} `xml:"copyright"`
			OtherAccess struct {
				Text string `xml:",chardata"` // url:http://0277.ch/ojs/in...
			} `xml:"other_access"`
			Keyword struct {
				Text string `xml:",chardata"`
			} `xml:"keyword"`
			Period []struct {
				Text string `xml:",chardata"`
			} `xml:"period"`
			Monitoring struct {
				Text string `xml:",chardata"`
			} `xml:"monitoring"`
			Language struct {
				Text string `xml:",chardata"` // de, en, de, de, de, de, d...
			} `xml:"language"`
			Abstract struct {
				Text string `xml:",chardata"` // DOI: 10.12685/027.7-1-1-1...
			} `xml:"abstract"`
		} `xml:"rfc1807"`
	} `xml:"metadata"`
	About struct {
		Text string `xml:",chardata"`
	} `xml:"about"`
}
