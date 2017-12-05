package misc

import "encoding/xml"

// Article was generated 2017-12-05 23:38:40 by tir on apollo.
type Article struct {
	XMLName    xml.Name `xml:"article"`
	Text       string   `xml:",chardata"`
	Xlink      string   `xml:"xlink,attr"`
	Mml        string   `xml:"mml,attr"`
	Oasis      string   `xml:"oasis,attr"`
	DtdVersion string   `xml:"dtd-version,attr"`
	Front      struct {
		Text        string `xml:",chardata"`
		JournalMeta struct {
			Text      string `xml:",chardata"`
			JournalID struct {
				Text          string `xml:",chardata"` // GMD
				JournalIDType string `xml:"journal-id-type,attr"`
			} `xml:"journal-id"`
			JournalTitleGroup struct {
				Text         string `xml:",chardata"`
				JournalTitle struct {
					Text string `xml:",chardata"` // Geoscientific Model Development
				} `xml:"journal-title"`
				AbbrevJournalTitle []struct {
					Text       string `xml:",chardata"` // GMD, Geosci. Model Dev.
					AbbrevType string `xml:"abbrev-type,attr"`
				} `xml:"abbrev-journal-title"`
			} `xml:"journal-title-group"`
			ISSN struct {
				Text    string `xml:",chardata"` // 1991-9603
				PubType string `xml:"pub-type,attr"`
			} `xml:"issn"`
			Publisher struct {
				Text          string `xml:",chardata"`
				PublisherName struct {
					Text string `xml:",chardata"` // Copernicus Publications
				} `xml:"publisher-name"`
				PublisherLoc struct {
					Text string `xml:",chardata"` // Göttingen, Germany
				} `xml:"publisher-loc"`
			} `xml:"publisher"`
		} `xml:"journal-meta"`
		ArticleMeta struct {
			Text      string `xml:",chardata"`
			ArticleID struct {
				Text      string `xml:",chardata"` // 10.5194/gmd-10-4393-2017
				PubIDType string `xml:"pub-id-type,attr"`
			} `xml:"article-id"`
			TitleGroup struct {
				Text         string `xml:",chardata"`
				ArticleTitle struct {
					Text string `xml:",chardata"` // A JavaScript API for the Ice Sheet System Model (ISSM) 4.11: towards an online interactive model for...
				} `xml:"article-title"`
			} `xml:"title-group"`
			ContribGroup struct {
				Text    string `xml:",chardata"`
				Contrib []struct {
					Text        string `xml:",chardata"`
					ContribType string `xml:"contrib-type,attr"`
					Corresp     string `xml:"corresp,attr"`
					Rid         string `xml:"rid,attr"`
					Name        struct {
						Text    string `xml:",chardata"`
						Surname struct {
							Text string `xml:",chardata"` // Larour, Cheng, Perez, Quinn, Morlighem, Duong, Nguyen, Petrie, Harounian, Halkides
						} `xml:"surname"`
						GivenNames struct {
							Text string `xml:",chardata"` // Eric, Daniel, Gilberto, Justin, Mathieu, Bao, Lan, Kit, Silva, Daria
						} `xml:"given-names"`
					} `xml:"name"`
					Email struct {
						Text string `xml:",chardata"` // eric.larour@jpl.nasa.gov
					} `xml:"email"`
				} `xml:"contrib"`
				Aff []struct {
					Text  string `xml:",chardata"`
					ID    string `xml:"id,attr"`
					Label struct {
						Text string `xml:",chardata"` // 1, 2, 3, 4, 5, 6, 7
					} `xml:"label"`
					Institution struct {
						Text string `xml:",chardata"` // Jet Propulsion Laboratory – California Institute of Technology, 4800 Oak Grove Drive 300-323,, Pas...
					} `xml:"institution"`
				} `xml:"aff"`
			} `xml:"contrib-group"`
			AuthorNotes struct {
				Text    string `xml:",chardata"`
				Corresp struct {
					Text string `xml:",chardata"` // Eric Larour (eric.larour@jpl.nasa.gov)
					ID   string `xml:"id,attr"`
				} `xml:"corresp"`
			} `xml:"author-notes"`
			PubDate struct {
				Text string `xml:",chardata"`
				Day  struct {
					Text string `xml:",chardata"` // 4
				} `xml:"day"`
				Month struct {
					Text string `xml:",chardata"` // December
				} `xml:"month"`
				Year struct {
					Text string `xml:",chardata"` // 2017
				} `xml:"year"`
			} `xml:"pub-date"`
			Volume struct {
				Text string `xml:",chardata"` // 10
			} `xml:"volume"`
			Issue struct {
				Text string `xml:",chardata"` // 12
			} `xml:"issue"`
			Fpage struct {
				Text string `xml:",chardata"` // 4393
			} `xml:"fpage"`
			Lpage struct {
				Text string `xml:",chardata"` // 4403
			} `xml:"lpage"`
			History struct {
				Text string `xml:",chardata"`
				Date []struct {
					Text     string `xml:",chardata"`
					DateType string `xml:"date-type,attr"`
					Day      struct {
						Text string `xml:",chardata"` // 9, 25, 1, 4
					} `xml:"day"`
					Month struct {
						Text string `xml:",chardata"` // July, August, February, February
					} `xml:"month"`
					Year struct {
						Text string `xml:",chardata"` // 2016, 2016, 2017, 2017
					} `xml:"year"`
				} `xml:"date"`
			} `xml:"history"`
			Permissions struct {
				Text    string `xml:",chardata"`
				License struct {
					Text        string `xml:",chardata"`
					LicenseType string `xml:"license-type,attr"`
					LicenseP    struct {
						Text    string `xml:",chardata"` // This work is licensed under the Creative Commons Attribution 3.0 Unported License. To view a copy of...
						ExtLink struct {
							Text        string `xml:",chardata"` // https://creativecommons.org/licenses/by/3.0/
							ExtLinkType string `xml:"ext-link-type,attr"`
							Href        string `xml:"href,attr"`
						} `xml:"ext-link"`
					} `xml:"license-p"`
				} `xml:"license"`
			} `xml:"permissions"`
			SelfURI []struct {
				Text string `xml:",chardata"` // This article is available from https://www.geosci-model-dev.net/10/4393/2017/gmd-10-4393-2017.html, ...
				Href string `xml:"href,attr"`
			} `xml:"self-uri"`
			Abstract struct {
				Text string `xml:",chardata"`
				P    struct {
					Text string `xml:",chardata"` // Earth system models (ESMs) are becoming increasingly complex, requiring extensive knowledge and expe...
					ID   string `xml:"id,attr"`
					URI  struct {
						Text string `xml:",chardata"` // http://vesl.jpl.nasa.gov
					} `xml:"uri"`
					Xref struct {
						Text    string `xml:",chardata"`
						RefType string `xml:"ref-type,attr"`
						Rid     string `xml:"rid,attr"`
						ID      string `xml:"id,attr"`
					} `xml:"xref"`
				} `xml:"p"`
			} `xml:"abstract"`
		} `xml:"article-meta"`
	} `xml:"front"`
	Body struct {
		Text string `xml:",chardata"`
		Sec  []struct {
			Text    string `xml:",chardata"`
			ID      string `xml:"id,attr"`
			SecType string `xml:"sec-type,attr"`
			Title   struct {
				Text string `xml:",chardata"` // Introduction, ISSM JavaScript API, HTTP/Python server, All-in-one design/simulations, Examples, Conc...
			} `xml:"title"`
			P []struct {
				Text string `xml:",chardata"` // Earth system models (ESMs) across the Earth science community have become increasingly sophisticated...
				ID   string `xml:"id,attr"`
				Xref []struct {
					Text         string `xml:",chardata"`
					RefType      string `xml:"ref-type,attr"`
					Rid          string `xml:"rid,attr"`
					ID           string `xml:"id,attr"`
					NamedContent struct {
						Text        string `xml:",chardata"` // CMIP-5;, CMIP-6;, e.g.,, ISSM;, ISMIP-6;, GRANTISM;
						ContentType string `xml:"content-type,attr"`
					} `xml:"named-content"`
				} `xml:"xref"`
				List struct {
					Text     string `xml:",chardata"`
					ListType string `xml:"list-type,attr"`
					ListItem []struct {
						Text string `xml:",chardata"`
						P    struct {
							Text string `xml:",chardata"` // Bridging the gap between ESM formulations of the physical cores and web technologies such as Hyperte...
							ID   string `xml:"id,attr"`
							Xref []struct {
								Text         string `xml:",chardata"`
								RefType      string `xml:"ref-type,attr"`
								Rid          string `xml:"rid,attr"`
								ID           string `xml:"id,attr"`
								NamedContent struct {
									Text        string `xml:",chardata"` // HTML;
									ContentType string `xml:"content-type,attr"`
								} `xml:"named-content"`
							} `xml:"xref"`
						} `xml:"p"`
					} `xml:"list-item"`
				} `xml:"list"`
				Preformat []struct {
					Text          string `xml:",chardata"` // model, mesh2d, x, y, numberofvertices, lat, long, elements, numberofelements, marshall
					PreformatType string `xml:"preformat-type,attr"`
				} `xml:"preformat"`
				URI struct {
					Text string `xml:",chardata"` // https://ross.ics.uci.edu:8080/
				} `xml:"uri"`
			} `xml:"p"`
			Fig []struct {
				Text        string `xml:",chardata"`
				ID          string `xml:"id,attr"`
				SpecificUse string `xml:"specific-use,attr"`
				Caption     struct {
					Text string `xml:",chardata"`
					P    struct {
						Text string `xml:",chardata"` // Line-by-line comparison of the code behind the mesh2d class, within the MATLAB ISSM API (upper frame...
						ID   string `xml:"id,attr"`
						Xref []struct {
							Text    string `xml:",chardata"`
							RefType string `xml:"ref-type,attr"`
							Rid     string `xml:"rid,attr"`
							ID      string `xml:"id,attr"`
						} `xml:"xref"`
						URI struct {
							Text string `xml:",chardata"` // https://vesl.jpl.nasa.gov, https://vesl.jpl.nasa.gov
						} `xml:"uri"`
						InlineFormula []struct {
							Text string `xml:",chardata"`
							Math struct {
								Text    string `xml:",chardata"`
								ID      string `xml:"id,attr"`
								Display string `xml:"display,attr"`
								Mrow    struct {
									Text string `xml:",chardata"`
									Mo   struct {
										Text string `xml:",chardata"` // -, +
									} `xml:"mo"`
									Mn struct {
										Text        string `xml:",chardata"` // 5, 5
										Mathvariant string `xml:"mathvariant,attr"`
									} `xml:"mn"`
								} `xml:"mrow"`
								Msup struct {
									Text string `xml:",chardata"`
									Mi   struct {
										Text string `xml:",chardata"`
									} `xml:"mi"`
									Mrow struct {
										Text string `xml:",chardata"`
										Mo   struct {
											Text string `xml:",chardata"` // -
										} `xml:"mo"`
										Mn struct {
											Text        string `xml:",chardata"` // 1
											Mathvariant string `xml:"mathvariant,attr"`
										} `xml:"mn"`
									} `xml:"mrow"`
								} `xml:"msup"`
							} `xml:"math"`
						} `xml:"inline-formula"`
					} `xml:"p"`
				} `xml:"caption"`
				Graphic struct {
					Text string `xml:",chardata"`
					Href string `xml:"href,attr"`
				} `xml:"graphic"`
			} `xml:"fig"`
		} `xml:"sec"`
	} `xml:"body"`
	Back struct {
		Text  string `xml:",chardata"`
		Notes []struct {
			Text      string `xml:",chardata"`
			NotesType string `xml:"notes-type,attr"`
			P         struct {
				Text string `xml:",chardata"` // The ISSM code and its JavaScript components are available at, . The instructions for the compilation...
				ID   string `xml:"id,attr"`
				URI  struct {
					Text string `xml:",chardata"` // http://issm.jpl.nasa.gov
				} `xml:"uri"`
			} `xml:"p"`
		} `xml:"notes"`
		AppGroup []struct {
			Text                  string `xml:",chardata"`
			ContentType           string `xml:"content-type,attr"`
			SupplementaryMaterial struct {
				Text     string `xml:",chardata"`
				Position string `xml:"position,attr"`
				P        struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Bold struct {
						Text                        string `xml:",chardata"` // The Supplement related to this article is available online at, .
						InlineSupplementaryMaterial struct {
							Text  string `xml:",chardata"` // https://doi.org/10.5194/gmd-10-4393-2017-supplement
							Href  string `xml:"href,attr"`
							Title string `xml:"title,attr"`
						} `xml:"inline-supplementary-material"`
					} `xml:"bold"`
				} `xml:"p"`
			} `xml:"supplementary-material"`
			App struct {
				Text  string `xml:",chardata"`
				Title struct {
					Text string `xml:",chardata"`
				} `xml:"title"`
			} `xml:"app"`
		} `xml:"app-group"`
		Ack struct {
			Text  string `xml:",chardata"`
			Title struct {
				Text string `xml:",chardata"` // Acknowledgements
			} `xml:"title"`
			P struct {
				Text string `xml:",chardata"` // This work was performed at the Jet Propulsion Laboratory (JPL), California Institute of Technology, ...
				ID   string `xml:"id,attr"`
			} `xml:"p"`
		} `xml:"ack"`
		RefList struct {
			Text  string `xml:",chardata"`
			Title struct {
				Text string `xml:",chardata"` // References
			} `xml:"title"`
			Ref []struct {
				Text  string `xml:",chardata"`
				ID    string `xml:"id,attr"`
				Label struct {
					Text string `xml:",chardata"` // Adhikari and Marshall(2011), Adhikari et al.(2016), Amazon Web Services, Inc.(2008), Bindschadler e...
				} `xml:"label"`
				MixedCitation struct {
					Text    string `xml:",chardata"` // Adhikari, S. and Marshall, S. J.: Improvements to shear-deformational models of glacier dynamics th...
					ExtLink struct {
						Text        string `xml:",chardata"` // 10.5194/gmd-9-1087-2016, 10.3189/2013JoG12J125, 10.5194/gmd-9-1937-2016, 10.1016/j.cageo.2016.08.007...
						Href        string `xml:"href,attr"`
						ExtLinkType string `xml:"ext-link-type,attr"`
					} `xml:"ext-link"`
					URI struct {
						Text string `xml:",chardata"` // http://aws.amazon.com/ec2/#pricing, http://www.ecma-international.org/publications/files/ECMA-ST/Ecm...
					} `xml:"uri"`
					Monospace struct {
						Text string `xml:",chardata"` // mpich
					} `xml:"monospace"`
				} `xml:"mixed-citation"`
			} `xml:"ref"`
		} `xml:"ref-list"`
	} `xml:"back"`
}
