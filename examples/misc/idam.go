package misc

import "encoding/xml"

// IdamPublication was generated 2017-12-18 16:15:28 by tir on hayiti.
type IdamPublication struct {
	XMLName xml.Name `xml:"publication"`
	Text    string   `xml:",chardata"`
	Title   struct {
		Text string `xml:",chardata"` // IET Generation, Transmiss...
	} `xml:"title"`
	Normtitle struct {
		Text string `xml:",chardata"` // Generation, Transmission ...
	} `xml:"normtitle"`
	Publicationinfo struct {
		Text    string `xml:",chardata"`
		Idamsid struct {
			Text string `xml:",chardata"` // 0b000064806f4938, 0b00006...
		} `xml:"idamsid"`
		Publicationtype struct {
			Text string `xml:",chardata"` // Periodical, Periodical, P...
		} `xml:"publicationtype"`
		Publicationsubtype struct {
			Text string `xml:",chardata"` // IEE Periodical, IEE Perio...
		} `xml:"publicationsubtype"`
		Pubstatus struct {
			Text string `xml:",chardata"` // Active, Active, Active, A...
		} `xml:"pubstatus"`
		Publicationopenaccess struct {
			Text string `xml:",chardata"` // F, F, F, F, F, F, F, F, F...
		} `xml:"publicationopenaccess"`
		StandardID struct {
			Text string `xml:",chardata"` // 0, 0, 0, 0, 0, 0, 0, 0, 0...
		} `xml:"standard_id"`
		ISSN []struct {
			Text      string `xml:",chardata"` // 1751-8687, 1751-8695, 175...
			Mediatype string `xml:"mediatype,attr"`
		} `xml:"issn"`
		Keywordset struct {
			Text        string `xml:",chardata"`
			Keywordtype string `xml:"keywordtype,attr"`
			Keyword     struct {
				Text        string `xml:",chardata"`
				Keywordterm struct {
					Text string `xml:",chardata"` // Changed "Copyright" from ...
				} `xml:"keywordterm"`
			} `xml:"keyword"`
		} `xml:"keywordset"`
		Pubtopicalbrowseset struct {
			Text             string `xml:",chardata"`
			Pubtopicalbrowse []struct {
				Text string `xml:",chardata"` // Power, Energy and Industr...
			} `xml:"pubtopicalbrowse"`
		} `xml:"pubtopicalbrowseset"`
		Copyrightgroup struct {
			Text      string `xml:",chardata"`
			Copyright struct {
				Text string `xml:",chardata"`
				Year struct {
					Text string `xml:",chardata"` // 2007-2012, 2007-2012, 200...
				} `xml:"year"`
				Holder struct {
					Text string `xml:",chardata"` // IET, IET, IET, IET, IET, ...
				} `xml:"holder"`
			} `xml:"copyright"`
		} `xml:"copyrightgroup"`
		Publisher struct {
			Text          string `xml:",chardata"`
			Publishername struct {
				Text string `xml:",chardata"` // IET, IET, IET, IET, IET, ...
			} `xml:"publishername"`
			Address struct {
				Text    string `xml:",chardata"`
				Country struct {
					Text string `xml:",chardata"` // UK, UK, UK, UK, UK, UK, U...
				} `xml:"country"`
			} `xml:"address"`
		} `xml:"publisher"`
		Holdstatus struct {
			Text string `xml:",chardata"` // Publish, Publish, Publish...
		} `xml:"holdstatus"`
		Confgroup struct {
			Text          string `xml:",chardata"`
			DoiPermission struct {
				Text string `xml:",chardata"` // F, F, F, F, F, F, F, F, F...
			} `xml:"doi_permission"`
			Confdate []struct {
				Text         string `xml:",chardata"`
				Confdatetype string `xml:"confdatetype,attr"`
				Year         struct {
					Text string `xml:",chardata"` // 1994, 1994, 1994, 1994, 1...
				} `xml:"year"`
				Month struct {
					Text string `xml:",chardata"` // Aug., Aug., Aug., Aug., A...
				} `xml:"month"`
				Day struct {
					Text string `xml:",chardata"` // 24, 26, 24, 26, 24, 26, 2...
				} `xml:"day"`
			} `xml:"confdate"`
			Confcountry struct {
				Text string `xml:",chardata"` // UK, UK, UK, UK, UK, UK, U...
			} `xml:"confcountry"`
		} `xml:"confgroup"`
		Amsid struct {
			Text string `xml:",chardata"` // 4082359, 4082359, 4082359...
		} `xml:"amsid"`
		Coden struct {
			Text string `xml:",chardata"` // IGTDAW, IGTDAW, IGTDAW, I...
		} `xml:"coden"`
		Packagememberset struct {
			Text          string `xml:",chardata"`
			Packagemember struct {
				Text string `xml:",chardata"` // None, None, None, None, N...
			} `xml:"packagemember"`
		} `xml:"packagememberset"`
		Ieeeabbrev struct {
			Text string `xml:",chardata"` // IEE, IEE, IEE, IEE, IEE, ...
		} `xml:"ieeeabbrev"`
	} `xml:"publicationinfo"`
	Volume struct {
		Text       string `xml:",chardata"`
		Volumeinfo struct {
			Text string `xml:",chardata"`
			Year struct {
				Text string `xml:",chardata"` // 2017, 2017, 2017, 2017, 2...
			} `xml:"year"`
			Volumenum struct {
				Text string `xml:",chardata"` // 11, 11, 11, 11, 11, 11, 1...
			} `xml:"volumenum"`
			Idamsid struct {
				Text string `xml:",chardata"` // 0b00006485b9045f, 0b00006...
			} `xml:"idamsid"`
			Issue struct {
				Text     string `xml:",chardata"`
				Issuenum struct {
					Text string `xml:",chardata"` // 2, 2, 2, 2, 2, 2, 2, 2, 2...
				} `xml:"issuenum"`
				Amsid struct {
					Text string `xml:",chardata"` // 7834146, 7834146, 7834146...
				} `xml:"amsid"`
				Issuestatus struct {
					Text string `xml:",chardata"` // Complete, Complete, Compl...
				} `xml:"issuestatus"`
			} `xml:"issue"`
		} `xml:"volumeinfo"`
		Article struct {
			Text  string `xml:",chardata"`
			Title struct {
				Text string `xml:",chardata"` // Power-dependent droop-bas...
			} `xml:"title"`
			Articleinfo struct {
				Text          string `xml:",chardata"`
				Articleseqnum struct {
					Text string `xml:",chardata"` // 3831, 3391, 3301, 5401, 4...
				} `xml:"articleseqnum"`
				Csarticlesortorder struct {
					Text string `xml:",chardata"` // 0, 0, 0, 0, 0, 0, 0, 0, 0...
				} `xml:"csarticlesortorder"`
				Articledoi struct {
					Text string `xml:",chardata"` // 10.1049/iet-gtd.2016.0764...
				} `xml:"articledoi"`
				Idamsid struct {
					Text string `xml:",chardata"` // 0b00006485b94e5b, 0b00006...
				} `xml:"idamsid"`
				Articlestatus struct {
					Text string `xml:",chardata"` // Active, Active, Active, A...
				} `xml:"articlestatus"`
				Contenttype struct {
					Text string `xml:",chardata"` // orig-research, orig-resea...
				} `xml:"contenttype"`
				Othercontenttype struct {
					Text string `xml:",chardata"` // research-article, researc...
				} `xml:"othercontenttype"`
				Articleopenaccess struct {
					Text string `xml:",chardata"` // F, F, F, F, F, F, F, F, F...
				} `xml:"articleopenaccess"`
				Articleshowflag struct {
					Text string `xml:",chardata"` // T, T, T, T, T, T, T, T, T...
				} `xml:"articleshowflag"`
				Articleplagiarizedflag struct {
					Text string `xml:",chardata"` // F, F, F, F, F, F, F, F, F...
				} `xml:"articleplagiarizedflag"`
				Articlenodoiflag struct {
					Text string `xml:",chardata"` // F, F, F, F, F, F, F, F, F...
				} `xml:"articlenodoiflag"`
				Articlecoverimageflag struct {
					Text string `xml:",chardata"` // F, F, F, F, F, F, F, F, F...
				} `xml:"articlecoverimageflag"`
				Csarticlehtmlflag struct {
					Text string `xml:",chardata"` // F, F, F, F, F, F, F, F, F...
				} `xml:"csarticlehtmlflag"`
				Articlereferenceflag struct {
					Text string `xml:",chardata"` // F, F, F, F, F, F, F, F, F...
				} `xml:"articlereferenceflag"`
				Articlepeerreviewflag struct {
					Text string `xml:",chardata"` // F, F, F, F, F, F, F, F, F...
				} `xml:"articlepeerreviewflag"`
				Holdstatus struct {
					Text string `xml:",chardata"` // Publish, Publish, Publish...
				} `xml:"holdstatus"`
				Issuenum struct {
					Text string `xml:",chardata"` // 2, 2, 2, 2, 2, 2, 2, 2, 2...
				} `xml:"issuenum"`
				Articlecopyright struct {
					Text         string `xml:",chardata"` // © The Institution of Eng...
					Holderisieee string `xml:"holderisieee,attr"`
					Year         string `xml:"year,attr"`
				} `xml:"articlecopyright"`
				Abstract struct {
					Text         string `xml:",chardata"` // The concept of voltage so...
					Abstracttype string `xml:"abstracttype,attr"`
				} `xml:"abstract"`
				Authorgroup struct {
					Text   string `xml:",chardata"`
					Author []struct {
						Text        string `xml:",chardata"`
						Authororder struct {
							Text string `xml:",chardata"` // 0, 0, 0, 0, 0, 0, 0, 0, 0...
						} `xml:"authororder"`
						Normname struct {
							Text string `xml:",chardata"` // Stamatiou, G., Bongiorno,...
						} `xml:"normname"`
						Nonnormname struct {
							Text string `xml:",chardata"` // Georgios Stamatiou, Massi...
						} `xml:"nonnormname"`
						Authorrefid struct {
							Text string `xml:",chardata"` // A16585073.1, A16585073.1,...
						} `xml:"authorrefid"`
						Firstname struct {
							Text string `xml:",chardata"` // Georgios, Massimo, Mikel,...
						} `xml:"firstname"`
						Surname struct {
							Text string `xml:",chardata"` // Stamatiou, Bongiorno, Arm...
						} `xml:"surname"`
						Affiliation struct {
							Text string `xml:",chardata"` // Dept. of Energy &amp; Env...
						} `xml:"affiliation"`
						Email struct {
							Text string `xml:",chardata"` // georgios.stamatiou@chalme...
						} `xml:"email"`
						Unicodefirstname struct {
							Text string `xml:",chardata"` // Abdulatif, Heshan, Ibrahi...
						} `xml:"unicodefirstname"`
						Unicodesurname struct {
							Text string `xml:",chardata"` // Alabdulatif, Kumarage, Kh...
						} `xml:"unicodesurname"`
						Authortype struct {
							Text string `xml:",chardata"` // Author, Author, Author, A...
						} `xml:"authortype"`
					} `xml:"author"`
				} `xml:"authorgroup"`
				Date []struct {
					Text     string `xml:",chardata"`
					Datetype string `xml:"datetype,attr"`
					Year     struct {
						Text string `xml:",chardata"` // 2017, 2017, 2017, 2017, 2...
					} `xml:"year"`
					Month struct {
						Text string `xml:",chardata"` // 1, 1, 1, 1, 1, 1, 1, 1, 1...
					} `xml:"month"`
					Day struct {
						Text string `xml:",chardata"` // 26, 26, 27, 26, 26, 27, 2...
					} `xml:"day"`
				} `xml:"date"`
				Numpages struct {
					Text string `xml:",chardata"` // 0, 0, 0, 0, 0, 0, 0, 0, 0...
				} `xml:"numpages"`
				Size struct {
					Text string `xml:",chardata"` // 872075, 716350, 362541, 1...
				} `xml:"size"`
				Filename []struct {
					Text         string `xml:",chardata"` // 07835087.pdf, 07834938.pd...
					Docpartition string `xml:"docpartition,attr"`
					Filetype     string `xml:"filetype,attr"`
				} `xml:"filename"`
				Artpagenums struct {
					Text      string `xml:",chardata"` // 13-13, 30-31, 26-29, 4-4,...
					Endpage   string `xml:"endpage,attr"`
					Startpage string `xml:"startpage,attr"`
				} `xml:"artpagenums"`
				Pubsnumber struct {
					Text      string `xml:",chardata"` // 16585073, 16585068, 16585...
					Pubidtype string `xml:"pubidtype,attr"`
				} `xml:"pubsnumber"`
				Numreferences struct {
					Text string `xml:",chardata"` // 24, 33, 34, 22, 25, 30, 2...
				} `xml:"numreferences"`
				Amsid struct {
					Text string `xml:",chardata"` // 7835087, 7834938, 7834955...
				} `xml:"amsid"`
				Keywordset []struct {
					Text        string `xml:",chardata"`
					Keywordtype string `xml:"keywordtype,attr"`
					Keyword     []struct {
						Text        string `xml:",chardata"`
						Keywordterm struct {
							Text string `xml:",chardata"` // voltage control, HVDC pow...
						} `xml:"keywordterm"`
					} `xml:"keyword"`
				} `xml:"keywordset"`
				Indexclassificationset struct {
					Text                string `xml:",chardata"`
					Indexclassification []struct {
						Text               string `xml:",chardata"` // B8120G, B8110C, C3340H, C...
						Classificationcode string `xml:"classificationcode,attr"`
					} `xml:"indexclassification"`
				} `xml:"indexclassificationset"`
				Treatmentcodeset struct {
					Text          string `xml:",chardata"`
					Treatmentcode []struct {
						Text string `xml:",chardata"` // P, T, P, P, T, P, T, P, P...
					} `xml:"treatmentcode"`
				} `xml:"treatmentcodeset"`
				Numericalindexset struct {
					Text           string `xml:",chardata"`
					Numericalindex []struct {
						Text             string `xml:",chardata"`
						Physicalquantity struct {
							Text string `xml:",chardata"` // power, voltage, size, pow...
						} `xml:"physicalquantity"`
						Numericvalue []struct {
							Text string `xml:",chardata"` // 1.0E+04, 1.2E+03, 4.5E-08...
						} `xml:"numericvalue"`
						Standardunit struct {
							Text string `xml:",chardata"` // W, V, m, W, bit/s, Hz, Hz...
						} `xml:"standardunit"`
					} `xml:"numericalindex"`
				} `xml:"numericalindexset"`
				Chemicalindexset struct {
					Text          string `xml:",chardata"`
					Chemicalindex []struct {
						Text     string `xml:",chardata"` // C, C, GaAs, As, Ga, GaAs,...
						Chemrole string `xml:"chemrole,attr"`
					} `xml:"chemicalindex"`
				} `xml:"chemicalindexset"`
			} `xml:"articleinfo"`
		} `xml:"article"`
	} `xml:"volume"`
	Titleabbrev struct {
		Text string `xml:",chardata"` // IET Radar Sonar Navig., I...
	} `xml:"titleabbrev"`
}
