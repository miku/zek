package xmdissplus

import "encoding/xml"

type Record struct {
	XMLName xml.Name `xml:"Record"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text       string `xml:",chardata"`
		Status     string `xml:"status,attr"`
		Identifier struct {
			Text string `xml:",chardata"` // oai:hss-opus.ub.ruhr-uni-...
		} `xml:"identifier"`
		Datestamp struct {
			Text string `xml:",chardata"` // 2016-07-25, 2016-07-25, 2...
		} `xml:"datestamp"`
		SetSpec []struct {
			Text string `xml:",chardata"` // doc-type:doctoralthesis, ...
		} `xml:"setSpec"`
	} `xml:"header"`
	Metadata struct {
		Text      string `xml:",chardata"`
		XMetaDiss struct {
			Text           string `xml:",chardata"`
			XMetaDiss      string `xml:"xMetaDiss,attr"`
			Cc             string `xml:"cc,attr"`
			Dc             string `xml:"dc,attr"`
			Dcmitype       string `xml:"dcmitype,attr"`
			Dcterms        string `xml:"dcterms,attr"`
			Pc             string `xml:"pc,attr"`
			URN            string `xml:"urn,attr"`
			Hdl            string `xml:"hdl,attr"`
			Doi            string `xml:"doi,attr"`
			Thesis         string `xml:"thesis,attr"`
			Ddb            string `xml:"ddb,attr"`
			Dini           string `xml:"dini,attr"`
			Xmlns          string `xml:"xmlns,attr"`
			SchemaLocation string `xml:"schemaLocation,attr"`
			Title          []struct {
				Text string `xml:",chardata"` // Micromechanical modeling ...
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
							Text string `xml:",chardata"` // Satyapriya, Thomas, Matth...
						} `xml:"foreName"`
						SurName struct {
							Text string `xml:",chardata"` // Gupta, Tacke, Lübbert, H...
						} `xml:"surName"`
					} `xml:"name"`
					AcademicTitle struct {
						Text string `xml:",chardata"` // M. Tech., Dipl., M. Sc., ...
					} `xml:"academicTitle"`
				} `xml:"person"`
			} `xml:"creator"`
			Subject []struct {
				Text string `xml:",chardata"` // TRIP-Stahl, Korngrenze, P...
				Type string `xml:"type,attr"`
			} `xml:"subject"`
			Abstract []struct {
				Text string `xml:",chardata"` // Transformation induzierte...
				Type string `xml:"type,attr"`
				// Type string `xml:"type,attr"`
				Lang string `xml:"lang,attr"`
			} `xml:"abstract"`
			Publisher struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
				// Type                    string `xml:"type,attr"`
				UniversityOrInstitution struct {
					Text string `xml:",chardata"`
					Name struct {
						Text string `xml:",chardata"` // Ruhr-Universität Bochum,...
					} `xml:"name"`
					Place struct {
						Text string `xml:",chardata"` // Bochum, Bochum, Bochum, B...
					} `xml:"place"`
				} `xml:"universityOrInstitution"`
				Address struct {
					Text   string `xml:",chardata"` // Universitätsstr. 150 D-...
					Scheme string `xml:"Scheme,attr"`
				} `xml:"address"`
			} `xml:"publisher"`
			Contributor []struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
				// Type   string `xml:"type,attr"`
				Role   string `xml:"role,attr"`
				Person struct {
					Text string `xml:",chardata"`
					Name struct {
						Text     string `xml:",chardata"`
						Type     string `xml:"type,attr"`
						ForeName struct {
							Text string `xml:",chardata"` // Alexander, Dierk, Rainer,...
						} `xml:"foreName"`
						SurName struct {
							Text string `xml:",chardata"` // Hartmaier, Raabe, Grauer,...
						} `xml:"surName"`
					} `xml:"name"`
					AcademicTitle struct {
						Text string `xml:",chardata"` // Prof. Dr. rer. nat., Prof...
					} `xml:"academicTitle"`
				} `xml:"person"`
			} `xml:"contributor"`
			DateAccepted struct {
				Text string `xml:",chardata"` // 2015-02-17, 2015-01-29, 2...
				Type string `xml:"type,attr"`
			} `xml:"dateAccepted"`
			Issued struct {
				Text string `xml:",chardata"` // 2015-09-29, 2015-11-05, 2...
				Type string `xml:"type,attr"`
			} `xml:"issued"`
			Type struct {
				Text string `xml:",chardata"` // doctoralThesis, doctoralT...
				Type string `xml:"type,attr"`
			} `xml:"type"`
			Identifier struct {
				Text string `xml:",chardata"` // urn:nbn:de:hbz:294-43966,...
				Type string `xml:"type,attr"`
				// Type string `xml:"type,attr"`
			} `xml:"identifier"`
			Medium struct {
				Text string `xml:",chardata"` // application/pdf, applicat...
				Type string `xml:"type,attr"`
			} `xml:"medium"`
			Language struct {
				Text string `xml:",chardata"` // eng, eng, eng, eng, ger, ...
				Type string `xml:"type,attr"`
			} `xml:"language"`
			Rights []struct {
				Text string `xml:",chardata"` // Keine Creative Commons Li...
				Kind string `xml:"kind,attr"`
			} `xml:"rights"`
			Degree struct {
				Text  string `xml:",chardata"`
				Level struct {
					Text string `xml:",chardata"` // thesis.doctoral, thesis.d...
				} `xml:"level"`
				Grantor struct {
					Text                    string `xml:",chardata"`
					Type                    string `xml:"type,attr"`
					UniversityOrInstitution struct {
						Text string `xml:",chardata"`
						Name struct {
							Text string `xml:",chardata"` // Ruhr-Universität Bochum,...
						} `xml:"name"`
						Place struct {
							Text string `xml:",chardata"` // Bochum, Bochum, Bochum, B...
						} `xml:"place"`
						Department struct {
							Text string `xml:",chardata"`
							Name struct {
								Text string `xml:",chardata"` // Fakultät für Maschinenb...
							} `xml:"name"`
						} `xml:"department"`
					} `xml:"universityOrInstitution"`
				} `xml:"grantor"`
			} `xml:"degree"`
			FileNumber struct {
				Text string `xml:",chardata"` // 1, 1, 1, 1, 1, 1, 1, 1, 1...
			} `xml:"fileNumber"`
			FileProperties []struct {
				Text     string `xml:",chardata"`
				FileName string `xml:"fileName,attr"`
				FileSize string `xml:"fileSize,attr"`
				FileID   string `xml:"fileID,attr"`
			} `xml:"fileProperties"`
			Transfer struct {
				Text string `xml:",chardata"` // http://hss-opus.ub.ruhr-u...
				Type string `xml:"type,attr"`
			} `xml:"transfer"`
			IsPartOf struct {
				Text string `xml:",chardata"` // Schriftenreihe des Lehrst...
				Type string `xml:"type,attr"`
			} `xml:"isPartOf"`
			Source struct {
				Text string `xml:",chardata"` // Functional and structural...
				Type string `xml:"type,attr"`
			} `xml:"source"`
		} `xml:"xMetaDiss"`
	} `xml:"metadata"`
	About struct {
		Text string `xml:",chardata"`
	} `xml:"about"`
}
