package fiperson

import "encoding/xml"

type Record struct {
	XMLName xml.Name `xml:"Record"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text       string `xml:",chardata"`
		Status     string `xml:"status,attr"`
		Identifier struct {
			Text string `xml:",chardata"` // oai:pure.atira.dk:persons...
		} `xml:"identifier"`
		Datestamp struct {
			Text string `xml:",chardata"` // 2014-08-05T13:48:03Z, 201...
		} `xml:"datestamp"`
		SetSpec struct {
			Text string `xml:",chardata"` // persons:all, persons:all,...
		} `xml:"setSpec"`
	} `xml:"header"`
	Metadata struct {
		Text    string `xml:",chardata"`
		Persons struct {
			Text           string `xml:",chardata"`
			Xmlns          string `xml:"xmlns,attr"`
			Xsi            string `xml:"xsi,attr"`
			SchemaLocation string `xml:"schemaLocation,attr"`
			Person         struct {
				Text       string `xml:",chardata"`
				RecUpdated string `xml:"rec_updated,attr"`
				RecCreated string `xml:"rec_created,attr"`
				RecSource  string `xml:"rec_source,attr"`
				RecID      string `xml:"rec_id,attr"`
				Name       struct {
					Text  string `xml:",chardata"`
					First struct {
						Text string `xml:",chardata"` // Kim, Helle Anita, Jeppe, ...
					} `xml:"first"`
					Last struct {
						Text string `xml:",chardata"` // Brasen, Brøns, Priess Ge...
					} `xml:"last"`
				} `xml:"name"`
				Affiliation []struct {
					Text    string `xml:",chardata"`
					Started struct {
						Text string `xml:",chardata"` // 2013-03-14+01:00, 2013-04...
					} `xml:"started"`
					Name []struct {
						Text   string `xml:",chardata"`
						Lang   string `xml:"lang,attr"`
						Level1 struct {
							Text string `xml:",chardata"` // Kulturministeriet, Minist...
						} `xml:"level1"`
						Level2 struct {
							Text string `xml:",chardata"` // SMK Statens Museum for Ku...
						} `xml:"level2"`
						Acronym struct {
							Text string `xml:",chardata"` // SMK, SMK, SMK, SMK, SMK, ...
						} `xml:"acronym"`
						Level3 struct {
							Text string `xml:",chardata"` // CATS, CATS, CATS, CATS, C...
						} `xml:"level3"`
					} `xml:"name"`
					Ended struct {
						Text string `xml:",chardata"` // 2004-12-31+01:00, 2013-12...
					} `xml:"ended"`
				} `xml:"affiliation"`
				Email struct {
					Text string `xml:",chardata"` // bavn@kb.dk, lgot@kb.dk, l...
				} `xml:"email"`
				NameAlternative struct {
					Text  string `xml:",chardata"`
					First struct {
						Text string `xml:",chardata"` // Henrik, Mikkel, Lasse
					} `xml:"first"`
					Last struct {
						Text string `xml:",chardata"` // Holm, Kirkedahl, Bendtsen...
					} `xml:"last"`
				} `xml:"name_alternative"`
			} `xml:"person"`
		} `xml:"persons"`
	} `xml:"metadata"`
	About struct {
		Text string `xml:",chardata"`
	} `xml:"about"`
}
