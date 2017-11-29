package xmdiss

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
			Xmlns          string `xml:"xmlns,attr"`
			Cc             string `xml:"cc,attr"`
			Dc             string `xml:"dc,attr"`
			Dcmitype       string `xml:"dcmitype,attr"`
			Dcterms        string `xml:"dcterms,attr"`
			Pc             string `xml:"pc,attr"`
			URN            string `xml:"urn,attr"`
			Thesis         string `xml:"thesis,attr"`
			Ddb            string `xml:"ddb,attr"`
			SchemaLocation string `xml:"schemaLocation,attr"`
			Title          []struct {
				Text string `xml:",chardata"` // Micromechanical modeling ...
				Type string `xml:"type,attr"`
				Lang string `xml:"lang,attr"`
				// Type string `xml:"type,attr"`
			} `xml:"title"`
			Creator struct {
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
				} `xml:"person"`
			} `xml:"creator"`
			Subject []struct {
				Text string `xml:",chardata"` // TRIP-Stahl, Korngrenze, P...
				Type string `xml:"type,attr"`
			} `xml:"subject"`
			Abstract []struct {
				Text string `xml:",chardata"` // Transformation induzierte...
				Type string `xml:"type,attr"`
				Lang string `xml:"lang,attr"`
				// Type string `xml:"type,attr"`
			} `xml:"abstract"`
			Publisher struct {
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
				} `xml:"universityOrInstitution"`
				Address struct {
					Text   string `xml:",chardata"` // Universitätsstr. 150 D-...
					Scheme string `xml:"Scheme,attr"`
				} `xml:"address"`
			} `xml:"publisher"`
			Contributor []struct {
				Text        string `xml:",chardata"`
				Type        string `xml:"type,attr"`
				Role        string `xml:"role,attr"`
				CountryCode string `xml:"countryCode,attr"`
				Person      struct {
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
			Type struct {
				Text string `xml:",chardata"` // ElectronicThesisandDisser...
				Type string `xml:"type,attr"`
			} `xml:"type"`
			Identifier struct {
				Text string `xml:",chardata"` // urn:nbn:de:hbz:294-43966,...
				Type string `xml:"type,attr"`
			} `xml:"identifier"`
			Medium struct {
				Text string `xml:",chardata"` // application/pdf, applicat...
				Type string `xml:"type,attr"`
			} `xml:"medium"`
			Language struct {
				Text string `xml:",chardata"` // eng, eng, eng, eng, ger, ...
				Type string `xml:"type,attr"`
			} `xml:"language"`
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
								Text string `xml:",chardata"` // Ruhr-Universität Bochum,...
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
				Text string `xml:",chardata"` // 1, 1, 1, 1, 1, 1, 1, 1, 1...
			} `xml:"fileNumber"`
			FileProperties []struct {
				Text          string `xml:",chardata"` // aus:Praesentationsformat,...
				FileName      string `xml:"fileName,attr"`
				FileID        string `xml:"fileID,attr"`
				FileSize      string `xml:"fileSize,attr"`
				FileDirectory string `xml:"fileDirectory,attr"`
			} `xml:"fileProperties"`
			Rights struct {
				Text string `xml:",chardata"`
				Kind string `xml:"kind,attr"`
			} `xml:"rights"`
		} `xml:"xMetaDiss"`
	} `xml:"metadata"`
	About struct {
		Text string `xml:",chardata"`
	} `xml:"about"`
}
