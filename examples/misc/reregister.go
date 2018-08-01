package misc

import "encoding/xml"

// RERegister was generated 2018-08-01 20:46:38 by tir on sol.
type RERegister struct {
	XMLName        xml.Name `xml:"RE_Register"`
	Text           string   `xml:",chardata"`
	Grg            string   `xml:"grg,attr"`
	Xmlns          string   `xml:"xmlns,attr"`
	Gco            string   `xml:"gco,attr"`
	Gmd            string   `xml:"gmd,attr"`
	Xlink          string   `xml:"xlink,attr"`
	Xs             string   `xml:"xs,attr"`
	Xsi            string   `xml:"xsi,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Name           struct {
		Text            string `xml:",chardata"`
		CharacterString struct {
			Text string `xml:",chardata"` // North American Profile of...
		} `xml:"CharacterString"`
	} `xml:"name"`
	ContentSummary struct {
		Text            string `xml:",chardata"`
		CharacterString struct {
			Text string `xml:",chardata"` // This register contains it...
		} `xml:"CharacterString"`
	} `xml:"contentSummary"`
	UniformResourceIdentifier []struct {
		Text             string `xml:",chardata"`
		CIOnlineResource struct {
			Text    string `xml:",chardata"`
			Linkage struct {
				Text string `xml:",chardata"`
				URL  struct {
					Text string `xml:",chardata"` // www.fgdc.gov/nap/metadata...
				} `xml:"URL"`
			} `xml:"linkage"`
			Protocol struct {
				Text            string `xml:",chardata"`
				CharacterString struct {
					Text string `xml:",chardata"` // http, http
				} `xml:"CharacterString"`
			} `xml:"protocol"`
		} `xml:"CI_OnlineResource"`
	} `xml:"uniformResourceIdentifier"`
	OperatingLanguage struct {
		Text     string `xml:",chardata"`
		RELocale struct {
			Text string `xml:",chardata"`
			Name struct {
				Text            string `xml:",chardata"`
				CharacterString struct {
					Text string `xml:",chardata"` // English
				} `xml:"CharacterString"`
			} `xml:"name"`
			Language struct {
				Text         string `xml:",chardata"`
				LanguageCode struct {
					Text          string `xml:",chardata"` // English
					CodeList      string `xml:"codeList,attr"`
					CodeListValue string `xml:"codeListValue,attr"`
				} `xml:"LanguageCode"`
			} `xml:"language"`
			Country struct {
				Text            string `xml:",chardata"`
				CharacterString struct {
					Text string `xml:",chardata"` // 840
				} `xml:"CharacterString"`
			} `xml:"country"`
			CharacterEncoding struct {
				Text               string `xml:",chardata"`
				MDCharacterSetCode struct {
					Text          string `xml:",chardata"` // utf8
					CodeList      string `xml:"codeList,attr"`
					CodeListValue string `xml:"codeListValue,attr"`
				} `xml:"MD_CharacterSetCode"`
			} `xml:"characterEncoding"`
		} `xml:"RE_Locale"`
	} `xml:"operatingLanguage"`
	AlternativeLanguages struct {
		Text     string `xml:",chardata"`
		RELocale struct {
			Text string `xml:",chardata"`
			Name struct {
				Text            string `xml:",chardata"`
				CharacterString struct {
					Text string `xml:",chardata"` // Français
				} `xml:"CharacterString"`
			} `xml:"name"`
			Language struct {
				Text         string `xml:",chardata"`
				LanguageCode struct {
					Text          string `xml:",chardata"` // French
					CodeList      string `xml:"codeList,attr"`
					CodeListValue string `xml:"codeListValue,attr"`
				} `xml:"LanguageCode"`
			} `xml:"language"`
			Country struct {
				Text            string `xml:",chardata"`
				CharacterString struct {
					Text string `xml:",chardata"` // 124
				} `xml:"CharacterString"`
			} `xml:"country"`
			CharacterEncoding struct {
				Text               string `xml:",chardata"`
				MDCharacterSetCode struct {
					Text          string `xml:",chardata"` // utf8
					CodeList      string `xml:"codeList,attr"`
					CodeListValue string `xml:"codeListValue,attr"`
				} `xml:"MD_CharacterSetCode"`
			} `xml:"characterEncoding"`
		} `xml:"RE_Locale"`
	} `xml:"alternativeLanguages"`
	Version struct {
		Text      string `xml:",chardata"`
		REVersion struct {
			Text          string `xml:",chardata"`
			VersionNumber struct {
				Text            string `xml:",chardata"`
				CharacterString struct {
					Text string `xml:",chardata"` // 1.3.1
				} `xml:"CharacterString"`
			} `xml:"versionNumber"`
			VersionDate struct {
				Text string `xml:",chardata"`
				Date struct {
					Text string `xml:",chardata"` // 2015-02-12
				} `xml:"Date"`
			} `xml:"versionDate"`
		} `xml:"RE_Version"`
	} `xml:"version"`
	DateOfLastChange struct {
		Text string `xml:",chardata"`
		Date struct {
			Text string `xml:",chardata"` // 2015-02-12
		} `xml:"Date"`
	} `xml:"dateOfLastChange"`
	Owner struct {
		Text            string `xml:",chardata"`
		RERegisterOwner struct {
			Text string `xml:",chardata"`
			Name struct {
				Text            string `xml:",chardata"`
				CharacterString struct {
					Text string `xml:",chardata"` // NAP - Metadata Working Gr...
				} `xml:"CharacterString"`
			} `xml:"name"`
			Contact struct {
				Text               string `xml:",chardata"`
				CIResponsibleParty struct {
					Text           string `xml:",chardata"`
					IndividualName struct {
						Text            string `xml:",chardata"`
						CharacterString struct {
							Text string `xml:",chardata"` // Jean Brodeur
						} `xml:"CharacterString"`
					} `xml:"individualName"`
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
											Text string `xml:",chardata"` // 1.819.564.5600 										...
										} `xml:"CharacterString"`
									} `xml:"voice"`
								} `xml:"CI_Telephone"`
							} `xml:"phone"`
							Address struct {
								Text      string `xml:",chardata"`
								CIAddress struct {
									Text          string `xml:",chardata"`
									DeliveryPoint struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // 2144-010 King West 						...
										} `xml:"CharacterString"`
									} `xml:"deliveryPoint"`
									City struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // Sherbrooke
										} `xml:"CharacterString"`
									} `xml:"city"`
									AdministrativeArea struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // Quebec
										} `xml:"CharacterString"`
									} `xml:"administrativeArea"`
									PostalCode struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // J1J 2E8
										} `xml:"CharacterString"`
									} `xml:"postalCode"`
									Country struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // Canada
										} `xml:"CharacterString"`
									} `xml:"country"`
									ElectronicMailAddress struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // brodeur@rncan.gc.ca
										} `xml:"CharacterString"`
									} `xml:"electronicMailAddress"`
								} `xml:"CI_Address"`
							} `xml:"address"`
						} `xml:"CI_Contact"`
					} `xml:"contactInfo"`
					Role struct {
						Text       string `xml:",chardata"`
						CIRoleCode struct {
							Text          string `xml:",chardata"` // owner
							CodeList      string `xml:"codeList,attr"`
							CodeListValue string `xml:"codeListValue,attr"`
						} `xml:"CI_RoleCode"`
					} `xml:"role"`
				} `xml:"CI_ResponsibleParty"`
			} `xml:"contact"`
		} `xml:"RE_RegisterOwner"`
	} `xml:"owner"`
	Submitter struct {
		Text                     string `xml:",chardata"`
		RESubmittingOrganization struct {
			Text string `xml:",chardata"`
			Name struct {
				Text            string `xml:",chardata"`
				CharacterString struct {
					Text string `xml:",chardata"` // NAP - Metadata Working Gr...
				} `xml:"CharacterString"`
			} `xml:"name"`
			Contact struct {
				Text               string `xml:",chardata"`
				CIResponsibleParty struct {
					Text           string `xml:",chardata"`
					IndividualName struct {
						Text            string `xml:",chardata"`
						CharacterString struct {
							Text string `xml:",chardata"` // Jean Brodeur
						} `xml:"CharacterString"`
					} `xml:"individualName"`
					OrganisationName struct {
						Text            string `xml:",chardata"`
						CharacterString struct {
							Text string `xml:",chardata"` // Natural Resources Canada
						} `xml:"CharacterString"`
					} `xml:"organisationName"`
					PositionName struct {
						Text            string `xml:",chardata"`
						CharacterString struct {
							Text string `xml:",chardata"` // Research scientist
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
											Text string `xml:",chardata"` // 1.819.564.5600 										...
										} `xml:"CharacterString"`
									} `xml:"voice"`
								} `xml:"CI_Telephone"`
							} `xml:"phone"`
							Address struct {
								Text      string `xml:",chardata"`
								CIAddress struct {
									Text          string `xml:",chardata"`
									DeliveryPoint struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // 2144-010 King West 						...
										} `xml:"CharacterString"`
									} `xml:"deliveryPoint"`
									City struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // Sherbrooke
										} `xml:"CharacterString"`
									} `xml:"city"`
									AdministrativeArea struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // Québec
										} `xml:"CharacterString"`
									} `xml:"administrativeArea"`
									PostalCode struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // J1J 2E8
										} `xml:"CharacterString"`
									} `xml:"postalCode"`
									Country struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // Canada
										} `xml:"CharacterString"`
									} `xml:"country"`
									ElectronicMailAddress struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // brodeur@rncan.gc.ca
										} `xml:"CharacterString"`
									} `xml:"electronicMailAddress"`
								} `xml:"CI_Address"`
							} `xml:"address"`
						} `xml:"CI_Contact"`
					} `xml:"contactInfo"`
					Role struct {
						Text       string `xml:",chardata"`
						CIRoleCode struct {
							Text          string `xml:",chardata"` // custodian
							CodeList      string `xml:"codeList,attr"`
							CodeListValue string `xml:"codeListValue,attr"`
						} `xml:"CI_RoleCode"`
					} `xml:"role"`
				} `xml:"CI_ResponsibleParty"`
			} `xml:"contact"`
		} `xml:"RE_SubmittingOrganization"`
	} `xml:"submitter"`
	ContainedItemClass []struct {
		Text        string `xml:",chardata"`
		REItemClass struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Name struct {
				Text            string `xml:",chardata"`
				CharacterString struct {
					Text string `xml:",chardata"` // MD_Metadata, MD_Identific...
				} `xml:"CharacterString"`
			} `xml:"name"`
			TechnicalStandard struct {
				Text       string `xml:",chardata"`
				CICitation struct {
					Text  string `xml:",chardata"`
					Title struct {
						Text            string `xml:",chardata"`
						CharacterString struct {
							Text string `xml:",chardata"` // ISO 19115:2003 Geographic...
						} `xml:"CharacterString"`
					} `xml:"title"`
					Date struct {
						Text   string `xml:",chardata"`
						CIDate struct {
							Text string `xml:",chardata"`
							Date struct {
								Text string `xml:",chardata"`
								Date struct {
									Text string `xml:",chardata"` // 2003-05-01, 2003-05-01, 2...
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
					Edition struct {
						Text            string `xml:",chardata"`
						CharacterString struct {
							Text string `xml:",chardata"` // 1.0, 1.0, 1.0, 1.0, 1.0, ...
						} `xml:"CharacterString"`
					} `xml:"edition"`
					EditionDate struct {
						Text string `xml:",chardata"`
						Date struct {
							Text string `xml:",chardata"` // 2003-05-01, 2003-05-01, 2...
						} `xml:"Date"`
					} `xml:"editionDate"`
					Identifier struct {
						Text         string `xml:",chardata"`
						MDIdentifier struct {
							Text      string `xml:",chardata"`
							Authority struct {
								Text       string `xml:",chardata"`
								CICitation struct {
									Text  string `xml:",chardata"`
									Title struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // International organizatio...
										} `xml:"CharacterString"`
									} `xml:"title"`
									AlternateTitle struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // ISO, ISO, ISO, ISO, ISO, ...
										} `xml:"CharacterString"`
									} `xml:"alternateTitle"`
									Date struct {
										Text   string `xml:",chardata"`
										CIDate struct {
											Text string `xml:",chardata"`
											Date struct {
												Text string `xml:",chardata"`
												Date struct {
													Text string `xml:",chardata"` // 1947-02-23, 1947-02-23, 1...
												} `xml:"Date"`
											} `xml:"date"`
											DateType struct {
												Text           string `xml:",chardata"`
												CIDateTypeCode struct {
													Text          string `xml:",chardata"` // creation, creation, creat...
													CodeList      string `xml:"codeList,attr"`
													CodeListValue string `xml:"codeListValue,attr"`
												} `xml:"CI_DateTypeCode"`
											} `xml:"dateType"`
										} `xml:"CI_Date"`
									} `xml:"date"`
								} `xml:"CI_Citation"`
							} `xml:"authority"`
							Code struct {
								Text            string `xml:",chardata"`
								CharacterString struct {
									Text string `xml:",chardata"` // ISO19115:2003(E), ISO1911...
								} `xml:"CharacterString"`
							} `xml:"code"`
						} `xml:"MD_Identifier"`
					} `xml:"identifier"`
				} `xml:"CI_Citation"`
			} `xml:"technicalStandard"`
			AlternativeNames struct {
				Text string `xml:",chardata"`
			} `xml:"alternativeNames"`
			DescribedItem []struct {
				Text string `xml:",chardata"`
				Href string `xml:"href,attr"`
			} `xml:"describedItem"`
		} `xml:"RE_ItemClass"`
	} `xml:"containedItemClass"`
	Manager struct {
		Text              string `xml:",chardata"`
		RERegisterManager struct {
			Text string `xml:",chardata"`
			Name struct {
				Text            string `xml:",chardata"`
				CharacterString struct {
					Text string `xml:",chardata"` // Jean Brodeur
				} `xml:"CharacterString"`
			} `xml:"name"`
			Contact struct {
				Text               string `xml:",chardata"`
				CIResponsibleParty struct {
					Text           string `xml:",chardata"`
					IndividualName struct {
						Text            string `xml:",chardata"`
						CharacterString struct {
							Text string `xml:",chardata"` // Jean Brodeur
						} `xml:"CharacterString"`
					} `xml:"individualName"`
					OrganisationName struct {
						Text            string `xml:",chardata"`
						CharacterString struct {
							Text string `xml:",chardata"` // Natural Resources Canada
						} `xml:"CharacterString"`
					} `xml:"organisationName"`
					PositionName struct {
						Text            string `xml:",chardata"`
						CharacterString struct {
							Text string `xml:",chardata"` // Research scientist
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
											Text string `xml:",chardata"` // 1.819.564.5600 										...
										} `xml:"CharacterString"`
									} `xml:"voice"`
								} `xml:"CI_Telephone"`
							} `xml:"phone"`
							Address struct {
								Text      string `xml:",chardata"`
								CIAddress struct {
									Text          string `xml:",chardata"`
									DeliveryPoint struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // 2144-010 King West 						...
										} `xml:"CharacterString"`
									} `xml:"deliveryPoint"`
									City struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // Sherbrooke
										} `xml:"CharacterString"`
									} `xml:"city"`
									AdministrativeArea struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // Québec
										} `xml:"CharacterString"`
									} `xml:"administrativeArea"`
									PostalCode struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // J1J 2E8
										} `xml:"CharacterString"`
									} `xml:"postalCode"`
									Country struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // Canada
										} `xml:"CharacterString"`
									} `xml:"country"`
									ElectronicMailAddress struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // brodeur@rncan.gc.ca
										} `xml:"CharacterString"`
									} `xml:"electronicMailAddress"`
								} `xml:"CI_Address"`
							} `xml:"address"`
						} `xml:"CI_Contact"`
					} `xml:"contactInfo"`
					Role struct {
						Text       string `xml:",chardata"`
						CIRoleCode struct {
							Text          string `xml:",chardata"` // custodian
							CodeList      string `xml:"codeList,attr"`
							CodeListValue string `xml:"codeListValue,attr"`
						} `xml:"CI_RoleCode"`
					} `xml:"role"`
				} `xml:"CI_ResponsibleParty"`
			} `xml:"contact"`
		} `xml:"RE_RegisterManager"`
	} `xml:"manager"`
	ContainedItem []struct {
		Text           string `xml:",chardata"`
		RERegisterItem struct {
			Text           string `xml:",chardata"`
			ID             string `xml:"id,attr"`
			ItemIdentifier struct {
				Text    string `xml:",chardata"`
				Integer struct {
					Text string `xml:",chardata"` // 1, 2, 3, 4, 5, 6, 7, 8, 9...
				} `xml:"Integer"`
			} `xml:"itemIdentifier"`
			Name struct {
				Text            string `xml:",chardata"`
				CharacterString struct {
					Text string `xml:",chardata"` // MetadataRecordInformation...
				} `xml:"CharacterString"`
			} `xml:"name"`
			Status struct {
				Text         string `xml:",chardata"`
				REItemStatus struct {
					Text string `xml:",chardata"` // valid, valid, valid, vali...
				} `xml:"RE_ItemStatus"`
			} `xml:"status"`
			DateAccepted struct {
				Text string `xml:",chardata"`
				Date struct {
					Text string `xml:",chardata"` // 2007-08-01, 2007-08-01, 2...
				} `xml:"Date"`
			} `xml:"dateAccepted"`
			DateAmended struct {
				Text string `xml:",chardata"`
				Date struct {
					Text string `xml:",chardata"` // 2012-03-31, 2012-03-31, 2...
				} `xml:"Date"`
			} `xml:"dateAmended"`
			Definition struct {
				Text            string `xml:",chardata"`
				CharacterString struct {
					Text string `xml:",chardata"` // root entity which defines...
				} `xml:"CharacterString"`
			} `xml:"definition"`
			Description struct {
				Text            string `xml:",chardata"`
				CharacterString struct {
					Text string `xml:",chardata"` // Information which describ...
				} `xml:"CharacterString"`
			} `xml:"description"`
			FieldOfApplication struct {
				Text                 string `xml:",chardata"`
				REFieldOfApplication struct {
					Text string `xml:",chardata"`
					Name struct {
						Text            string `xml:",chardata"`
						CharacterString struct {
							Text string `xml:",chardata"` // geographic information, g...
						} `xml:"CharacterString"`
					} `xml:"name"`
				} `xml:"RE_FieldOfApplication"`
			} `xml:"fieldOfApplication"`
			AlternativeExpressions struct {
				Text                    string `xml:",chardata"`
				REAlternativeExpression struct {
					Text string `xml:",chardata"`
					Name struct {
						Text            string `xml:",chardata"`
						CharacterString struct {
							Text string `xml:",chardata"` // InformationEnregistrement...
						} `xml:"CharacterString"`
					} `xml:"name"`
					Definition struct {
						Text            string `xml:",chardata"`
						CharacterString struct {
							Text string `xml:",chardata"` // entité maîtresse qui re...
						} `xml:"CharacterString"`
					} `xml:"definition"`
					Description struct {
						Text            string `xml:",chardata"`
						CharacterString struct {
							Text string `xml:",chardata"` // Information qui décrit l...
						} `xml:"CharacterString"`
					} `xml:"description"`
					FieldOfApplication struct {
						Text                 string `xml:",chardata"`
						REFieldOfApplication struct {
							Text string `xml:",chardata"`
							Name struct {
								Text            string `xml:",chardata"`
								CharacterString struct {
									Text string `xml:",chardata"` // information géographique...
								} `xml:"CharacterString"`
							} `xml:"name"`
						} `xml:"RE_FieldOfApplication"`
					} `xml:"fieldOfApplication"`
					Locale struct {
						Text     string `xml:",chardata"`
						RELocale struct {
							Text string `xml:",chardata"`
							Name struct {
								Text            string `xml:",chardata"`
								CharacterString struct {
									Text string `xml:",chardata"` // Français, Français, Fra...
								} `xml:"CharacterString"`
							} `xml:"name"`
							Language struct {
								Text         string `xml:",chardata"`
								LanguageCode struct {
									Text          string `xml:",chardata"` // French, French, French, F...
									CodeList      string `xml:"codeList,attr"`
									CodeListValue string `xml:"codeListValue,attr"`
								} `xml:"LanguageCode"`
							} `xml:"language"`
							Country struct {
								Text            string `xml:",chardata"`
								CharacterString struct {
									Text string `xml:",chardata"` // 124, 124, 124, 124, 124, ...
								} `xml:"CharacterString"`
							} `xml:"country"`
							CharacterEncoding struct {
								Text               string `xml:",chardata"`
								MDCharacterSetCode struct {
									Text          string `xml:",chardata"` // utf8, utf8, utf8, utf8, u...
									CodeList      string `xml:"codeList,attr"`
									CodeListValue string `xml:"codeListValue,attr"`
								} `xml:"MD_CharacterSetCode"`
							} `xml:"characterEncoding"`
						} `xml:"RE_Locale"`
					} `xml:"locale"`
				} `xml:"RE_AlternativeExpression"`
			} `xml:"alternativeExpressions"`
			AdditionInformation []struct {
				Text                  string `xml:",chardata"`
				REAdditionInformation struct {
					Text         string `xml:",chardata"`
					DateProposed struct {
						Text string `xml:",chardata"`
						Date struct {
							Text string `xml:",chardata"` // 2007-08-01, 2012-03-01, 2...
						} `xml:"Date"`
					} `xml:"dateProposed"`
					Justification struct {
						Text            string `xml:",chardata"`
						CharacterString struct {
							Text string `xml:",chardata"` // creation of the registry ...
						} `xml:"CharacterString"`
					} `xml:"justification"`
					Status struct {
						Text             string `xml:",chardata"`
						REDecisionStatus struct {
							Text string `xml:",chardata"` // final, final, final, fina...
						} `xml:"RE_DecisionStatus"`
					} `xml:"status"`
					Disposition struct {
						Text          string `xml:",chardata"`
						REDisposition struct {
							Text string `xml:",chardata"` // accepted, accepted, accep...
						} `xml:"RE_Disposition"`
					} `xml:"disposition"`
					Sponsor struct {
						Text                     string `xml:",chardata"`
						RESubmittingOrganization struct {
							Text string `xml:",chardata"`
							Name struct {
								Text            string `xml:",chardata"`
								CharacterString struct {
									Text string `xml:",chardata"` // NAP - Metadata Working 		...
								} `xml:"CharacterString"`
							} `xml:"name"`
							Contact struct {
								Text               string `xml:",chardata"`
								CIResponsibleParty struct {
									Text           string `xml:",chardata"`
									IndividualName struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // Jean Brodeur, Jean Brodeu...
										} `xml:"CharacterString"`
									} `xml:"individualName"`
									OrganisationName struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // Natural Resources 							...
										} `xml:"CharacterString"`
									} `xml:"organisationName"`
									PositionName struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // Research 											scien...
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
															Text string `xml:",chardata"` // 1.819.564.5600 										...
														} `xml:"CharacterString"`
													} `xml:"voice"`
												} `xml:"CI_Telephone"`
											} `xml:"phone"`
											Address struct {
												Text      string `xml:",chardata"`
												CIAddress struct {
													Text          string `xml:",chardata"`
													DeliveryPoint struct {
														Text            string `xml:",chardata"`
														CharacterString struct {
															Text string `xml:",chardata"` // 2144-010 King West 						...
														} `xml:"CharacterString"`
													} `xml:"deliveryPoint"`
													City struct {
														Text            string `xml:",chardata"`
														CharacterString struct {
															Text string `xml:",chardata"` // Sherbrooke, Sherbrooke, S...
														} `xml:"CharacterString"`
													} `xml:"city"`
													AdministrativeArea struct {
														Text            string `xml:",chardata"`
														CharacterString struct {
															Text string `xml:",chardata"` // Québec, Québec, Québec...
														} `xml:"CharacterString"`
													} `xml:"administrativeArea"`
													PostalCode struct {
														Text            string `xml:",chardata"`
														CharacterString struct {
															Text string `xml:",chardata"` // J1J 2E8, J1J 2E8, J1J 2E8...
														} `xml:"CharacterString"`
													} `xml:"postalCode"`
													Country struct {
														Text            string `xml:",chardata"`
														CharacterString struct {
															Text string `xml:",chardata"` // Canada, Canada, Canada, C...
														} `xml:"CharacterString"`
													} `xml:"country"`
													ElectronicMailAddress struct {
														Text            string `xml:",chardata"`
														CharacterString struct {
															Text string `xml:",chardata"` // brodeur@rncan.gc.ca, brod...
														} `xml:"CharacterString"`
													} `xml:"electronicMailAddress"`
												} `xml:"CI_Address"`
											} `xml:"address"`
										} `xml:"CI_Contact"`
									} `xml:"contactInfo"`
									Role struct {
										Text       string `xml:",chardata"`
										CIRoleCode struct {
											Text          string `xml:",chardata"` // custodian, custodian, cus...
											CodeList      string `xml:"codeList,attr"`
											CodeListValue string `xml:"codeListValue,attr"`
										} `xml:"CI_RoleCode"`
									} `xml:"role"`
								} `xml:"CI_ResponsibleParty"`
							} `xml:"contact"`
						} `xml:"RE_SubmittingOrganization"`
					} `xml:"sponsor"`
					RegisterManagerNotes struct {
						Text            string `xml:",chardata"`
						CharacterString struct {
							Text string `xml:",chardata"` // change from "SpatialTempo...
						} `xml:"CharacterString"`
					} `xml:"registerManagerNotes"`
					DateDisposed struct {
						Text string `xml:",chardata"`
						Date struct {
							Text string `xml:",chardata"` // 2015-02-12
						} `xml:"Date"`
					} `xml:"dateDisposed"`
				} `xml:"RE_AdditionInformation"`
			} `xml:"additionInformation"`
			ItemClass struct {
				Text string `xml:",chardata"`
				Href string `xml:"href,attr"`
			} `xml:"itemClass"`
			SpecificationSource struct {
				Text        string `xml:",chardata"`
				REReference struct {
					Text                   string `xml:",chardata"`
					ItemIdentifierAtSource struct {
						Text            string `xml:",chardata"`
						CharacterString struct {
							Text string `xml:",chardata"` // 1, 2, 3, 4, 5, 6, 8, 9, 1...
						} `xml:"CharacterString"`
					} `xml:"itemIdentifierAtSource"`
					Similarity struct {
						Text                 string `xml:",chardata"`
						RESimilarityToSource struct {
							Text          string `xml:",chardata"` // identical, identical, ide...
							CodeList      string `xml:"codeList,attr"`
							CodeListValue string `xml:"codeListValue,attr"`
						} `xml:"RE_SimilarityToSource"`
					} `xml:"similarity"`
					SourceCitation struct {
						Text              string `xml:",chardata"`
						REReferenceSource struct {
							Text     string `xml:",chardata"`
							Citation struct {
								Text       string `xml:",chardata"`
								CICitation struct {
									Text  string `xml:",chardata"`
									Title struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // ISO 19115:2003 Geographic...
										} `xml:"CharacterString"`
									} `xml:"title"`
									Date struct {
										Text   string `xml:",chardata"`
										CIDate struct {
											Text string `xml:",chardata"`
											Date struct {
												Text string `xml:",chardata"`
												Date struct {
													Text string `xml:",chardata"` // 2003-05-01, 2003-05-01, 2...
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
									Edition struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // 1.0, 1.0, 1.0, 1.0, 1.0, ...
										} `xml:"CharacterString"`
									} `xml:"edition"`
									EditionDate struct {
										Text string `xml:",chardata"`
										Date struct {
											Text string `xml:",chardata"` // 2003-05-01, 2003-05-01, 2...
										} `xml:"Date"`
									} `xml:"editionDate"`
									Identifier struct {
										Text         string `xml:",chardata"`
										MDIdentifier struct {
											Text      string `xml:",chardata"`
											Authority struct {
												Text       string `xml:",chardata"`
												CICitation struct {
													Text  string `xml:",chardata"`
													Title struct {
														Text            string `xml:",chardata"`
														CharacterString struct {
															Text string `xml:",chardata"` // International organizatio...
														} `xml:"CharacterString"`
													} `xml:"title"`
													AlternateTitle struct {
														Text            string `xml:",chardata"`
														CharacterString struct {
															Text string `xml:",chardata"` // ISO, ISO, ISO, ISO, ISO, ...
														} `xml:"CharacterString"`
													} `xml:"alternateTitle"`
													Date struct {
														Text   string `xml:",chardata"`
														CIDate struct {
															Text string `xml:",chardata"`
															Date struct {
																Text string `xml:",chardata"`
																Date struct {
																	Text string `xml:",chardata"` // 1947-02-23, 1947-02-23, 1...
																} `xml:"Date"`
															} `xml:"date"`
															DateType struct {
																Text           string `xml:",chardata"`
																CIDateTypeCode struct {
																	Text          string `xml:",chardata"` // creation, creation, creat...
																	CodeList      string `xml:"codeList,attr"`
																	CodeListValue string `xml:"codeListValue,attr"`
																} `xml:"CI_DateTypeCode"`
															} `xml:"dateType"`
														} `xml:"CI_Date"`
													} `xml:"date"`
												} `xml:"CI_Citation"`
											} `xml:"authority"`
											Code struct {
												Text            string `xml:",chardata"`
												CharacterString struct {
													Text string `xml:",chardata"` // ISO19115:2003(E), ISO1911...
												} `xml:"CharacterString"`
											} `xml:"code"`
										} `xml:"MD_Identifier"`
									} `xml:"identifier"`
								} `xml:"CI_Citation"`
							} `xml:"citation"`
						} `xml:"RE_ReferenceSource"`
					} `xml:"sourceCitation"`
				} `xml:"RE_Reference"`
			} `xml:"specificationSource"`
			SpecificationLineage struct {
				Text        string `xml:",chardata"`
				REReference struct {
					Text                   string `xml:",chardata"`
					ItemIdentifierAtSource struct {
						Text            string `xml:",chardata"`
						CharacterString struct {
							Text string `xml:",chardata"` // 91
						} `xml:"CharacterString"`
					} `xml:"itemIdentifierAtSource"`
					Similarity struct {
						Text                 string `xml:",chardata"`
						RESimilarityToSource struct {
							Text          string `xml:",chardata"` // identical
							CodeList      string `xml:"codeList,attr"`
							CodeListValue string `xml:"codeListValue,attr"`
						} `xml:"RE_SimilarityToSource"`
					} `xml:"similarity"`
					SourceCitation struct {
						Text              string `xml:",chardata"`
						REReferenceSource struct {
							Text     string `xml:",chardata"`
							Citation struct {
								Text       string `xml:",chardata"`
								CICitation struct {
									Text  string `xml:",chardata"`
									Title struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // ISO 19115:2003 Geographic...
										} `xml:"CharacterString"`
									} `xml:"title"`
									Date struct {
										Text   string `xml:",chardata"`
										CIDate struct {
											Text string `xml:",chardata"`
											Date struct {
												Text string `xml:",chardata"`
												Date struct {
													Text string `xml:",chardata"` // 2003-05-01
												} `xml:"Date"`
											} `xml:"date"`
											DateType struct {
												Text           string `xml:",chardata"`
												CIDateTypeCode struct {
													Text          string `xml:",chardata"` // publication
													CodeList      string `xml:"codeList,attr"`
													CodeListValue string `xml:"codeListValue,attr"`
												} `xml:"CI_DateTypeCode"`
											} `xml:"dateType"`
										} `xml:"CI_Date"`
									} `xml:"date"`
									Edition struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // 1.0
										} `xml:"CharacterString"`
									} `xml:"edition"`
									EditionDate struct {
										Text string `xml:",chardata"`
										Date struct {
											Text string `xml:",chardata"` // 2003-05-01
										} `xml:"Date"`
									} `xml:"editionDate"`
									Identifier struct {
										Text         string `xml:",chardata"`
										MDIdentifier struct {
											Text      string `xml:",chardata"`
											Authority struct {
												Text       string `xml:",chardata"`
												CICitation struct {
													Text  string `xml:",chardata"`
													Title struct {
														Text            string `xml:",chardata"`
														CharacterString struct {
															Text string `xml:",chardata"` // International organizatio...
														} `xml:"CharacterString"`
													} `xml:"title"`
													AlternateTitle struct {
														Text            string `xml:",chardata"`
														CharacterString struct {
															Text string `xml:",chardata"` // ISO
														} `xml:"CharacterString"`
													} `xml:"alternateTitle"`
													Date struct {
														Text   string `xml:",chardata"`
														CIDate struct {
															Text string `xml:",chardata"`
															Date struct {
																Text string `xml:",chardata"`
																Date struct {
																	Text string `xml:",chardata"` // 1947-02-23
																} `xml:"Date"`
															} `xml:"date"`
															DateType struct {
																Text           string `xml:",chardata"`
																CIDateTypeCode struct {
																	Text          string `xml:",chardata"` // creation
																	CodeList      string `xml:"codeList,attr"`
																	CodeListValue string `xml:"codeListValue,attr"`
																} `xml:"CI_DateTypeCode"`
															} `xml:"dateType"`
														} `xml:"CI_Date"`
													} `xml:"date"`
												} `xml:"CI_Citation"`
											} `xml:"authority"`
											Code struct {
												Text            string `xml:",chardata"`
												CharacterString struct {
													Text string `xml:",chardata"` // ISO19115:2003(E)
												} `xml:"CharacterString"`
											} `xml:"code"`
										} `xml:"MD_Identifier"`
									} `xml:"identifier"`
								} `xml:"CI_Citation"`
							} `xml:"citation"`
						} `xml:"RE_ReferenceSource"`
					} `xml:"sourceCitation"`
				} `xml:"RE_Reference"`
			} `xml:"specificationLineage"`
		} `xml:"RE_RegisterItem"`
		RESubregisterDescription struct {
			Text           string `xml:",chardata"`
			ID             string `xml:"id,attr"`
			ItemIdentifier struct {
				Text    string `xml:",chardata"`
				Integer struct {
					Text string `xml:",chardata"` // 733, 734, 735
				} `xml:"Integer"`
			} `xml:"itemIdentifier"`
			Name struct {
				Text            string `xml:",chardata"`
				CharacterString struct {
					Text string `xml:",chardata"` // ISO639-2_LanguageNameCode...
				} `xml:"CharacterString"`
			} `xml:"name"`
			Status struct {
				Text         string `xml:",chardata"`
				REItemStatus struct {
					Text string `xml:",chardata"` // valid, valid, valid
				} `xml:"RE_ItemStatus"`
			} `xml:"status"`
			DateAccepted struct {
				Text string `xml:",chardata"`
				Date struct {
					Text string `xml:",chardata"` // 2010-10-08, 2010-10-08, 2...
				} `xml:"Date"`
			} `xml:"dateAccepted"`
			Definition struct {
				Text            string `xml:",chardata"`
				CharacterString struct {
					Text string `xml:",chardata"` // ISO standard codes for th...
				} `xml:"CharacterString"`
			} `xml:"definition"`
			FieldOfApplication struct {
				Text                 string `xml:",chardata"`
				REFieldOfApplication struct {
					Text string `xml:",chardata"`
					Name struct {
						Text            string `xml:",chardata"`
						CharacterString struct {
							Text string `xml:",chardata"` // geographic information, g...
						} `xml:"CharacterString"`
					} `xml:"name"`
				} `xml:"RE_FieldOfApplication"`
			} `xml:"fieldOfApplication"`
			AlternativeExpressions struct {
				Text                    string `xml:",chardata"`
				REAlternativeExpression struct {
					Text string `xml:",chardata"`
					Name struct {
						Text            string `xml:",chardata"`
						CharacterString struct {
							Text string `xml:",chardata"` // ISO639-2_CodesNomLangues,...
						} `xml:"CharacterString"`
					} `xml:"name"`
					Definition struct {
						Text            string `xml:",chardata"`
						CharacterString struct {
							Text string `xml:",chardata"` // codes normalisés ISO des...
						} `xml:"CharacterString"`
					} `xml:"definition"`
					FieldOfApplication struct {
						Text                 string `xml:",chardata"`
						REFieldOfApplication struct {
							Text string `xml:",chardata"`
							Name struct {
								Text            string `xml:",chardata"`
								CharacterString struct {
									Text string `xml:",chardata"` // information géographique...
								} `xml:"CharacterString"`
							} `xml:"name"`
						} `xml:"RE_FieldOfApplication"`
					} `xml:"fieldOfApplication"`
					Locale struct {
						Text     string `xml:",chardata"`
						RELocale struct {
							Text string `xml:",chardata"`
							Name struct {
								Text            string `xml:",chardata"`
								CharacterString struct {
									Text string `xml:",chardata"` // Français, Français, Fra...
								} `xml:"CharacterString"`
							} `xml:"name"`
							Language struct {
								Text         string `xml:",chardata"`
								LanguageCode struct {
									Text          string `xml:",chardata"` // French, French, French
									CodeList      string `xml:"codeList,attr"`
									CodeListValue string `xml:"codeListValue,attr"`
								} `xml:"LanguageCode"`
							} `xml:"language"`
							Country struct {
								Text            string `xml:",chardata"`
								CharacterString struct {
									Text string `xml:",chardata"` // 124, 124, 124
								} `xml:"CharacterString"`
							} `xml:"country"`
							CharacterEncoding struct {
								Text               string `xml:",chardata"`
								MDCharacterSetCode struct {
									Text          string `xml:",chardata"` // utf8, utf8, utf8
									CodeList      string `xml:"codeList,attr"`
									CodeListValue string `xml:"codeListValue,attr"`
								} `xml:"MD_CharacterSetCode"`
							} `xml:"characterEncoding"`
						} `xml:"RE_Locale"`
					} `xml:"locale"`
					Description struct {
						Text            string `xml:",chardata"`
						CharacterString struct {
							Text string `xml:",chardata"` // La liste donne les noms d...
						} `xml:"CharacterString"`
					} `xml:"description"`
				} `xml:"RE_AlternativeExpression"`
			} `xml:"alternativeExpressions"`
			AdditionInformation []struct {
				Text                  string `xml:",chardata"`
				REAdditionInformation struct {
					Text         string `xml:",chardata"`
					DateProposed struct {
						Text string `xml:",chardata"`
						Date struct {
							Text string `xml:",chardata"` // 2007-08-01, 2012-05-16, 2...
						} `xml:"Date"`
					} `xml:"dateProposed"`
					Justification struct {
						Text            string `xml:",chardata"`
						CharacterString struct {
							Text string `xml:",chardata"` // creation of the registry ...
						} `xml:"CharacterString"`
					} `xml:"justification"`
					Status struct {
						Text             string `xml:",chardata"`
						REDecisionStatus struct {
							Text string `xml:",chardata"` // final, final, final, fina...
						} `xml:"RE_DecisionStatus"`
					} `xml:"status"`
					Disposition struct {
						Text          string `xml:",chardata"`
						REDisposition struct {
							Text string `xml:",chardata"` // accepted, accepted, accep...
						} `xml:"RE_Disposition"`
					} `xml:"disposition"`
					Sponsor struct {
						Text                     string `xml:",chardata"`
						RESubmittingOrganization struct {
							Text string `xml:",chardata"`
							Name struct {
								Text            string `xml:",chardata"`
								CharacterString struct {
									Text string `xml:",chardata"` // NAP - Metadata Working 		...
								} `xml:"CharacterString"`
							} `xml:"name"`
							Contact struct {
								Text               string `xml:",chardata"`
								CIResponsibleParty struct {
									Text           string `xml:",chardata"`
									IndividualName struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // Jean Brodeur, Jean Brodeu...
										} `xml:"CharacterString"`
									} `xml:"individualName"`
									OrganisationName struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // Natural Resources 							...
										} `xml:"CharacterString"`
									} `xml:"organisationName"`
									PositionName struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // Research 											scien...
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
															Text string `xml:",chardata"` // 1.819.564.5600 										...
														} `xml:"CharacterString"`
													} `xml:"voice"`
												} `xml:"CI_Telephone"`
											} `xml:"phone"`
											Address struct {
												Text      string `xml:",chardata"`
												CIAddress struct {
													Text          string `xml:",chardata"`
													DeliveryPoint struct {
														Text            string `xml:",chardata"`
														CharacterString struct {
															Text string `xml:",chardata"` // 2144-010 King West 						...
														} `xml:"CharacterString"`
													} `xml:"deliveryPoint"`
													City struct {
														Text            string `xml:",chardata"`
														CharacterString struct {
															Text string `xml:",chardata"` // Sherbrooke, Sherbrooke, S...
														} `xml:"CharacterString"`
													} `xml:"city"`
													AdministrativeArea struct {
														Text            string `xml:",chardata"`
														CharacterString struct {
															Text string `xml:",chardata"` // Québec, Québec, Québec...
														} `xml:"CharacterString"`
													} `xml:"administrativeArea"`
													PostalCode struct {
														Text            string `xml:",chardata"`
														CharacterString struct {
															Text string `xml:",chardata"` // J1J 2E8, J1J 2E8, J1J 2E8...
														} `xml:"CharacterString"`
													} `xml:"postalCode"`
													Country struct {
														Text            string `xml:",chardata"`
														CharacterString struct {
															Text string `xml:",chardata"` // Canada, Canada, Canada, C...
														} `xml:"CharacterString"`
													} `xml:"country"`
													ElectronicMailAddress struct {
														Text            string `xml:",chardata"`
														CharacterString struct {
															Text string `xml:",chardata"` // brodeur@rncan.gc.ca, brod...
														} `xml:"CharacterString"`
													} `xml:"electronicMailAddress"`
												} `xml:"CI_Address"`
											} `xml:"address"`
										} `xml:"CI_Contact"`
									} `xml:"contactInfo"`
									Role struct {
										Text       string `xml:",chardata"`
										CIRoleCode struct {
											Text          string `xml:",chardata"` // custodian, custodian, cus...
											CodeList      string `xml:"codeList,attr"`
											CodeListValue string `xml:"codeListValue,attr"`
										} `xml:"CI_RoleCode"`
									} `xml:"role"`
								} `xml:"CI_ResponsibleParty"`
							} `xml:"contact"`
						} `xml:"RE_SubmittingOrganization"`
					} `xml:"sponsor"`
				} `xml:"RE_AdditionInformation"`
			} `xml:"additionInformation"`
			ItemClass struct {
				Text string `xml:",chardata"`
				Href string `xml:"href,attr"`
			} `xml:"itemClass"`
			SpecificationLineage struct {
				Text        string `xml:",chardata"`
				REReference struct {
					Text                   string `xml:",chardata"`
					ItemIdentifierAtSource struct {
						Text            string `xml:",chardata"`
						CharacterString struct {
							Text string `xml:",chardata"` // ISO639-2
						} `xml:"CharacterString"`
					} `xml:"itemIdentifierAtSource"`
					Similarity struct {
						Text                 string `xml:",chardata"`
						RESimilarityToSource struct {
							Text          string `xml:",chardata"` // identical
							CodeList      string `xml:"codeList,attr"`
							CodeListValue string `xml:"codeListValue,attr"`
						} `xml:"RE_SimilarityToSource"`
					} `xml:"similarity"`
					SourceCitation struct {
						Text              string `xml:",chardata"`
						REReferenceSource struct {
							Text     string `xml:",chardata"`
							Citation struct {
								Text       string `xml:",chardata"`
								CICitation struct {
									Text  string `xml:",chardata"`
									Title struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // ISO 639-2:1998 Codes for ...
										} `xml:"CharacterString"`
									} `xml:"title"`
									Date struct {
										Text   string `xml:",chardata"`
										CIDate struct {
											Text string `xml:",chardata"`
											Date struct {
												Text string `xml:",chardata"`
												Date struct {
													Text string `xml:",chardata"` // 2003-07-30
												} `xml:"Date"`
											} `xml:"date"`
											DateType struct {
												Text           string `xml:",chardata"`
												CIDateTypeCode struct {
													Text          string `xml:",chardata"` // publication
													CodeList      string `xml:"codeList,attr"`
													CodeListValue string `xml:"codeListValue,attr"`
												} `xml:"CI_DateTypeCode"`
											} `xml:"dateType"`
										} `xml:"CI_Date"`
									} `xml:"date"`
									Identifier struct {
										Text         string `xml:",chardata"`
										MDIdentifier struct {
											Text      string `xml:",chardata"`
											Authority struct {
												Text       string `xml:",chardata"`
												CICitation struct {
													Text  string `xml:",chardata"`
													Title struct {
														Text            string `xml:",chardata"`
														CharacterString struct {
															Text string `xml:",chardata"` // International organizatio...
														} `xml:"CharacterString"`
													} `xml:"title"`
													AlternateTitle struct {
														Text            string `xml:",chardata"`
														CharacterString struct {
															Text string `xml:",chardata"` // ISO
														} `xml:"CharacterString"`
													} `xml:"alternateTitle"`
													Date struct {
														Text   string `xml:",chardata"`
														CIDate struct {
															Text string `xml:",chardata"`
															Date struct {
																Text string `xml:",chardata"`
																Date struct {
																	Text string `xml:",chardata"` // 1947-02-23
																} `xml:"Date"`
															} `xml:"date"`
															DateType struct {
																Text           string `xml:",chardata"`
																CIDateTypeCode struct {
																	Text          string `xml:",chardata"` // creation
																	CodeList      string `xml:"codeList,attr"`
																	CodeListValue string `xml:"codeListValue,attr"`
																} `xml:"CI_DateTypeCode"`
															} `xml:"dateType"`
														} `xml:"CI_Date"`
													} `xml:"date"`
												} `xml:"CI_Citation"`
											} `xml:"authority"`
											Code struct {
												Text            string `xml:",chardata"`
												CharacterString struct {
													Text string `xml:",chardata"` // ISO 												639-2:199...
												} `xml:"CharacterString"`
											} `xml:"code"`
										} `xml:"MD_Identifier"`
									} `xml:"identifier"`
								} `xml:"CI_Citation"`
							} `xml:"citation"`
						} `xml:"RE_ReferenceSource"`
					} `xml:"sourceCitation"`
				} `xml:"RE_Reference"`
			} `xml:"specificationLineage"`
			UniformResourceIdentifier struct {
				Text             string `xml:",chardata"`
				CIOnlineResource struct {
					Text    string `xml:",chardata"`
					Linkage struct {
						Text string `xml:",chardata"`
						URL  struct {
							Text string `xml:",chardata"` // www.loc.gov/standards/iso...
						} `xml:"URL"`
					} `xml:"linkage"`
					Protocol struct {
						Text            string `xml:",chardata"`
						CharacterString struct {
							Text string `xml:",chardata"` // http, http, http
						} `xml:"CharacterString"`
					} `xml:"protocol"`
					Name struct {
						Text            string `xml:",chardata"`
						CharacterString struct {
							Text string `xml:",chardata"` // Codes for the representat...
						} `xml:"CharacterString"`
					} `xml:"name"`
					Description struct {
						Text            string `xml:",chardata"`
						CharacterString struct {
							Text string `xml:",chardata"` // alpha-3 codes arranged al...
						} `xml:"CharacterString"`
					} `xml:"description"`
					Function struct {
						Text                 string `xml:",chardata"`
						CIOnLineFunctionCode struct {
							Text          string `xml:",chardata"` // information, information,...
							CodeList      string `xml:"codeList,attr"`
							CodeListValue string `xml:"codeListValue,attr"`
						} `xml:"CI_OnLineFunctionCode"`
					} `xml:"function"`
				} `xml:"CI_OnlineResource"`
			} `xml:"uniformResourceIdentifier"`
			OperatingLanguage struct {
				Text     string `xml:",chardata"`
				RELocale struct {
					Text string `xml:",chardata"`
					Name struct {
						Text            string `xml:",chardata"`
						CharacterString struct {
							Text string `xml:",chardata"` // English, English, English...
						} `xml:"CharacterString"`
					} `xml:"name"`
					Language struct {
						Text         string `xml:",chardata"`
						LanguageCode struct {
							Text          string `xml:",chardata"` // English, English, English...
							CodeList      string `xml:"codeList,attr"`
							CodeListValue string `xml:"codeListValue,attr"`
						} `xml:"LanguageCode"`
					} `xml:"language"`
					Country struct {
						Text            string `xml:",chardata"`
						CharacterString struct {
							Text string `xml:",chardata"` // 840, 840, 840
						} `xml:"CharacterString"`
					} `xml:"country"`
					CharacterEncoding struct {
						Text               string `xml:",chardata"`
						MDCharacterSetCode struct {
							Text          string `xml:",chardata"` // utf8, utf8, utf8
							CodeList      string `xml:"codeList,attr"`
							CodeListValue string `xml:"codeListValue,attr"`
						} `xml:"MD_CharacterSetCode"`
					} `xml:"characterEncoding"`
				} `xml:"RE_Locale"`
			} `xml:"operatingLanguage"`
			ContainedItemClass struct {
				Text string `xml:",chardata"`
				Href string `xml:"href,attr"`
			} `xml:"containedItemClass"`
			SubregisterManager struct {
				Text              string `xml:",chardata"`
				RERegisterManager struct {
					Text string `xml:",chardata"`
					Name struct {
						Text            string `xml:",chardata"`
						CharacterString struct {
							Text string `xml:",chardata"` // The Library of Congress N...
						} `xml:"CharacterString"`
					} `xml:"name"`
					Contact struct {
						Text               string `xml:",chardata"`
						CIResponsibleParty struct {
							Text             string `xml:",chardata"`
							OrganisationName struct {
								Text            string `xml:",chardata"`
								CharacterString struct {
									Text string `xml:",chardata"` // The Library of Congress N...
								} `xml:"CharacterString"`
							} `xml:"organisationName"`
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
													Text string `xml:",chardata"` // 1.202.707.6237, +41 22 74...
												} `xml:"CharacterString"`
											} `xml:"voice"`
										} `xml:"CI_Telephone"`
									} `xml:"phone"`
									Address struct {
										Text      string `xml:",chardata"`
										CIAddress struct {
											Text string `xml:",chardata"`
											City struct {
												Text            string `xml:",chardata"`
												CharacterString struct {
													Text string `xml:",chardata"` // Washington, Genève, Lond...
												} `xml:"CharacterString"`
											} `xml:"city"`
											AdministrativeArea struct {
												Text            string `xml:",chardata"`
												CharacterString struct {
													Text string `xml:",chardata"` // DC, 20
												} `xml:"CharacterString"`
											} `xml:"administrativeArea"`
											PostalCode struct {
												Text            string `xml:",chardata"`
												CharacterString struct {
													Text string `xml:",chardata"` // 20540-4402, CH-1211, W4 4...
												} `xml:"CharacterString"`
											} `xml:"postalCode"`
											Country struct {
												Text            string `xml:",chardata"`
												CharacterString struct {
													Text string `xml:",chardata"` // USA, Switzerland, United ...
												} `xml:"CharacterString"`
											} `xml:"country"`
											ElectronicMailAddress struct {
												Text            string `xml:",chardata"`
												CharacterString struct {
													Text string `xml:",chardata"` // iso639-2@loc.gov, country...
												} `xml:"CharacterString"`
											} `xml:"electronicMailAddress"`
											DeliveryPoint []struct {
												Text            string `xml:",chardata"`
												CharacterString struct {
													Text string `xml:",chardata"` // 1, rue de 												Var...
												} `xml:"CharacterString"`
											} `xml:"deliveryPoint"`
										} `xml:"CI_Address"`
									} `xml:"address"`
									OnlineResource struct {
										Text             string `xml:",chardata"`
										CIOnlineResource struct {
											Text    string `xml:",chardata"`
											Linkage struct {
												Text string `xml:",chardata"`
												URL  struct {
													Text string `xml:",chardata"` // www.loc.gov/standards/iso...
												} `xml:"URL"`
											} `xml:"linkage"`
											Protocol struct {
												Text            string `xml:",chardata"`
												CharacterString struct {
													Text string `xml:",chardata"` // http, http, http
												} `xml:"CharacterString"`
											} `xml:"protocol"`
											Name struct {
												Text            string `xml:",chardata"`
												CharacterString struct {
													Text string `xml:",chardata"` // ISO 639-2 Registration 		...
												} `xml:"CharacterString"`
											} `xml:"name"`
											Description struct {
												Text            string `xml:",chardata"`
												CharacterString struct {
													Text string `xml:",chardata"` // ISO 639-2 Registration 		...
												} `xml:"CharacterString"`
											} `xml:"description"`
											Function struct {
												Text                 string `xml:",chardata"`
												CIOnLineFunctionCode struct {
													Text          string `xml:",chardata"` // information, information
													CodeList      string `xml:"codeList,attr"`
													CodeListValue string `xml:"codeListValue,attr"`
												} `xml:"CI_OnLineFunctionCode"`
											} `xml:"function"`
										} `xml:"CI_OnlineResource"`
									} `xml:"onlineResource"`
								} `xml:"CI_Contact"`
							} `xml:"contactInfo"`
							Role struct {
								Text       string `xml:",chardata"`
								CIRoleCode struct {
									Text          string `xml:",chardata"` // custodian, custodian, cus...
									CodeList      string `xml:"codeList,attr"`
									CodeListValue string `xml:"codeListValue,attr"`
								} `xml:"CI_RoleCode"`
							} `xml:"role"`
						} `xml:"CI_ResponsibleParty"`
					} `xml:"contact"`
				} `xml:"RE_RegisterManager"`
			} `xml:"subregisterManager"`
			Description struct {
				Text            string `xml:",chardata"`
				CharacterString struct {
					Text string `xml:",chardata"` // The list states the count...
				} `xml:"CharacterString"`
			} `xml:"description"`
			SpecificationSource struct {
				Text        string `xml:",chardata"`
				REReference struct {
					Text                   string `xml:",chardata"`
					ItemIdentifierAtSource struct {
						Text            string `xml:",chardata"`
						CharacterString struct {
							Text string `xml:",chardata"` // ISO3166-1, ISO4217
						} `xml:"CharacterString"`
					} `xml:"itemIdentifierAtSource"`
					Similarity struct {
						Text                 string `xml:",chardata"`
						RESimilarityToSource struct {
							Text          string `xml:",chardata"` // identical, identical
							CodeList      string `xml:"codeList,attr"`
							CodeListValue string `xml:"codeListValue,attr"`
						} `xml:"RE_SimilarityToSource"`
					} `xml:"similarity"`
					SourceCitation struct {
						Text              string `xml:",chardata"`
						REReferenceSource struct {
							Text     string `xml:",chardata"`
							Citation struct {
								Text       string `xml:",chardata"`
								CICitation struct {
									Text  string `xml:",chardata"`
									Title struct {
										Text            string `xml:",chardata"`
										CharacterString struct {
											Text string `xml:",chardata"` // ISO 3166-1:2006 Codes for...
										} `xml:"CharacterString"`
									} `xml:"title"`
									Date struct {
										Text   string `xml:",chardata"`
										CIDate struct {
											Text string `xml:",chardata"`
											Date struct {
												Text string `xml:",chardata"`
												Date struct {
													Text string `xml:",chardata"` // 2006-11-20, 2001
												} `xml:"Date"`
											} `xml:"date"`
											DateType struct {
												Text           string `xml:",chardata"`
												CIDateTypeCode struct {
													Text          string `xml:",chardata"` // publication, publication
													CodeList      string `xml:"codeList,attr"`
													CodeListValue string `xml:"codeListValue,attr"`
												} `xml:"CI_DateTypeCode"`
											} `xml:"dateType"`
										} `xml:"CI_Date"`
									} `xml:"date"`
									Identifier struct {
										Text         string `xml:",chardata"`
										MDIdentifier struct {
											Text      string `xml:",chardata"`
											Authority struct {
												Text       string `xml:",chardata"`
												CICitation struct {
													Text  string `xml:",chardata"`
													Title struct {
														Text            string `xml:",chardata"`
														CharacterString struct {
															Text string `xml:",chardata"` // International organizatio...
														} `xml:"CharacterString"`
													} `xml:"title"`
													AlternateTitle struct {
														Text            string `xml:",chardata"`
														CharacterString struct {
															Text string `xml:",chardata"` // ISO, ISO
														} `xml:"CharacterString"`
													} `xml:"alternateTitle"`
													Date struct {
														Text   string `xml:",chardata"`
														CIDate struct {
															Text string `xml:",chardata"`
															Date struct {
																Text string `xml:",chardata"`
																Date struct {
																	Text string `xml:",chardata"` // 1947-02-23, 1947-02-23
																} `xml:"Date"`
															} `xml:"date"`
															DateType struct {
																Text           string `xml:",chardata"`
																CIDateTypeCode struct {
																	Text          string `xml:",chardata"` // creation, creation
																	CodeList      string `xml:"codeList,attr"`
																	CodeListValue string `xml:"codeListValue,attr"`
																} `xml:"CI_DateTypeCode"`
															} `xml:"dateType"`
														} `xml:"CI_Date"`
													} `xml:"date"`
												} `xml:"CI_Citation"`
											} `xml:"authority"`
											Code struct {
												Text            string `xml:",chardata"`
												CharacterString struct {
													Text string `xml:",chardata"` // ISO 												3166-1:20...
												} `xml:"CharacterString"`
											} `xml:"code"`
										} `xml:"MD_Identifier"`
									} `xml:"identifier"`
								} `xml:"CI_Citation"`
							} `xml:"citation"`
						} `xml:"RE_ReferenceSource"`
					} `xml:"sourceCitation"`
				} `xml:"RE_Reference"`
			} `xml:"specificationSource"`
		} `xml:"RE_SubregisterDescription"`
	} `xml:"containedItem"`
}
