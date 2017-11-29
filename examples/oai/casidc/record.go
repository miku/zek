package casidc

import "encoding/xml"

type Record struct {
	XMLName xml.Name `xml:"Record"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text       string `xml:",chardata"`
		Status     string `xml:"status,attr"`
		Identifier struct {
			Text string `xml:",chardata"` // oai:casi.ntrs.nasa.gov:20...
		} `xml:"identifier"`
		Datestamp struct {
			Text string `xml:",chardata"` // 2013-04-10, 2013-04-10, 2...
		} `xml:"datestamp"`
	} `xml:"header"`
	Metadata struct {
		Text   string `xml:",chardata"`
		Casidc struct {
			Text             string `xml:",chardata"`
			CasiDc           string `xml:"casi_dc,attr"`
			Casiterms        string `xml:"casiterms,attr"`
			Xsi              string `xml:"xsi,attr"`
			Dc               string `xml:"dc,attr"`
			SchemaLocation   string `xml:"schemaLocation,attr"`
			IdentifierCasiID struct {
				Text string `xml:",chardata"` // 20120001369, 20120001368,...
			} `xml:"identifier.casi_id"`
			Title struct {
				Text string `xml:",chardata"` // Probabilistic Risk Assess...
			} `xml:"title"`
			Description struct {
				Text string `xml:",chardata"` // Probabilistic Risk Assess...
			} `xml:"description"`
			DateIssued []struct {
				Text string `xml:",chardata"` // 20111201, December 2011, ...
			} `xml:"date.issued"`
			Rights struct {
				Text string `xml:",chardata"` // Copyright, Distribution u...
			} `xml:"rights"`
			RightsAccessRights struct {
				Text string `xml:",chardata"` // Unclassified, Unlimited, ...
			} `xml:"rights.accessRights"`
			Identifier []struct {
				Text string `xml:",chardata"` // http://hdl.handle.net/206...
			} `xml:"identifier"`
			RelationRequires struct {
				Text string `xml:",chardata"` // CASI, CASI, CASI, CASI, C...
			} `xml:"relation.requires"`
			Format struct {
				Text string `xml:",chardata"` // application/pdf, applicat...
			} `xml:"format"`
			Creator []struct {
				Text string `xml:",chardata"` // Apostolakis, George, Ever...
			} `xml:"creator"`
			Language struct {
				Text string `xml:",chardata"` // English, English, English...
			} `xml:"language"`
			Type struct {
				Text string `xml:",chardata"` // Technical Report, Oral/Vi...
			} `xml:"type"`
			ContributorOriginator []struct {
				Text string `xml:",chardata"` // NASA, United Space Allian...
			} `xml:"contributor.originator"`
			Subject struct {
				Text string `xml:",chardata"` // Administration and Manage...
			} `xml:"subject"`
			DateAvailable struct {
				Text string `xml:",chardata"` // 2013-04-08, 2013-04-08, 2...
			} `xml:"date.available"`
			IdentifierBibliographicCitation []struct {
				Text string `xml:",chardata"` // NASA Recycling and Affirm...
			} `xml:"identifier.bibliographicCitation"`
		} `xml:"casidc"`
	} `xml:"metadata"`
	About struct {
		Text string `xml:",chardata"`
	} `xml:"about"`
}
