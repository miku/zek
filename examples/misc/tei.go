package misc

import "encoding/xml"

// TEI was generated 2018-07-09 16:06:52 by tir on sol.
type TEI struct {
	XMLName   xml.Name `xml:"TEI"`
	Chardata  string   `xml:",chardata"`
	Xmlns     string   `xml:"xmlns,attr"`
	TeiHeader struct {
		Text     string `xml:",chardata"`
		FileDesc struct {
			Text      string `xml:",chardata"`
			TitleStmt struct {
				Text  string `xml:",chardata"`
				Title struct {
					Text string `xml:",chardata"` // Manuskriptlist
				} `xml:"title"`
				Author struct {
					Text     string `xml:",chardata"`
					PersName struct {
						Text string `xml:",chardata"` // Stefan Zweig
					} `xml:"persName"`
				} `xml:"author"`
			} `xml:"titleStmt"`
			PublicationStmt struct {
				Text      string `xml:",chardata"`
				Publisher []struct {
					Text    string `xml:",chardata"` // Literaturarchiv Salzburg
					OrgName struct {
						Text    string `xml:",chardata"` // Literaturarchiv Salzburg
						Corresp string `xml:"corresp,attr"`
						Ref     string `xml:"ref,attr"`
					} `xml:"orgName"`
				} `xml:"publisher"`
				Authority struct {
					Text    string `xml:",chardata"`
					OrgName struct {
						Text    string `xml:",chardata"` // Zentrum für Informations...
						Corresp string `xml:"corresp,attr"`
						Ref     string `xml:"ref,attr"`
					} `xml:"orgName"`
				} `xml:"authority"`
				Distributor struct {
					Text    string `xml:",chardata"`
					OrgName struct {
						Text string `xml:",chardata"` // GAMS - Geisteswissenschaf...
						Ref  string `xml:"ref,attr"`
					} `xml:"orgName"`
				} `xml:"distributor"`
				Availability struct {
					Text    string `xml:",chardata"`
					Licence struct {
						Text   string `xml:",chardata"` // Creative Commons BY-NC-SA...
						Target string `xml:"target,attr"`
					} `xml:"licence"`
				} `xml:"availability"`
				Idno struct {
					Text string `xml:",chardata"` // o:szd.werke
					Type string `xml:"type,attr"`
				} `xml:"idno"`
				Date struct {
					Text string `xml:",chardata"` // 15.12.2017
					When string `xml:"when,attr"`
				} `xml:"date"`
			} `xml:"publicationStmt"`
			SeriesStmt struct {
				Text  string `xml:",chardata"`
				Title struct {
					Text string `xml:",chardata"` // Stefan Zweig Digitale Nac...
					Ref  string `xml:"ref,attr"`
				} `xml:"title"`
				RespStmt []struct {
					Text string `xml:",chardata"`
					Resp struct {
						Text string `xml:",chardata"` // Projektleitung, Datenerfa...
					} `xml:"resp"`
					PersName struct {
						Text     string `xml:",chardata"`
						Forename struct {
							Text string `xml:",chardata"` // Manfred, Lina, Christophe...
						} `xml:"forename"`
						Surname struct {
							Text string `xml:",chardata"` // Mittermayer, Zangerl, Pol...
						} `xml:"surname"`
					} `xml:"persName"`
				} `xml:"respStmt"`
			} `xml:"seriesStmt"`
			SourceDesc struct {
				Text string `xml:",chardata"`
				P    struct {
					Text string `xml:",chardata"` // ToDo: Manuskripte von Ste...
				} `xml:"p"`
			} `xml:"sourceDesc"`
		} `xml:"fileDesc"`
		EncodingDesc struct {
			Text          string `xml:",chardata"`
			EditorialDecl struct {
				Text string `xml:",chardata"`
				P    struct {
					Text string `xml:",chardata"` // ToDo: was über die editi...
				} `xml:"p"`
			} `xml:"editorialDecl"`
			ProjectDesc struct {
				Text string `xml:",chardata"`
				Ab   struct {
					Text string `xml:",chardata"`
					Ref  struct {
						Text   string `xml:",chardata"`
						Target string `xml:"target,attr"`
						Type   string `xml:"type,attr"`
					} `xml:"ref"`
				} `xml:"ab"`
				P struct {
					Text string `xml:",chardata"` // Das Projekt verfolgt das ...
				} `xml:"p"`
			} `xml:"projectDesc"`
		} `xml:"encodingDesc"`
	} `xml:"teiHeader"`
	Text struct {
		Text string `xml:",chardata"`
		Body struct {
			Text     string `xml:",chardata"`
			ListBibl struct {
				Text     string `xml:",chardata"`
				BiblFull []struct {
					Text     string `xml:",chardata"`
					ID       string `xml:"id,attr"`
					Ana      string `xml:"ana,attr"`
					FileDesc struct {
						Text      string `xml:",chardata"`
						TitleStmt struct {
							Text  string `xml:",chardata"`
							Title []struct {
								Text string `xml:",chardata"` // Clarissa, Clarissa, Stefa...
								Ana  string `xml:"ana,attr"`
								Type string `xml:"type,attr"`
								Ref  string `xml:"ref,attr"`
							} `xml:"title"`
							Author struct {
								Text string `xml:",chardata"` // Zweig, Stefan, Zweig, Ste...
								Ref  string `xml:"ref,attr"`
								Role string `xml:"role,attr"`
							} `xml:"author"`
							Editor struct {
								Text string `xml:",chardata"` // Zweig, Lotte, Zweig, Lott...
								Role string `xml:"role,attr"`
							} `xml:"editor"`
						} `xml:"titleStmt"`
						PublicationStmt struct {
							Text string `xml:",chardata"`
							Ab   struct {
								Text string `xml:",chardata"` // Archivmaterial, Archivmat...
							} `xml:"ab"`
						} `xml:"publicationStmt"`
						SourceDesc struct {
							Text   string `xml:",chardata"`
							MsDesc struct {
								Text         string `xml:",chardata"`
								MsIdentifier struct {
									Text    string `xml:",chardata"`
									Country struct {
										Text string `xml:",chardata"` // Österreich, Österreich,...
									} `xml:"country"`
									Settlement struct {
										Text string `xml:",chardata"` // Salzburg, Salzburg, Salzb...
									} `xml:"settlement"`
									Repository struct {
										Text string `xml:",chardata"` // Literaturarchiv Salzburg,...
										Ref  string `xml:"ref,attr"`
									} `xml:"repository"`
									Idno struct {
										Text string `xml:",chardata"` // SZ-AAP/W1, SZ-AAP/W2.1, S...
										Type string `xml:"type,attr"`
									} `xml:"idno"`
									AltIdentifier []struct {
										Text string `xml:",chardata"`
										Idno struct {
											Text string `xml:",chardata"` // o:szd.6814, o:szd.6819, o...
											Type string `xml:"type,attr"`
										} `xml:"idno"`
									} `xml:"altIdentifier"`
								} `xml:"msIdentifier"`
								MsContents struct {
									Text     string `xml:",chardata"`
									TextLang struct {
										Text string `xml:",chardata"`
										Lang []struct {
											Text string `xml:",chardata"` // ger, ger, ger, ger, ger, ...
										} `xml:"lang"`
									} `xml:"textLang"`
									MsItem struct {
										Text    string `xml:",chardata"`
										Incipit struct {
											Text string `xml:",chardata"` // „Ein Dorfpostamt in Oes...
										} `xml:"incipit"`
										DocEdition struct {
											Text string `xml:",chardata"` // „II Teil noch nicht cop...
											Ana  string `xml:"ana,attr"`
										} `xml:"docEdition"`
									} `xml:"msItem"`
									Summary struct {
										Text string `xml:",chardata"` // Notizen und Entwürfe, No...
									} `xml:"summary"`
								} `xml:"msContents"`
								PhysDesc struct {
									Text       string `xml:",chardata"`
									ObjectDesc struct {
										Text        string `xml:",chardata"`
										SupportDesc struct {
											Text    string `xml:",chardata"`
											Support struct {
												Text     string `xml:",chardata"`
												Material []struct {
													Text string `xml:",chardata"` // Notizbuch mit Händlermar...
													Ana  string `xml:"ana,attr"`
												} `xml:"material"`
											} `xml:"support"`
											Extent struct {
												Text string `xml:",chardata"`
												Term struct {
													Text string `xml:",chardata"` // Notizbuch, Typoskriptdurc...
													Type string `xml:"type,attr"`
												} `xml:"term"`
												Measure []struct {
													Text string `xml:",chardata"` // 154 Blatt, 99 Blatt besch...
													Ana  string `xml:"ana,attr"`
													Type string `xml:"type,attr"`
												} `xml:"measure"`
											} `xml:"extent"`
											Foliation struct {
												Text string `xml:",chardata"`
												Ab   struct {
													Text string `xml:",chardata"` // durchgehend paginiert, du...
												} `xml:"ab"`
											} `xml:"foliation"`
										} `xml:"supportDesc"`
									} `xml:"objectDesc"`
									AccMat struct {
										Text string `xml:",chardata"`
										List struct {
											Text string `xml:",chardata"`
											Item []struct {
												Text string `xml:",chardata"`
												Ana  string `xml:"ana,attr"`
												Desc struct {
													Text string `xml:",chardata"` // , [korrigiert], ,] Beilag...
												} `xml:"desc"`
												Measure struct {
													Text string `xml:",chardata"` // 23x15 cm, 31x19 									...
													Ana  string `xml:"ana,attr"`
												} `xml:"measure"`
											} `xml:"item"`
										} `xml:"list"`
									} `xml:"accMat"`
									HandDesc struct {
										Text string `xml:",chardata"`
										Ab   struct {
											Text string `xml:",chardata"` // Stefan Zweig, Lotte Zweig...
										} `xml:"ab"`
									} `xml:"handDesc"`
									BindingDesc struct {
										Text    string `xml:",chardata"`
										Binding struct {
											Text string `xml:",chardata"`
											Ab   struct {
												Text string `xml:",chardata"` // Lederimitat mit schwarzem...
											} `xml:"ab"`
										} `xml:"binding"`
									} `xml:"bindingDesc"`
								} `xml:"physDesc"`
								History struct {
									Text   string `xml:",chardata"`
									Origin struct {
										Text      string `xml:",chardata"`
										OrigPlace struct {
											Text string `xml:",chardata"` // Petropolis, Lyncombe Hill...
										} `xml:"origPlace"`
										OrigDate struct {
											Text string `xml:",chardata"` // 1. Nov. 1941, Sommer 1939...
											Type string `xml:"type,attr"`
											When string `xml:"when,attr"`
										} `xml:"origDate"`
									} `xml:"origin"`
									Provenance struct {
										Text string `xml:",chardata"` // Archiv Atrium Press, Arch...
									} `xml:"provenance"`
									Acquisition struct {
										Text string `xml:",chardata"` // Ankauf Christie's London ...
									} `xml:"acquisition"`
								} `xml:"history"`
							} `xml:"msDesc"`
						} `xml:"sourceDesc"`
						NotesStmt struct {
							Text string `xml:",chardata"`
							Note struct {
								Text string `xml:",chardata"` // 1 Blatt mit Anmerkungen i...
							} `xml:"note"`
						} `xml:"notesStmt"`
					} `xml:"fileDesc"`
					ProfileDesc struct {
						Text      string `xml:",chardata"`
						TextClass struct {
							Text     string `xml:",chardata"`
							Keywords struct {
								Text string `xml:",chardata"`
								Term []struct {
									Text string `xml:",chardata"` // Romane/Erzählungen, Roma...
									Type string `xml:"type,attr"`
									Ref  string `xml:"ref,attr"`
								} `xml:"term"`
							} `xml:"keywords"`
						} `xml:"textClass"`
					} `xml:"profileDesc"`
				} `xml:"biblFull"`
			} `xml:"listBibl"`
		} `xml:"body"`
	} `xml:"text"`
}
