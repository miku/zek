package nlm

import "encoding/xml"

type Record struct {
	XMLName xml.Name `xml:"Record"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text       string `xml:",chardata"`
		Status     string `xml:"status,attr"`
		Identifier struct {
			Text string `xml:",chardata"` // oai:ojs.ub-ojs.ub.unibas....
		} `xml:"identifier"`
		Datestamp struct {
			Text string `xml:",chardata"` // 2013-03-21T09:05:39Z, 201...
		} `xml:"datestamp"`
		SetSpec struct {
			Text string `xml:",chardata"` // cdrs_0277:ART, cdrs_0277:...
		} `xml:"setSpec"`
	} `xml:"header"`
	Metadata struct {
		Text    string `xml:",chardata"`
		Article struct {
			Text           string `xml:",chardata"`
			Xmlns          string `xml:"xmlns,attr"`
			Xlink          string `xml:"xlink,attr"`
			Mml            string `xml:"mml,attr"`
			Xsi            string `xml:"xsi,attr"`
			SchemaLocation string `xml:"schemaLocation,attr"`
			Lang           string `xml:"lang,attr"`
			Front          struct {
				Text        string `xml:",chardata"`
				JournalMeta struct {
					Text      string `xml:",chardata"`
					JournalID struct {
						Text          string `xml:",chardata"` // cdrs_0277, cdrs_0277, cdr...
						JournalIDType string `xml:"journal-id-type,attr"`
					} `xml:"journal-id"`
					JournalTitle struct {
						Text string `xml:",chardata"` // 027.7 Zeitschrift für Bi...
					} `xml:"journal-title"`
					TransTitle struct {
						Text string `xml:",chardata"` // 027.7 Zeitschrift für Bi...
						Lang string `xml:"lang,attr"`
					} `xml:"trans-title"`
					ISSN struct {
						Text    string `xml:",chardata"` // 2296-0597, 2296-0597, 229...
						PubType string `xml:"pub-type,attr"`
					} `xml:"issn"`
				} `xml:"journal-meta"`
				ArticleMeta struct {
					Text      string `xml:",chardata"`
					ArticleID []struct {
						Text      string `xml:",chardata"` // 12, 10.12685/027.7-1-1-12...
						PubIDType string `xml:"pub-id-type,attr"`
					} `xml:"article-id"`
					ArticleCategories struct {
						Text      string `xml:",chardata"`
						SubjGroup struct {
							Text          string `xml:",chardata"`
							SubjGroupType string `xml:"subj-group-type,attr"`
							Subject       struct {
								Text string `xml:",chardata"` // Artikel / Articles, Stand...
							} `xml:"subject"`
						} `xml:"subj-group"`
					} `xml:"article-categories"`
					TitleGroup struct {
						Text         string `xml:",chardata"`
						ArticleTitle struct {
							Text string `xml:",chardata"` // Die Bibliothek 2.0 ist to...
						} `xml:"article-title"`
						TransTitle struct {
							Text string `xml:",chardata"` // Die Bibliothek 2.0 ist to...
							Lang string `xml:"lang,attr"`
						} `xml:"trans-title"`
					} `xml:"title-group"`
					ContribGroup struct {
						Text    string `xml:",chardata"`
						Contrib []struct {
							Text        string `xml:",chardata"`
							Corresp     string `xml:"corresp,attr"`
							ContribType string `xml:"contrib-type,attr"`
							Name        struct {
								Text      string `xml:",chardata"`
								NameStyle string `xml:"name-style,attr"`
								Surname   struct {
									Text string `xml:",chardata"` // Tréfás, Ledl, Tréfás,...
								} `xml:"surname"`
								GivenNames struct {
									Text string `xml:",chardata"` // David, Dr. Andreas, Dr. D...
								} `xml:"given-names"`
							} `xml:"name"`
							Aff struct {
								Text string `xml:",chardata"` // Universitätsbibliothek B...
							} `xml:"aff"`
							Email struct {
								Text string `xml:",chardata"` // david.trefas@unibas.ch, g...
							} `xml:"email"`
							Uri struct {
								Text string `xml:",chardata"` // https://logos.vision/logo...
							} `xml:"uri"`
						} `xml:"contrib"`
					} `xml:"contrib-group"`
					PubDate []struct {
						Text    string `xml:",chardata"`
						PubType string `xml:"pub-type,attr"`
						Day     struct {
							Text string `xml:",chardata"` // 12, 12, 12, 12, 24, 24, 2...
						} `xml:"day"`
						Month struct {
							Text string `xml:",chardata"` // 03, 03, 03, 03, 09, 09, 0...
						} `xml:"month"`
						Year struct {
							Text string `xml:",chardata"` // 2013, 2013, 2013, 2013, 2...
						} `xml:"year"`
					} `xml:"pub-date"`
					Volume struct {
						Text string `xml:",chardata"` // 1, 1, 1, 1, 1, 1, 1, 1, 1...
					} `xml:"volume"`
					Issue struct {
						Text string `xml:",chardata"` // 1, 1, 1, 1, 2, 2, 2, 1, 2...
						Seq  string `xml:"seq,attr"`
					} `xml:"issue"`
					IssueID struct {
						Text      string `xml:",chardata"` // 13, 13, 13, 13, 14, 14, 1...
						PubIDType string `xml:"pub-id-type,attr"`
					} `xml:"issue-id"`
					IssueTitle struct {
						Text string `xml:",chardata"` // Bibliothek 2.0 am Ende?! ...
					} `xml:"issue-title"`
					Permissions struct {
						Text               string `xml:",chardata"`
						CopyrightStatement struct {
							Text string `xml:",chardata"` // Copyright (c), Copyright ...
						} `xml:"copyright-statement"`
						CopyrightYear struct {
							Text string `xml:",chardata"` // 2016, 2016, 2016, 2016, 2...
						} `xml:"copyright-year"`
						License struct {
							Text string `xml:",chardata"`
							Href string `xml:"href,attr"`
						} `xml:"license"`
					} `xml:"permissions"`
					SelfUri []struct {
						Text        string `xml:",chardata"`
						Href        string `xml:"href,attr"`
						ContentType string `xml:"content-type,attr"`
					} `xml:"self-uri"`
					Abstract struct {
						Text string `xml:",chardata"`
						Lang string `xml:"lang,attr"`
						P    struct {
							Text string `xml:",chardata"` // DOI: 10.12685/027.7-1-1-1...
						} `xml:"p"`
					} `xml:"abstract"`
					AbstractTrans struct {
						Text string `xml:",chardata"`
						Lang string `xml:"lang,attr"`
						P    struct {
							Text string `xml:",chardata"` // DOI: 10.12685/027.7-1-1-1...
						} `xml:"p"`
					} `xml:"abstract-trans"`
				} `xml:"article-meta"`
			} `xml:"front"`
			Body struct {
				Text string `xml:",chardata"`
				P    struct {
					Text string `xml:",chardata"` // Tréfás       027.7 Zeit...
				} `xml:"p"`
			} `xml:"body"`
		} `xml:"article"`
	} `xml:"metadata"`
	About struct {
		Text string `xml:",chardata"`
	} `xml:"about"`
}
