package mods

import "encoding/xml"

type Record struct {
	XMLName xml.Name `xml:"Record"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text       string `xml:",chardata"`
		Status     string `xml:"status,attr"`
		Identifier struct {
			Text string `xml:",chardata"` // oai:berkeley.edu:13030-hb...
		} `xml:"identifier"`
		Datestamp struct {
			Text string `xml:",chardata"` // 2014-10-21T21:45:55Z, 201...
		} `xml:"datestamp"`
	} `xml:"header"`
	Metadata struct {
		Text   string `xml:",chardata"`
		Record struct {
			Text   string `xml:",chardata"`
			Header struct {
				Text       string `xml:",chardata"`
				Identifier struct {
					Text string `xml:",chardata"` // oai:library.berkeley.edu:...
				} `xml:"identifier"`
				Datestamp struct {
					Text string `xml:",chardata"` // 2008-02-19T14:06:03Z, 200...
				} `xml:"datestamp"`
				SetSpec struct {
					Text string `xml:",chardata"` // technicalReports:EECS, te...
				} `xml:"setSpec"`
			} `xml:"header"`
			Metadata struct {
				Text string `xml:",chardata"`
				Mods struct {
					Text           string `xml:",chardata"`
					Mods           string `xml:"mods,attr"`
					Xsi            string `xml:"xsi,attr"`
					SchemaLocation string `xml:"schemaLocation,attr"`
					TitleInfo      struct {
						Text  string `xml:",chardata"`
						Title struct {
							Text string `xml:",chardata"` // To Support Global Change ...
						} `xml:"title"`
					} `xml:"titleInfo"`
					Name []struct {
						Text     string `xml:",chardata"`
						Type     string `xml:"type,attr"`
						NamePart struct {
							Text string `xml:",chardata"` // Stonebraker, Michael, Doz...
						} `xml:"namePart"`
						Role struct {
							Text     string `xml:",chardata"`
							RoleTerm struct {
								Text      string `xml:",chardata"` // Author, Author, Author, A...
								Type      string `xml:"type,attr"`
								Authority string `xml:"authority,attr"`
							} `xml:"roleTerm"`
						} `xml:"role"`
					} `xml:"name"`
					TypeOfResource struct {
						Text string `xml:",chardata"` // text, text, text, text, t...
					} `xml:"typeOfResource"`
					Genre []struct {
						Text      string `xml:",chardata"` // technical reports, techni...
						Authority string `xml:"authority,attr"`
					} `xml:"genre"`
					OriginInfo struct {
						Text  string `xml:",chardata"`
						Place []struct {
							Text      string `xml:",chardata"`
							PlaceTerm struct {
								Text string `xml:",chardata"` // University of California ...
								Type string `xml:"type,attr"`
							} `xml:"placeTerm"`
						} `xml:"place"`
						Publisher []struct {
							Text string `xml:",chardata"` // Computer Science Division...
						} `xml:"publisher"`
						DateCreated []struct {
							Text     string `xml:",chardata"` // September 17, 1991, Augus...
							Encoding string `xml:"encoding,attr"`
							Point    string `xml:"point,attr"`
						} `xml:"dateCreated"`
						DateIssued []struct {
							Text     string `xml:",chardata"` // January 1990, 1990-01, Ju...
							Encoding string `xml:"encoding,attr"`
							Point    string `xml:"point,attr"`
						} `xml:"dateIssued"`
						DateModified []struct {
							Text     string `xml:",chardata"` // revised March 1992, 1992-...
							Encoding string `xml:"encoding,attr"`
							Point    string `xml:"point,attr"`
						} `xml:"dateModified"`
					} `xml:"originInfo"`
					Language []struct {
						Text         string `xml:",chardata"`
						LanguageTerm struct {
							Text      string `xml:",chardata"` // eng, eng, eng, eng, eng, ...
							Type      string `xml:"type,attr"`
							Authority string `xml:"authority,attr"`
						} `xml:"languageTerm"`
					} `xml:"language"`
					PhysicalDescription struct {
						Text   string `xml:",chardata"`
						Extent struct {
							Text string `xml:",chardata"` // 26 p, 77 p, 29 p, 268 p, ...
						} `xml:"extent"`
					} `xml:"physicalDescription"`
					Abstract []struct {
						Text string `xml:",chardata"` // Improved data management ...
					} `xml:"abstract"`
					RelatedItem []struct {
						Text         string `xml:",chardata"`
						DisplayLabel string `xml:"displayLabel,attr"`
						Type         string `xml:"type,attr"`
						TitleInfo    struct {
							Text  string `xml:",chardata"`
							Title struct {
								Text string `xml:",chardata"` // EECS, technical reports, ...
							} `xml:"title"`
						} `xml:"titleInfo"`
						Identifier struct {
							Text string `xml:",chardata"` // bk0000m744s, hb1k40068r, ...
							Type string `xml:"type,attr"`
						} `xml:"identifier"`
					} `xml:"relatedItem"`
					Identifier []struct {
						Text         string `xml:",chardata"` // S2K-91-01, EECS-2006-105,...
						Type         string `xml:"type,attr"`
						DisplayLabel string `xml:"displayLabel,attr"`
					} `xml:"identifier"`
					Location []struct {
						Text             string `xml:",chardata"`
						PhysicalLocation struct {
							Text string `xml:",chardata"` // Engineering Library, 110 ...
						} `xml:"physicalLocation"`
						URL struct {
							Text string `xml:",chardata"` // http://techreports.lib.be...
						} `xml:"url"`
					} `xml:"location"`
					Note []struct {
						Text string `xml:",chardata"` // Proc. Nat. Acad. Sci. U.S...
						Type string `xml:"type,attr"`
					} `xml:"note"`
				} `xml:"mods"`
			} `xml:"metadata"`
		} `xml:"record"`
	} `xml:"metadata"`
	About struct {
		Text string `xml:",chardata"`
	} `xml:"about"`
}
