package mets

import "encoding/xml"

type Record struct {
	XMLName xml.Name `xml:"Record"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text       string `xml:",chardata"`
		Status     string `xml:"status,attr"`
		Identifier struct {
			Text string `xml:",chardata"` // oai:generic.eprints.org:3...
		} `xml:"identifier"`
		Datestamp struct {
			Text string `xml:",chardata"` // 2010-01-02T04:42:29Z, 201...
		} `xml:"datestamp"`
		SetSpec []struct {
			Text string `xml:",chardata"` // 7374617475733D707562, 737...
		} `xml:"setSpec"`
	} `xml:"header"`
	Metadata struct {
		Text string `xml:",chardata"`
		Mets struct {
			Text           string `xml:",chardata"`
			OBJID          string `xml:"OBJID,attr"`
			Xsi            string `xml:"xsi,attr"`
			Mods           string `xml:"mods,attr"`
			LABEL          string `xml:"LABEL,attr"`
			SchemaLocation string `xml:"schemaLocation,attr"`
			Xlink          string `xml:"xlink,attr"`
			Mets           string `xml:"mets,attr"`
			MetsHdr        struct {
				Text       string `xml:",chardata"`
				CREATEDATA string `xml:"CREATEDATA,attr"`
				Agent      struct {
					Text string `xml:",chardata"`
					TYPE string `xml:"TYPE,attr"`
					ROLE string `xml:"ROLE,attr"`
					Name struct {
						Text string `xml:",chardata"` // Eprints@MDRF, Eprints@MDR...
					} `xml:"name"`
				} `xml:"agent"`
			} `xml:"metsHdr"`
			DmdSec struct {
				Text   string `xml:",chardata"`
				ID     string `xml:"ID,attr"`
				MdWrap struct {
					Text    string `xml:",chardata"`
					MDTYPE  string `xml:"MDTYPE,attr"`
					XmlData struct {
						Text      string `xml:",chardata"`
						TitleInfo struct {
							Text  string `xml:",chardata"`
							Title struct {
								Text string `xml:",chardata"` // Changes in plasma growth ...
							} `xml:"title"`
						} `xml:"titleInfo"`
						Name []struct {
							Text     string `xml:",chardata"`
							Type     string `xml:"type,attr"`
							NamePart []struct {
								Text string `xml:",chardata"` // P S, Sharp, V, Mohan, F, ...
								Type string `xml:"type,attr"`
							} `xml:"namePart"`
							Role struct {
								Text     string `xml:",chardata"`
								RoleTerm struct {
									Text string `xml:",chardata"` // author, author, author, a...
									Type string `xml:"type,attr"`
								} `xml:"roleTerm"`
							} `xml:"role"`
						} `xml:"name"`
						Abstract struct {
							Text string `xml:",chardata"` // A group of 22 newly diagn...
						} `xml:"abstract"`
						Classification []struct {
							Text      string `xml:",chardata"` // Genetics and Diabetes, Di...
							Authority string `xml:"authority,attr"`
						} `xml:"classification"`
						OriginInfo []struct {
							Text       string `xml:",chardata"`
							DateIssued struct {
								Text     string `xml:",chardata"` // 1987, 1987, 1987, 1987, 1...
								Encoding string `xml:"encoding,attr"`
							} `xml:"dateIssued"`
							Publisher struct {
								Text string `xml:",chardata"` // John Wiley & Sons, Americ...
							} `xml:"publisher"`
						} `xml:"originInfo"`
						Genre struct {
							Text string `xml:",chardata"` // Article, Article, Article...
						} `xml:"genre"`
					} `xml:"xmlData"`
				} `xml:"mdWrap"`
			} `xml:"dmdSec"`
			AmdSec struct {
				Text     string `xml:",chardata"`
				ID       string `xml:"ID,attr"`
				RightsMD struct {
					Text   string `xml:",chardata"`
					ID     string `xml:"ID,attr"`
					MdWrap struct {
						Text    string `xml:",chardata"`
						MDTYPE  string `xml:"MDTYPE,attr"`
						XmlData struct {
							Text               string `xml:",chardata"`
							UseAndReproduction struct {
								Text string `xml:",chardata"`
								P    []struct {
									Text   string `xml:",chardata"` // In self-archiving this co...
									Strong struct {
										Text string `xml:",chardata"` // For work being deposited ...
									} `xml:"strong"`
								} `xml:"p"`
							} `xml:"useAndReproduction"`
						} `xml:"xmlData"`
					} `xml:"mdWrap"`
				} `xml:"rightsMD"`
			} `xml:"amdSec"`
			FileSec struct {
				Text    string `xml:",chardata"`
				FileGrp []struct {
					Text string `xml:",chardata"`
					USE  string `xml:"USE,attr"`
					File struct {
						Text     string `xml:",chardata"`
						SIZE     string `xml:"SIZE,attr"`
						ID       string `xml:"ID,attr"`
						MIMETYPE string `xml:"MIMETYPE,attr"`
						OWNERID  string `xml:"OWNERID,attr"`
						FLocat   struct {
							Text    string `xml:",chardata"`
							LOCTYPE string `xml:"LOCTYPE,attr"`
							Href    string `xml:"href,attr"`
							Type    string `xml:"type,attr"`
						} `xml:"FLocat"`
					} `xml:"file"`
				} `xml:"fileGrp"`
			} `xml:"fileSec"`
			StructMap struct {
				Text string `xml:",chardata"`
				Div  struct {
					Text  string `xml:",chardata"`
					DMDID string `xml:"DMDID,attr"`
					AMDID string `xml:"AMDID,attr"`
					Fptr  []struct {
						Text   string `xml:",chardata"`
						FILEID string `xml:"FILEID,attr"`
					} `xml:"fptr"`
				} `xml:"div"`
			} `xml:"structMap"`
		} `xml:"mets"`
	} `xml:"metadata"`
	About struct {
		Text string `xml:",chardata"`
	} `xml:"about"`
}
