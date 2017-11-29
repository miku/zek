package iso19139

import "encoding/xml"

type Record struct {
	XMLName xml.Name `xml:"Record"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text       string `xml:",chardata"`
		Status     string `xml:"status,attr"`
		Identifier struct {
			Text string `xml:",chardata"` // urn:x-wmo:md:int.wmo.wis:...
		} `xml:"identifier"`
		Datestamp struct {
			Text string `xml:",chardata"` // 2012-06-20T06:58:35Z, 201...
		} `xml:"datestamp"`
		SetSpec struct {
			Text string `xml:",chardata"` // OKPR, OKPR, OKPR, OKPR, O...
		} `xml:"setSpec"`
	} `xml:"header"`
	Metadata struct {
		Text       string `xml:",chardata"`
		MDMetadata struct {
			Text           string `xml:",chardata"`
			Gco            string `xml:"gco,attr"`
			Gmd            string `xml:"gmd,attr"`
			Gml            string `xml:"gml,attr"`
			Gmx            string `xml:"gmx,attr"`
			Gts            string `xml:"gts,attr"`
			Xlink          string `xml:"xlink,attr"`
			Xsi            string `xml:"xsi,attr"`
			SchemaLocation string `xml:"schemaLocation,attr"`
			FileIdentifier struct {
				Text            string `xml:",chardata"`
				CharacterString struct {
					Text string `xml:",chardata"` // urn:x-wmo:md:int.wmo.wis:...
				} `xml:"CharacterString"`
			} `xml:"fileIdentifier"`
			Language struct {
				Text         string `xml:",chardata"`
				LanguageCode struct {
					Text          string `xml:",chardata"` // eng, eng, eng, eng, eng, ...
					CodeList      string `xml:"codeList,attr"`
					CodeListValue string `xml:"codeListValue,attr"`
				} `xml:"LanguageCode"`
			} `xml:"language"`
			CharacterSet struct {
				Text               string `xml:",chardata"`
				MDCharacterSetCode struct {
					Text          string `xml:",chardata"` // utf8, utf8, utf8, utf8, u...
					CodeList      string `xml:"codeList,attr"`
					CodeListValue string `xml:"codeListValue,attr"`
				} `xml:"MD_CharacterSetCode"`
			} `xml:"characterSet"`
			ParentIdentifier struct {
				Text            string `xml:",chardata"`
				CharacterString struct {
					Text string `xml:",chardata"` // urn:x-wmo:md:int.wmo.wis:...
				} `xml:"CharacterString"`
			} `xml:"parentIdentifier"`
			HierarchyLevel struct {
				Text        string `xml:",chardata"`
				MDScopeCode struct {
					Text          string `xml:",chardata"` // series, series, dataset, ...
					CodeList      string `xml:"codeList,attr"`
					CodeListValue string `xml:"codeListValue,attr"`
				} `xml:"MD_ScopeCode"`
			} `xml:"hierarchyLevel"`
			HierarchyLevelName struct {
				Text            string `xml:",chardata"`
				CharacterString struct {
					Text string `xml:",chardata"` // Series of WMO GTS bulleti...
				} `xml:"CharacterString"`
			} `xml:"hierarchyLevelName"`
			Contact struct {
				Text               string `xml:",chardata"`
				CIResponsibleParty struct {
					Text           string `xml:",chardata"`
					IndividualName struct {
						Text            string `xml:",chardata"`
						CharacterString struct {
							Text string `xml:",chardata"` // Pavel Gal, Pavel Gal, Pav...
						} `xml:"CharacterString"`
					} `xml:"individualName"`
					OrganisationName struct {
						Text            string `xml:",chardata"`
						CharacterString struct {
							Text string `xml:",chardata"` // Czech hydrometeorological...
						} `xml:"CharacterString"`
					} `xml:"organisationName"`
					PositionName struct {
						Text            string `xml:",chardata"`
						CharacterString struct {
							Text string `xml:",chardata"` // HEAD of RTH, HEAD of RTH,...
						} `xml:"CharacterString"`
					} `xml:"positionName"`
					ContactInfo struct {
						Text      string `xml:",chardata"`
						CIContact struct {
							Text  string `xml:",chardata"`
							Phone struct {
								Text        string `xml:",chardata"`
								CITelephone struct {
									Text  string `xml:",chardata"`
									Voice struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // +420 244 032 135, +420 24...
										} `xml:"CharacterString"`
									} `xml:"voice"`
									Facsimile struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // +420 241 760 689, +420 24...
										} `xml:"CharacterString"`
									} `xml:"facsimile"`
								} `xml:"CI_Telephone"`
							} `xml:"phone"`
							Address struct {
								Text      string `xml:",chardata"`
								CIAddress struct {
									Text          string `xml:",chardata"`
									DeliveryPoint struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // Na Sabatce 2050/17, Na Sa...
										} `xml:"CharacterString"`
									} `xml:"deliveryPoint"`
									City struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // Prague, Prague, Prague, P...
										} `xml:"CharacterString"`
									} `xml:"city"`
									PostalCode struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // 14306, 14306, 14306, 1430...
										} `xml:"CharacterString"`
									} `xml:"postalCode"`
									Country struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // Czech republic, Czech rep...
										} `xml:"CharacterString"`
									} `xml:"country"`
									ElectronicMailAddress struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // pavel.gal@chmi.cz, pavel....
										} `xml:"CharacterString"`
									} `xml:"electronicMailAddress"`
								} `xml:"CI_Address"`
							} `xml:"address"`
							OnlineResource struct {
								Text             string `xml:",chardata"`
								CIOnlineResource struct {
									Text    string `xml:",chardata"`
									Linkage struct {
										Text string `xml:",chardata"`
										URL  struct {
											Text string `xml:",chardata"` // www.chmi.cz, www.chmi.cz,...
										} `xml:"URL"`
									} `xml:"linkage"`
								} `xml:"CI_OnlineResource"`
							} `xml:"onlineResource"`
						} `xml:"CI_Contact"`
					} `xml:"contactInfo"`
					Role struct {
						Text       string `xml:",chardata"`
						CIRoleCode struct {
							Text          string `xml:",chardata"` // pointOfContact, pointOfCo...
							CodeList      string `xml:"codeList,attr"`
							CodeListValue string `xml:"codeListValue,attr"`
						} `xml:"CI_RoleCode"`
					} `xml:"role"`
				} `xml:"CI_ResponsibleParty"`
			} `xml:"contact"`
			DateStamp struct {
				Text     string `xml:",chardata"`
				DateTime struct {
					Text string `xml:",chardata"` // 2016-12-28T09:31:01Z, 201...
				} `xml:"DateTime"`
			} `xml:"dateStamp"`
			MetadataStandardName struct {
				Text            string `xml:",chardata"`
				CharacterString struct {
					Text string `xml:",chardata"` // WMO Core Metadata Profile...
				} `xml:"CharacterString"`
			} `xml:"metadataStandardName"`
			MetadataStandardVersion struct {
				Text            string `xml:",chardata"`
				CharacterString struct {
					Text string `xml:",chardata"` // 1.3, 1.3, 1.3, 1.3, 1.3, ...
				} `xml:"CharacterString"`
			} `xml:"metadataStandardVersion"`
			ReferenceSystemInfo struct {
				Text              string `xml:",chardata"`
				MDReferenceSystem struct {
					Text                      string `xml:",chardata"`
					ReferenceSystemIdentifier struct {
						Text         string `xml:",chardata"`
						RSIdentifier struct {
							Text string `xml:",chardata"`
							Code struct {
								Text            string `xml:",chardata"`
								CharacterString struct {
									Text string `xml:",chardata"` // WGS 84, WGS 84, WGS 84, W...
								} `xml:"CharacterString"`
							} `xml:"code"`
							CodeSpace struct {
								Text            string `xml:",chardata"`
								CharacterString struct {
									Text string `xml:",chardata"` // World Geodetic System, Wo...
								} `xml:"CharacterString"`
							} `xml:"codeSpace"`
						} `xml:"RS_Identifier"`
					} `xml:"referenceSystemIdentifier"`
				} `xml:"MD_ReferenceSystem"`
			} `xml:"referenceSystemInfo"`
			MetadataExtensionInfo struct {
				Text                           string `xml:",chardata"`
				MDMetadataExtensionInformation struct {
					Text                    string `xml:",chardata"`
					ExtensionOnLineResource struct {
						Text             string `xml:",chardata"`
						CIOnlineResource struct {
							Text    string `xml:",chardata"`
							Linkage struct {
								Text string `xml:",chardata"`
								URL  struct {
									Text string `xml:",chardata"` // http://www.wmo.int/pages/...
								} `xml:"URL"`
							} `xml:"linkage"`
							Name struct {
								Text            string `xml:",chardata"`
								CharacterString struct {
									Text string `xml:",chardata"` // WMO Core Profile version ...
								} `xml:"CharacterString"`
							} `xml:"name"`
						} `xml:"CI_OnlineResource"`
					} `xml:"extensionOnLineResource"`
				} `xml:"MD_MetadataExtensionInformation"`
			} `xml:"metadataExtensionInfo"`
			IdentificationInfo struct {
				Text                 string `xml:",chardata"`
				MDDataIdentification struct {
					Text     string `xml:",chardata"`
					ID       string `xml:"id,attr"`
					Citation struct {
						Text       string `xml:",chardata"`
						CICitation struct {
							Text  string `xml:",chardata"`
							Title struct {
								Text            string `xml:",chardata"`
								CharacterString struct {
									Text string `xml:",chardata"` // SACZ41 collection of META...
								} `xml:"CharacterString"`
							} `xml:"title"`
							Date []struct {
								Text   string `xml:",chardata"`
								CIDate struct {
									Text string `xml:",chardata"`
									Date struct {
										Text string `xml:",chardata"`
										Date struct {
											Text string `xml:",chardata"` // 2012-04-17Z, 2015-04-14Z,...
										} `xml:"Date"`
									} `xml:"date"`
									DateType struct {
										Text           string `xml:",chardata"`
										CIDateTypeCode struct {
											Text          string `xml:",chardata"` // publication, revision, pu...
											CodeList      string `xml:"codeList,attr"`
											CodeListValue string `xml:"codeListValue,attr"`
										} `xml:"CI_DateTypeCode"`
									} `xml:"dateType"`
								} `xml:"CI_Date"`
							} `xml:"date"`
							Identifier []struct {
								Text         string `xml:",chardata"`
								RSIdentifier struct {
									Text string `xml:",chardata"`
									ID   string `xml:"id,attr"`
									Code struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // urn:x-wmo:md:int.wmo.wis:...
										} `xml:"CharacterString"`
									} `xml:"code"`
									CodeSpace struct {
										Text      string `xml:",chardata"`
										NilReason string `xml:"nilReason,attr"`
									} `xml:"codeSpace"`
								} `xml:"RS_Identifier"`
							} `xml:"identifier"`
							PresentationForm struct {
								Text                   string `xml:",chardata"`
								CIPresentationFormCode struct {
									Text          string `xml:",chardata"` // documentDigital, document...
									CodeList      string `xml:"codeList,attr"`
									CodeListValue string `xml:"codeListValue,attr"`
								} `xml:"CI_PresentationFormCode"`
							} `xml:"presentationForm"`
							OtherCitationDetails struct {
								Text            string `xml:",chardata"`
								CharacterString struct {
									Text string `xml:",chardata"` // The dot notation recommen...
								} `xml:"CharacterString"`
							} `xml:"otherCitationDetails"`
						} `xml:"CI_Citation"`
					} `xml:"citation"`
					Abstract struct {
						Text            string `xml:",chardata"`
						CharacterString struct {
							Text string `xml:",chardata"` // ---- The bulletin collect...
						} `xml:"CharacterString"`
					} `xml:"abstract"`
					Purpose struct {
						Text      string `xml:",chardata"`
						NilReason string `xml:"nilReason,attr"`
					} `xml:"purpose"`
					PointOfContact struct {
						Text               string `xml:",chardata"`
						CIResponsibleParty struct {
							Text           string `xml:",chardata"`
							IndividualName struct {
								Text            string `xml:",chardata"`
								CharacterString struct {
									Text string `xml:",chardata"` // Pavel Gal, Pavel Gal, Pav...
								} `xml:"CharacterString"`
							} `xml:"individualName"`
							OrganisationName struct {
								Text            string `xml:",chardata"`
								CharacterString struct {
									Text string `xml:",chardata"` // Czech hydrometeorological...
								} `xml:"CharacterString"`
							} `xml:"organisationName"`
							PositionName struct {
								Text            string `xml:",chardata"`
								CharacterString struct {
									Text string `xml:",chardata"` // HEAD of RTH, HEAD of RTH,...
								} `xml:"CharacterString"`
							} `xml:"positionName"`
							ContactInfo struct {
								Text      string `xml:",chardata"`
								CIContact struct {
									Text  string `xml:",chardata"`
									Phone struct {
										Text        string `xml:",chardata"`
										CITelephone struct {
											Text  string `xml:",chardata"`
											Voice struct {
												Text            string `xml:",chardata"`
												CharacterString struct {
													Text string `xml:",chardata"` // +420 244 032 135, +420 24...
												} `xml:"CharacterString"`
											} `xml:"voice"`
											Facsimile struct {
												Text            string `xml:",chardata"`
												CharacterString struct {
													Text string `xml:",chardata"` // +420 241 760 689, +420 24...
												} `xml:"CharacterString"`
											} `xml:"facsimile"`
										} `xml:"CI_Telephone"`
									} `xml:"phone"`
									Address struct {
										Text      string `xml:",chardata"`
										CIAddress struct {
											Text          string `xml:",chardata"`
											DeliveryPoint struct {
												Text            string `xml:",chardata"`
												CharacterString struct {
													Text string `xml:",chardata"` // Na Sabatce 2050/17, Na Sa...
												} `xml:"CharacterString"`
											} `xml:"deliveryPoint"`
											City struct {
												Text            string `xml:",chardata"`
												CharacterString struct {
													Text string `xml:",chardata"` // Prague, Prague, Prague, P...
												} `xml:"CharacterString"`
											} `xml:"city"`
											PostalCode struct {
												Text            string `xml:",chardata"`
												CharacterString struct {
													Text string `xml:",chardata"` // 14306, 14306, 14306, 1430...
												} `xml:"CharacterString"`
											} `xml:"postalCode"`
											Country struct {
												Text            string `xml:",chardata"`
												CharacterString struct {
													Text string `xml:",chardata"` // Czech republic, Czech rep...
												} `xml:"CharacterString"`
											} `xml:"country"`
											ElectronicMailAddress struct {
												Text            string `xml:",chardata"`
												CharacterString struct {
													Text string `xml:",chardata"` // pavel.gal@chmi.cz, pavel....
												} `xml:"CharacterString"`
											} `xml:"electronicMailAddress"`
										} `xml:"CI_Address"`
									} `xml:"address"`
									OnlineResource struct {
										Text             string `xml:",chardata"`
										CIOnlineResource struct {
											Text    string `xml:",chardata"`
											Linkage struct {
												Text string `xml:",chardata"`
												URL  struct {
													Text string `xml:",chardata"` // www.chmi.cz, www.chmi.cz,...
												} `xml:"URL"`
											} `xml:"linkage"`
										} `xml:"CI_OnlineResource"`
									} `xml:"onlineResource"`
								} `xml:"CI_Contact"`
							} `xml:"contactInfo"`
							Role struct {
								Text       string `xml:",chardata"`
								CIRoleCode struct {
									Text          string `xml:",chardata"` // originator, originator, o...
									CodeList      string `xml:"codeList,attr"`
									CodeListValue string `xml:"codeListValue,attr"`
								} `xml:"CI_RoleCode"`
							} `xml:"role"`
						} `xml:"CI_ResponsibleParty"`
					} `xml:"pointOfContact"`
					ResourceMaintenance struct {
						Text                     string `xml:",chardata"`
						MDMaintenanceInformation struct {
							Text                          string `xml:",chardata"`
							MaintenanceAndUpdateFrequency struct {
								Text                       string `xml:",chardata"`
								MDMaintenanceFrequencyCode struct {
									Text          string `xml:",chardata"` // continual, continual, con...
									CodeList      string `xml:"codeList,attr"`
									CodeListValue string `xml:"codeListValue,attr"`
								} `xml:"MD_MaintenanceFrequencyCode"`
							} `xml:"maintenanceAndUpdateFrequency"`
							UpdateScope struct {
								Text        string `xml:",chardata"`
								MDScopeCode struct {
									Text          string `xml:",chardata"` // series, series, dataset, ...
									CodeList      string `xml:"codeList,attr"`
									CodeListValue string `xml:"codeListValue,attr"`
								} `xml:"MD_ScopeCode"`
							} `xml:"updateScope"`
							UpdateScopeDescription struct {
								Text               string `xml:",chardata"`
								MDScopeDescription struct {
									Text    string `xml:",chardata"`
									Dataset struct {
										Text      string `xml:",chardata"`
										NilReason string `xml:"nilReason,attr"`
									} `xml:"dataset"`
								} `xml:"MD_ScopeDescription"`
							} `xml:"updateScopeDescription"`
							MaintenanceNote struct {
								Text            string `xml:",chardata"`
								CharacterString struct {
									Text string `xml:",chardata"` // The details of the update...
								} `xml:"CharacterString"`
							} `xml:"maintenanceNote"`
						} `xml:"MD_MaintenanceInformation"`
					} `xml:"resourceMaintenance"`
					DescriptiveKeywords []struct {
						Text       string `xml:",chardata"`
						Uuidref    string `xml:"uuidref,attr"`
						MDKeywords struct {
							Text    string `xml:",chardata"`
							Keyword []struct {
								Text            string `xml:",chardata"`
								NilReason       string `xml:"nilReason,attr"`
								CharacterString struct {
									Text string `xml:",chardata"` // hourly, atmospheric condi...
								} `xml:"CharacterString"`
							} `xml:"keyword"`
							Type struct {
								Text              string `xml:",chardata"`
								MDKeywordTypeCode struct {
									Text          string `xml:",chardata"` // temporal, theme, theme, p...
									CodeList      string `xml:"codeList,attr"`
									CodeListValue string `xml:"codeListValue,attr"`
								} `xml:"MD_KeywordTypeCode"`
							} `xml:"type"`
							ThesaurusName struct {
								Text       string `xml:",chardata"`
								NilReason  string `xml:"nilReason,attr"`
								CICitation struct {
									Text  string `xml:",chardata"`
									Title struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // WMO_CategoryCode, GEMET -...
										} `xml:"CharacterString"`
									} `xml:"title"`
									Date struct {
										Text   string `xml:",chardata"`
										CIDate struct {
											Text string `xml:",chardata"`
											Date struct {
												Text string `xml:",chardata"`
												Date struct {
													Text string `xml:",chardata"` // 2012-06-27, 2008-06-01, 2...
												} `xml:"Date"`
											} `xml:"date"`
											DateType struct {
												Text           string `xml:",chardata"`
												CIDateTypeCode struct {
													Text          string `xml:",chardata"` // revision, publication, re...
													CodeList      string `xml:"codeList,attr"`
													CodeListValue string `xml:"codeListValue,attr"`
												} `xml:"CI_DateTypeCode"`
											} `xml:"dateType"`
										} `xml:"CI_Date"`
									} `xml:"date"`
									CitedResponsibleParty struct {
										Text               string `xml:",chardata"`
										CIResponsibleParty struct {
											Text             string `xml:",chardata"`
											OrganisationName struct {
												Text            string `xml:",chardata"`
												CharacterString struct {
													Text string `xml:",chardata"` // WMO Secretariat, WMO Secr...
												} `xml:"CharacterString"`
											} `xml:"organisationName"`
											Role struct {
												Text       string `xml:",chardata"`
												CIRoleCode struct {
													Text          string `xml:",chardata"` // publisher, publisher, pub...
													CodeList      string `xml:"codeList,attr"`
													CodeListValue string `xml:"codeListValue,attr"`
												} `xml:"CI_RoleCode"`
											} `xml:"role"`
										} `xml:"CI_ResponsibleParty"`
									} `xml:"citedResponsibleParty"`
								} `xml:"CI_Citation"`
							} `xml:"thesaurusName"`
						} `xml:"MD_Keywords"`
					} `xml:"descriptiveKeywords"`
					ResourceConstraints []struct {
						Text          string `xml:",chardata"`
						MDConstraints struct {
							Text          string `xml:",chardata"`
							UseLimitation struct {
								Text            string `xml:",chardata"`
								CharacterString struct {
									Text string `xml:",chardata"` // WMOAdditional, WMOEssenti...
								} `xml:"CharacterString"`
							} `xml:"useLimitation"`
						} `xml:"MD_Constraints"`
						MDLegalConstraints struct {
							Text              string `xml:",chardata"`
							AccessConstraints struct {
								Text              string `xml:",chardata"`
								MDRestrictionCode struct {
									Text          string `xml:",chardata"` // otherRestrictions, otherR...
									CodeList      string `xml:"codeList,attr"`
									CodeListValue string `xml:"codeListValue,attr"`
								} `xml:"MD_RestrictionCode"`
							} `xml:"accessConstraints"`
							UseConstraints struct {
								Text              string `xml:",chardata"`
								MDRestrictionCode struct {
									Text          string `xml:",chardata"` // otherRestrictions, otherR...
									CodeList      string `xml:"codeList,attr"`
									CodeListValue string `xml:"codeListValue,attr"`
								} `xml:"MD_RestrictionCode"`
							} `xml:"useConstraints"`
							OtherConstraints []struct {
								Text   string `xml:",chardata"`
								Anchor struct {
									Text string `xml:",chardata"` // WMOAdditional, GTSPriorit...
									Href string `xml:"href,attr"`
								} `xml:"Anchor"`
							} `xml:"otherConstraints"`
						} `xml:"MD_LegalConstraints"`
					} `xml:"resourceConstraints"`
					Language struct {
						Text         string `xml:",chardata"`
						LanguageCode struct {
							Text          string `xml:",chardata"` // eng, eng, eng, eng, eng, ...
							CodeList      string `xml:"codeList,attr"`
							CodeListValue string `xml:"codeListValue,attr"`
						} `xml:"LanguageCode"`
					} `xml:"language"`
					CharacterSet struct {
						Text               string `xml:",chardata"`
						MDCharacterSetCode struct {
							Text          string `xml:",chardata"` // utf8, utf8, utf8, utf8, u...
							CodeList      string `xml:"codeList,attr"`
							CodeListValue string `xml:"codeListValue,attr"`
						} `xml:"MD_CharacterSetCode"`
					} `xml:"characterSet"`
					TopicCategory struct {
						Text                string `xml:",chardata"`
						MDTopicCategoryCode struct {
							Text string `xml:",chardata"` // climatologyMeteorologyAtm...
						} `xml:"MD_TopicCategoryCode"`
					} `xml:"topicCategory"`
					Extent struct {
						Text     string `xml:",chardata"`
						EXExtent struct {
							Text        string `xml:",chardata"`
							ID          string `xml:"id,attr"`
							Description struct {
								Text            string `xml:",chardata"`
								CharacterString struct {
									Text string `xml:",chardata"` // Geographic bounding box t...
								} `xml:"CharacterString"`
							} `xml:"description"`
							GeographicElement struct {
								Text                    string `xml:",chardata"`
								EXGeographicBoundingBox struct {
									Text               string `xml:",chardata"`
									ID                 string `xml:"id,attr"`
									WestBoundLongitude struct {
										Text    string `xml:",chardata"`
										Decimal struct {
											Text string `xml:",chardata"` // 12.1000, 12.1000, 12.1000...
										} `xml:"Decimal"`
									} `xml:"westBoundLongitude"`
									EastBoundLongitude struct {
										Text    string `xml:",chardata"`
										Decimal struct {
											Text string `xml:",chardata"` // 18.9000, 18.9000, 18.9000...
										} `xml:"Decimal"`
									} `xml:"eastBoundLongitude"`
									SouthBoundLatitude struct {
										Text    string `xml:",chardata"`
										Decimal struct {
											Text string `xml:",chardata"` // 48.6000, 48.6000, 48.6000...
										} `xml:"Decimal"`
									} `xml:"southBoundLatitude"`
									NorthBoundLatitude struct {
										Text    string `xml:",chardata"`
										Decimal struct {
											Text string `xml:",chardata"` // 51.1000, 51.1000, 51.1000...
										} `xml:"Decimal"`
									} `xml:"northBoundLatitude"`
								} `xml:"EX_GeographicBoundingBox"`
							} `xml:"geographicElement"`
							TemporalElement struct {
								Text string `xml:",chardata"`
							} `xml:"temporalElement"`
						} `xml:"EX_Extent"`
					} `xml:"extent"`
				} `xml:"MD_DataIdentification"`
			} `xml:"identificationInfo"`
			ContentInfo struct {
				Text                  string `xml:",chardata"`
				MDCoverageDescription struct {
					Text                 string `xml:",chardata"`
					AttributeDescription struct {
						Text       string `xml:",chardata"`
						RecordType struct {
							Text string `xml:",chardata"` // SACZ41 : SACZ41 collectio...
						} `xml:"RecordType"`
					} `xml:"attributeDescription"`
					ContentType struct {
						Text                      string `xml:",chardata"`
						MDCoverageContentTypeCode struct {
							Text          string `xml:",chardata"` // thematicClassification, t...
							CodeList      string `xml:"codeList,attr"`
							CodeListValue string `xml:"codeListValue,attr"`
						} `xml:"MD_CoverageContentTypeCode"`
					} `xml:"contentType"`
				} `xml:"MD_CoverageDescription"`
			} `xml:"contentInfo"`
			DistributionInfo struct {
				Text           string `xml:",chardata"`
				MDDistribution struct {
					Text               string `xml:",chardata"`
					DistributionFormat struct {
						Text     string `xml:",chardata"`
						MDFormat struct {
							Text string `xml:",chardata"`
							Name struct {
								Text            string `xml:",chardata"`
								CharacterString struct {
									Text string `xml:",chardata"` // METAR-5-X EXT., BUFR-4-XI...
								} `xml:"CharacterString"`
							} `xml:"name"`
							Version struct {
								Text            string `xml:",chardata"`
								CharacterString struct {
									Text string `xml:",chardata"` // XI, 12, 12, 12, 89
								} `xml:"CharacterString"`
							} `xml:"version"`
							Specification struct {
								Text            string `xml:",chardata"`
								CharacterString struct {
									Text string `xml:",chardata"` // http://www.wmo.int/pages/...
								} `xml:"CharacterString"`
							} `xml:"specification"`
						} `xml:"MD_Format"`
					} `xml:"distributionFormat"`
					TransferOptions struct {
						Text                     string `xml:",chardata"`
						MDDigitalTransferOptions struct {
							Text   string `xml:",chardata"`
							OnLine struct {
								Text             string `xml:",chardata"`
								CIOnlineResource struct {
									Text    string `xml:",chardata"`
									Linkage struct {
										Text string `xml:",chardata"`
										URL  struct {
											Text string `xml:",chardata"` // http://dcpc.chmi.cz/dw/in...
										} `xml:"URL"`
									} `xml:"linkage"`
								} `xml:"CI_OnlineResource"`
							} `xml:"onLine"`
						} `xml:"MD_DigitalTransferOptions"`
					} `xml:"transferOptions"`
				} `xml:"MD_Distribution"`
			} `xml:"distributionInfo"`
			DataQualityInfo struct {
				Text          string `xml:",chardata"`
				DQDataQuality struct {
					Text  string `xml:",chardata"`
					Scope struct {
						Text    string `xml:",chardata"`
						DQScope struct {
							Text  string `xml:",chardata"`
							Level struct {
								Text        string `xml:",chardata"`
								MDScopeCode struct {
									Text          string `xml:",chardata"` // dataset, dataset, dataset...
									CodeList      string `xml:"codeList,attr"`
									CodeListValue string `xml:"codeListValue,attr"`
								} `xml:"MD_ScopeCode"`
							} `xml:"level"`
						} `xml:"DQ_Scope"`
					} `xml:"scope"`
					Report struct {
						Text                string `xml:",chardata"`
						DQDomainConsistency struct {
							Text   string `xml:",chardata"`
							Type   string `xml:"type,attr"`
							Result struct {
								Text                string `xml:",chardata"`
								DQConformanceResult struct {
									Text          string `xml:",chardata"`
									Type          string `xml:"type,attr"`
									Specification struct {
										Text       string `xml:",chardata"`
										CICitation struct {
											Text  string `xml:",chardata"`
											Title struct {
												Text            string `xml:",chardata"`
												CharacterString struct {
													Text string `xml:",chardata"` // COMMISSION REGULATION (EU...
												} `xml:"CharacterString"`
											} `xml:"title"`
											Date struct {
												Text   string `xml:",chardata"`
												CIDate struct {
													Text string `xml:",chardata"`
													Date struct {
														Text string `xml:",chardata"`
														Date struct {
															Text string `xml:",chardata"` // 2010-12-08, 2010-12-08, 2...
														} `xml:"Date"`
													} `xml:"date"`
													DateType struct {
														Text           string `xml:",chardata"`
														CIDateTypeCode struct {
															Text          string `xml:",chardata"` // publication, publication,...
															CodeList      string `xml:"codeList,attr"`
															CodeListValue string `xml:"codeListValue,attr"`
														} `xml:"CI_DateTypeCode"`
													} `xml:"dateType"`
												} `xml:"CI_Date"`
											} `xml:"date"`
										} `xml:"CI_Citation"`
									} `xml:"specification"`
									Explanation struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // See the referenced specif...
										} `xml:"CharacterString"`
									} `xml:"explanation"`
									Pass struct {
										Text    string `xml:",chardata"`
										Boolean struct {
											Text string `xml:",chardata"` // true, true, true, true, t...
										} `xml:"Boolean"`
									} `xml:"pass"`
								} `xml:"DQ_ConformanceResult"`
							} `xml:"result"`
						} `xml:"DQ_DomainConsistency"`
					} `xml:"report"`
					Lineage struct {
						Text      string `xml:",chardata"`
						LILineage struct {
							Text      string `xml:",chardata"`
							Statement struct {
								Text            string `xml:",chardata"`
								CharacterString struct {
									Text string `xml:",chardata"` // High data quality control...
								} `xml:"CharacterString"`
							} `xml:"statement"`
							ProcessStep struct {
								Text          string `xml:",chardata"`
								LIProcessStep struct {
									Text        string `xml:",chardata"`
									Description struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // This metadata record was ...
										} `xml:"CharacterString"`
									} `xml:"description"`
								} `xml:"LI_ProcessStep"`
							} `xml:"processStep"`
						} `xml:"LI_Lineage"`
					} `xml:"lineage"`
				} `xml:"DQ_DataQuality"`
			} `xml:"dataQualityInfo"`
		} `xml:"MD_Metadata"`
	} `xml:"metadata"`
	About struct {
		Text string `xml:",chardata"`
	} `xml:"about"`
}
