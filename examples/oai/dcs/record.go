package dcs

import "encoding/xml"

type Record struct {
	XMLName xml.Name `xml:"Record"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text       string `xml:",chardata"`
		Status     string `xml:"status,attr"`
		Identifier struct {
			Text string `xml:",chardata"` // oai:digitalcommons.risd.e...
		} `xml:"identifier"`
		Datestamp struct {
			Text string `xml:",chardata"` // 2015-06-12T16:15:26Z, 201...
		} `xml:"datestamp"`
		SetSpec []struct {
			Text string `xml:",chardata"` // publication:library, publ...
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
				Text string `xml:",chardata"` // Materials Collection Crea...
			} `xml:"title"`
			Creator []struct {
				Text string `xml:",chardata"` // Pompelia, Mark, Pompelia,...
			} `xml:"creator"`
			Subject []struct {
				Text string `xml:",chardata"` // Adult and Continuing Educ...
			} `xml:"subject"`
			Description []struct {
				Text string `xml:",chardata"` // <p>The Problem/Opportunit...
			} `xml:"description"`
			Identifier []struct {
				Text string `xml:",chardata"` // https://digitalcommons.ri...
			} `xml:"identifier"`
			Date struct {
				Text string `xml:",chardata"` // 1969-07-01T07:00:00Z, 197...
			} `xml:"date"`
			Coverage []struct {
				Text string `xml:",chardata"` // 41.826737, -71.407485, 41...
			} `xml:"coverage"`
			Thesis []struct {
				Text string `xml:",chardata"` // Thesis, Master of Fine Ar...
			} `xml:"thesis"`
			Contributor []struct {
				Text string `xml:",chardata"` // Eva Sutton, Ann Fessler, ...
			} `xml:"contributor"`
			Rights struct {
				Text string `xml:",chardata"` // http://creativecommons.or...
			} `xml:"rights"`
			Relation struct {
				Text string `xml:",chardata"` // Type 5 Design D Port Side...
			} `xml:"relation"`
			Type struct {
				Text string `xml:",chardata"` // Report, Report, Report, R...
			} `xml:"type"`
		} `xml:"dc"`
	} `xml:"metadata"`
	About struct {
		Text string `xml:",chardata"`
	} `xml:"about"`
}
