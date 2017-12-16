package oaf

import "encoding/xml"

// Record was generated 2017-12-16 13:32:04 by tir on hayiti.
type Record struct {
	XMLName xml.Name `xml:"record"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text       string `xml:",chardata"`
		Status     string `xml:"status,attr"`
		Identifier struct {
			Text string `xml:",chardata"` // oai:dnet:od________65::0f...
		} `xml:"identifier"`
		Datestamp struct {
			Text string `xml:",chardata"` // 2017-12-12T10:09:21Z, 201...
		} `xml:"datestamp"`
		SetSpec []struct {
			Text string `xml:",chardata"` // CERN_Document_Server, CER...
		} `xml:"setSpec"`
	} `xml:"header"`
	Metadata struct {
		Text   string `xml:",chardata"`
		Entity struct {
			Text           string `xml:",chardata"`
			Oaf            string `xml:"oaf,attr"`
			Xsi            string `xml:"xsi,attr"`
			SchemaLocation string `xml:"schemaLocation,attr"`
			Result         struct {
				Text    string `xml:",chardata"`
				Journal struct {
					Text  string `xml:",chardata"` // R.E.M. (Rekayasa Energi M...
					ISSN  string `xml:"issn,attr"`
					EISSN string `xml:"eissn,attr"`
					Lissn string `xml:"lissn,attr"`
					Ep    string `xml:"ep,attr"`
					Iss   string `xml:"iss,attr"`
					Sp    string `xml:"sp,attr"`
					Vol   string `xml:"vol,attr"`
				} `xml:"journal"`
				Description []struct {
					Text string `xml:",chardata"` // In this paper we present ...
				} `xml:"description"`
				Title []struct {
					Text       string `xml:",chardata"` // D-Rank: a framework for s...
					Classid    string `xml:"classid,attr"`
					Classname  string `xml:"classname,attr"`
					Schemeid   string `xml:"schemeid,attr"`
					Schemename string `xml:"schemename,attr"`
				} `xml:"title"`
				Subject []struct {
					Text                string `xml:",chardata"` // Computing and Computers, ...
					Classid             string `xml:"classid,attr"`
					Classname           string `xml:"classname,attr"`
					Schemeid            string `xml:"schemeid,attr"`
					Schemename          string `xml:"schemename,attr"`
					Inferred            string `xml:"inferred,attr"`
					Inferenceprovenance string `xml:"inferenceprovenance,attr"`
					Provenanceaction    string `xml:"provenanceaction,attr"`
					Trust               string `xml:"trust,attr"`
				} `xml:"subject"`
				Embargoenddate struct {
					Text string `xml:",chardata"` // info:eu-repo/date/publica...
				} `xml:"embargoenddate"`
				Language struct {
					Text       string `xml:",chardata"`
					Classid    string `xml:"classid,attr"`
					Classname  string `xml:"classname,attr"`
					Schemeid   string `xml:"schemeid,attr"`
					Schemename string `xml:"schemename,attr"`
				} `xml:"language"`
				Resulttype struct {
					Text       string `xml:",chardata"`
					Classid    string `xml:"classid,attr"`
					Classname  string `xml:"classname,attr"`
					Schemeid   string `xml:"schemeid,attr"`
					Schemename string `xml:"schemename,attr"`
				} `xml:"resulttype"`
				Dateofacceptance struct {
					Text string `xml:",chardata"` // 2010-11-10, 2013-01-01, 2...
				} `xml:"dateofacceptance"`
				Country struct {
					Text       string `xml:",chardata"`
					Classid    string `xml:"classid,attr"`
					Classname  string `xml:"classname,attr"`
					Schemeid   string `xml:"schemeid,attr"`
					Schemename string `xml:"schemename,attr"`
				} `xml:"country"`
				Relevantdate []struct {
					Text       string `xml:",chardata"` // 1971/2100, 2011-10-14, 20...
					Classid    string `xml:"classid,attr"`
					Classname  string `xml:"classname,attr"`
					Schemeid   string `xml:"schemeid,attr"`
					Schemename string `xml:"schemename,attr"`
				} `xml:"relevantdate"`
				Publisher struct {
					Text string `xml:",chardata"` // Marseille, CPT, SISSA, CE...
				} `xml:"publisher"`
				Source []struct {
					Text string `xml:",chardata"` // Nucl. Instrum. Methods Ph...
				} `xml:"source"`
				Fulltext []struct {
					Text string `xml:",chardata"` // http://cogprints.org/4009...
				} `xml:"fulltext"`
				Format []struct {
					Text string `xml:",chardata"` // text/html, application/pd...
				} `xml:"format"`
				Contributor []struct {
					Text string `xml:",chardata"` // Jones, Byron C., Mormède...
				} `xml:"contributor"`
				Storagedate struct {
					Text string `xml:",chardata"` // 2017-10-20, 2017-10-20, 2...
				} `xml:"storagedate"`
				Resourcetype struct {
					Text       string `xml:",chardata"`
					Classid    string `xml:"classid,attr"`
					Classname  string `xml:"classname,attr"`
					Schemeid   string `xml:"schemeid,attr"`
					Schemename string `xml:"schemename,attr"`
				} `xml:"resourcetype"`
				Device struct {
					Text string `xml:",chardata"`
				} `xml:"device"`
				Size struct {
					Text string `xml:",chardata"` // PEARL_High-PressureData_O...
				} `xml:"size"`
				Version struct {
					Text string `xml:",chardata"` // 1, 1, 1, 1, 1, 1, 1, 1, 1...
				} `xml:"version"`
				Lastmetadataupdate struct {
					Text string `xml:",chardata"` // 2015-12-02, 2016-10-04, 2...
				} `xml:"lastmetadataupdate"`
				Metadataversionnumber struct {
					Text string `xml:",chardata"`
				} `xml:"metadataversionnumber"`
				DocumentationUrl struct {
					Text string `xml:",chardata"`
				} `xml:"documentationUrl"`
				License struct {
					Text       string `xml:",chardata"`
					Classid    string `xml:"classid,attr"`
					Classname  string `xml:"classname,attr"`
					Schemeid   string `xml:"schemeid,attr"`
					Schemename string `xml:"schemename,attr"`
				} `xml:"license"`
				Type struct {
					Text string `xml:",chardata"`
				} `xml:"type"`
				CodeRepositoryUrl struct {
					Text string `xml:",chardata"`
				} `xml:"codeRepositoryUrl"`
				ProgrammingLanguage struct {
					Text       string `xml:",chardata"`
					Classid    string `xml:"classid,attr"`
					Classname  string `xml:"classname,attr"`
					Schemeid   string `xml:"schemeid,attr"`
					Schemename string `xml:"schemename,attr"`
				} `xml:"programmingLanguage"`
				OriginalId []struct {
					Text string `xml:",chardata"` // oai:cds.cern.ch:1306230, ...
				} `xml:"originalId"`
				Collectedfrom []struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
					ID   string `xml:"id,attr"`
				} `xml:"collectedfrom"`
				Pid []struct {
					Text       string `xml:",chardata"` // 10.1088/0031-8949/2013/T1...
					Classid    string `xml:"classid,attr"`
					Classname  string `xml:"classname,attr"`
					Schemeid   string `xml:"schemeid,attr"`
					Schemename string `xml:"schemename,attr"`
				} `xml:"pid"`
				Bestlicense struct {
					Text       string `xml:",chardata"`
					Classid    string `xml:"classid,attr"`
					Classname  string `xml:"classname,attr"`
					Schemeid   string `xml:"schemeid,attr"`
					Schemename string `xml:"schemename,attr"`
				} `xml:"bestlicense"`
				Datainfo struct {
					Text     string `xml:",chardata"`
					Inferred struct {
						Text string `xml:",chardata"` // false, false, false, fals...
					} `xml:"inferred"`
					Deletedbyinference struct {
						Text string `xml:",chardata"` // false, false, false, fals...
					} `xml:"deletedbyinference"`
					Trust struct {
						Text string `xml:",chardata"` // 0.9, 0.9, 0.9, 0.9, 0.9, ...
					} `xml:"trust"`
					Inferenceprovenance struct {
						Text string `xml:",chardata"` // dedup-similarity-result-l...
					} `xml:"inferenceprovenance"`
					Provenanceaction struct {
						Text       string `xml:",chardata"`
						Classid    string `xml:"classid,attr"`
						Classname  string `xml:"classname,attr"`
						Schemeid   string `xml:"schemeid,attr"`
						Schemename string `xml:"schemename,attr"`
					} `xml:"provenanceaction"`
				} `xml:"datainfo"`
				Rels struct {
					Text string `xml:",chardata"`
					Rel  []struct {
						Text                string `xml:",chardata"`
						Inferred            string `xml:"inferred,attr"`
						Trust               string `xml:"trust,attr"`
						Inferenceprovenance string `xml:"inferenceprovenance,attr"`
						Provenanceaction    string `xml:"provenanceaction,attr"`
						To                  struct {
							Text   string `xml:",chardata"` // dedup_wf_001::c8d4a41a905...
							Class  string `xml:"class,attr"`
							Scheme string `xml:"scheme,attr"`
							Type   string `xml:"type,attr"`
						} `xml:"to"`
						Ranking struct {
							Text string `xml:",chardata"` // 1, 5, 2, 3, 4, 1, 1, 1, 2...
						} `xml:"ranking"`
						Fullname struct {
							Text string `xml:",chardata"` // Veselý, M., Caffaro, Jer...
						} `xml:"fullname"`
						Resulttype struct {
							Text       string `xml:",chardata"`
							Classid    string `xml:"classid,attr"`
							Classname  string `xml:"classname,attr"`
							Schemeid   string `xml:"schemeid,attr"`
							Schemename string `xml:"schemename,attr"`
						} `xml:"resulttype"`
						Similarity struct {
							Text string `xml:",chardata"` // 0.70587885, 0.70587885, 0...
						} `xml:"similarity"`
						Type struct {
							Text string `xml:",chardata"` // STANDARD, STANDARD, STAND...
						} `xml:"type"`
						Dateofacceptance []struct {
							Text string `xml:",chardata"` // 1973-01-01, 1973-01-01, 1...
						} `xml:"dateofacceptance"`
						Title []struct {
							Text       string `xml:",chardata"` // The gauge properties of t...
							Classid    string `xml:"classid,attr"`
							Classname  string `xml:"classname,attr"`
							Schemeid   string `xml:"schemeid,attr"`
							Schemename string `xml:"schemename,attr"`
						} `xml:"title"`
						Publisher struct {
							Text string `xml:",chardata"` // Elsevier BV, Elsevier BV,...
						} `xml:"publisher"`
						Funding []struct {
							Text   string `xml:",chardata"`
							Funder struct {
								Text         string `xml:",chardata"`
								ID           string `xml:"id,attr"`
								Shortname    string `xml:"shortname,attr"`
								Name         string `xml:"name,attr"`
								Jurisdiction string `xml:"jurisdiction,attr"`
							} `xml:"funder"`
							FundingLevel0 struct {
								Text string `xml:",chardata"` // ec__________::EC::FP7, ec...
								Name string `xml:"name,attr"`
							} `xml:"funding_level_0"`
							FundingLevel1 struct {
								Text string `xml:",chardata"` // ec__________::EC::FP7::SP...
								Name string `xml:"name,attr"`
							} `xml:"funding_level_1"`
							FundingLevel2 struct {
								Text string `xml:",chardata"` // ec__________::EC::FP7::SP...
								Name string `xml:"name,attr"`
							} `xml:"funding_level_2"`
						} `xml:"funding"`
						Code struct {
							Text string `xml:",chardata"` // 237920, 262025, 227579, 6...
						} `xml:"code"`
						Contracttype struct {
							Text       string `xml:",chardata"`
							Classid    string `xml:"classid,attr"`
							Classname  string `xml:"classname,attr"`
							Schemeid   string `xml:"schemeid,attr"`
							Schemename string `xml:"schemename,attr"`
						} `xml:"contracttype"`
						Acronym struct {
							Text string `xml:",chardata"` // UNILHC, AIDA, EUCARD, BEA...
						} `xml:"acronym"`
						Websiteurl struct {
							Text string `xml:",chardata"` // http://www.symbiosis-eu.n...
						} `xml:"websiteurl"`
					} `xml:"rel"`
				} `xml:"rels"`
				Children struct {
					Text     string `xml:",chardata"`
					Instance []struct {
						Text     string `xml:",chardata"`
						ID       string `xml:"id,attr"`
						Hostedby struct {
							Text string `xml:",chardata"`
							Name string `xml:"name,attr"`
							ID   string `xml:"id,attr"`
						} `xml:"hostedby"`
						Dateofacceptance struct {
							Text string `xml:",chardata"` // 2010-11-10, 2013-01-01, 2...
						} `xml:"dateofacceptance"`
						Licence struct {
							Text       string `xml:",chardata"`
							Classid    string `xml:"classid,attr"`
							Classname  string `xml:"classname,attr"`
							Schemeid   string `xml:"schemeid,attr"`
							Schemename string `xml:"schemename,attr"`
						} `xml:"licence"`
						Collectedfrom struct {
							Text string `xml:",chardata"`
							Name string `xml:"name,attr"`
							ID   string `xml:"id,attr"`
						} `xml:"collectedfrom"`
						Instancetype struct {
							Text       string `xml:",chardata"`
							Classid    string `xml:"classid,attr"`
							Classname  string `xml:"classname,attr"`
							Schemeid   string `xml:"schemeid,attr"`
							Schemename string `xml:"schemename,attr"`
						} `xml:"instancetype"`
						Webresource []struct {
							Text string `xml:",chardata"`
							URL  struct {
								Text string `xml:",chardata"` // http://cds.cern.ch/record...
							} `xml:"url"`
						} `xml:"webresource"`
					} `xml:"instance"`
					Externalreference []struct {
						Text     string `xml:",chardata"`
						Sitename struct {
							Text string `xml:",chardata"` // SourceForge, GitHub, Prot...
						} `xml:"sitename"`
						URL struct {
							Text string `xml:",chardata"` // http://sourceforge.net/pr...
						} `xml:"url"`
						Qualifier struct {
							Text       string `xml:",chardata"`
							Classid    string `xml:"classid,attr"`
							Classname  string `xml:"classname,attr"`
							Schemeid   string `xml:"schemeid,attr"`
							Schemename string `xml:"schemename,attr"`
						} `xml:"qualifier"`
						Refidentifier struct {
							Text string `xml:",chardata"` // 1gt0, 1o4x, 1f45, 3d85, 3...
						} `xml:"refidentifier"`
					} `xml:"externalreference"`
					Result []struct {
						Text          string `xml:",chardata"`
						Objidentifier string `xml:"objidentifier,attr"`
						Title         []struct {
							Text       string `xml:",chardata"` // Disorder, Promiscuity, an...
							Classid    string `xml:"classid,attr"`
							Classname  string `xml:"classname,attr"`
							Schemeid   string `xml:"schemeid,attr"`
							Schemename string `xml:"schemename,attr"`
						} `xml:"title"`
						Publisher struct {
							Text string `xml:",chardata"` // Elsevier BV, Hindawi Publ...
						} `xml:"publisher"`
						Dateofacceptance struct {
							Text string `xml:",chardata"` // 2009-07-01, 2009-07-10, 2...
						} `xml:"dateofacceptance"`
						Resulttype struct {
							Text       string `xml:",chardata"`
							Classid    string `xml:"classid,attr"`
							Classname  string `xml:"classname,attr"`
							Schemeid   string `xml:"schemeid,attr"`
							Schemename string `xml:"schemename,attr"`
						} `xml:"resulttype"`
					} `xml:"result"`
				} `xml:"children"`
				Context []struct {
					Text     string `xml:",chardata"`
					ID       string `xml:"id,attr"`
					Label    string `xml:"label,attr"`
					Type     string `xml:"type,attr"`
					Category []struct {
						Text    string `xml:",chardata"`
						ID      string `xml:"id,attr"`
						Label   string `xml:"label,attr"`
						Concept []struct {
							Text  string `xml:",chardata"`
							ID    string `xml:"id,attr"`
							Label string `xml:"label,attr"`
						} `xml:"concept"`
					} `xml:"category"`
				} `xml:"context"`
			} `xml:"result"`
			ExtraInfo struct {
				Text       string `xml:",chardata"`
				Name       string `xml:"name,attr"`
				Typology   string `xml:"typology,attr"`
				Provenance string `xml:"provenance,attr"`
				Trust      string `xml:"trust,attr"`
				Citations  struct {
					Text     string `xml:",chardata"`
					Citation []struct {
						Text     string `xml:",chardata"`
						Position string `xml:"position,attr"`
						RawText  struct {
							Text string `xml:",chardata"` // G,Q; I de dybere Lag~ und...
						} `xml:"rawText"`
						ID []struct {
							Text            string `xml:",chardata"`
							Value           string `xml:"value,attr"`
							Type            string `xml:"type,attr"`
							ConfidenceLevel string `xml:"confidenceLevel,attr"`
						} `xml:"id"`
					} `xml:"citation"`
				} `xml:"citations"`
			} `xml:"extraInfo"`
			Datasource struct {
				Text             string `xml:",chardata"`
				Datasourcetypeui struct {
					Text       string `xml:",chardata"`
					Classid    string `xml:"classid,attr"`
					Classname  string `xml:"classname,attr"`
					Schemeid   string `xml:"schemeid,attr"`
					Schemename string `xml:"schemename,attr"`
				} `xml:"datasourcetypeui"`
				Datasourcetype struct {
					Text       string `xml:",chardata"`
					Classid    string `xml:"classid,attr"`
					Classname  string `xml:"classname,attr"`
					Schemeid   string `xml:"schemeid,attr"`
					Schemename string `xml:"schemename,attr"`
				} `xml:"datasourcetype"`
				Odcontenttypes []struct {
					Text string `xml:",chardata"` // Journal articles, Journal...
				} `xml:"odcontenttypes"`
				Openairecompatibility struct {
					Text       string `xml:",chardata"`
					Classid    string `xml:"classid,attr"`
					Classname  string `xml:"classname,attr"`
					Schemeid   string `xml:"schemeid,attr"`
					Schemename string `xml:"schemename,attr"`
				} `xml:"openairecompatibility"`
				Englishname struct {
					Text string `xml:",chardata"` // Drinking Water Engineerin...
				} `xml:"englishname"`
				Officialname struct {
					Text string `xml:",chardata"` // Drinking Water Engineerin...
				} `xml:"officialname"`
				Namespaceprefix struct {
					Text string `xml:",chardata"` // ojs_19969465, doaj1076980...
				} `xml:"namespaceprefix"`
				Latitude struct {
					Text string `xml:",chardata"` // 0.0, 0.0, 0.0, 0.0, 0.0, ...
				} `xml:"latitude"`
				Longitude struct {
					Text string `xml:",chardata"` // 0.0, 0.0, 0.0, 0.0, 0.0, ...
				} `xml:"longitude"`
				Websiteurl struct {
					Text string `xml:",chardata"` // http://www.drinking-water...
				} `xml:"websiteurl"`
				Logourl struct {
					Text string `xml:",chardata"` // http://www.uran.net.ua/im...
				} `xml:"logourl"`
				Contactemail struct {
					Text string `xml:",chardata"` // glavcheva@khpi.edu.ua, ch...
				} `xml:"contactemail"`
				Dateofvalidation struct {
					Text string `xml:",chardata"` // 2015-02-05, 2015-11-18, 2...
				} `xml:"dateofvalidation"`
				Description struct {
					Text string `xml:",chardata"` // mathematics, exercise phy...
				} `xml:"description"`
				Subjects []struct {
					Text       string `xml:",chardata"` // Business and Economics, L...
					Classid    string `xml:"classid,attr"`
					Classname  string `xml:"classname,attr"`
					Schemeid   string `xml:"schemeid,attr"`
					Schemename string `xml:"schemename,attr"`
				} `xml:"subjects"`
				Odnumberofitems struct {
					Text string `xml:",chardata"` // 171, 531, 13054, 1301, 26...
				} `xml:"odnumberofitems"`
				Odnumberofitemsdate struct {
					Text string `xml:",chardata"` // 2014-03-11, 2017-11-30, 2...
				} `xml:"odnumberofitemsdate"`
				Odpolicies struct {
					Text string `xml:",chardata"`
				} `xml:"odpolicies"`
				Odlanguages []struct {
					Text string `xml:",chardata"` // eng, eng, eng, eng, Engli...
				} `xml:"odlanguages"`
				Accessinfopackage []struct {
					Text string `xml:",chardata"` // http://eee.khpi.edu.ua/in...
				} `xml:"accessinfopackage"`
				Releasestartdate struct {
					Text string `xml:",chardata"` // 2015-08-09
				} `xml:"releasestartdate"`
				Releaseenddate struct {
					Text string `xml:",chardata"`
				} `xml:"releaseenddate"`
				Missionstatementurl struct {
					Text string `xml:",chardata"` // http://sldr.org/index.php...
				} `xml:"missionstatementurl"`
				Dataprovider struct {
					Text string `xml:",chardata"` // false, false, true, true,...
				} `xml:"dataprovider"`
				Serviceprovider struct {
					Text string `xml:",chardata"` // false, false, false, true...
				} `xml:"serviceprovider"`
				Databaseaccesstype struct {
					Text string `xml:",chardata"` // open, open, open, open, o...
				} `xml:"databaseaccesstype"`
				Datauploadtype struct {
					Text string `xml:",chardata"` // restricted, restricted, r...
				} `xml:"datauploadtype"`
				Databaseaccessrestriction struct {
					Text string `xml:",chardata"` // registration
				} `xml:"databaseaccessrestriction"`
				Datauploadrestriction struct {
					Text string `xml:",chardata"` // registration, registratio...
				} `xml:"datauploadrestriction"`
				Versioning struct {
					Text string `xml:",chardata"` // false, false, false, fals...
				} `xml:"versioning"`
				Citationguidelineurl struct {
					Text string `xml:",chardata"` // http://libguides.anu.edu....
				} `xml:"citationguidelineurl"`
				Qualitymanagementkind struct {
					Text string `xml:",chardata"` // unknown, yes, yes, yes, y...
				} `xml:"qualitymanagementkind"`
				Pidsystems struct {
					Text string `xml:",chardata"` // none, hdl, none, DOI, DOI...
				} `xml:"pidsystems"`
				Certificates struct {
					Text string `xml:",chardata"`
				} `xml:"certificates"`
				Policies []struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
					ID   string `xml:"id,attr"`
				} `xml:"policies"`
				Collectedfrom struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
					ID   string `xml:"id,attr"`
				} `xml:"collectedfrom"`
				OriginalId []struct {
					Text string `xml:",chardata"` // copernicuspu::1996-9465, ...
				} `xml:"originalId"`
				Pid struct {
					Text       string `xml:",chardata"`
					Classid    string `xml:"classid,attr"`
					Classname  string `xml:"classname,attr"`
					Schemeid   string `xml:"schemeid,attr"`
					Schemename string `xml:"schemename,attr"`
				} `xml:"pid"`
				Datainfo struct {
					Text     string `xml:",chardata"`
					Inferred struct {
						Text string `xml:",chardata"` // false, false, false, fals...
					} `xml:"inferred"`
					Deletedbyinference struct {
						Text string `xml:",chardata"` // false, false, false, fals...
					} `xml:"deletedbyinference"`
					Trust struct {
						Text string `xml:",chardata"` // 0.9, 0.9, 0.9, 0.9, 0.9, ...
					} `xml:"trust"`
					Inferenceprovenance struct {
						Text string `xml:",chardata"`
					} `xml:"inferenceprovenance"`
					Provenanceaction struct {
						Text       string `xml:",chardata"`
						Classid    string `xml:"classid,attr"`
						Classname  string `xml:"classname,attr"`
						Schemeid   string `xml:"schemeid,attr"`
						Schemename string `xml:"schemename,attr"`
					} `xml:"provenanceaction"`
				} `xml:"datainfo"`
				Rels struct {
					Text string `xml:",chardata"`
					Rel  []struct {
						Text                string `xml:",chardata"`
						Inferred            string `xml:"inferred,attr"`
						Trust               string `xml:"trust,attr"`
						Inferenceprovenance string `xml:"inferenceprovenance,attr"`
						Provenanceaction    string `xml:"provenanceaction,attr"`
						To                  struct {
							Text   string `xml:",chardata"` // openaire____::9709c677b07...
							Class  string `xml:"class,attr"`
							Scheme string `xml:"scheme,attr"`
							Type   string `xml:"type,attr"`
						} `xml:"to"`
						Country struct {
							Text       string `xml:",chardata"`
							Classid    string `xml:"classid,attr"`
							Classname  string `xml:"classname,attr"`
							Schemeid   string `xml:"schemeid,attr"`
							Schemename string `xml:"schemename,attr"`
						} `xml:"country"`
						Legalname struct {
							Text string `xml:",chardata"` // Ubiquity Press, Politihø...
						} `xml:"legalname"`
						Websiteurl struct {
							Text string `xml:",chardata"` // http://www.phs.no/, http:...
						} `xml:"websiteurl"`
						Legalshortname struct {
							Text string `xml:",chardata"` // INERIS, KU, ULP, Universi...
						} `xml:"legalshortname"`
					} `xml:"rel"`
				} `xml:"rels"`
				Children struct {
					Text string `xml:",chardata"`
				} `xml:"children"`
			} `xml:"datasource"`
			Organization struct {
				Text       string `xml:",chardata"`
				Websiteurl struct {
					Text string `xml:",chardata"` // http://www.dainst.de, htt...
				} `xml:"websiteurl"`
				Legalshortname struct {
					Text string `xml:",chardata"` // DAI, LINKSPACE, HOMAG, NO...
				} `xml:"legalshortname"`
				Eclegalbody struct {
					Text string `xml:",chardata"` // true, false, false, true,...
				} `xml:"eclegalbody"`
				Country struct {
					Text       string `xml:",chardata"`
					Classid    string `xml:"classid,attr"`
					Classname  string `xml:"classname,attr"`
					Schemeid   string `xml:"schemeid,attr"`
					Schemename string `xml:"schemename,attr"`
				} `xml:"country"`
				Ecnonprofit struct {
					Text string `xml:",chardata"` // true, false, false, true,...
				} `xml:"ecnonprofit"`
				Ecnutscode struct {
					Text string `xml:",chardata"` // false, false, false, fals...
				} `xml:"ecnutscode"`
				Echighereducation struct {
					Text string `xml:",chardata"` // false, false, false, fals...
				} `xml:"echighereducation"`
				Ecenterprise struct {
					Text string `xml:",chardata"` // false, true, true, false,...
				} `xml:"ecenterprise"`
				Eclegalperson struct {
					Text string `xml:",chardata"` // true, true, true, true, t...
				} `xml:"eclegalperson"`
				Ecresearchorganization struct {
					Text string `xml:",chardata"` // true, false, false, true,...
				} `xml:"ecresearchorganization"`
				Legalname struct {
					Text string `xml:",chardata"` // DEUTSCHES ARCHAOLOGISCHES...
				} `xml:"legalname"`
				Ecinternationalorganizationeurinterests struct {
					Text string `xml:",chardata"` // false, false, false, fals...
				} `xml:"ecinternationalorganizationeurinterests"`
				Ecinternationalorganization struct {
					Text string `xml:",chardata"` // false, false, false, fals...
				} `xml:"ecinternationalorganization"`
				Logourl struct {
					Text string `xml:",chardata"`
				} `xml:"logourl"`
				Ecsmevalidated struct {
					Text string `xml:",chardata"` // true, false, true, true, ...
				} `xml:"ecsmevalidated"`
				Collectedfrom []struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
					ID   string `xml:"id,attr"`
				} `xml:"collectedfrom"`
				OriginalId []struct {
					Text string `xml:",chardata"` // corda_______::998543415, ...
				} `xml:"originalId"`
				Pid struct {
					Text       string `xml:",chardata"`
					Classid    string `xml:"classid,attr"`
					Classname  string `xml:"classname,attr"`
					Schemeid   string `xml:"schemeid,attr"`
					Schemename string `xml:"schemename,attr"`
				} `xml:"pid"`
				Datainfo struct {
					Text     string `xml:",chardata"`
					Inferred struct {
						Text string `xml:",chardata"` // false, false, false, fals...
					} `xml:"inferred"`
					Deletedbyinference struct {
						Text string `xml:",chardata"` // false, false, false, fals...
					} `xml:"deletedbyinference"`
					Trust struct {
						Text string `xml:",chardata"` // 0.91, 0.91, 0.91, 0.91, 0...
					} `xml:"trust"`
					Inferenceprovenance struct {
						Text string `xml:",chardata"` // dedup-similarity-organiza...
					} `xml:"inferenceprovenance"`
					Provenanceaction struct {
						Text       string `xml:",chardata"`
						Classid    string `xml:"classid,attr"`
						Classname  string `xml:"classname,attr"`
						Schemeid   string `xml:"schemeid,attr"`
						Schemename string `xml:"schemename,attr"`
					} `xml:"provenanceaction"`
				} `xml:"datainfo"`
				Rels struct {
					Text string `xml:",chardata"`
					Rel  []struct {
						Text                string `xml:",chardata"`
						Inferred            string `xml:"inferred,attr"`
						Trust               string `xml:"trust,attr"`
						Inferenceprovenance string `xml:"inferenceprovenance,attr"`
						Provenanceaction    string `xml:"provenanceaction,attr"`
						To                  struct {
							Text   string `xml:",chardata"` // corda_______::31d8c335179...
							Class  string `xml:"class,attr"`
							Scheme string `xml:"scheme,attr"`
							Type   string `xml:"type,attr"`
						} `xml:"to"`
						Acronym struct {
							Text string `xml:",chardata"` // ROMP, FORGING IDENTITIES,...
						} `xml:"acronym"`
						Funding []struct {
							Text   string `xml:",chardata"`
							Funder struct {
								Text         string `xml:",chardata"`
								ID           string `xml:"id,attr"`
								Shortname    string `xml:"shortname,attr"`
								Name         string `xml:"name,attr"`
								Jurisdiction string `xml:"jurisdiction,attr"`
							} `xml:"funder"`
							FundingLevel0 struct {
								Text string `xml:",chardata"` // ec__________::EC::FP7, ec...
								Name string `xml:"name,attr"`
							} `xml:"funding_level_0"`
							FundingLevel1 struct {
								Text string `xml:",chardata"` // ec__________::EC::FP7::SP...
								Name string `xml:"name,attr"`
							} `xml:"funding_level_1"`
							FundingLevel2 struct {
								Text string `xml:",chardata"` // ec__________::EC::FP7::SP...
								Name string `xml:"name,attr"`
							} `xml:"funding_level_2"`
						} `xml:"funding"`
						Code struct {
							Text string `xml:",chardata"` // 339123, 212402, 313193, 3...
						} `xml:"code"`
						Contracttype struct {
							Text       string `xml:",chardata"`
							Classid    string `xml:"classid,attr"`
							Classname  string `xml:"classname,attr"`
							Schemeid   string `xml:"schemeid,attr"`
							Schemename string `xml:"schemename,attr"`
						} `xml:"contracttype"`
						Title struct {
							Text string `xml:",chardata"` // Rome's Mediterranean Port...
						} `xml:"title"`
						Websiteurl struct {
							Text string `xml:",chardata"` // http://www.bovinose.eu/, ...
						} `xml:"websiteurl"`
						Datasourcetype struct {
							Text       string `xml:",chardata"`
							Classid    string `xml:"classid,attr"`
							Classname  string `xml:"classname,attr"`
							Schemeid   string `xml:"schemeid,attr"`
							Schemename string `xml:"schemename,attr"`
						} `xml:"datasourcetype"`
						Officialname struct {
							Text string `xml:",chardata"` // RiFEUP - Repositório Ins...
						} `xml:"officialname"`
						Datasourcetypeui struct {
							Text       string `xml:",chardata"`
							Classid    string `xml:"classid,attr"`
							Classname  string `xml:"classname,attr"`
							Schemeid   string `xml:"schemeid,attr"`
							Schemename string `xml:"schemename,attr"`
						} `xml:"datasourcetypeui"`
						Openairecompatibility struct {
							Text       string `xml:",chardata"`
							Classid    string `xml:"classid,attr"`
							Classname  string `xml:"classname,attr"`
							Schemeid   string `xml:"schemeid,attr"`
							Schemename string `xml:"schemename,attr"`
						} `xml:"openairecompatibility"`
					} `xml:"rel"`
				} `xml:"rels"`
				Children struct {
					Text         string `xml:",chardata"`
					Organization []struct {
						Text          string `xml:",chardata"`
						Objidentifier string `xml:"objidentifier,attr"`
						Websiteurl    struct {
							Text string `xml:",chardata"` // http://www.up.pt, http://...
						} `xml:"websiteurl"`
						Legalshortname struct {
							Text string `xml:",chardata"` // UPORTO, UPORTO, UPORTO, U...
						} `xml:"legalshortname"`
						Country struct {
							Text       string `xml:",chardata"`
							Classid    string `xml:"classid,attr"`
							Classname  string `xml:"classname,attr"`
							Schemeid   string `xml:"schemeid,attr"`
							Schemename string `xml:"schemename,attr"`
						} `xml:"country"`
						Legalname struct {
							Text string `xml:",chardata"` // UNIVERSIDADE DO PORTO, UN...
						} `xml:"legalname"`
					} `xml:"organization"`
				} `xml:"children"`
			} `xml:"organization"`
			Project struct {
				Text        string `xml:",chardata"`
				Fundingtree []struct {
					Text   string `xml:",chardata"`
					Funder struct {
						Text string `xml:",chardata"`
						ID   struct {
							Text string `xml:",chardata"` // arc_________::ARC, arc___...
						} `xml:"id"`
						Shortname struct {
							Text string `xml:",chardata"` // ARC, ARC, ARC, ARC, ARC, ...
						} `xml:"shortname"`
						Name struct {
							Text string `xml:",chardata"` // Australian Research Counc...
						} `xml:"name"`
						Jurisdiction struct {
							Text string `xml:",chardata"` // AU, AU, AU, AU, AU, AU, A...
						} `xml:"jurisdiction"`
						Originalname struct {
							Text string `xml:",chardata"` // Fonds zur Förderung der ...
						} `xml:"originalname"`
					} `xml:"funder"`
					FundingLevel0 struct {
						Text string `xml:",chardata"`
						ID   struct {
							Text string `xml:",chardata"` // arc_________::ARC::Specia...
						} `xml:"id"`
						Name struct {
							Text string `xml:",chardata"` // Special Research initiati...
						} `xml:"name"`
						Description struct {
							Text string `xml:",chardata"` // Special Research initiati...
						} `xml:"description"`
						Parent struct {
							Text string `xml:",chardata"`
						} `xml:"parent"`
						Class struct {
							Text string `xml:",chardata"` // arc:fundingStream, arc:fu...
						} `xml:"class"`
					} `xml:"funding_level_0"`
					FundingLevel2 struct {
						Text string `xml:",chardata"`
						ID   struct {
							Text string `xml:",chardata"` // ec__________::EC::FP7::SP...
						} `xml:"id"`
						Description struct {
							Text string `xml:",chardata"` // Energy, Health, ERC, Mari...
						} `xml:"description"`
						Name struct {
							Text string `xml:",chardata"` // ENERGY, HEALTH, ERC, PEOP...
						} `xml:"name"`
						Class struct {
							Text string `xml:",chardata"` // ec:program, ec:program, e...
						} `xml:"class"`
						Parent struct {
							Text          string `xml:",chardata"`
							FundingLevel1 struct {
								Text string `xml:",chardata"`
								ID   struct {
									Text string `xml:",chardata"` // ec__________::EC::FP7::SP...
								} `xml:"id"`
								Description struct {
									Text string `xml:",chardata"` // SP1-Cooperation, SP1-Coop...
								} `xml:"description"`
								Name struct {
									Text string `xml:",chardata"` // SP1, SP1, SP2, SP3, SP3, ...
								} `xml:"name"`
								Class struct {
									Text string `xml:",chardata"` // ec:specificprogram, ec:sp...
								} `xml:"class"`
								Parent struct {
									Text          string `xml:",chardata"`
									FundingLevel0 struct {
										Text string `xml:",chardata"`
										ID   struct {
											Text string `xml:",chardata"` // ec__________::EC::FP7, ec...
										} `xml:"id"`
										Description struct {
											Text string `xml:",chardata"` // SEVENTH FRAMEWORK PROGRAM...
										} `xml:"description"`
										Name struct {
											Text string `xml:",chardata"` // FP7, FP7, FP7, FP7, FP7, ...
										} `xml:"name"`
										Parent struct {
											Text string `xml:",chardata"`
										} `xml:"parent"`
										Class struct {
											Text string `xml:",chardata"` // ec:frameworkprogram, ec:f...
										} `xml:"class"`
									} `xml:"funding_level_0"`
								} `xml:"parent"`
							} `xml:"funding_level_1"`
						} `xml:"parent"`
					} `xml:"funding_level_2"`
					FundingLevel1 struct {
						Text string `xml:",chardata"`
						ID   struct {
							Text string `xml:",chardata"` // ec__________::EC::H2020::...
						} `xml:"id"`
						Description struct {
							Text string `xml:",chardata"` // Research and Innovation a...
						} `xml:"description"`
						Name struct {
							Text string `xml:",chardata"` // RIA, CSA, IA, SME-1, MSCA...
						} `xml:"name"`
						Class struct {
							Text string `xml:",chardata"` // ec:h2020toas, ec:h2020toa...
						} `xml:"class"`
						Parent struct {
							Text          string `xml:",chardata"`
							FundingLevel0 struct {
								Text string `xml:",chardata"`
								ID   struct {
									Text string `xml:",chardata"` // ec__________::EC::H2020, ...
								} `xml:"id"`
								Name struct {
									Text string `xml:",chardata"` // H2020, H2020, H2020, H202...
								} `xml:"name"`
								Description struct {
									Text string `xml:",chardata"` // Horizon 2020 Framework Pr...
								} `xml:"description"`
								Parent struct {
									Text string `xml:",chardata"`
								} `xml:"parent"`
								Class struct {
									Text string `xml:",chardata"` // ec:h2020fundings, ec:h202...
								} `xml:"class"`
							} `xml:"funding_level_0"`
						} `xml:"parent"`
					} `xml:"funding_level_1"`
				} `xml:"fundingtree"`
				Keywords struct {
					Text string `xml:",chardata"` // Eduction,learning,neurosc...
				} `xml:"keywords"`
				Websiteurl struct {
					Text string `xml:",chardata"` // http://purl.org/au-resear...
				} `xml:"websiteurl"`
				Enddate struct {
					Text string `xml:",chardata"` // 2016-12-31, 2010-12-31, 2...
				} `xml:"enddate"`
				Oamandatepublications struct {
					Text string `xml:",chardata"` // false, false, false, fals...
				} `xml:"oamandatepublications"`
				Title struct {
					Text string `xml:",chardata"` // Special Research initiati...
				} `xml:"title"`
				Startdate struct {
					Text string `xml:",chardata"` // 2012-01-01, 2008-01-01, 2...
				} `xml:"startdate"`
				Code struct {
					Text string `xml:",chardata"` // SR120300015, DP0878623, D...
				} `xml:"code"`
				Acronym struct {
					Text string `xml:",chardata"` // HAWE, MILESTONE, TERNANOM...
				} `xml:"acronym"`
				Callidentifier struct {
					Text string `xml:",chardata"` // FP7-ENERGY-2010-FET, FP7-...
				} `xml:"callidentifier"`
				Duration struct {
					Text string `xml:",chardata"`
				} `xml:"duration"`
				Ecsc39 struct {
					Text string `xml:",chardata"` // true, true, false, false,...
				} `xml:"ecsc39"`
				Ecarticle293 struct {
					Text string `xml:",chardata"` // true, true, false, false,...
				} `xml:"ecarticle29_3"`
				Subjects []struct {
					Text       string `xml:",chardata"` // Photonics KET, Supporting...
					Classid    string `xml:"classid,attr"`
					Classname  string `xml:"classname,attr"`
					Schemeid   string `xml:"schemeid,attr"`
					Schemename string `xml:"schemename,attr"`
				} `xml:"subjects"`
				Contracttype struct {
					Text       string `xml:",chardata"`
					Classid    string `xml:"classid,attr"`
					Classname  string `xml:"classname,attr"`
					Schemeid   string `xml:"schemeid,attr"`
					Schemename string `xml:"schemename,attr"`
				} `xml:"contracttype"`
				Collectedfrom struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
					ID   string `xml:"id,attr"`
				} `xml:"collectedfrom"`
				OriginalId struct {
					Text string `xml:",chardata"` // arc_________::SR120300015...
				} `xml:"originalId"`
				Pid struct {
					Text       string `xml:",chardata"`
					Classid    string `xml:"classid,attr"`
					Classname  string `xml:"classname,attr"`
					Schemeid   string `xml:"schemeid,attr"`
					Schemename string `xml:"schemename,attr"`
				} `xml:"pid"`
				Datainfo struct {
					Text     string `xml:",chardata"`
					Inferred struct {
						Text string `xml:",chardata"` // false, false, false, fals...
					} `xml:"inferred"`
					Deletedbyinference struct {
						Text string `xml:",chardata"` // false, false, false, fals...
					} `xml:"deletedbyinference"`
					Trust struct {
						Text string `xml:",chardata"` // 0.9, 0.9, 0.9, 0.9, 0.9, ...
					} `xml:"trust"`
					Inferenceprovenance struct {
						Text string `xml:",chardata"`
					} `xml:"inferenceprovenance"`
					Provenanceaction struct {
						Text       string `xml:",chardata"`
						Classid    string `xml:"classid,attr"`
						Classname  string `xml:"classname,attr"`
						Schemeid   string `xml:"schemeid,attr"`
						Schemename string `xml:"schemename,attr"`
					} `xml:"provenanceaction"`
				} `xml:"datainfo"`
				Rels struct {
					Text string `xml:",chardata"`
					Rel  []struct {
						Text                string `xml:",chardata"`
						Inferred            string `xml:"inferred,attr"`
						Trust               string `xml:"trust,attr"`
						Inferenceprovenance string `xml:"inferenceprovenance,attr"`
						Provenanceaction    string `xml:"provenanceaction,attr"`
						To                  struct {
							Text   string `xml:",chardata"` // dedup_wf_001::0aff0e2c079...
							Class  string `xml:"class,attr"`
							Scheme string `xml:"scheme,attr"`
							Type   string `xml:"type,attr"`
						} `xml:"to"`
						Phone struct {
							Text string `xml:",chardata"` // +351211913169, +44 2476 1...
						} `xml:"phone"`
						Fullname struct {
							Text string `xml:",chardata"` // Fernandes, Nuno, Singh, S...
						} `xml:"fullname"`
						Fax struct {
							Text string `xml:",chardata"` // +351212948531, +44 2476 5...
						} `xml:"fax"`
						Email struct {
							Text string `xml:",chardata"` // nuno.fernandes@omnidea.ne...
						} `xml:"email"`
						Websiteurl struct {
							Text string `xml:",chardata"` // http://edp.pt, http://www...
						} `xml:"websiteurl"`
						Country struct {
							Text       string `xml:",chardata"`
							Classid    string `xml:"classid,attr"`
							Classname  string `xml:"classname,attr"`
							Schemeid   string `xml:"schemeid,attr"`
							Schemename string `xml:"schemename,attr"`
						} `xml:"country"`
						Legalshortname struct {
							Text string `xml:",chardata"` // EDP, UNIZAG FSB, DTU, LAN...
						} `xml:"legalshortname"`
						Legalname struct {
							Text string `xml:",chardata"` // EDP INOVACAO SA, Sveucili...
						} `xml:"legalname"`
					} `xml:"rel"`
				} `xml:"rels"`
				Children struct {
					Text string `xml:",chardata"`
				} `xml:"children"`
			} `xml:"project"`
		} `xml:"entity"`
	} `xml:"metadata"`
	About struct {
		Text string `xml:",chardata"`
	} `xml:"about"`
}
