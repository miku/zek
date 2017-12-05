package misc

import "encoding/xml"

// Legislation was generated 2017-12-05 23:33:07 by tir on apollo.
type Legislation struct {
	XMLName            xml.Name `xml:"Legislation"`
	Text               string   `xml:",chardata"`
	Xmlns              string   `xml:"xmlns,attr"`
	Xsi                string   `xml:"xsi,attr"`
	DocumentURI        string   `xml:"DocumentURI,attr"`
	IdURI              string   `xml:"IdURI,attr"`
	NumberOfProvisions string   `xml:"NumberOfProvisions,attr"`
	SchemaLocation     string   `xml:"schemaLocation,attr"`
	SchemaVersion      string   `xml:"SchemaVersion,attr"`
	Base               string   `xml:"base,attr"`
	Metadata           struct {
		Text       string `xml:",chardata"`
		Dc         string `xml:"dc,attr"`
		Dct        string `xml:"dct,attr"`
		Atom       string `xml:"atom,attr"`
		Ukm        string `xml:"ukm,attr"`
		Identifier struct {
			Text string `xml:",chardata"` // http://www.legislation.gov.uk/nidsr/2004/0337954364
		} `xml:"identifier"`
		Title struct {
			Text string `xml:",chardata"` // Lay Magistrates (Eligibility) (Northern Ireland) Order 2004
		} `xml:"title"`
		Subject []struct {
			Text   string `xml:",chardata"` // JUSTICE, JUSTICE
			Scheme string `xml:"scheme,attr"`
		} `xml:"subject"`
		Publisher struct {
			Text string `xml:",chardata"` // Government Printer for Northern Ireland
		} `xml:"publisher"`
		Modified struct {
			Text string `xml:",chardata"` // 2016-11-21
		} `xml:"modified"`
		Valid struct {
			Text string `xml:",chardata"` // 2004-01-01
		} `xml:"valid"`
		Description struct {
			Text string `xml:",chardata"` // This Order provides that, unless the Lord Chancellor otherwise determines in the case of a particula...
		} `xml:"description"`
		Link []struct {
			Text     string `xml:",chardata"`
			Rel      string `xml:"rel,attr"`
			Type     string `xml:"type,attr"`
			Href     string `xml:"href,attr"`
			Title    string `xml:"title,attr"`
			Hreflang string `xml:"hreflang,attr"`
		} `xml:"link"`
		SecondaryMetadata struct {
			Text                   string `xml:",chardata"`
			DocumentClassification struct {
				Text             string `xml:",chardata"`
				DocumentCategory struct {
					Text  string `xml:",chardata"`
					Value string `xml:"Value,attr"`
				} `xml:"DocumentCategory"`
				DocumentMainType struct {
					Text  string `xml:",chardata"`
					Value string `xml:"Value,attr"`
				} `xml:"DocumentMainType"`
				DocumentStatus struct {
					Text  string `xml:",chardata"`
					Value string `xml:"Value,attr"`
				} `xml:"DocumentStatus"`
				DocumentMinorType struct {
					Text  string `xml:",chardata"`
					Value string `xml:"Value,attr"`
				} `xml:"DocumentMinorType"`
			} `xml:"DocumentClassification"`
			Year struct {
				Text  string `xml:",chardata"`
				Value string `xml:"Value,attr"`
			} `xml:"Year"`
			ISBN struct {
				Text  string `xml:",chardata"`
				Value string `xml:"Value,attr"`
			} `xml:"ISBN"`
		} `xml:"SecondaryMetadata"`
		SupersededBy struct {
			Text  string `xml:",chardata"`
			URI   string `xml:"URI,attr"`
			Title struct {
				Text string `xml:",chardata"` // Lay Magistrates (Eligibility) (Northern Ireland) Order 2004
			} `xml:"title"`
			DocumentMainType struct {
				Text  string `xml:",chardata"`
				Value string `xml:"Value,attr"`
			} `xml:"DocumentMainType"`
			Year struct {
				Text  string `xml:",chardata"`
				Value string `xml:"Value,attr"`
			} `xml:"Year"`
			Number struct {
				Text  string `xml:",chardata"`
				Value string `xml:"Value,attr"`
			} `xml:"Number"`
			ISBN struct {
				Text  string `xml:",chardata"`
				Value string `xml:"Value,attr"`
			} `xml:"ISBN"`
		} `xml:"SupersededBy"`
		Statistics struct {
			Text            string `xml:",chardata"`
			TotalParagraphs struct {
				Text  string `xml:",chardata"`
				Value string `xml:"Value,attr"`
			} `xml:"TotalParagraphs"`
			BodyParagraphs struct {
				Text  string `xml:",chardata"`
				Value string `xml:"Value,attr"`
			} `xml:"BodyParagraphs"`
			ScheduleParagraphs struct {
				Text  string `xml:",chardata"`
				Value string `xml:"Value,attr"`
			} `xml:"ScheduleParagraphs"`
			TotalImages struct {
				Text  string `xml:",chardata"`
				Value string `xml:"Value,attr"`
			} `xml:"TotalImages"`
		} `xml:"Statistics"`
	} `xml:"Metadata"`
	Secondary struct {
		Text             string `xml:",chardata"`
		SecondaryPrelims struct {
			Text        string `xml:",chardata"`
			DocumentURI string `xml:"DocumentURI,attr"`
			IdURI       string `xml:"IdURI,attr"`
			Draft       struct {
				Text string `xml:",chardata"`
				Para struct {
					Chardata string `xml:",chardata"`
					Text     struct {
						Text     string `xml:",chardata"` // Draft Order laid before Parliament under section 90(4) of the, , for approval by resolution of each ...
						Citation struct {
							Text   string `xml:",chardata"` // Justice (Northern Ireland) Act 2002
							URI    string `xml:"URI,attr"`
							ID     string `xml:"id,attr"`
							Class  string `xml:"Class,attr"`
							Year   string `xml:"Year,attr"`
							Number string `xml:"Number,attr"`
						} `xml:"Citation"`
					} `xml:"Text"`
				} `xml:"Para"`
			} `xml:"Draft"`
			Number struct {
				Text string `xml:",chardata"` // 2004 No.
			} `xml:"Number"`
			SubjectInformation struct {
				Text    string `xml:",chardata"`
				Subject struct {
					Text  string `xml:",chardata"`
					Title struct {
						Text string `xml:",chardata"` // JUSTICE
					} `xml:"Title"`
				} `xml:"Subject"`
			} `xml:"SubjectInformation"`
			Title struct {
				Text string `xml:",chardata"` // Lay Magistrates (Eligibility) (Northern Ireland) Order 2004
			} `xml:"Title"`
			MadeDate struct {
				Chardata string `xml:",chardata"`
				Text     struct {
					Text string `xml:",chardata"` // Made
				} `xml:"Text"`
				DateText struct {
					Text string `xml:",chardata"` // 2004
				} `xml:"DateText"`
			} `xml:"MadeDate"`
			ComingIntoForce struct {
				Chardata string `xml:",chardata"`
				Text     struct {
					Text string `xml:",chardata"` // Coming into operation in accordance with Article 1(1) of this Order
				} `xml:"Text"`
			} `xml:"ComingIntoForce"`
			SecondaryPreamble struct {
				Text         string `xml:",chardata"`
				EnactingText struct {
					Text string `xml:",chardata"`
					Para struct {
						Chardata string `xml:",chardata"`
						Text     struct {
							Text     string `xml:",chardata"` // The Lord Chancellor, in exercise of the powers conferred on him by sections 9(4), (5) and (6) of the...
							Citation struct {
								Text   string `xml:",chardata"` // Justice (Northern Ireland) Act 2002
								URI    string `xml:"URI,attr"`
								ID     string `xml:"id,attr"`
								Class  string `xml:"Class,attr"`
								Year   string `xml:"Year,attr"`
								Number string `xml:"Number,attr"`
							} `xml:"Citation"`
							FootnoteRef struct {
								Text string `xml:",chardata"`
								Ref  string `xml:"Ref,attr"`
							} `xml:"FootnoteRef"`
						} `xml:"Text"`
					} `xml:"Para"`
				} `xml:"EnactingText"`
			} `xml:"SecondaryPreamble"`
		} `xml:"SecondaryPrelims"`
		Body struct {
			Text               string `xml:",chardata"`
			DocumentURI        string `xml:"DocumentURI,attr"`
			IdURI              string `xml:"IdURI,attr"`
			NumberOfProvisions string `xml:"NumberOfProvisions,attr"`
			NumberFormat       string `xml:"NumberFormat,attr"`
			P1group            []struct {
				Text  string `xml:",chardata"`
				Title struct {
					Text string `xml:",chardata"` // Title, commencement and interpretation, Eligibility of persons for appointment as Lay Magistrates
				} `xml:"Title"`
				P1 struct {
					Text        string `xml:",chardata"`
					DocumentURI string `xml:"DocumentURI,attr"`
					IdURI       string `xml:"IdURI,attr"`
					ID          string `xml:"id,attr"`
					Pnumber     struct {
						Text string `xml:",chardata"` // 1, 2
					} `xml:"Pnumber"`
					P1para struct {
						Chardata string `xml:",chardata"`
						P2       []struct {
							Text        string `xml:",chardata"`
							DocumentURI string `xml:"DocumentURI,attr"`
							IdURI       string `xml:"IdURI,attr"`
							ID          string `xml:"id,attr"`
							Pnumber     struct {
								Text string `xml:",chardata"` // 1, 2
							} `xml:"Pnumber"`
							P2para struct {
								Chardata string `xml:",chardata"`
								Text     struct {
									Text     string `xml:",chardata"` // This Order may be cited as the Lay Magistrates (Eligibility) (Northern Ireland) Order 2004 and shall...
									Citation struct {
										Text   string `xml:",chardata"` // Justice (Northern Ireland) Act 2002
										URI    string `xml:"URI,attr"`
										ID     string `xml:"id,attr"`
										Class  string `xml:"Class,attr"`
										Year   string `xml:"Year,attr"`
										Number string `xml:"Number,attr"`
									} `xml:"Citation"`
								} `xml:"Text"`
								UnorderedList struct {
									Text       string `xml:",chardata"`
									Decoration string `xml:"Decoration,attr"`
									Class      string `xml:"Class,attr"`
									ListItem   []struct {
										Text string `xml:",chardata"`
										Para struct {
											Chardata string `xml:",chardata"`
											Text     struct {
												Text string `xml:",chardata"` // “close relative” means, in relation to a particular person, the father, father-in-law, mother, m...
											} `xml:"Text"`
										} `xml:"Para"`
									} `xml:"ListItem"`
								} `xml:"UnorderedList"`
							} `xml:"P2para"`
						} `xml:"P2"`
						Text struct {
							Text string `xml:",chardata"` // Unless the Lord Chancellor otherwise determines in the case of a particular person, no person shall ...
						} `xml:"Text"`
						P3 []struct {
							Text        string `xml:",chardata"`
							DocumentURI string `xml:"DocumentURI,attr"`
							IdURI       string `xml:"IdURI,attr"`
							ID          string `xml:"id,attr"`
							Pnumber     struct {
								Text string `xml:",chardata"` // a, b, c, d, e, f, g, h, i, j
							} `xml:"Pnumber"`
							P3para struct {
								Chardata string `xml:",chardata"`
								Text     []struct {
									Text string `xml:",chardata"` // if he does not reside or work in, or within 15 miles of, the county court division to which the appo...
								} `xml:"Text"`
								P4 []struct {
									Text        string `xml:",chardata"`
									DocumentURI string `xml:"DocumentURI,attr"`
									IdURI       string `xml:"IdURI,attr"`
									ID          string `xml:"id,attr"`
									Pnumber     struct {
										Text string `xml:",chardata"` // i, ii, iii, iv, v, vi, i, ii, iii, iv
									} `xml:"Pnumber"`
									P4para struct {
										Chardata string `xml:",chardata"`
										Text     struct {
											Text        string `xml:",chardata"` // the House of Commons or the House of Lords,, the European Parliament,, the Scottish Parliament,, the...
											FootnoteRef struct {
												Text string `xml:",chardata"`
												Ref  string `xml:"Ref,attr"`
											} `xml:"FootnoteRef"`
										} `xml:"Text"`
									} `xml:"P4para"`
								} `xml:"P4"`
							} `xml:"P3para"`
						} `xml:"P3"`
					} `xml:"P1para"`
				} `xml:"P1"`
			} `xml:"P1group"`
			SignedSection struct {
				Text        string `xml:",chardata"`
				DocumentURI string `xml:"DocumentURI,attr"`
				IdURI       string `xml:"IdURI,attr"`
				Signatory   struct {
					Text   string `xml:",chardata"`
					Signee struct {
						Text       string `xml:",chardata"`
						PersonName struct {
							Text string `xml:",chardata"`
						} `xml:"PersonName"`
						DateSigned struct {
							Text     string `xml:",chardata"`
							DateText struct {
								Text string `xml:",chardata"` // Dated 2004.
							} `xml:"DateText"`
						} `xml:"DateSigned"`
					} `xml:"Signee"`
				} `xml:"Signatory"`
			} `xml:"SignedSection"`
		} `xml:"Body"`
		ExplanatoryNotes struct {
			Text        string `xml:",chardata"`
			DocumentURI string `xml:"DocumentURI,attr"`
			IdURI       string `xml:"IdURI,attr"`
			Comment     struct {
				Text string `xml:",chardata"`
				Para struct {
					Chardata string `xml:",chardata"`
					Text     struct {
						Text string `xml:",chardata"` // (This note is not part of the Order.)
					} `xml:"Text"`
				} `xml:"Para"`
			} `xml:"Comment"`
			P struct {
				Chardata string `xml:",chardata"`
				Text     struct {
					Text     string `xml:",chardata"` // This Order provides that, unless the Lord Chancellor otherwise determines in the case of a particula...
					Citation struct {
						Text   string `xml:",chardata"` // Justice (Northern Ireland) Act 2002
						URI    string `xml:"URI,attr"`
						ID     string `xml:"id,attr"`
						Class  string `xml:"Class,attr"`
						Year   string `xml:"Year,attr"`
						Number string `xml:"Number,attr"`
					} `xml:"Citation"`
				} `xml:"Text"`
				P2 []struct {
					Text    string `xml:",chardata"`
					Pnumber struct {
						Text string `xml:",chardata"` // 1, 2, 3, 4, 5, 6
					} `xml:"Pnumber"`
					P2para struct {
						Chardata string `xml:",chardata"`
						Text     struct {
							Text string `xml:",chardata"` // he does not reside or work in (or within 15 miles of) the county court division to which the appoint...
						} `xml:"Text"`
					} `xml:"P2para"`
				} `xml:"P2"`
			} `xml:"P"`
		} `xml:"ExplanatoryNotes"`
	} `xml:"Secondary"`
	Footnotes struct {
		Text     string `xml:",chardata"`
		Footnote []struct {
			Text         string `xml:",chardata"`
			ID           string `xml:"id,attr"`
			FootnoteText struct {
				Text string `xml:",chardata"`
				Para struct {
					Chardata string `xml:",chardata"`
					Text     struct {
						Text     string `xml:",chardata"` // Section 1(1) of the Police (Northern Ireland) Act, Section 1(3) of the Police (Northern Ireland) Act...
						Citation struct {
							Text   string `xml:",chardata"` // 2002 c.26, 2000 (c. 32), 2000 (c. 32), 2000 (c. 32), 1998 (c. 32), 2002 (c. 29), 1972/538 (N.I. 1), ...
							URI    string `xml:"URI,attr"`
							ID     string `xml:"id,attr"`
							Class  string `xml:"Class,attr"`
							Year   string `xml:"Year,attr"`
							Number string `xml:"Number,attr"`
						} `xml:"Citation"`
						Acronym struct {
							Text      string `xml:",chardata"` // S.I.
							Expansion string `xml:"Expansion,attr"`
						} `xml:"Acronym"`
					} `xml:"Text"`
				} `xml:"Para"`
			} `xml:"FootnoteText"`
		} `xml:"Footnote"`
	} `xml:"Footnotes"`
}
