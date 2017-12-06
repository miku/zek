package misc

import "encoding/xml"

// TcgaBcr was generated 2017-12-06 01:06:50 by tir on apollo.
type TcgaBcr struct {
	XMLName        xml.Name `xml:"tcga_bcr"`
	Text           string   `xml:",chardata"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	SchemaVersion  string   `xml:"schemaVersion,attr"`
	Bio            string   `xml:"bio,attr"`
	Xsi            string   `xml:"xsi,attr"`
	AttrAdmin      string   `xml:"admin,attr"`
	Shared         string   `xml:"shared,attr"`
	BioShared      string   `xml:"bio_shared,attr"`
	Cqcf           string   `xml:"cqcf,attr"`
	Admin          struct {
		Text string `xml:",chardata"`
		Bcr  struct {
			Text   string `xml:",chardata"` // Nationwide Children's Hos...
			XsdVer string `xml:"xsd_ver,attr"`
		} `xml:"bcr"`
		BatchNumber struct {
			Text   string `xml:",chardata"` // 37.36.0
			XsdVer string `xml:"xsd_ver,attr"`
		} `xml:"batch_number"`
		DiseaseCode struct {
			Text   string `xml:",chardata"` // LUAD
			XsdVer string `xml:"xsd_ver,attr"`
		} `xml:"disease_code"`
		DayOfDccUpload struct {
			Text   string `xml:",chardata"` // 17
			XsdVer string `xml:"xsd_ver,attr"`
		} `xml:"day_of_dcc_upload"`
		MonthOfDccUpload struct {
			Text   string `xml:",chardata"` // 1
			XsdVer string `xml:"xsd_ver,attr"`
		} `xml:"month_of_dcc_upload"`
		YearOfDccUpload struct {
			Text   string `xml:",chardata"` // 2014
			XsdVer string `xml:"xsd_ver,attr"`
		} `xml:"year_of_dcc_upload"`
	} `xml:"admin"`
	Patient struct {
		Text              string `xml:",chardata"`
		BcrPatientBarcode struct {
			Text              string `xml:",chardata"` // TCGA-44-2668
			Cde               string `xml:"cde,attr"`
			XsdVer            string `xml:"xsd_ver,attr"`
			ProcurementStatus string `xml:"procurement_status,attr"`
			Owner             string `xml:"owner,attr"`
		} `xml:"bcr_patient_barcode"`
		BcrPatientUuid struct {
			Text              string `xml:",chardata"` // bab43415-d413-40be-a4c0-2...
			Cde               string `xml:"cde,attr"`
			XsdVer            string `xml:"xsd_ver,attr"`
			ProcurementStatus string `xml:"procurement_status,attr"`
			Owner             string `xml:"owner,attr"`
		} `xml:"bcr_patient_uuid"`
		TissueSourceSite struct {
			Text              string `xml:",chardata"` // 44
			Cde               string `xml:"cde,attr"`
			XsdVer            string `xml:"xsd_ver,attr"`
			ProcurementStatus string `xml:"procurement_status,attr"`
			Owner             string `xml:"owner,attr"`
		} `xml:"tissue_source_site"`
		PatientID struct {
			Text              string `xml:",chardata"` // 2668
			Cde               string `xml:"cde,attr"`
			XsdVer            string `xml:"xsd_ver,attr"`
			ProcurementStatus string `xml:"procurement_status,attr"`
			Owner             string `xml:"owner,attr"`
		} `xml:"patient_id"`
		DaysToIndex struct {
			Text              string `xml:",chardata"` // 0
			Precision         string `xml:"precision,attr"`
			Cde               string `xml:"cde,attr"`
			XsdVer            string `xml:"xsd_ver,attr"`
			ProcurementStatus string `xml:"procurement_status,attr"`
			Owner             string `xml:"owner,attr"`
		} `xml:"days_to_index"`
		BcrCanonicalCheck struct {
			Text                      string `xml:",chardata"`
			BcrPatientCanonicalStatus struct {
				Text              string `xml:",chardata"` // Canonical - Plus
				Cde               string `xml:"cde,attr"`
				XsdVer            string `xml:"xsd_ver,attr"`
				ProcurementStatus string `xml:"procurement_status,attr"`
				Owner             string `xml:"owner,attr"`
			} `xml:"bcr_patient_canonical_status"`
			BcrPatientCanonicalReasons struct {
				Text               string `xml:",chardata"`
				BcrCanonicalReason struct {
					Text              string `xml:",chardata"` // FFPE Validation
					Cde               string `xml:"cde,attr"`
					XsdVer            string `xml:"xsd_ver,attr"`
					ProcurementStatus string `xml:"procurement_status,attr"`
					Owner             string `xml:"owner,attr"`
				} `xml:"bcr_canonical_reason"`
			} `xml:"bcr_patient_canonical_reasons"`
		} `xml:"bcr_canonical_check"`
		Samples struct {
			Text   string `xml:",chardata"`
			Sample []struct {
				Text         string `xml:",chardata"`
				SampleTypeID struct {
					Text              string `xml:",chardata"` // 01, 01, 10, 11
					Cde               string `xml:"cde,attr"`
					XsdVer            string `xml:"xsd_ver,attr"`
					ProcurementStatus string `xml:"procurement_status,attr"`
					Owner             string `xml:"owner,attr"`
				} `xml:"sample_type_id"`
				VialNumber struct {
					Text              string `xml:",chardata"` // A, B, A, A
					Cde               string `xml:"cde,attr"`
					XsdVer            string `xml:"xsd_ver,attr"`
					ProcurementStatus string `xml:"procurement_status,attr"`
					Owner             string `xml:"owner,attr"`
				} `xml:"vial_number"`
				SampleType struct {
					Text              string `xml:",chardata"` // Primary Tumor, Primary Tu...
					Cde               string `xml:"cde,attr"`
					XsdVer            string `xml:"xsd_ver,attr"`
					ProcurementStatus string `xml:"procurement_status,attr"`
					Owner             string `xml:"owner,attr"`
				} `xml:"sample_type"`
				LongestDimension struct {
					Text              string `xml:",chardata"` // 1.2, 1.5
					Cde               string `xml:"cde,attr"`
					XsdVer            string `xml:"xsd_ver,attr"`
					ProcurementStatus string `xml:"procurement_status,attr"`
					Owner             string `xml:"owner,attr"`
				} `xml:"longest_dimension"`
				IntermediateDimension struct {
					Text              string `xml:",chardata"` // 1, 1
					Cde               string `xml:"cde,attr"`
					XsdVer            string `xml:"xsd_ver,attr"`
					ProcurementStatus string `xml:"procurement_status,attr"`
					Owner             string `xml:"owner,attr"`
				} `xml:"intermediate_dimension"`
				ShortestDimension struct {
					Text              string `xml:",chardata"` // 0.3, 0.5
					Cde               string `xml:"cde,attr"`
					XsdVer            string `xml:"xsd_ver,attr"`
					ProcurementStatus string `xml:"procurement_status,attr"`
					Owner             string `xml:"owner,attr"`
				} `xml:"shortest_dimension"`
				InitialWeight struct {
					Text              string `xml:",chardata"`
					Cde               string `xml:"cde,attr"`
					XsdVer            string `xml:"xsd_ver,attr"`
					ProcurementStatus string `xml:"procurement_status,attr"`
					Owner             string `xml:"owner,attr"`
				} `xml:"initial_weight"`
				CurrentWeight struct {
					Text              string `xml:",chardata"`
					Cde               string `xml:"cde,attr"`
					XsdVer            string `xml:"xsd_ver,attr"`
					ProcurementStatus string `xml:"procurement_status,attr"`
					Owner             string `xml:"owner,attr"`
				} `xml:"current_weight"`
				FreezingMethod struct {
					Text              string `xml:",chardata"`
					Cde               string `xml:"cde,attr"`
					XsdVer            string `xml:"xsd_ver,attr"`
					ProcurementStatus string `xml:"procurement_status,attr"`
					Owner             string `xml:"owner,attr"`
				} `xml:"freezing_method"`
				OctEmbedded struct {
					Text              string `xml:",chardata"` // false
					Cde               string `xml:"cde,attr"`
					XsdVer            string `xml:"xsd_ver,attr"`
					ProcurementStatus string `xml:"procurement_status,attr"`
					Owner             string `xml:"owner,attr"`
				} `xml:"oct_embedded"`
				DaysToCollection struct {
					Text              string `xml:",chardata"` // 776
					Cde               string `xml:"cde,attr"`
					XsdVer            string `xml:"xsd_ver,attr"`
					ProcurementStatus string `xml:"procurement_status,attr"`
					Owner             string `xml:"owner,attr"`
					Precision         string `xml:"precision,attr"`
				} `xml:"days_to_collection"`
				TimeBetweenClampingAndFreezing struct {
					Text              string `xml:",chardata"`
					Cde               string `xml:"cde,attr"`
					XsdVer            string `xml:"xsd_ver,attr"`
					ProcurementStatus string `xml:"procurement_status,attr"`
					Owner             string `xml:"owner,attr"`
				} `xml:"time_between_clamping_and_freezing"`
				TimeBetweenExcisionAndFreezing struct {
					Text              string `xml:",chardata"`
					Cde               string `xml:"cde,attr"`
					XsdVer            string `xml:"xsd_ver,attr"`
					ProcurementStatus string `xml:"procurement_status,attr"`
					Owner             string `xml:"owner,attr"`
				} `xml:"time_between_excision_and_freezing"`
				BcrSampleBarcode struct {
					Text              string `xml:",chardata"` // TCGA-44-2668-01A, TCGA-44...
					Cde               string `xml:"cde,attr"`
					XsdVer            string `xml:"xsd_ver,attr"`
					ProcurementStatus string `xml:"procurement_status,attr"`
					Owner             string `xml:"owner,attr"`
				} `xml:"bcr_sample_barcode"`
				BcrSampleUuid struct {
					Text              string `xml:",chardata"` // dd9a6c68-b8b4-4168-9ff9-7...
					Cde               string `xml:"cde,attr"`
					XsdVer            string `xml:"xsd_ver,attr"`
					ProcurementStatus string `xml:"procurement_status,attr"`
					Owner             string `xml:"owner,attr"`
				} `xml:"bcr_sample_uuid"`
				Portions struct {
					Text            string `xml:",chardata"`
					ShipmentPortion struct {
						Text          string `xml:",chardata"`
						PortionNumber struct {
							Text              string `xml:",chardata"` // 21
							Cde               string `xml:"cde,attr"`
							XsdVer            string `xml:"xsd_ver,attr"`
							ProcurementStatus string `xml:"procurement_status,attr"`
							Owner             string `xml:"owner,attr"`
						} `xml:"portion_number"`
						PlateID struct {
							Text              string `xml:",chardata"` // 2190
							Cde               string `xml:"cde,attr"`
							XsdVer            string `xml:"xsd_ver,attr"`
							ProcurementStatus string `xml:"procurement_status,attr"`
							Owner             string `xml:"owner,attr"`
						} `xml:"plate_id"`
						CenterID struct {
							Text              string `xml:",chardata"` // 20
							Cde               string `xml:"cde,attr"`
							XsdVer            string `xml:"xsd_ver,attr"`
							ProcurementStatus string `xml:"procurement_status,attr"`
							Owner             string `xml:"owner,attr"`
						} `xml:"center_id"`
						ShipmentPortionDayOfShipment struct {
							Text              string `xml:",chardata"` // 5
							Cde               string `xml:"cde,attr"`
							XsdVer            string `xml:"xsd_ver,attr"`
							ProcurementStatus string `xml:"procurement_status,attr"`
							Owner             string `xml:"owner,attr"`
						} `xml:"shipment_portion_day_of_shipment"`
						ShipmentPortionMonthOfShipment struct {
							Text              string `xml:",chardata"` // 3
							Cde               string `xml:"cde,attr"`
							XsdVer            string `xml:"xsd_ver,attr"`
							ProcurementStatus string `xml:"procurement_status,attr"`
							Owner             string `xml:"owner,attr"`
						} `xml:"shipment_portion_month_of_shipment"`
						ShipmentPortionYearOfShipment struct {
							Text              string `xml:",chardata"` // 2012
							Cde               string `xml:"cde,attr"`
							XsdVer            string `xml:"xsd_ver,attr"`
							ProcurementStatus string `xml:"procurement_status,attr"`
							Owner             string `xml:"owner,attr"`
						} `xml:"shipment_portion_year_of_shipment"`
						ShipmentPortionBcrAliquotBarcode struct {
							Text              string `xml:",chardata"` // TCGA-44-2668-01A-21-2190-...
							Cde               string `xml:"cde,attr"`
							XsdVer            string `xml:"xsd_ver,attr"`
							ProcurementStatus string `xml:"procurement_status,attr"`
							Owner             string `xml:"owner,attr"`
						} `xml:"shipment_portion_bcr_aliquot_barcode"`
						BcrShipmentPortionUuid struct {
							Text              string `xml:",chardata"` // 2e058de2-768c-4a90-9105-a...
							Cde               string `xml:"cde,attr"`
							XsdVer            string `xml:"xsd_ver,attr"`
							ProcurementStatus string `xml:"procurement_status,attr"`
							Owner             string `xml:"owner,attr"`
						} `xml:"bcr_shipment_portion_uuid"`
					} `xml:"shipment_portion"`
					Portion struct {
						Text          string `xml:",chardata"`
						PortionNumber struct {
							Text              string `xml:",chardata"` // 01, 02, 01, 01
							Cde               string `xml:"cde,attr"`
							XsdVer            string `xml:"xsd_ver,attr"`
							ProcurementStatus string `xml:"procurement_status,attr"`
							Owner             string `xml:"owner,attr"`
						} `xml:"portion_number"`
						DayOfCreation struct {
							Text              string `xml:",chardata"` // 30, 30, 22, 16
							Cde               string `xml:"cde,attr"`
							XsdVer            string `xml:"xsd_ver,attr"`
							ProcurementStatus string `xml:"procurement_status,attr"`
							Owner             string `xml:"owner,attr"`
						} `xml:"day_of_creation"`
						MonthOfCreation struct {
							Text              string `xml:",chardata"` // 12, 01, 01, 12
							Cde               string `xml:"cde,attr"`
							XsdVer            string `xml:"xsd_ver,attr"`
							ProcurementStatus string `xml:"procurement_status,attr"`
							Owner             string `xml:"owner,attr"`
						} `xml:"month_of_creation"`
						YearOfCreation struct {
							Text              string `xml:",chardata"` // 2009, 2013, 2010, 2010
							Cde               string `xml:"cde,attr"`
							XsdVer            string `xml:"xsd_ver,attr"`
							ProcurementStatus string `xml:"procurement_status,attr"`
							Owner             string `xml:"owner,attr"`
						} `xml:"year_of_creation"`
						Weight struct {
							Text              string `xml:",chardata"` // 173.00, 400.00, 67.00
							Cde               string `xml:"cde,attr"`
							XsdVer            string `xml:"xsd_ver,attr"`
							ProcurementStatus string `xml:"procurement_status,attr"`
							Owner             string `xml:"owner,attr"`
						} `xml:"weight"`
						BcrPortionBarcode struct {
							Text              string `xml:",chardata"` // TCGA-44-2668-01A-01, TCGA...
							Cde               string `xml:"cde,attr"`
							XsdVer            string `xml:"xsd_ver,attr"`
							ProcurementStatus string `xml:"procurement_status,attr"`
							Owner             string `xml:"owner,attr"`
						} `xml:"bcr_portion_barcode"`
						BcrPortionUuid struct {
							Text              string `xml:",chardata"` // 12fbb582-96c7-4dfc-b6ef-8...
							Cde               string `xml:"cde,attr"`
							XsdVer            string `xml:"xsd_ver,attr"`
							ProcurementStatus string `xml:"procurement_status,attr"`
							Owner             string `xml:"owner,attr"`
						} `xml:"bcr_portion_uuid"`
						Analytes struct {
							Text    string `xml:",chardata"`
							Analyte []struct {
								Text          string `xml:",chardata"`
								AnalyteTypeID struct {
									Text              string `xml:",chardata"` // D, R, T, W, X, D, R, D, W...
									Cde               string `xml:"cde,attr"`
									XsdVer            string `xml:"xsd_ver,attr"`
									ProcurementStatus string `xml:"procurement_status,attr"`
									Owner             string `xml:"owner,attr"`
								} `xml:"analyte_type_id"`
								AnalyteType struct {
									Text              string `xml:",chardata"` // DNA, RNA, Total RNA, Repl...
									Cde               string `xml:"cde,attr"`
									XsdVer            string `xml:"xsd_ver,attr"`
									ProcurementStatus string `xml:"procurement_status,attr"`
									Owner             string `xml:"owner,attr"`
								} `xml:"analyte_type"`
								Concentration struct {
									Text              string `xml:",chardata"` // 0.14, 0.14, 0.10, 0.50, 0...
									Cde               string `xml:"cde,attr"`
									XsdVer            string `xml:"xsd_ver,attr"`
									ProcurementStatus string `xml:"procurement_status,attr"`
									Owner             string `xml:"owner,attr"`
								} `xml:"concentration"`
								Amount struct {
									Text              string `xml:",chardata"` // 172.93, 570.30, 66.00, 42...
									Cde               string `xml:"cde,attr"`
									XsdVer            string `xml:"xsd_ver,attr"`
									ProcurementStatus string `xml:"procurement_status,attr"`
									Owner             string `xml:"owner,attr"`
								} `xml:"amount"`
								A260A280Ratio struct {
									Text              string `xml:",chardata"` // 1.90, 1.80, 1.83, 1.90, 1...
									Cde               string `xml:"cde,attr"`
									XsdVer            string `xml:"xsd_ver,attr"`
									ProcurementStatus string `xml:"procurement_status,attr"`
									Owner             string `xml:"owner,attr"`
								} `xml:"a260_a280_ratio"`
								GelImageFile struct {
									Text              string `xml:",chardata"` // https://sharedoc.nchri.or...
									Cde               string `xml:"cde,attr"`
									XsdVer            string `xml:"xsd_ver,attr"`
									ProcurementStatus string `xml:"procurement_status,attr"`
									Owner             string `xml:"owner,attr"`
								} `xml:"gel_image_file"`
								WellNumber struct {
									Text              string `xml:",chardata"`
									Cde               string `xml:"cde,attr"`
									XsdVer            string `xml:"xsd_ver,attr"`
									ProcurementStatus string `xml:"procurement_status,attr"`
									Owner             string `xml:"owner,attr"`
								} `xml:"well_number"`
								BcrAnalyteBarcode struct {
									Text              string `xml:",chardata"` // TCGA-44-2668-01A-01D, TCG...
									Cde               string `xml:"cde,attr"`
									XsdVer            string `xml:"xsd_ver,attr"`
									ProcurementStatus string `xml:"procurement_status,attr"`
									Owner             string `xml:"owner,attr"`
								} `xml:"bcr_analyte_barcode"`
								BcrAnalyteUuid struct {
									Text              string `xml:",chardata"` // 55c2032b-ad6e-4791-93a0-f...
									Cde               string `xml:"cde,attr"`
									XsdVer            string `xml:"xsd_ver,attr"`
									ProcurementStatus string `xml:"procurement_status,attr"`
									Owner             string `xml:"owner,attr"`
								} `xml:"bcr_analyte_uuid"`
								Aliquots struct {
									Text    string `xml:",chardata"`
									Aliquot []struct {
										Text    string `xml:",chardata"`
										PlateID struct {
											Text              string `xml:",chardata"` // 0944, 0945, 0969, 0970, 1...
											Cde               string `xml:"cde,attr"`
											XsdVer            string `xml:"xsd_ver,attr"`
											ProcurementStatus string `xml:"procurement_status,attr"`
											Owner             string `xml:"owner,attr"`
										} `xml:"plate_id"`
										CenterID struct {
											Text              string `xml:",chardata"` // 01, 05, 08, 09, 02, 01, 0...
											Cde               string `xml:"cde,attr"`
											XsdVer            string `xml:"xsd_ver,attr"`
											ProcurementStatus string `xml:"procurement_status,attr"`
											Owner             string `xml:"owner,attr"`
										} `xml:"center_id"`
										Amount struct {
											Text              string `xml:",chardata"` // 6.70, 26.70, 6.70, 6.70, ...
											Cde               string `xml:"cde,attr"`
											XsdVer            string `xml:"xsd_ver,attr"`
											ProcurementStatus string `xml:"procurement_status,attr"`
											Owner             string `xml:"owner,attr"`
										} `xml:"amount"`
										DayOfShipment struct {
											Text              string `xml:",chardata"` // 21, 21, 28, 28, 14, 14, 1...
											Cde               string `xml:"cde,attr"`
											XsdVer            string `xml:"xsd_ver,attr"`
											ProcurementStatus string `xml:"procurement_status,attr"`
											Owner             string `xml:"owner,attr"`
										} `xml:"day_of_shipment"`
										MonthOfShipment struct {
											Text              string `xml:",chardata"` // 04, 04, 04, 04, 02, 02, 0...
											Cde               string `xml:"cde,attr"`
											XsdVer            string `xml:"xsd_ver,attr"`
											ProcurementStatus string `xml:"procurement_status,attr"`
											Owner             string `xml:"owner,attr"`
										} `xml:"month_of_shipment"`
										YearOfShipment struct {
											Text              string `xml:",chardata"` // 2010, 2010, 2010, 2010, 2...
											Cde               string `xml:"cde,attr"`
											XsdVer            string `xml:"xsd_ver,attr"`
											ProcurementStatus string `xml:"procurement_status,attr"`
											Owner             string `xml:"owner,attr"`
										} `xml:"year_of_shipment"`
										BcrAliquotBarcode struct {
											Text              string `xml:",chardata"` // TCGA-44-2668-01A-01D-0944...
											Cde               string `xml:"cde,attr"`
											XsdVer            string `xml:"xsd_ver,attr"`
											ProcurementStatus string `xml:"procurement_status,attr"`
											Owner             string `xml:"owner,attr"`
										} `xml:"bcr_aliquot_barcode"`
										BcrAliquotUuid struct {
											Text              string `xml:",chardata"` // ce8f9ffd-122d-465d-a677-7...
											Cde               string `xml:"cde,attr"`
											XsdVer            string `xml:"xsd_ver,attr"`
											ProcurementStatus string `xml:"procurement_status,attr"`
											Owner             string `xml:"owner,attr"`
										} `xml:"bcr_aliquot_uuid"`
										Concentration struct {
											Text              string `xml:",chardata"` // 0.14, 0.14, 0.14, 0.14, 0...
											Cde               string `xml:"cde,attr"`
											XsdVer            string `xml:"xsd_ver,attr"`
											ProcurementStatus string `xml:"procurement_status,attr"`
											Owner             string `xml:"owner,attr"`
										} `xml:"concentration"`
										PlateRow struct {
											Text              string `xml:",chardata"` // H, H, A, A, E, E, E, E, E...
											Cde               string `xml:"cde,attr"`
											XsdVer            string `xml:"xsd_ver,attr"`
											ProcurementStatus string `xml:"procurement_status,attr"`
											Owner             string `xml:"owner,attr"`
										} `xml:"plate_row"`
										PlateColumn struct {
											Text              string `xml:",chardata"` // 1, 1, 2, 2, 2, 2, 2, 2, 2...
											Cde               string `xml:"cde,attr"`
											XsdVer            string `xml:"xsd_ver,attr"`
											ProcurementStatus string `xml:"procurement_status,attr"`
											Owner             string `xml:"owner,attr"`
										} `xml:"plate_column"`
										BiospecimenBarcodeBottom struct {
											Text              string `xml:",chardata"` // 0096491719, 0097517860, 0...
											Cde               string `xml:"cde,attr"`
											XsdVer            string `xml:"xsd_ver,attr"`
											ProcurementStatus string `xml:"procurement_status,attr"`
											Owner             string `xml:"owner,attr"`
										} `xml:"biospecimen_barcode_bottom"`
										SourceCenter struct {
											Text              string `xml:",chardata"` // 22, 22, 22, 22, 22, 22, 2...
											Cde               string `xml:"cde,attr"`
											XsdVer            string `xml:"xsd_ver,attr"`
											ProcurementStatus string `xml:"procurement_status,attr"`
											Owner             string `xml:"owner,attr"`
										} `xml:"source_center"`
										BcrBiospecimenCanonicalReasons struct {
											Text               string `xml:",chardata"`
											BcrCanonicalReason struct {
												Text              string `xml:",chardata"` // FFPE Validation, FFPE Val...
												Cde               string `xml:"cde,attr"`
												XsdVer            string `xml:"xsd_ver,attr"`
												ProcurementStatus string `xml:"procurement_status,attr"`
												Owner             string `xml:"owner,attr"`
											} `xml:"bcr_canonical_reason"`
										} `xml:"bcr_biospecimen_canonical_reasons"`
									} `xml:"aliquot"`
								} `xml:"aliquots"`
								Protocols struct {
									Text     string `xml:",chardata"`
									Protocol struct {
										Text                     string `xml:",chardata"`
										ExperimentalProtocolType struct {
											Text              string `xml:",chardata"` // aDNA Preparation Type, Al...
											Cde               string `xml:"cde,attr"`
											XsdVer            string `xml:"xsd_ver,attr"`
											ProcurementStatus string `xml:"procurement_status,attr"`
											Owner             string `xml:"owner,attr"`
										} `xml:"experimental_protocol_type"`
										ProtocolName struct {
											Text              string `xml:",chardata"`
											Cde               string `xml:"cde,attr"`
											XsdVer            string `xml:"xsd_ver,attr"`
											ProcurementStatus string `xml:"procurement_status,attr"`
											Owner             string `xml:"owner,attr"`
										} `xml:"protocol_name"`
										ProtocolFileName struct {
											Text              string `xml:",chardata"`
											Cde               string `xml:"cde,attr"`
											XsdVer            string `xml:"xsd_ver,attr"`
											ProcurementStatus string `xml:"procurement_status,attr"`
											Owner             string `xml:"owner,attr"`
										} `xml:"protocol_file_name"`
									} `xml:"protocol"`
								} `xml:"protocols"`
								Dna struct {
									Text                     string `xml:",chardata"`
									NormalTumorGenotypeMatch struct {
										Text              string `xml:",chardata"` // YES, YES, YES, YES, YES, ...
										Cde               string `xml:"cde,attr"`
										XsdVer            string `xml:"xsd_ver,attr"`
										ProcurementStatus string `xml:"procurement_status,attr"`
										Owner             string `xml:"owner,attr"`
									} `xml:"normal_tumor_genotype_match"`
									PcrAmplificationSuccessful struct {
										Text              string `xml:",chardata"` // YES, NO, NO, YES, YES, NO...
										Cde               string `xml:"cde,attr"`
										XsdVer            string `xml:"xsd_ver,attr"`
										ProcurementStatus string `xml:"procurement_status,attr"`
										Owner             string `xml:"owner,attr"`
									} `xml:"pcr_amplification_successful"`
								} `xml:"dna"`
								SpectrophotometerMethod struct {
									Text              string `xml:",chardata"` // UV Spec, UV Spec, UV Spec...
									Cde               string `xml:"cde,attr"`
									XsdVer            string `xml:"xsd_ver,attr"`
									ProcurementStatus string `xml:"procurement_status,attr"`
									Owner             string `xml:"owner,attr"`
								} `xml:"spectrophotometer_method"`
								Rna struct {
									Text        string `xml:",chardata"`
									Ratio28s18s struct {
										Text              string `xml:",chardata"` // 1.60, 0.00, 1.85
										Cde               string `xml:"cde,attr"`
										XsdVer            string `xml:"xsd_ver,attr"`
										ProcurementStatus string `xml:"procurement_status,attr"`
										Owner             string `xml:"owner,attr"`
									} `xml:"ratio_28s_18s"`
									Rinvalue struct {
										Text              string `xml:",chardata"` // 9.00, 2.40, 8.40
										Cde               string `xml:"cde,attr"`
										XsdVer            string `xml:"xsd_ver,attr"`
										ProcurementStatus string `xml:"procurement_status,attr"`
										Owner             string `xml:"owner,attr"`
									} `xml:"rinvalue"`
								} `xml:"rna"`
							} `xml:"analyte"`
						} `xml:"analytes"`
						Slides struct {
							Text  string `xml:",chardata"`
							Slide []struct {
								Text            string `xml:",chardata"`
								SectionLocation struct {
									Text              string `xml:",chardata"` // BOTTOM, TOP, TOP, TOP
									Cde               string `xml:"cde,attr"`
									XsdVer            string `xml:"xsd_ver,attr"`
									ProcurementStatus string `xml:"procurement_status,attr"`
									Owner             string `xml:"owner,attr"`
								} `xml:"section_location"`
								NumberProliferatingCells struct {
									Text              string `xml:",chardata"`
									Cde               string `xml:"cde,attr"`
									XsdVer            string `xml:"xsd_ver,attr"`
									ProcurementStatus string `xml:"procurement_status,attr"`
									Owner             string `xml:"owner,attr"`
								} `xml:"number_proliferating_cells"`
								PercentTumorCells struct {
									Text              string `xml:",chardata"` // 80, 70, 50, 0
									Cde               string `xml:"cde,attr"`
									XsdVer            string `xml:"xsd_ver,attr"`
									ProcurementStatus string `xml:"procurement_status,attr"`
									Owner             string `xml:"owner,attr"`
								} `xml:"percent_tumor_cells"`
								PercentTumorNuclei struct {
									Text              string `xml:",chardata"` // 85, 70, 60, 0
									Cde               string `xml:"cde,attr"`
									XsdVer            string `xml:"xsd_ver,attr"`
									ProcurementStatus string `xml:"procurement_status,attr"`
									Owner             string `xml:"owner,attr"`
								} `xml:"percent_tumor_nuclei"`
								PercentNormalCells struct {
									Text              string `xml:",chardata"` // 0, 0, 0, 0
									Cde               string `xml:"cde,attr"`
									XsdVer            string `xml:"xsd_ver,attr"`
									ProcurementStatus string `xml:"procurement_status,attr"`
									Owner             string `xml:"owner,attr"`
								} `xml:"percent_normal_cells"`
								PercentNecrosis struct {
									Text              string `xml:",chardata"` // 0, 0, 20, 0
									Cde               string `xml:"cde,attr"`
									XsdVer            string `xml:"xsd_ver,attr"`
									ProcurementStatus string `xml:"procurement_status,attr"`
									Owner             string `xml:"owner,attr"`
								} `xml:"percent_necrosis"`
								PercentStromalCells struct {
									Text              string `xml:",chardata"` // 20, 30, 30, 0
									Cde               string `xml:"cde,attr"`
									XsdVer            string `xml:"xsd_ver,attr"`
									ProcurementStatus string `xml:"procurement_status,attr"`
									Owner             string `xml:"owner,attr"`
								} `xml:"percent_stromal_cells"`
								PercentInflamInfiltration struct {
									Text              string `xml:",chardata"`
									Cde               string `xml:"cde,attr"`
									XsdVer            string `xml:"xsd_ver,attr"`
									ProcurementStatus string `xml:"procurement_status,attr"`
									Owner             string `xml:"owner,attr"`
								} `xml:"percent_inflam_infiltration"`
								PercentLymphocyteInfiltration struct {
									Text              string `xml:",chardata"` // 10, 20, 50, 0
									Cde               string `xml:"cde,attr"`
									XsdVer            string `xml:"xsd_ver,attr"`
									ProcurementStatus string `xml:"procurement_status,attr"`
									Owner             string `xml:"owner,attr"`
								} `xml:"percent_lymphocyte_infiltration"`
								PercentMonocyteInfiltration struct {
									Text              string `xml:",chardata"` // 0, 0, 20, 0
									Cde               string `xml:"cde,attr"`
									XsdVer            string `xml:"xsd_ver,attr"`
									ProcurementStatus string `xml:"procurement_status,attr"`
									Owner             string `xml:"owner,attr"`
								} `xml:"percent_monocyte_infiltration"`
								PercentGranulocyteInfiltration struct {
									Text              string `xml:",chardata"`
									Cde               string `xml:"cde,attr"`
									XsdVer            string `xml:"xsd_ver,attr"`
									ProcurementStatus string `xml:"procurement_status,attr"`
									Owner             string `xml:"owner,attr"`
								} `xml:"percent_granulocyte_infiltration"`
								PercentNeutrophilInfiltration struct {
									Text              string `xml:",chardata"` // 0, 0, 30, 0
									Cde               string `xml:"cde,attr"`
									XsdVer            string `xml:"xsd_ver,attr"`
									ProcurementStatus string `xml:"procurement_status,attr"`
									Owner             string `xml:"owner,attr"`
								} `xml:"percent_neutrophil_infiltration"`
								PercentEosinophilInfiltration struct {
									Text              string `xml:",chardata"`
									Cde               string `xml:"cde,attr"`
									XsdVer            string `xml:"xsd_ver,attr"`
									ProcurementStatus string `xml:"procurement_status,attr"`
									Owner             string `xml:"owner,attr"`
								} `xml:"percent_eosinophil_infiltration"`
								BcrSlideBarcode struct {
									Text              string `xml:",chardata"` // TCGA-44-2668-01A-01-BS1, ...
									Cde               string `xml:"cde,attr"`
									XsdVer            string `xml:"xsd_ver,attr"`
									ProcurementStatus string `xml:"procurement_status,attr"`
									Owner             string `xml:"owner,attr"`
								} `xml:"bcr_slide_barcode"`
								BcrSlideUuid struct {
									Text              string `xml:",chardata"` // 98d005af-6c76-4072-9141-7...
									Cde               string `xml:"cde,attr"`
									XsdVer            string `xml:"xsd_ver,attr"`
									ProcurementStatus string `xml:"procurement_status,attr"`
									Owner             string `xml:"owner,attr"`
								} `xml:"bcr_slide_uuid"`
							} `xml:"slide"`
						} `xml:"slides"`
						LCE struct {
							Text              string `xml:",chardata"`
							Cde               string `xml:"cde,attr"`
							XsdVer            string `xml:"xsd_ver,attr"`
							ProcurementStatus string `xml:"procurement_status,attr"`
							Owner             string `xml:"owner,attr"`
						} `xml:"LCE"`
					} `xml:"portion"`
				} `xml:"portions"`
				TumorPathology struct {
					Text string `xml:",chardata"`
				} `xml:"tumor_pathology"`
				MethodOfSampleProcurement struct {
					Text              string `xml:",chardata"`
					Cde               string `xml:"cde,attr"`
					XsdVer            string `xml:"xsd_ver,attr"`
					ProcurementStatus string `xml:"procurement_status,attr"`
					Owner             string `xml:"owner,attr"`
				} `xml:"method_of_sample_procurement"`
				NCNNCTOthMethONSP struct {
					Text              string `xml:",chardata"`
					Cde               string `xml:"cde,attr"`
					XsdVer            string `xml:"xsd_ver,attr"`
					ProcurementStatus string `xml:"procurement_status,attr"`
					Owner             string `xml:"owner,attr"`
				} `xml:"NCNNCT_OthMethONSP"`
				OtherMethodOfSampleProcurement struct {
					Text              string `xml:",chardata"`
					Cde               string `xml:"cde,attr"`
					XsdVer            string `xml:"xsd_ver,attr"`
					ProcurementStatus string `xml:"procurement_status,attr"`
					Owner             string `xml:"owner,attr"`
				} `xml:"other_method_of_sample_procurement"`
				DaysToSampleProcurement struct {
					Text              string `xml:",chardata"`
					Cde               string `xml:"cde,attr"`
					XsdVer            string `xml:"xsd_ver,attr"`
					ProcurementStatus string `xml:"procurement_status,attr"`
					Owner             string `xml:"owner,attr"`
				} `xml:"days_to_sample_procurement"`
				PathologyReportUuid struct {
					Text              string `xml:",chardata"` // d5ebaec8-79c6-493b-ae2a-b...
					Cde               string `xml:"cde,attr"`
					XsdVer            string `xml:"xsd_ver,attr"`
					ProcurementStatus string `xml:"procurement_status,attr"`
					Owner             string `xml:"owner,attr"`
				} `xml:"pathology_report_uuid"`
				DiagnosticSlides struct {
					Text          string `xml:",chardata"`
					FfpeSlideUuid struct {
						Text              string `xml:",chardata"` // DBD6EEC8-071A-43E0-9A98-D...
						Cde               string `xml:"cde,attr"`
						XsdVer            string `xml:"xsd_ver,attr"`
						ProcurementStatus string `xml:"procurement_status,attr"`
						Owner             string `xml:"owner,attr"`
					} `xml:"ffpe_slide_uuid"`
				} `xml:"diagnostic_slides"`
			} `xml:"sample"`
		} `xml:"samples"`
		BiospecimenCqcf struct {
			Text                  string `xml:",chardata"`
			MaximumTumorDimension struct {
				Text              string `xml:",chardata"`
				Cde               string `xml:"cde,attr"`
				XsdVer            string `xml:"xsd_ver,attr"`
				ProcurementStatus string `xml:"procurement_status,attr"`
				Owner             string `xml:"owner,attr"`
			} `xml:"maximum_tumor_dimension"`
			DaysToSampleProcurement struct {
				Text              string `xml:",chardata"` // 0
				Precision         string `xml:"precision,attr"`
				Cde               string `xml:"cde,attr"`
				XsdVer            string `xml:"xsd_ver,attr"`
				ProcurementStatus string `xml:"procurement_status,attr"`
				Owner             string `xml:"owner,attr"`
			} `xml:"days_to_sample_procurement"`
			MethodOfSampleProcurement struct {
				Text              string `xml:",chardata"` // Tumor Resection
				Cde               string `xml:"cde,attr"`
				XsdVer            string `xml:"xsd_ver,attr"`
				ProcurementStatus string `xml:"procurement_status,attr"`
				Owner             string `xml:"owner,attr"`
			} `xml:"method_of_sample_procurement"`
			OtherMethodOfSampleProcurement struct {
				Text              string `xml:",chardata"`
				Cde               string `xml:"cde,attr"`
				XsdVer            string `xml:"xsd_ver,attr"`
				ProcurementStatus string `xml:"procurement_status,attr"`
				Owner             string `xml:"owner,attr"`
			} `xml:"other_method_of_sample_procurement"`
			Country struct {
				Text              string `xml:",chardata"` // United States
				Cde               string `xml:"cde,attr"`
				XsdVer            string `xml:"xsd_ver,attr"`
				ProcurementStatus string `xml:"procurement_status,attr"`
				Owner             string `xml:"owner,attr"`
			} `xml:"country"`
			VesselUsed struct {
				Text              string `xml:",chardata"` // Cryovial
				Cde               string `xml:"cde,attr"`
				XsdVer            string `xml:"xsd_ver,attr"`
				ProcurementStatus string `xml:"procurement_status,attr"`
				Owner             string `xml:"owner,attr"`
			} `xml:"vessel_used"`
			OtherVesselUsed struct {
				Text              string `xml:",chardata"`
				Cde               string `xml:"cde,attr"`
				XsdVer            string `xml:"xsd_ver,attr"`
				ProcurementStatus string `xml:"procurement_status,attr"`
				Owner             string `xml:"owner,attr"`
			} `xml:"other_vessel_used"`
			SubmittedForLce struct {
				Text              string `xml:",chardata"`
				Cde               string `xml:"cde,attr"`
				XsdVer            string `xml:"xsd_ver,attr"`
				ProcurementStatus string `xml:"procurement_status,attr"`
				Owner             string `xml:"owner,attr"`
			} `xml:"submitted_for_lce"`
			SamplePrescreened struct {
				Text              string `xml:",chardata"` // YES
				Cde               string `xml:"cde,attr"`
				XsdVer            string `xml:"xsd_ver,attr"`
				ProcurementStatus string `xml:"procurement_status,attr"`
				Owner             string `xml:"owner,attr"`
			} `xml:"sample_prescreened"`
			TopSlideSubmitted struct {
				Text              string `xml:",chardata"` // YES
				Cde               string `xml:"cde,attr"`
				XsdVer            string `xml:"xsd_ver,attr"`
				ProcurementStatus string `xml:"procurement_status,attr"`
				Owner             string `xml:"owner,attr"`
			} `xml:"top_slide_submitted"`
			DigitalImageSubmitted struct {
				Text              string `xml:",chardata"` // NO
				Cde               string `xml:"cde,attr"`
				XsdVer            string `xml:"xsd_ver,attr"`
				ProcurementStatus string `xml:"procurement_status,attr"`
				Owner             string `xml:"owner,attr"`
			} `xml:"digital_image_submitted"`
			FfpeTumorSlideSubmitted struct {
				Text              string `xml:",chardata"`
				Cde               string `xml:"cde,attr"`
				XsdVer            string `xml:"xsd_ver,attr"`
				ProcurementStatus string `xml:"procurement_status,attr"`
				Owner             string `xml:"owner,attr"`
			} `xml:"ffpe_tumor_slide_submitted"`
			BCellTumorSlideSubmitted struct {
				Text              string `xml:",chardata"`
				Cde               string `xml:"cde,attr"`
				XsdVer            string `xml:"xsd_ver,attr"`
				ProcurementStatus string `xml:"procurement_status,attr"`
				Owner             string `xml:"owner,attr"`
			} `xml:"b_cell_tumor_slide_submitted"`
			TCellTumorSlideSubmitted struct {
				Text              string `xml:",chardata"`
				Cde               string `xml:"cde,attr"`
				XsdVer            string `xml:"xsd_ver,attr"`
				ProcurementStatus string `xml:"procurement_status,attr"`
				Owner             string `xml:"owner,attr"`
			} `xml:"t_cell_tumor_slide_submitted"`
			TumorSamples struct {
				Text        string `xml:",chardata"`
				TumorSample struct {
					Text       string `xml:",chardata"`
					VialNumber struct {
						Text              string `xml:",chardata"` // 1
						Cde               string `xml:"cde,attr"`
						XsdVer            string `xml:"xsd_ver,attr"`
						ProcurementStatus string `xml:"procurement_status,attr"`
						Owner             string `xml:"owner,attr"`
					} `xml:"vial_number"`
					BcrSampleUuid struct {
						Text              string `xml:",chardata"` // dd9a6c68-b8b4-4168-9ff9-7...
						Cde               string `xml:"cde,attr"`
						XsdVer            string `xml:"xsd_ver,attr"`
						ProcurementStatus string `xml:"procurement_status,attr"`
						Owner             string `xml:"owner,attr"`
					} `xml:"bcr_sample_uuid"`
					TumorWeight struct {
						Text              string `xml:",chardata"` // 400
						Cde               string `xml:"cde,attr"`
						XsdVer            string `xml:"xsd_ver,attr"`
						ProcurementStatus string `xml:"procurement_status,attr"`
						Owner             string `xml:"owner,attr"`
					} `xml:"tumor_weight"`
					TumorNucleiPercent struct {
						Text              string `xml:",chardata"` // 75
						Cde               string `xml:"cde,attr"`
						XsdVer            string `xml:"xsd_ver,attr"`
						ProcurementStatus string `xml:"procurement_status,attr"`
						Owner             string `xml:"owner,attr"`
					} `xml:"tumor_nuclei_percent"`
					TumorNecrosisPercent struct {
						Text              string `xml:",chardata"` // 10
						Cde               string `xml:"cde,attr"`
						XsdVer            string `xml:"xsd_ver,attr"`
						ProcurementStatus string `xml:"procurement_status,attr"`
						Owner             string `xml:"owner,attr"`
					} `xml:"tumor_necrosis_percent"`
					TumorPathology struct {
						Text string `xml:",chardata"`
					} `xml:"tumor_pathology"`
				} `xml:"tumor_sample"`
			} `xml:"tumor_samples"`
			NormalControls struct {
				Text          string `xml:",chardata"`
				NormalControl []struct {
					Text       string `xml:",chardata"`
					VialNumber struct {
						Text              string `xml:",chardata"` // 1, 2
						Cde               string `xml:"cde,attr"`
						XsdVer            string `xml:"xsd_ver,attr"`
						ProcurementStatus string `xml:"procurement_status,attr"`
						Owner             string `xml:"owner,attr"`
					} `xml:"vial_number"`
					BcrSampleUuid struct {
						Text              string `xml:",chardata"` // 275b1d31-bc89-4e56-9861-6...
						Cde               string `xml:"cde,attr"`
						XsdVer            string `xml:"xsd_ver,attr"`
						ProcurementStatus string `xml:"procurement_status,attr"`
						Owner             string `xml:"owner,attr"`
					} `xml:"bcr_sample_uuid"`
					NormalControlType struct {
						Text              string `xml:",chardata"`
						Cde               string `xml:"cde,attr"`
						XsdVer            string `xml:"xsd_ver,attr"`
						ProcurementStatus string `xml:"procurement_status,attr"`
						Owner             string `xml:"owner,attr"`
					} `xml:"normal_control_type"`
					MethodOfNormalSampleProcurement struct {
						Text              string `xml:",chardata"` // Tumor Resection
						Cde               string `xml:"cde,attr"`
						XsdVer            string `xml:"xsd_ver,attr"`
						ProcurementStatus string `xml:"procurement_status,attr"`
						Owner             string `xml:"owner,attr"`
					} `xml:"method_of_normal_sample_procurement"`
					DaysToNormalSampleProcurement struct {
						Text              string `xml:",chardata"` // 0
						Precision         string `xml:"precision,attr"`
						Cde               string `xml:"cde,attr"`
						XsdVer            string `xml:"xsd_ver,attr"`
						ProcurementStatus string `xml:"procurement_status,attr"`
						Owner             string `xml:"owner,attr"`
					} `xml:"days_to_normal_sample_procurement"`
				} `xml:"normal_control"`
			} `xml:"normal_controls"`
			DaysToPathologyReview struct {
				Text              string `xml:",chardata"` // 6
				Precision         string `xml:"precision,attr"`
				Cde               string `xml:"cde,attr"`
				XsdVer            string `xml:"xsd_ver,attr"`
				ProcurementStatus string `xml:"procurement_status,attr"`
				Owner             string `xml:"owner,attr"`
			} `xml:"days_to_pathology_review"`
			PathConfirmTumorNucleiMetrics struct {
				Text              string `xml:",chardata"` // YES
				Cde               string `xml:"cde,attr"`
				XsdVer            string `xml:"xsd_ver,attr"`
				ProcurementStatus string `xml:"procurement_status,attr"`
				Owner             string `xml:"owner,attr"`
			} `xml:"path_confirm_tumor_nuclei_metrics"`
			PathConfirmTumorNecrosisMetrics struct {
				Text              string `xml:",chardata"` // YES
				Cde               string `xml:"cde,attr"`
				XsdVer            string `xml:"xsd_ver,attr"`
				ProcurementStatus string `xml:"procurement_status,attr"`
				Owner             string `xml:"owner,attr"`
			} `xml:"path_confirm_tumor_necrosis_metrics"`
			PathConfirmReportAttached struct {
				Text              string `xml:",chardata"` // YES
				Cde               string `xml:"cde,attr"`
				XsdVer            string `xml:"xsd_ver,attr"`
				ProcurementStatus string `xml:"procurement_status,attr"`
				Owner             string `xml:"owner,attr"`
			} `xml:"path_confirm_report_attached"`
			PathConfirmDiagnosisMatching struct {
				Text              string `xml:",chardata"` // YES
				Cde               string `xml:"cde,attr"`
				XsdVer            string `xml:"xsd_ver,attr"`
				ProcurementStatus string `xml:"procurement_status,attr"`
				Owner             string `xml:"owner,attr"`
			} `xml:"path_confirm_diagnosis_matching"`
			ReasonPathConfirmDiagnosisNotMatching struct {
				Text              string `xml:",chardata"`
				Cde               string `xml:"cde,attr"`
				XsdVer            string `xml:"xsd_ver,attr"`
				ProcurementStatus string `xml:"procurement_status,attr"`
				Owner             string `xml:"owner,attr"`
			} `xml:"reason_path_confirm_diagnosis_not_matching"`
			CytogeneticReportSubmitted struct {
				Text              string `xml:",chardata"`
				Cde               string `xml:"cde,attr"`
				XsdVer            string `xml:"xsd_ver,attr"`
				ProcurementStatus string `xml:"procurement_status,attr"`
				Owner             string `xml:"owner,attr"`
			} `xml:"cytogenetic_report_submitted"`
			DifferentialReportSubmitted struct {
				Text              string `xml:",chardata"`
				Cde               string `xml:"cde,attr"`
				XsdVer            string `xml:"xsd_ver,attr"`
				ProcurementStatus string `xml:"procurement_status,attr"`
				Owner             string `xml:"owner,attr"`
			} `xml:"differential_report_submitted"`
			HivPositiveStatus struct {
				Text              string `xml:",chardata"`
				Cde               string `xml:"cde,attr"`
				XsdVer            string `xml:"xsd_ver,attr"`
				ProcurementStatus string `xml:"procurement_status,attr"`
				Owner             string `xml:"owner,attr"`
			} `xml:"hiv_positive_status"`
			TotalCellsSubmitted struct {
				Text              string `xml:",chardata"`
				Cde               string `xml:"cde,attr"`
				XsdVer            string `xml:"xsd_ver,attr"`
				ProcurementStatus string `xml:"procurement_status,attr"`
				Owner             string `xml:"owner,attr"`
			} `xml:"total_cells_submitted"`
			PercentMyeloblastsForSubmittedSpecimen struct {
				Text              string `xml:",chardata"`
				Cde               string `xml:"cde,attr"`
				XsdVer            string `xml:"xsd_ver,attr"`
				ProcurementStatus string `xml:"procurement_status,attr"`
				Owner             string `xml:"owner,attr"`
			} `xml:"percent_myeloblasts_for_submitted_specimen"`
			Race struct {
				Text              string `xml:",chardata"`
				Cde               string `xml:"cde,attr"`
				XsdVer            string `xml:"xsd_ver,attr"`
				ProcurementStatus string `xml:"procurement_status,attr"`
				Owner             string `xml:"owner,attr"`
			} `xml:"race"`
			Ethnicity struct {
				Text              string `xml:",chardata"`
				Cde               string `xml:"cde,attr"`
				XsdVer            string `xml:"xsd_ver,attr"`
				ProcurementStatus string `xml:"procurement_status,attr"`
				Owner             string `xml:"owner,attr"`
			} `xml:"ethnicity"`
			HistologicalType struct {
				Text              string `xml:",chardata"` // Lung Adenocarcinoma- Not ...
				Cde               string `xml:"cde,attr"`
				XsdVer            string `xml:"xsd_ver,attr"`
				ProcurementStatus string `xml:"procurement_status,attr"`
				Owner             string `xml:"owner,attr"`
			} `xml:"histological_type"`
			HistologicalTypeOther struct {
				Text              string `xml:",chardata"`
				Cde               string `xml:"cde,attr"`
				XsdVer            string `xml:"xsd_ver,attr"`
				ProcurementStatus string `xml:"procurement_status,attr"`
				Owner             string `xml:"owner,attr"`
			} `xml:"histological_type_other"`
			HistoryOfNeoadjuvantTreatment struct {
				Text              string `xml:",chardata"` // No
				Cde               string `xml:"cde,attr"`
				XsdVer            string `xml:"xsd_ver,attr"`
				ProcurementStatus string `xml:"procurement_status,attr"`
				Owner             string `xml:"owner,attr"`
			} `xml:"history_of_neoadjuvant_treatment"`
			ConsentOrDeathStatus struct {
				Text              string `xml:",chardata"` // Consented
				Cde               string `xml:"cde,attr"`
				XsdVer            string `xml:"xsd_ver,attr"`
				ProcurementStatus string `xml:"procurement_status,attr"`
				Owner             string `xml:"owner,attr"`
			} `xml:"consent_or_death_status"`
		} `xml:"biospecimen_cqcf"`
	} `xml:"patient"`
}
