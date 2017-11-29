package rioxx

import "encoding/xml"

type Record struct {
	XMLName xml.Name `xml:"Record"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text       string `xml:",chardata"`
		Status     string `xml:"status,attr"`
		Identifier struct {
			Text string `xml:",chardata"` // oai:wrap.warwick.ac.uk:6,...
		} `xml:"identifier"`
		Datestamp struct {
			Text string `xml:",chardata"` // 2013-09-11T07:00:02Z, 201...
		} `xml:"datestamp"`
		SetSpec []struct {
			Text string `xml:",chardata"` // 7375626A656374733D48:4842...
		} `xml:"setSpec"`
	} `xml:"header"`
	Metadata struct {
		Text  string `xml:",chardata"`
		Rioxx struct {
			Text           string `xml:",chardata"`
			Xmlns          string `xml:"xmlns,attr"`
			Ali            string `xml:"ali,attr"`
			Dc             string `xml:"dc,attr"`
			Dcterms        string `xml:"dcterms,attr"`
			Rioxxterms     string `xml:"rioxxterms,attr"`
			SchemaLocation string `xml:"schemaLocation,attr"`
			Xsi            string `xml:"xsi,attr"`
			FreeToRead     struct {
				Text      string `xml:",chardata"`
				StartDate string `xml:"start_date,attr"`
			} `xml:"free_to_read"`
			Description struct {
				Text string `xml:",chardata"` // We offer a new proof that...
			} `xml:"description"`
			Format struct {
				Text string `xml:",chardata"` // application/pdf, applicat...
			} `xml:"format"`
			Identifier struct {
				Text string `xml:",chardata"` // http://wrap.warwick.ac.uk...
			} `xml:"identifier"`
			Language struct {
				Text string `xml:",chardata"` // en, en, en, en, en, en, e...
			} `xml:"language"`
			Publisher struct {
				Text string `xml:",chardata"` // Berkeley Electronic Press...
			} `xml:"publisher"`
			Relation []struct {
				Text string `xml:",chardata"` // http://www.bepress.com/cg...
			} `xml:"relation"`
			Source struct {
				Text string `xml:",chardata"` // 1935-1704, ALISS Quarterl...
			} `xml:"source"`
			Subject []struct {
				Text string `xml:",chardata"` // HB, ZA4050, QA, HB, HC, H...
			} `xml:"subject"`
			Title struct {
				Text string `xml:",chardata"` // Identification of prefere...
			} `xml:"title"`
			Author []struct {
				Text string `xml:",chardata"` // Carvajal, Andrés M., Ri...
			} `xml:"author"`
			Type struct {
				Text string `xml:",chardata"` // Journal Article/Review, J...
			} `xml:"type"`
			Version struct {
				Text string `xml:",chardata"` // NA, NA, NA, NA, NA, NA, N...
			} `xml:"version"`
			VersionOfRecord struct {
				Text string `xml:",chardata"` // http://dx.doi.org/10.2202...
			} `xml:"version_of_record"`
			PublicationDate struct {
				Text string `xml:",chardata"` // 2008-01, 2008, 2006-09-10...
			} `xml:"publication_date"`
			Contributor []struct {
				Text string `xml:",chardata"` // Shah, Ghanshyam, Rutten ,...
			} `xml:"contributor"`
			LicenseRef struct {
				Text      string `xml:",chardata"` // http://creativecommons.or...
				StartDate string `xml:"start_date,attr"`
			} `xml:"license_ref"`
		} `xml:"rioxx"`
	} `xml:"metadata"`
	About struct {
		Text string `xml:",chardata"`
	} `xml:"about"`
}
