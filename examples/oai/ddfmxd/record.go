package ddfmxd

import "encoding/xml"

type Record struct {
	XMLName xml.Name `xml:"Record"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text       string `xml:",chardata"`
		Status     string `xml:"status,attr"`
		Identifier struct {
			Text string `xml:",chardata"` // oai:pure.atira.dk:publica...
		} `xml:"identifier"`
		Datestamp struct {
			Text string `xml:",chardata"` // 2015-10-29T10:56:55Z, 201...
		} `xml:"datestamp"`
		SetSpec []struct {
			Text string `xml:",chardata"` // publications:all, publica...
		} `xml:"setSpec"`
	} `xml:"header"`
	Metadata struct {
		Text   string `xml:",chardata"`
		DdfDoc struct {
			Text           string `xml:",chardata"`
			Mxd            string `xml:"mxd,attr"`
			Xsi            string `xml:"xsi,attr"`
			FormatVersion  string `xml:"format_version,attr"`
			DocType        string `xml:"doc_type,attr"`
			DocLang        string `xml:"doc_lang,attr"`
			DocYear        string `xml:"doc_year,attr"`
			DocReview      string `xml:"doc_review,attr"`
			DocLevel       string `xml:"doc_level,attr"`
			RecSource      string `xml:"rec_source,attr"`
			RecID          string `xml:"rec_id,attr"`
			RecCreated     string `xml:"rec_created,attr"`
			RecUpd         string `xml:"rec_upd,attr"`
			RecStatus      string `xml:"rec_status,attr"`
			SchemaLocation string `xml:"schemaLocation,attr"`
			Title          struct {
				Text     string `xml:",chardata"`
				Original struct {
					Text string `xml:",chardata"`
					Lang string `xml:"lang,attr"`
					Main struct {
						Text string `xml:",chardata"` // Mattia Preti’s vault in...
						Xs   string `xml:"xs,attr"`
						Type string `xml:"type,attr"`
					} `xml:"main"`
					Sub struct {
						Text string `xml:",chardata"` // Tegninger, fotografier og...
					} `xml:"sub"`
				} `xml:"original"`
				Translated []struct {
					Text string `xml:",chardata"`
					Lang string `xml:"lang,attr"`
					Main struct {
						Text string `xml:",chardata"` // Hammershøiana, Musikalsk...
						Xs   string `xml:"xs,attr"`
						Type string `xml:"type,attr"`
					} `xml:"main"`
					Sub struct {
						Text string `xml:",chardata"` // Drawings, photographs and...
					} `xml:"sub"`
				} `xml:"translated"`
			} `xml:"title"`
			Description struct {
				Text     string `xml:",chardata"`
				Abstract []struct {
					Text string `xml:",chardata"` // The article aims at casti...
					Lang string `xml:"lang,attr"`
				} `xml:"abstract"`
				ResearchArea struct {
					Text     string `xml:",chardata"` // Humanities, Humanities, H...
					Lang     string `xml:"lang,attr"`
					AreaCode string `xml:"area_code,attr"`
				} `xml:"research_area"`
				Subject struct {
					Text    string `xml:",chardata"`
					Keyword []struct {
						Text    string `xml:",chardata"` // værkfortegnelse, digital...
						KeyType string `xml:"key_type,attr"`
					} `xml:"keyword"`
				} `xml:"subject"`
				Note struct {
					Text string `xml:",chardata"` // Lennart Gottlieb: "Langel...
					Lang string `xml:"lang,attr"`
				} `xml:"note"`
			} `xml:"description"`
			Person []struct {
				Text     string `xml:",chardata"`
				PersRole string `xml:"pers_role,attr"`
				AffNo    string `xml:"aff_no,attr"`
				Name     struct {
					Text  string `xml:",chardata"`
					First struct {
						Text string `xml:",chardata"` // Jesper, Jesper, Marina, M...
					} `xml:"first"`
					Last struct {
						Text string `xml:",chardata"` // Svenningsen, Svenningsen,...
					} `xml:"last"`
				} `xml:"name"`
				ID []struct {
					Text     string `xml:",chardata"` // 89f77890-6424-4f8a-8004-c...
					IDType   string `xml:"id_type,attr"`
					IDSource string `xml:"id_source,attr"`
				} `xml:"id"`
				Title struct {
					Text string `xml:",chardata"` // Projektforsker, Projektfo...
				} `xml:"title"`
				Country struct {
					Text string `xml:",chardata"` // dk, dk, dk, dk, dk, dk, d...
				} `xml:"country"`
				Address struct {
					Text string `xml:",chardata"` // Sølvgade 48-50, 1307 Kø...
				} `xml:"address"`
				URI struct {
					Text string `xml:",chardata"` // http://pure-01.kb.dk/port...
				} `xml:"uri"`
				Email struct {
					Text string `xml:",chardata"` // atge@kb.dk, atge@kb.dk, a...
				} `xml:"email"`
				Birthdate struct {
					Text string `xml:",chardata"` // 1953-12-28+01:00, 1953-12...
				} `xml:"birthdate"`
			} `xml:"person"`
			Organisation []struct {
				Text    string `xml:",chardata"`
				OrgRole string `xml:"org_role,attr"`
				AffNo   string `xml:"aff_no,attr"`
				Name    []struct {
					Text   string `xml:",chardata"`
					Lang   string `xml:"lang,attr"`
					Level1 struct {
						Text string `xml:",chardata"` // Kulturministeriet, Minist...
					} `xml:"level1"`
					Level2 struct {
						Text string `xml:",chardata"` // SMK Statens Museum for Ku...
					} `xml:"level2"`
					Acronym struct {
						Text string `xml:",chardata"` // SMK, SMK, SMK, SMK, KB, K...
					} `xml:"acronym"`
					Level3 struct {
						Text string `xml:",chardata"` // CATS, CATS, CATS, CATS, C...
					} `xml:"level3"`
				} `xml:"name"`
				ID struct {
					Text   string `xml:",chardata"` // 213321, 213321, 213113, 2...
					IDType string `xml:"id_type,attr"`
				} `xml:"id"`
				URI struct {
					Text string `xml:",chardata"` // http://www.smk.dk/, http:...
				} `xml:"uri"`
			} `xml:"organisation"`
			LocalField struct {
				Text    string `xml:",chardata"`
				TagType string `xml:"tag_type,attr"`
				Code    struct {
					Text string `xml:",chardata"` // handleNet, handleNet, han...
				} `xml:"code"`
				Data struct {
					Text string `xml:",chardata"` // http://hdl.handle.net/109...
				} `xml:"data"`
			} `xml:"local_field"`
			Publication struct {
				Text      string `xml:",chardata"`
				InJournal struct {
					Text      string `xml:",chardata"`
					PubStatus string `xml:"pub_status,attr"`
					Title     struct {
						Text string `xml:",chardata"` // Artibus et Historiae, Stu...
						Xs   string `xml:"xs,attr"`
						Type string `xml:"type,attr"`
					} `xml:"title"`
					ISSN []struct {
						Text string `xml:",chardata"` // 03919064, 01481029, 18633...
						Type string `xml:"type,attr"`
					} `xml:"issn"`
					Year struct {
						Text string `xml:",chardata"` // 2015, 2008, 2011, 2013, 2...
					} `xml:"year"`
					Vol struct {
						Text string `xml:",chardata"` // 71, 29, 5, 02, 20, 45, 31...
					} `xml:"vol"`
					Pages struct {
						Text string `xml:",chardata"` // 282-295, 136-175, 289-294...
					} `xml:"pages"`
					Issue struct {
						Text string `xml:",chardata"` // 3, 2, 1, 1, 1, 1, 1, 1, I...
					} `xml:"issue"`
					URI struct {
						Text string `xml:",chardata"` // http://www.kb.dk/da/dia/o...
					} `xml:"uri"`
					Doi struct {
						Text string `xml:",chardata"` // 10.3402/polar.v31i0.20122...
					} `xml:"doi"`
				} `xml:"in_journal"`
				Book struct {
					Text      string `xml:",chardata"`
					PubStatus string `xml:"pub_status,attr"`
					ISBN      []struct {
						Text string `xml:",chardata"` // 9788790597153, 9788779557...
						Type string `xml:"type,attr"`
					} `xml:"isbn"`
					Publisher struct {
						Text string `xml:",chardata"` // Den Hirschsprungske Samli...
					} `xml:"publisher"`
					Year struct {
						Text string `xml:",chardata"` // 2011, 2013, 2009, 1992, 2...
					} `xml:"year"`
					Pages struct {
						Text string `xml:",chardata"` // 32, 315, 250, 183, 308, 1...
					} `xml:"pages"`
					Series []struct {
						Text string `xml:",chardata"` // 'Music Theory in Britain,...
					} `xml:"series"`
					URI struct {
						Text string `xml:",chardata"` // http://www.mtp.hum.ku.dk/...
					} `xml:"uri"`
					Edition struct {
						Text string `xml:",chardata"` // 1, 1., 1, 1, 1, 1., 1., W...
					} `xml:"edition"`
					Vol struct {
						Text string `xml:",chardata"` // 1, 1, 11, 1, 2, 6, 1, 1, ...
					} `xml:"vol"`
					ISSN struct {
						Text string `xml:",chardata"` // 01058746, 01058746, 15691...
						Type string `xml:"type,attr"`
					} `xml:"issn"`
					Issue struct {
						Text string `xml:",chardata"` // Bind XVI, hft. 58/61, Bin...
					} `xml:"issue"`
					Place struct {
						Text string `xml:",chardata"` // Kbh., Kbh., Kbh., Kbh.
					} `xml:"place"`
					Doi struct {
						Text string `xml:",chardata"` // 10.1093/acprof:oso/978019...
					} `xml:"doi"`
				} `xml:"book"`
				InBook struct {
					Text      string `xml:",chardata"`
					PubStatus string `xml:"pub_status,attr"`
					Title     struct {
						Text string `xml:",chardata"` // Tributes to Jonathan J.G....
						Xs   string `xml:"xs,attr"`
						Type string `xml:"type,attr"`
					} `xml:"title"`
					Publisher struct {
						Text string `xml:",chardata"` // Brepols Publishers, Klim,...
					} `xml:"publisher"`
					Year struct {
						Text string `xml:",chardata"` // 2006, 2012, 2012, 2012, 2...
					} `xml:"year"`
					Pages struct {
						Text string `xml:",chardata"` // 203-212, 9-25, 119-172, 2...
					} `xml:"pages"`
					ISBN []struct {
						Text string `xml:",chardata"` // 9788779557598, 9788763504...
						Type string `xml:"type,attr"`
					} `xml:"isbn"`
					URI struct {
						Text string `xml:",chardata"` // http://www.mtp.hum.ku.dk/...
					} `xml:"uri"`
					SubTitle struct {
						Text string `xml:",chardata"` // Sentieri liguri per viagg...
						Xs   string `xml:"xs,attr"`
						Type string `xml:"type,attr"`
					} `xml:"sub_title"`
					ISSN struct {
						Text string `xml:",chardata"` // 04096231, 01058746, 00650...
						Type string `xml:"type,attr"`
					} `xml:"issn"`
					Vol struct {
						Text string `xml:",chardata"` // 61, 2, 16, 1, 1, II, 1, E...
					} `xml:"vol"`
					Series []struct {
						Text string `xml:",chardata"` // Lares. Biblioteca, Danish...
					} `xml:"series"`
					Place struct {
						Text string `xml:",chardata"` // Sophienholm, Assens, Berl...
					} `xml:"place"`
					Edition struct {
						Text string `xml:",chardata"` // CATS Series of Technical ...
					} `xml:"edition"`
					Doi struct {
						Text string `xml:",chardata"` // 10.1163/9789004290839_008...
					} `xml:"doi"`
					Issue struct {
						Text string `xml:",chardata"` // 16, 2, 9, 10, 6, 168, 15,...
					} `xml:"issue"`
				} `xml:"in_book"`
				Other struct {
					// Text string `xml:",chardata"`
					Text struct {
						Text string `xml:",chardata"` // publisher: Dansk Center f...
					} `xml:"text"`
					Year struct {
						Text string `xml:",chardata"` // 2010, 2013, 2013, 2014, 2...
					} `xml:"year"`
				} `xml:"other"`
				Inetpub []struct {
					// Text string `xml:",chardata"`
					Text struct {
						Text string `xml:",chardata"` // PDF in institutional repo...
					} `xml:"text"`
					URI struct {
						Text   string `xml:",chardata"` // http://pure-01.kb.dk/ws/f...
						Access string `xml:"access,attr"`
					} `xml:"uri"`
				} `xml:"inetpub"`
				DigitalObject []struct {
					Text   string `xml:",chardata"`
					ID     string `xml:"id,attr"`
					Access string `xml:"access,attr"`
					Role   string `xml:"role,attr"`
					File   struct {
						Text      string `xml:",chardata"`
						Size      string `xml:"size,attr"`
						MimeType  string `xml:"mime_type,attr"`
						Timestamp string `xml:"timestamp,attr"`
						Filename  string `xml:"filename,attr"`
					} `xml:"file"`
					URI struct {
						Text string `xml:",chardata"` // http://pure-01.kb.dk/ws/f...
					} `xml:"uri"`
				} `xml:"digital_object"`
				InReport struct {
					Text      string `xml:",chardata"`
					PubStatus string `xml:"pub_status,attr"`
					Title     struct {
						Text string `xml:",chardata"` // Det Kongelige Bibliotek. ...
						Xs   string `xml:"xs,attr"`
						Type string `xml:"type,attr"`
					} `xml:"title"`
					Year struct {
						Text string `xml:",chardata"` // 2003, 2006, 2006, 2004, 2...
					} `xml:"year"`
					Pages struct {
						Text string `xml:",chardata"` // 54-56, 15-16, 59-64, 23-6...
					} `xml:"pages"`
					Place struct {
						Text string `xml:",chardata"` // Bergen, København, Højb...
					} `xml:"place"`
					SubTitle struct {
						Text string `xml:",chardata"` // arbejdsrapport fra semina...
						Xs   string `xml:"xs,attr"`
						Type string `xml:"type,attr"`
					} `xml:"sub_title"`
					Publisher struct {
						Text string `xml:",chardata"` // Kulturarvsstyrelsen, Jutl...
					} `xml:"publisher"`
					ISBN struct {
						Text string `xml:",chardata"` // 9788788415599, 9788789500...
						Type string `xml:"type,attr"`
					} `xml:"isbn"`
					Series struct {
						Text string `xml:",chardata"` // Jysk Arkæologisk Selskab...
					} `xml:"series"`
					ISSN struct {
						Text string `xml:",chardata"` // 01072854, 01054503, 16011...
						Type string `xml:"type,attr"`
					} `xml:"issn"`
					Issue struct {
						Text string `xml:",chardata"` // 69, 122, 9, 15, 15
					} `xml:"issue"`
					Vol struct {
						Text string `xml:",chardata"` // 4, 2, 7, 1996-97, 30
					} `xml:"vol"`
					URI struct {
						Text string `xml:",chardata"` // http://www.nwo.nl/demayer...
					} `xml:"uri"`
				} `xml:"in_report"`
				Report struct {
					Text      string `xml:",chardata"`
					PubStatus string `xml:"pub_status,attr"`
					Series    struct {
						Text string `xml:",chardata"` // SILA Feltrapport, NNU ra...
					} `xml:"series"`
					RepNo struct {
						Text string `xml:",chardata"` // SILA Feltrapport, No. 16...
					} `xml:"rep_no"`
					Publisher struct {
						Text string `xml:",chardata"` // Nationalmuseet, Nationalm...
					} `xml:"publisher"`
					Year struct {
						Text string `xml:",chardata"` // 2004, 2009, 2008, 2008, 2...
					} `xml:"year"`
					Issue struct {
						Text string `xml:",chardata"` // 16, 2009:15, CB 2:2008, 2...
					} `xml:"issue"`
					Pages struct {
						Text string `xml:",chardata"` // 72, 9, 0, 0, 0, 0, 0, 9, ...
					} `xml:"pages"`
					Vol struct {
						Text string `xml:",chardata"` // 30, 16, 14, 34, 33, NNU r...
					} `xml:"vol"`
					ISBN struct {
						Text string `xml:",chardata"` // 9788792825278, 9788770731...
						Type string `xml:"type,attr"`
					} `xml:"isbn"`
					URI struct {
						Text string `xml:",chardata"` // file:///C:/Users/jwa/Down...
					} `xml:"uri"`
					ISSN struct {
						Text string `xml:",chardata"` // 13997335
						Type string `xml:"type,attr"`
					} `xml:"issn"`
				} `xml:"report"`
			} `xml:"publication"`
			OaLink []struct {
				Text         string `xml:",chardata"`
				Type         string `xml:"type,attr"`
				Version      string `xml:"version,attr"`
				PublicAccess string `xml:"public_access,attr"`
				URL          string `xml:"url,attr"`
				License      string `xml:"license,attr"`
				EmbargoStart string `xml:"embargo_start,attr"`
				EmbargoEnd   string `xml:"embargo_end,attr"`
			} `xml:"oa_link"`
			Event struct {
				Text      string `xml:",chardata"`
				EventRole string `xml:"event_role,attr"`
				Title     struct {
					Text string `xml:",chardata"`
					Lang string `xml:"lang,attr"`
					Full struct {
						Text string `xml:",chardata"` // Music Encoding Conference...
					} `xml:"full"`
				} `xml:"title"`
				Dates struct {
					Text  string `xml:",chardata"`
					Start struct {
						Text string `xml:",chardata"` // 2013-05-21, 2013-05-21, 2...
					} `xml:"start"`
					End struct {
						Text string `xml:",chardata"` // 2013-05-24, 2013-05-24, 2...
					} `xml:"end"`
				} `xml:"dates"`
				Place struct {
					Text string `xml:",chardata"` // Mainz, de, Mainz, de, Hel...
				} `xml:"place"`
			} `xml:"event"`
			Project []struct {
				Text     string `xml:",chardata"`
				ProjRole string `xml:"proj_role,attr"`
				Title    []struct {
					Text string `xml:",chardata"`
					Lang string `xml:"lang,attr"`
				} `xml:"title"`
				URI struct {
					Text string `xml:",chardata"` // http://hintme.dk/, http:/...
				} `xml:"uri"`
			} `xml:"project"`
		} `xml:"ddf_doc"`
	} `xml:"metadata"`
	About struct {
		Text string `xml:",chardata"`
	} `xml:"about"`
}
