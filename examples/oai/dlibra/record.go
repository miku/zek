package dlibra

import "encoding/xml"

type Record struct {
	XMLName xml.Name `xml:"Record"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text       string `xml:",chardata"`
		Status     string `xml:"status,attr"`
		Identifier struct {
			Text string `xml:",chardata"` // oai:www.wbc.poznan.pl:420...
		} `xml:"identifier"`
		Datestamp struct {
			Text string `xml:",chardata"` // 2004-10-28T13:48:47Z, 200...
		} `xml:"datestamp"`
		SetSpec []struct {
			Text string `xml:",chardata"` // rootCollection:wbc:Digita...
		} `xml:"setSpec"`
	} `xml:"header"`
	Metadata struct {
		Text                string `xml:",chardata"`
		MetadataDescription struct {
			Text           string `xml:",chardata"`
			DlibraAvs      string `xml:"dlibra_avs,attr"`
			Xsi            string `xml:"xsi,attr"`
			SchemaLocation string `xml:"schemaLocation,attr"`
			Title          struct {
				Text  string `xml:",chardata"`
				Value struct {
					Text string `xml:",chardata"` // Berliner St. Bonifacius-K...
					Lang string `xml:"lang,attr"`
				} `xml:"value"`
			} `xml:"Title"`
			Creator struct {
				Text        string `xml:",chardata"`
				Contributor []struct {
					Text string `xml:",chardata"` // Müller, E., Boehn von, M...
					Lang string `xml:"lang,attr"`
				} `xml:"Contributor"`
			} `xml:"Creator"`
			Subject struct {
				Text  string `xml:",chardata"`
				Value struct {
					Text string `xml:",chardata"` // Kalendarze, 19 w., Kalend...
					Lang string `xml:"lang,attr"`
				} `xml:"value"`
				Tags struct {
					Text string `xml:",chardata"`
				} `xml:"Tags"`
			} `xml:"Subject"`
			Description struct {
				Text              string `xml:",chardata"`
				DegreeInformation struct {
					Text string `xml:",chardata"`
				} `xml:"DegreeInformation"`
				Value struct {
					Text string `xml:",chardata"` // Estr. XXXI, 287, defekt: ...
					Lang string `xml:"lang,attr"`
				} `xml:"value"`
			} `xml:"Description"`
			Publisher struct {
				Text  string `xml:",chardata"`
				Value []struct {
					Text string `xml:",chardata"` // Berlin, G. Jansen, F. Bru...
					Lang string `xml:"lang,attr"`
				} `xml:"value"`
			} `xml:"Publisher"`
			Type struct {
				Text string `xml:",chardata"` // kalendarze, kalendarze, k...
				Lang string `xml:"lang,attr"`
			} `xml:"Type"`
			Identifier struct {
				Text string `xml:",chardata"`
			} `xml:"Identifier"`
			Language struct {
				Text string `xml:",chardata"` // ger, ger, eng, eng, ger, ...
				Lang string `xml:"lang,attr"`
			} `xml:"Language"`
			Relation struct {
				Text string `xml:",chardata"`
			} `xml:"Relation"`
			Rights struct {
				Text string `xml:",chardata"` // PAN Biblioteka Kórnicka,...
				Lang string `xml:"lang,attr"`
			} `xml:"Rights"`
			CreationPublicationDate struct {
				Text string `xml:",chardata"`
				Date struct {
					Text string `xml:",chardata"` // 1869, 1912, 1880, 1868, 1...
					Lang string `xml:"lang,attr"`
				} `xml:"Date"`
			} `xml:"CreationPublicationDate"`
			AccessRights struct {
				Text string `xml:",chardata"`
			} `xml:"AccessRights"`
			RightsHolder struct {
				Text string `xml:",chardata"` // PAN Biblioteka Kórnicka,...
				Lang string `xml:"lang,attr"`
			} `xml:"RightsHolder"`
			Digitisation struct {
				Text  string `xml:",chardata"`
				Value struct {
					Text string `xml:",chardata"` // PAN Biblioteka Kórnicka,...
					Lang string `xml:"lang,attr"`
				} `xml:"value"`
			} `xml:"Digitisation"`
			LocationOfPhysicalObject struct {
				Text  string `xml:",chardata"`
				Value struct {
					Text string `xml:",chardata"` // PAN Biblioteka Kórnicka,...
					Lang string `xml:"lang,attr"`
				} `xml:"value"`
			} `xml:"LocationOfPhysicalObject"`
			Format struct {
				Text string `xml:",chardata"` // image/x.djvu, image/x.djv...
			} `xml:"Format"`
		} `xml:"metadataDescription"`
	} `xml:"metadata"`
	About struct {
		Text string `xml:",chardata"`
	} `xml:"about"`
}
