package misc

import "encoding/xml"

// Uketddc was generated 2017-12-06 00:04:02 by tir on apollo.
type Uketddc struct {
	XMLName        xml.Name `xml:"uketddc"`
	Text           string   `xml:",chardata"`
	Xmlns          string   `xml:"xmlns,attr"`
	Dc             string   `xml:"dc,attr"`
	Dcterms        string   `xml:"dcterms,attr"`
	Uketdterms     string   `xml:"uketdterms,attr"`
	Xsi            string   `xml:"xsi,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Title          struct {
		Text string `xml:",chardata"` // Title
	} `xml:"title"`
	Alternative struct {
		Text string `xml:",chardata"` // Alternative Title
	} `xml:"alternative"`
	Creator struct {
		Text string `xml:",chardata"` // Author
	} `xml:"creator"`
	Authoridentifier []struct {
		Text string `xml:",chardata"` // A HUSID, An ORCID, An ISN...
		Type string `xml:"type,attr"`
	} `xml:"authoridentifier"`
	Advisor struct {
		Text string `xml:",chardata"` // Supervisor(s)/advisor
	} `xml:"advisor"`
	Sponsor struct {
		Text string `xml:",chardata"` // Sponsor/Funding body
	} `xml:"sponsor"`
	Grantnumber struct {
		Text string `xml:",chardata"` // Grant number/funding code...
	} `xml:"grantnumber"`
	Subject []struct {
		Text string `xml:",chardata"` // Subject keywords, Subject...
		Type string `xml:"type,attr"`
	} `xml:"subject"`
	Abstract struct {
		Text string `xml:",chardata"` // Abstract
	} `xml:"abstract"`
	Institution struct {
		Text string `xml:",chardata"` // Institution
	} `xml:"institution"`
	Department struct {
		Text string `xml:",chardata"` // College, Department, etc....
	} `xml:"department"`
	Commercial struct {
		Text string `xml:",chardata"` // Commercial publisher
	} `xml:"commercial"`
	Issued struct {
		Text string `xml:",chardata"` // Date of submission (date ...
	} `xml:"issued"`
	DateAccepted struct {
		Text string `xml:",chardata"` // Date of Award
	} `xml:"dateAccepted"`
	Type struct {
		Text string `xml:",chardata"` // Thesis or dissertation
	} `xml:"type"`
	Qualificationlevel struct {
		Text string `xml:",chardata"` // Qualification level
	} `xml:"qualificationlevel"`
	Qualificationname struct {
		Text string `xml:",chardata"` // Qualification name
	} `xml:"qualificationname"`
	Language struct {
		Text string `xml:",chardata"` // Language
		Type string `xml:"type,attr"`
	} `xml:"language"`
	HasVersion struct {
		Text string `xml:",chardata"` // Citations
	} `xml:"hasVersion"`
	References struct {
		Text string `xml:",chardata"` // Included/Quoted work
	} `xml:"references"`
	IsReferencedBy struct {
		Text string `xml:",chardata"` // Metadata jump-off page
	} `xml:"isReferencedBy"`
	Rights struct {
		Text string `xml:",chardata"` // Rights
	} `xml:"rights"`
	Embargotype struct {
		Text string `xml:",chardata"` // Embargo type, e.g. Partia...
	} `xml:"embargotype"`
	Embargodate struct {
		Text string `xml:",chardata"` // Embargo date
	} `xml:"embargodate"`
	Embargoreason struct {
		Text string `xml:",chardata"` // Embargo reason
	} `xml:"embargoreason"`
	Identifier []struct {
		Text string `xml:",chardata"` // http://some.url, 10.1234/...
		Type string `xml:"type,attr"`
	} `xml:"identifier"`
}
