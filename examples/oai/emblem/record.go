package emblem

import "encoding/xml"

type Record struct {
	XMLName xml.Name `xml:"Record"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text       string `xml:",chardata"`
		Status     string `xml:"status,attr"`
		Identifier struct {
			Text string `xml:",chardata"` // oai:diglib.hab.de:ppn_318...
		} `xml:"identifier"`
		Datestamp struct {
			Text string `xml:",chardata"` // 2017-11-23, 2017-11-23, 2...
		} `xml:"datestamp"`
		SetSpec []struct {
			Text string `xml:",chardata"` // ddb, embl:iconclass, fko,...
		} `xml:"setSpec"`
	} `xml:"header"`
	Metadata struct {
		Text       string `xml:",chardata"`
		BiblioDesc struct {
			Text       string `xml:",chardata"`
			AttrEmblem string `xml:"emblem,attr"`
			AttrMods   string `xml:"mods,attr"`
			Xlink      string `xml:"xlink,attr"`
			Tei        string `xml:"tei,attr"`
			Skos       string `xml:"skos,attr"`
			Mods       struct {
				Text      string `xml:",chardata"`
				TitleInfo []struct {
					Text  string `xml:",chardata"`
					Type  string `xml:"type,attr"`
					Title struct {
						Text string `xml:",chardata"` // Frauenzimmer Gesprechspie...
					} `xml:"title"`
					SubTitle struct {
						Text string `xml:",chardata"` // Von Jhme vor diesem in Sp...
					} `xml:"subTitle"`
					NonSort struct {
						Text string `xml:",chardata"` // ..., Der, Der, Der, Die, ...
					} `xml:"nonSort"`
				} `xml:"titleInfo"`
				Identifier []struct {
					Text string `xml:",chardata"` // d.ch t.t, l-de Keun 3 164...
					Type string `xml:"type,attr"`
				} `xml:"identifier"`
				OriginInfo []struct {
					Text       string `xml:",chardata"`
					DateIssued struct {
						Text     string `xml:",chardata"` // 1644, 1698, 1625, 1674, 1...
						KeyDate  string `xml:"keyDate,attr"`
						Encoding string `xml:"encoding,attr"`
					} `xml:"dateIssued"`
					Place struct {
						Text      string `xml:",chardata"`
						PlaceTerm struct {
							Text string `xml:",chardata"` // Nürnberg, Berlin, Züric...
							Type string `xml:"type,attr"`
						} `xml:"placeTerm"`
					} `xml:"place"`
					Publisher struct {
						Text string `xml:",chardata"` // Endter, HAB Wolfenbüttel...
					} `xml:"publisher"`
					DateCaptured struct {
						Text     string `xml:",chardata"` // 2000, 2001, 2001, 2001, 2...
						Encoding string `xml:"encoding,attr"`
					} `xml:"dateCaptured"`
					Edition struct {
						Text string `xml:",chardata"` // [Electronic ed.], [Electr...
					} `xml:"edition"`
				} `xml:"originInfo"`
				Name []struct {
					Text     string `xml:",chardata"`
					Type     string `xml:"type,attr"`
					ValueURI string `xml:"valueURI,attr"`
					NamePart []struct {
						Text string `xml:",chardata"` // Georg Philipp, Harsdoerff...
						Type string `xml:"type,attr"`
					} `xml:"namePart"`
					DisplayForm struct {
						Text string `xml:",chardata"` // Harsdoerffer, Georg Phili...
					} `xml:"displayForm"`
					Role struct {
						Text     string `xml:",chardata"`
						RoleTerm struct {
							Text      string `xml:",chardata"` // aut, prt, aut, oth, prt, ...
							Authority string `xml:"authority,attr"`
							Type      string `xml:"type,attr"`
							ValueURI  string `xml:"valueURI,attr"`
						} `xml:"roleTerm"`
					} `xml:"role"`
				} `xml:"name"`
				Language []struct {
					Text         string `xml:",chardata"`
					LanguageTerm struct {
						Text      string `xml:",chardata"` // ger, ger, ger, ger, fre, ...
						Type      string `xml:"type,attr"`
						Authority string `xml:"authority,attr"`
					} `xml:"languageTerm"`
				} `xml:"language"`
				RelatedItem []struct {
					Text         string `xml:",chardata"`
					Type         string `xml:"type,attr"`
					DisplayLabel string `xml:"displayLabel,attr"`
					TitleInfo    []struct {
						Text  string `xml:",chardata"`
						Type  string `xml:"type,attr"`
						Title struct {
							Text string `xml:",chardata"` // Festkultur Online, Projek...
						} `xml:"title"`
						SubTitle struct {
							Text string `xml:",chardata"` // aus italiänischen, frant...
						} `xml:"subTitle"`
						NonSort struct {
							Text string `xml:",chardata"` // ...
						} `xml:"nonSort"`
					} `xml:"titleInfo"`
					RecordInfo struct {
						Text             string `xml:",chardata"`
						RecordIdentifier struct {
							Text string `xml:",chardata"` // GBV 09217907X, GBV 344846...
						} `xml:"recordIdentifier"`
					} `xml:"recordInfo"`
				} `xml:"relatedItem"`
				PhysicalDescription struct {
					Text   string `xml:",chardata"`
					Extent []struct {
						Text string `xml:",chardata"` // [22] Bl., 622 S., [41] Bl...
					} `xml:"extent"`
				} `xml:"physicalDescription"`
				Part struct {
					Text   string `xml:",chardata"`
					Order  string `xml:"order,attr"`
					Detail struct {
						Text   string `xml:",chardata"`
						Type   string `xml:"type,attr"`
						Number struct {
							Text string `xml:",chardata"` // Theil 4, Classis 1, Ps. 1...
						} `xml:"number"`
					} `xml:"detail"`
				} `xml:"part"`
				RecordInfo struct {
					Text             string `xml:",chardata"`
					RecordIdentifier struct {
						Text   string `xml:",chardata"` // 318435721, 331774453, 331...
						Source string `xml:"source,attr"`
					} `xml:"recordIdentifier"`
					RecordOrigin struct {
						Text string `xml:",chardata"` // Converted from PICA using...
						Lang string `xml:"lang,attr"`
					} `xml:"recordOrigin"`
					RecordContentSource struct {
						Text      string `xml:",chardata"` // DE-23, DE-23, DE-23, DE-2...
						Authority string `xml:"authority,attr"`
					} `xml:"recordContentSource"`
				} `xml:"recordInfo"`
				Classification []struct {
					Text         string `xml:",chardata"` // 08.24, 17.87, 11.63, 20.2...
					Authority    string `xml:"authority,attr"`
					ValueURI     string `xml:"valueURI,attr"`
					DisplayLabel string `xml:"displayLabel,attr"`
				} `xml:"classification"`
				Genre []struct {
					Text      string `xml:",chardata"` // Emblembuch, Emblembuch, E...
					Authority string `xml:"authority,attr"`
				} `xml:"genre"`
				Subject struct {
					Text      string `xml:",chardata"`
					Authority string `xml:"authority,attr"`
					Topic     struct {
						Text     string `xml:",chardata"` // Fruchtbringende Gesellsch...
						ValueURI string `xml:"valueURI,attr"`
					} `xml:"topic"`
					Geographic []struct {
						Text     string `xml:",chardata"` // Augsburg, Bayern, Deutsch...
						ValueURI string `xml:"valueURI,attr"`
					} `xml:"geographic"`
				} `xml:"subject"`
			} `xml:"mods"`
			CopyDesc struct {
				Text   string `xml:",chardata"`
				CopyID struct {
					Text string `xml:",chardata"` // http://diglib.hab.de/druc...
				} `xml:"copyID"`
				Owner struct {
					Text            string `xml:",chardata"` // Herzog August Bibliothek ...
					CountryCode     string `xml:"countryCode,attr"`
					InstitutionCode string `xml:"institutionCode,attr"`
				} `xml:"owner"`
			} `xml:"copyDesc"`
			Emblem []struct {
				Text     string `xml:",chardata"`
				GlobalID string `xml:"globalID,attr"`
				Motto    struct {
					Text          string `xml:",chardata"`
					Transcription []struct {
						Text string `xml:",chardata"`
						Href string `xml:"href,attr"`
						Lang string `xml:"lang,attr"`
						Page string `xml:"page,attr"`
						P    struct {
							Text  string `xml:",chardata"` // Mit Freuden Befreyet., Ei...
							Xmlns string `xml:"xmlns,attr"`
							Lang  string `xml:"lang,attr"`
							Lg    struct {
								Text string `xml:",chardata"`
								Lang string `xml:"lang,attr"`
								L    []struct {
									Text string `xml:",chardata"` // Gott sey gelobet innigKli...
								} `xml:"l"`
								Head struct {
									Text string `xml:",chardata"` // Vom Seidenwurm.
								} `xml:"head"`
							} `xml:"lg"`
							L struct {
								Text string `xml:",chardata"` // Erkandtnuß deß wahren D...
							} `xml:"l"`
						} `xml:"p"`
					} `xml:"transcription"`
				} `xml:"motto"`
				Pictura []struct {
					Text      string `xml:",chardata"`
					Medium    string `xml:"medium,attr"`
					Href      string `xml:"href,attr"`
					Iconclass []struct {
						Text     string `xml:",chardata"`
						Notation struct {
							Text string `xml:",chardata"` // 86(MIT FREUDEN BEFREYET),...
						} `xml:"notation"`
						PrefLabel struct {
							Text string `xml:",chardata"` // Sprichwörter, Redewendun...
							Lang string `xml:"lang,attr"`
						} `xml:"prefLabel"`
						Keyword []struct {
							Text string `xml:",chardata"` // Sprichwort, Redewendung, ...
							Lang string `xml:"lang,attr"`
						} `xml:"keyword"`
					} `xml:"iconclass"`
				} `xml:"pictura"`
				Subscriptio struct {
					Text          string `xml:",chardata"`
					Transcription []struct {
						Text string `xml:",chardata"`
						Href string `xml:"href,attr"`
						Lang string `xml:"lang,attr"`
						Page string `xml:"page,attr"`
						P    struct {
							Text  string `xml:",chardata"` // Nichts desto weniger wird...
							Xmlns string `xml:"xmlns,attr"`
							Lang  string `xml:"lang,attr"`
							Lg    []struct {
								Text string `xml:",chardata"`
								Lang string `xml:"lang,attr"`
								L    []struct {
									Text string `xml:",chardata"` // Wie alle Flüß vom Meer ...
								} `xml:"l"`
								Head struct {
									Text string `xml:",chardata"` // Gedenk Sprüchlein
								} `xml:"head"`
								Lg struct {
									Text string `xml:",chardata"`
									Head struct {
										Text string `xml:",chardata"` // Hertzlicher Seüfftzer.
									} `xml:"head"`
									L []struct {
										Text string `xml:",chardata"` // Lass dich durch mein Gebe...
									} `xml:"l"`
								} `xml:"lg"`
							} `xml:"lg"`
							L []struct {
								Text string `xml:",chardata"` // NB, Gott ist im Wesen Ein...
							} `xml:"l"`
						} `xml:"p"`
					} `xml:"transcription"`
				} `xml:"subscriptio"`
			} `xml:"emblem"`
		} `xml:"biblioDesc"`
	} `xml:"metadata"`
	About struct {
		Text string `xml:",chardata"`
	} `xml:"about"`
}
