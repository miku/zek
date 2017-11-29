package rematom

import "encoding/xml"

type Record struct {
	XMLName xml.Name `xml:"Record"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text       string `xml:",chardata"`
		Status     string `xml:"status,attr"`
		Identifier struct {
			Text string `xml:",chardata"` // oai:edtb.euskomedia.org:1...
		} `xml:"identifier"`
		Datestamp struct {
			Text string `xml:",chardata"` // 2010-05-24T15:00:27Z, 201...
		} `xml:"datestamp"`
		SetSpec []struct {
			Text string `xml:",chardata"` // 7374617475733D707562, 737...
		} `xml:"setSpec"`
	} `xml:"header"`
	Metadata struct {
		Text string `xml:",chardata"`
		Atom struct {
			Text           string `xml:",chardata"`
			Atom           string `xml:"atom,attr"`
			Xmlns          string `xml:"xmlns,attr"`
			Dcterms        string `xml:"dcterms,attr"`
			Rdf            string `xml:"rdf,attr"`
			SchemaLocation string `xml:"schemaLocation,attr"`
			Feed           struct {
				Text    string `xml:",chardata"`
				Xmlns   string `xml:"xmlns,attr"`
				Rdf     string `xml:"rdf,attr"`
				Rdfs    string `xml:"rdfs,attr"`
				Ore     string `xml:"ore,attr"`
				Dcterms string `xml:"dcterms,attr"`
				ID      struct {
					Text string `xml:",chardata"` // http://edtb.euskomedia.or...
				} `xml:"id"`
				Link struct {
					Text string `xml:",chardata"`
					Href string `xml:"href,attr"`
					Rel  string `xml:"rel,attr"`
					Type string `xml:"type,attr"`
				} `xml:"link"`
				Updated struct {
					Text string `xml:",chardata"` // 2010-05-24T15:00:27Z, 201...
				} `xml:"updated"`
				Generator struct {
					Text string `xml:",chardata"` // EDTB, EDTB, EDTB, EDTB, E...
					URL  string `xml:"url,attr"`
				} `xml:"generator"`
				Category struct {
					Text   string `xml:",chardata"`
					Scheme string `xml:"scheme,attr"`
					Term   string `xml:"term,attr"`
					Label  string `xml:"label,attr"`
				} `xml:"category"`
				Title struct {
					Text string `xml:",chardata"` // Sistemas de planificació...
				} `xml:"title"`
				Author struct {
					Text string `xml:",chardata"`
					Name struct {
						Text string `xml:",chardata"` // edtb EPrints Repository @...
					} `xml:"name"`
				} `xml:"author"`
				Entry []struct {
					Text string `xml:",chardata"`
					ID   struct {
						Text string `xml:",chardata"` // http://oreproxy.org/r?wha...
					} `xml:"id"`
					Link []struct {
						Text string `xml:",chardata"`
						Href string `xml:"href,attr"`
						Rel  string `xml:"rel,attr"`
						Type string `xml:"type,attr"`
					} `xml:"link"`
					Title struct {
						Text string `xml:",chardata"` // Splash Page for "Sistemas...
					} `xml:"title"`
					Updated struct {
						Text string `xml:",chardata"` // 2010-05-24T15:00:27Z, 201...
					} `xml:"updated"`
					Category struct {
						Text   string `xml:",chardata"`
						Scheme string `xml:"scheme,attr"`
						Term   string `xml:"term,attr"`
						Label  string `xml:"label,attr"`
					} `xml:"category"`
					Description struct {
						Text  string `xml:",chardata"`
						About string `xml:"about,attr"`
						Title struct {
							Text string `xml:",chardata"` // EP3 XML, EP3 XML, EP3 XML...
						} `xml:"title"`
						Format struct {
							Text string `xml:",chardata"` // text/xml; charset=utf-8, ...
						} `xml:"format"`
						ConformsTo struct {
							Text string `xml:",chardata"` // http://eprints.org/ep2/da...
						} `xml:"conformsTo"`
					} `xml:"Description"`
				} `xml:"entry"`
			} `xml:"feed"`
		} `xml:"atom"`
	} `xml:"metadata"`
	About struct {
		Text string `xml:",chardata"`
	} `xml:"about"`
}
