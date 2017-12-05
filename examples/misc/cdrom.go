package misc

import "encoding/xml"

// CDROM was generated 2017-12-05 23:41:13 by tir on apollo.
type CDROM struct {
	XMLName  xml.Name `xml:"CDROM"`
	Text     string   `xml:",chardata"`
	CHAPITRE struct {
		Text    string `xml:",chardata"`
		Dossier string `xml:"dossier,attr"`
		Ref     string `xml:"ref,attr"`
		Titre   string `xml:"titre,attr"`
		THEME   []struct {
			Text  string `xml:",chardata"`
			Ref   string `xml:"ref,attr"`
			Titre string `xml:"titre,attr"`
			PAGE  []struct {
				Text  string `xml:",chardata"`
				Ref   string `xml:"ref,attr"`
				Icone string `xml:"icone,attr"`
				Titre string `xml:"titre,attr"`
				IMAGE []struct {
					Text    string `xml:",chardata"`
					Adresse string `xml:"adresse,attr"`
					LEGENDE struct {
						Text string `xml:",chardata"` // Vue schématique du syst�...
					} `xml:"LEGENDE"`
					CREDITS struct {
						Text string `xml:",chardata"` // Observatoire de Paris-Meu...
					} `xml:"CREDITS"`
				} `xml:"IMAGE"`
				TEXTE struct {
					Text    string `xml:",chardata"` // Les astres du système so...
					Ref     string `xml:"ref,attr"`
					MOTBLEU []struct {
						Text    string `xml:",chardata"` // plus, Distance au Soleil,...
						Adresse string `xml:"adresse,attr"`
					} `xml:"MOTBLEU"`
				} `xml:"TEXTE"`
			} `xml:"PAGE"`
		} `xml:"THEME"`
	} `xml:"CHAPITRE"`
}
