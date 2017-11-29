package ore

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
		Text  string `xml:",chardata"`
		Entry struct {
			Text           string `xml:",chardata"`
			Atom           string `xml:"atom,attr"`
			Doc            string `xml:"doc,attr"`
			Dcterms        string `xml:"dcterms,attr"`
			Oreatom        string `xml:"oreatom,attr"`
			Ore            string `xml:"ore,attr"`
			Xsi            string `xml:"xsi,attr"`
			SchemaLocation string `xml:"schemaLocation,attr"`
			ID             struct {
				Text string `xml:",chardata"` // http://hdl.handle.net/114...
			} `xml:"id"`
			Link []struct {
				Text   string `xml:",chardata"`
				Rel    string `xml:"rel,attr"`
				Href   string `xml:"href,attr"`
				Type   string `xml:"type,attr"`
				Title  string `xml:"title,attr"`
				Length string `xml:"length,attr"`
			} `xml:"link"`
			Published struct {
				Text string `xml:",chardata"` // 2014-11-26T13:23:23Z, 201...
			} `xml:"published"`
			Updated struct {
				Text string `xml:",chardata"` // 2014-11-26T13:23:23Z, 201...
			} `xml:"updated"`
			Source struct {
				Text      string `xml:",chardata"`
				Generator struct {
					Text string `xml:",chardata"` // idUS - Universidad de Sev...
				} `xml:"generator"`
			} `xml:"source"`
			Title []struct {
				Text string `xml:",chardata"` // Democracia asamblearia y ...
			} `xml:"title"`
			Category []struct {
				Text   string `xml:",chardata"`
				Label  string `xml:"label,attr"`
				Term   string `xml:"term,attr"`
				Scheme string `xml:"scheme,attr"`
			} `xml:"category"`
			Triples struct {
				Text        string `xml:",chardata"`
				Description []struct {
					Text  string `xml:",chardata"`
					Rdf   string `xml:"rdf,attr"`
					About string `xml:"about,attr"`
					Type  struct {
						Text     string `xml:",chardata"`
						Resource string `xml:"resource,attr"`
					} `xml:"type"`
					Modified struct {
						Text string `xml:",chardata"` // 2014-11-26T13:23:23Z, 201...
					} `xml:"modified"`
					Description struct {
						Text string `xml:",chardata"` // ORIGINAL, ORIGINAL, ORIGI...
					} `xml:"description"`
				} `xml:"Description"`
			} `xml:"triples"`
			Author struct {
				Text string `xml:",chardata"`
				Name struct {
					Text string `xml:",chardata"` // Oña Cots, José Manuel D...
				} `xml:"name"`
			} `xml:"author"`
		} `xml:"entry"`
	} `xml:"metadata"`
	About struct {
		Text string `xml:",chardata"`
	} `xml:"about"`
}
