package dcq

import "encoding/xml"

type Record struct {
	XMLName xml.Name `xml:"Record"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text       string `xml:",chardata"`
		Status     string `xml:"status,attr"`
		Identifier struct {
			Text string `xml:",chardata"` // oai:scholarworks.harding....
		} `xml:"identifier"`
		Datestamp struct {
			Text string `xml:",chardata"` // 2016-10-26T20:36:07Z, 201...
		} `xml:"datestamp"`
		SetSpec []struct {
			Text string `xml:",chardata"` // publication:spring-sing, ...
		} `xml:"setSpec"`
	} `xml:"header"`
	Metadata struct {
		Text string `xml:",chardata"`
		Dc   struct {
			Text           string `xml:",chardata"`
			OaiDc          string `xml:"oai_dc,attr"`
			Dc             string `xml:"dc,attr"`
			Xsi            string `xml:"xsi,attr"`
			SchemaLocation string `xml:"schemaLocation,attr"`
			Title          struct {
				Text string `xml:",chardata"` // Harding College Spring Si...
			} `xml:"title"`
			Creator []struct {
				Text string `xml:",chardata"` // Plumark, Inc., Palmer & P...
			} `xml:"creator"`
			DescriptionAbstract struct {
				Text string `xml:",chardata"` // <p>Program for the 1974 S...
			} `xml:"description.abstract"`
			DateCreated struct {
				Text string `xml:",chardata"` // 1974-04-01T07:00:00Z, 197...
			} `xml:"date.created"`
			Type struct {
				Text string `xml:",chardata"` // Book, Book, Book, Book, B...
			} `xml:"type"`
			Publisher struct {
				Text string `xml:",chardata"` // Harding College, Harding ...
			} `xml:"publisher"`
			Subject []struct {
				Text string `xml:",chardata"` // Spring Sing (Harding Univ...
			} `xml:"subject"`
			Description []struct {
				Text string `xml:",chardata"` // From the collection of Ha...
			} `xml:"description"`
			Identifier struct {
				Text string `xml:",chardata"` // https://scholarworks.hard...
			} `xml:"identifier"`
			Rights struct {
				Text string `xml:",chardata"` // Harding University, Hardi...
			} `xml:"rights"`
			CoverageSpatial struct {
				Text string `xml:",chardata"` // Searcy, Searcy, Searcy, S...
			} `xml:"coverage.spatial"`
		} `xml:"dc"`
	} `xml:"metadata"`
	About struct {
		Text string `xml:",chardata"`
	} `xml:"about"`
}
