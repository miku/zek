package misc

import "encoding/xml"

// Statute was generated 2017-12-05 23:59:01 by tir on apollo.
type Statute struct {
	XMLName        xml.Name `xml:"Statute"`
	Text           string   `xml:",chardata"`
	BillOrigin     string   `xml:"bill-origin,attr"`
	BillType       string   `xml:"bill-type,attr"`
	Lang           string   `xml:"lang,attr"`
	InForce        string   `xml:"in-force,attr"`
	Startdate      string   `xml:"startdate,attr"`
	Identification struct {
		Text               string `xml:",chardata"`
		Code               string `xml:"Code,attr"`
		HasPreviousVersion string `xml:"hasPreviousVersion,attr"`
		BillNumber         struct {
			Text string `xml:",chardata"` // S-207
		} `xml:"BillNumber"`
		Parliament struct {
			Text    string `xml:",chardata"`
			Session struct {
				Text string `xml:",chardata"` // 2
			} `xml:"Session"`
			Number struct {
				Text string `xml:",chardata"` // 39
			} `xml:"Number"`
			RegnalYear struct {
				Text  string `xml:",chardata"`
				YearS struct {
					Text string `xml:",chardata"` // 56-57
				} `xml:"Year-s"`
				Monarch struct {
					Text string `xml:",chardata"` // Elizabeth II
				} `xml:"Monarch"`
			} `xml:"RegnalYear"`
			YearS struct {
				Text string `xml:",chardata"` // 2007-2008
			} `xml:"Year-s"`
		} `xml:"Parliament"`
		LongTitle struct {
			Text string `xml:",chardata"` // Loi prévoyant l’abroga...
			Code string `xml:"Code,attr"`
		} `xml:"LongTitle"`
		ShortTitle struct {
			Text   string `xml:",chardata"` // Loi sur l’abrogation de...
			Status string `xml:"status,attr"`
			Code   string `xml:"Code,attr"`
		} `xml:"ShortTitle"`
		RunningHead struct {
			Text string `xml:",chardata"` // Abrogation des lois
		} `xml:"RunningHead"`
		BillHistory struct {
			Text   string `xml:",chardata"`
			Stages []struct {
				Text  string `xml:",chardata"`
				Stage string `xml:"stage,attr"`
				Date  struct {
					Text string `xml:",chardata"`
					YYYY struct {
						Text string `xml:",chardata"` // 2017, 2008
					} `xml:"YYYY"`
					MM struct {
						Text string `xml:",chardata"` // 04, 6
					} `xml:"MM"`
					DD struct {
						Text string `xml:",chardata"` // 13, 18
					} `xml:"DD"`
				} `xml:"Date"`
			} `xml:"Stages"`
		} `xml:"BillHistory"`
		Chapter struct {
			Text               string `xml:",chardata"`
			Code               string `xml:"Code,attr"`
			ConsolidatedNumber struct {
				Text     string `xml:",chardata"` // S-21.5
				Official string `xml:"official,attr"`
				Code     string `xml:"Code,attr"`
			} `xml:"ConsolidatedNumber"`
			AnnualStatuteId struct {
				Text                string `xml:",chardata"`
				AnnualStatuteNumber struct {
					Text string `xml:",chardata"` // 20
				} `xml:"AnnualStatuteNumber"`
				YYYY struct {
					Text string `xml:",chardata"` // 2008
				} `xml:"YYYY"`
			} `xml:"AnnualStatuteId"`
		} `xml:"Chapter"`
		BillRefNumber struct {
			Text     string `xml:",chardata"` // 399998
			DateTime string `xml:"date-time,attr"`
		} `xml:"BillRefNumber"`
	} `xml:"Identification"`
	Introduction struct {
		Text   string `xml:",chardata"`
		Code   string `xml:"Code,attr"`
		Enacts struct {
			Text      string `xml:",chardata"`
			Code      string `xml:"Code,attr"`
			Provision struct {
				Chardata      string `xml:",chardata"`
				FormatRef     string `xml:"format-ref,attr"`
				LanguageAlign string `xml:"language-align,attr"`
				Code          string `xml:"Code,attr"`
				Text          struct {
					Text string `xml:",chardata"` // Sa Majesté, sur l’avis...
				} `xml:"Text"`
			} `xml:"Provision"`
		} `xml:"Enacts"`
	} `xml:"Introduction"`
	Body struct {
		Text    string `xml:",chardata"`
		Section []struct {
			Chardata     string `xml:",chardata"`
			Code         string `xml:"Code,attr"`
			MarginalNote struct {
				Text         string `xml:",chardata"` // Titre abrégé, Rapport a...
				Code         string `xml:"Code,attr"`
				XRefExternal struct {
					Text          string `xml:",chardata"` // Gazette du Canada
					ReferenceType string `xml:"reference-type,attr"`
					Link          string `xml:"link,attr"`
				} `xml:"XRefExternal"`
			} `xml:"MarginalNote"`
			Label struct {
				Text        string `xml:",chardata"` // 1, 2, 3, 4, 5, 6
				FootnoteRef struct {
					Text  string `xml:",chardata"` // *
					Idref string `xml:"idref,attr"`
				} `xml:"FootnoteRef"`
			} `xml:"Label"`
			Text struct {
				Text         string `xml:",chardata"` // Titre abrégé :, ., Le ...
				XRefExternal struct {
					Text          string `xml:",chardata"` // Loi sur l’abrogation de...
					ReferenceType string `xml:"reference-type,attr"`
					Link          string `xml:"link,attr"`
				} `xml:"XRefExternal"`
			} `xml:"Text"`
			Paragraph []struct {
				Chardata string `xml:",chardata"`
				Code     string `xml:"Code,attr"`
				Label    struct {
					Text string `xml:",chardata"` // a), b)
				} `xml:"Label"`
				Text struct {
					Text string `xml:",chardata"` // d’une part, ont été s...
				} `xml:"Text"`
			} `xml:"Paragraph"`
			Footnote struct {
				Chardata  string `xml:",chardata"`
				ID        string `xml:"id,attr"`
				Placement string `xml:"placement,attr"`
				Status    string `xml:"status,attr"`
				Label     struct {
					Text string `xml:",chardata"` // *
				} `xml:"Label"`
				Text struct {
					Text string `xml:",chardata"` // [Note : Loi en vigueur l...
				} `xml:"Text"`
			} `xml:"Footnote"`
		} `xml:"Section"`
	} `xml:"Body"`
}
