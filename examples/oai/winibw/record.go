package winibw

import "encoding/xml"

type Record struct {
	XMLName xml.Name `xml:"Record"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text       string `xml:",chardata"`
		Status     string `xml:"status,attr"`
		Identifier struct {
			Text string `xml:",chardata"` // oai:qucosa.de:bsz:ch1-199...
		} `xml:"identifier"`
		Datestamp struct {
			Text string `xml:",chardata"` // 2010-08-25, 2010-08-25, 2...
		} `xml:"datestamp"`
		SetSpec []struct {
			Text string `xml:",chardata"` // pub-type:paper, has-sourc...
		} `xml:"setSpec"`
	} `xml:"header"`
	Metadata struct {
		Text      string `xml:",chardata"`
		XMetaDiss struct {
			Text           string `xml:",chardata"`
			Xmlns          string `xml:"xmlns,attr"`
			XMetaDiss      string `xml:"xMetaDiss,attr"`
			Cc             string `xml:"cc,attr"`
			Dc             string `xml:"dc,attr"`
			Dcmitype       string `xml:"dcmitype,attr"`
			Dcterms        string `xml:"dcterms,attr"`
			Pc             string `xml:"pc,attr"`
			URN            string `xml:"urn,attr"`
			Thesis         string `xml:"thesis,attr"`
			Dini           string `xml:"dini,attr"`
			Ddb            string `xml:"ddb,attr"`
			Fn             string `xml:"fn,attr"`
			Str            string `xml:"str,attr"`
			SchemaLocation string `xml:"schemaLocation,attr"`
			Xsi            string `xml:"xsi,attr"`
			Title          []struct {
				Text string `xml:",chardata"` // Realisierung eines Batchm...
				Type string `xml:"type,attr"`
				Lang string `xml:"lang,attr"`
				// Type string `xml:"type,attr"`
			} `xml:"title"`
			Creator []struct {
				Text   string `xml:",chardata"`
				Type   string `xml:"type,attr"`
				Person struct {
					Text string `xml:",chardata"`
					Name struct {
						Text     string `xml:",chardata"`
						Type     string `xml:"type,attr"`
						ForeName struct {
							Text string `xml:",chardata"` // Jens, Michael, Anne-Corin...
						} `xml:"foreName"`
						SurName struct {
							Text string `xml:",chardata"` // Truetzschler, Schaefer, B...
						} `xml:"surName"`
					} `xml:"name"`
				} `xml:"person"`
			} `xml:"creator"`
			Subject []struct {
				Text string `xml:",chardata"` // 004, 004, Parallelrechner...
				Type string `xml:"type,attr"`
			} `xml:"subject"`
			Abstract []struct {
				Text string `xml:",chardata"` // Es werden sowohl die Char...
				Type string `xml:"type,attr"`
				Lang string `xml:"lang,attr"`
				// Type string `xml:"type,attr"`
			} `xml:"abstract"`
			Publisher []struct {
				Text                    string `xml:",chardata"`
				Type                    string `xml:"type,attr"`
				UniversityOrInstitution struct {
					Text string `xml:",chardata"`
					Name struct {
						Text string `xml:",chardata"` // Universitätsbibliothek C...
					} `xml:"name"`
					Place struct {
						Text string `xml:",chardata"` // Chemnitz, Chemnitz, Chemn...
					} `xml:"place"`
				} `xml:"universityOrInstitution"`
				Address struct {
					Text   string `xml:",chardata"` // 09107 Chemnitz, 09107 Che...
					Scheme string `xml:"Scheme,attr"`
				} `xml:"address"`
			} `xml:"publisher"`
			Issued struct {
				Text string `xml:",chardata"` // 1995-06-20, 1995-06-20, 1...
				Type string `xml:"type,attr"`
			} `xml:"issued"`
			Type struct {
				Text string `xml:",chardata"` // StudyThesis, masterThesis...
				Type string `xml:"type,attr"`
			} `xml:"type"`
			Identifier struct {
				Text string `xml:",chardata"` // urn:nbn:de:bsz:ch1-199500...
				Type string `xml:"type,attr"`
				// Type string `xml:"type,attr"`
			} `xml:"identifier"`
			Language []struct {
				Text string `xml:",chardata"` // ger, ger, ger, ger, ger, ...
				Type string `xml:"type,attr"`
			} `xml:"language"`
			Degree struct {
				Text  string `xml:",chardata"`
				Level struct {
					Text string `xml:",chardata"` // other, Diplom, Diplom, ot...
				} `xml:"level"`
				Grantor []struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
					// Type                    string `xml:"type,attr"`
					CountryCode             string `xml:"countryCode,attr"`
					UniversityOrInstitution struct {
						Text string `xml:",chardata"`
						Name struct {
							Text string `xml:",chardata"` // TU Chemnitz, TU Chemnitz,...
						} `xml:"name"`
						Place struct {
							Text string `xml:",chardata"` // Chemnitz, Chemnitz, Chemn...
						} `xml:"place"`
						Department struct {
							Text string `xml:",chardata"`
							Name struct {
								Text string `xml:",chardata"` // Fakultät für Informatik...
							} `xml:"name"`
						} `xml:"department"`
					} `xml:"universityOrInstitution"`
				} `xml:"grantor"`
			} `xml:"degree"`
			Contact struct {
				Text      string `xml:",chardata"`
				ContactID string `xml:"contactID,attr"`
			} `xml:"contact"`
			FileNumber struct {
				Text string `xml:",chardata"` // 3, 3, 3, 3, 3, 3, 3, 3, 3...
			} `xml:"fileNumber"`
			FileProperties []struct {
				Text          string `xml:",chardata"`
				FileName      string `xml:"fileName,attr"`
				FileSize      string `xml:"fileSize,attr"`
				FileDirectory string `xml:"fileDirectory,attr"`
			} `xml:"fileProperties"`
			Transfer struct {
				Text string `xml:",chardata"` // http://www.qucosa.de/file...
				Type string `xml:"type,attr"`
			} `xml:"transfer"`
			Rights struct {
				Text string `xml:",chardata"` // Dokument ist für Print o...
				Kind string `xml:"kind,attr"`
			} `xml:"rights"`
			DateSubmitted struct {
				Text string `xml:",chardata"` // 1995-06-20, 1995-06-21, 1...
				Type string `xml:"type,attr"`
			} `xml:"dateSubmitted"`
			Source []struct {
				Text string `xml:",chardata"` // Engineering Management sp...
				Type string `xml:"type,attr"`
			} `xml:"source"`
			Contributor []struct {
				Text        string `xml:",chardata"`
				Role        string `xml:"role,attr"`
				Type        string `xml:"type,attr"`
				CountryCode string `xml:"countryCode,attr"`
				Person      struct {
					Text string `xml:",chardata"`
					Name struct {
						Text     string `xml:",chardata"`
						Type     string `xml:"type,attr"`
						ForeName struct {
							Text string `xml:",chardata"` // Christian, Christian., Di...
						} `xml:"foreName"`
						SurName struct {
							Text string `xml:",chardata"` // v. Borczyskowski, v. Borc...
						} `xml:"surName"`
					} `xml:"name"`
					AcademicTitle struct {
						Text string `xml:",chardata"` // Prof. Dr., Prof. Dr., Pro...
					} `xml:"academicTitle"`
				} `xml:"person"`
			} `xml:"contributor"`
			DateAccepted struct {
				Text string `xml:",chardata"` // 1997-07-03, 1998-12-14, 2...
				Type string `xml:"type,attr"`
			} `xml:"dateAccepted"`
			IsPartOf []struct {
				Text string `xml:",chardata"` // 1439-1597, http://nbn-res...
				Type string `xml:"type,attr"`
			} `xml:"isPartOf"`
			TableOfContents struct {
				Text string `xml:",chardata"` // Klassifizierung von P2P-A...
				Type string `xml:"type,attr"`
				Lang string `xml:"lang,attr"`
				// Type string `xml:"type,attr"`
			} `xml:"tableOfContents"`
			Replaces struct {
				Text string `xml:",chardata"` // http://nbn-resolving.de/u...
				Type string `xml:"type,attr"`
			} `xml:"replaces"`
			Alternative []struct {
				Text string `xml:",chardata"` // Das Vampir-Genre im Wande...
				Type string `xml:"type,attr"`
				Lang string `xml:"lang,attr"`
				// Type string `xml:"type,attr"`
			} `xml:"alternative"`
			Created struct {
				Text string `xml:",chardata"` // 2010, 2010, 2010, 2010, 2...
				Type string `xml:"type,attr"`
			} `xml:"created"`
		} `xml:"xMetaDiss"`
	} `xml:"metadata"`
	About struct {
		Text string `xml:",chardata"`
	} `xml:"about"`
}
