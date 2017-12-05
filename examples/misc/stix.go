package misc

import "encoding/xml"

// STIXPackage was generated 2017-12-05 23:52:32 by tir on apollo.
type STIXPackage struct {
	XMLName        xml.Name `xml:"STIX_Package"`
	Text           string   `xml:",chardata"`
	CyboxCommon    string   `xml:"cyboxCommon,attr"`
	Cybox          string   `xml:"cybox,attr"`
	CyboxVocabs    string   `xml:"cyboxVocabs,attr"`
	AddressObj     string   `xml:"AddressObj,attr"`
	FileObj        string   `xml:"FileObj,attr"`
	Marking        string   `xml:"marking,attr"`
	TlpMarking     string   `xml:"tlpMarking,attr"`
	TOUMarking     string   `xml:"TOUMarking,attr"`
	Indicator      string   `xml:"indicator,attr"`
	StixCommon     string   `xml:"stixCommon,attr"`
	StixVocabs     string   `xml:"stixVocabs,attr"`
	Stix           string   `xml:"stix,attr"`
	NCCIC          string   `xml:"NCCIC,attr"`
	CISCP          string   `xml:"CISCP,attr"`
	Xsi            string   `xml:"xsi,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	ID             string   `xml:"id,attr"`
	Version        string   `xml:"version,attr"`
	Timestamp      string   `xml:"timestamp,attr"`
	STIXHeader     struct {
		Text  string `xml:",chardata"`
		Title struct {
			Text string `xml:",chardata"` // HIDDEN COBRA TROJAN: Volgmer
		} `xml:"Title"`
		PackageIntent struct {
			Text string `xml:",chardata"` // Indicators - Watchlist
			Type string `xml:"type,attr"`
		} `xml:"Package_Intent"`
		Description struct {
			Text string `xml:",chardata"` // According to DHS & FBI analysis, these IOCs are associated with the Volgmer Trojan and were used by ...
		} `xml:"Description"`
		Handling struct {
			Text    string `xml:",chardata"`
			Marking struct {
				Text                string `xml:",chardata"`
				ControlledStructure struct {
					Text string `xml:",chardata"` // //node() | //@*
				} `xml:"Controlled_Structure"`
				MarkingStructure []struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr"`
					Color      string `xml:"color,attr"`
					TermsOfUse struct {
						Text string `xml:",chardata"` // DISCLAIMER: This report is provided "as is" for informational purposes only. The Department of Homel...
					} `xml:"Terms_Of_Use"`
				} `xml:"Marking_Structure"`
			} `xml:"Marking"`
		} `xml:"Handling"`
		InformationSource struct {
			Text string `xml:",chardata"`
			Time struct {
				Text         string `xml:",chardata"`
				ProducedTime struct {
					Text string `xml:",chardata"` // 2017-11-06T13:59:34
				} `xml:"Produced_Time"`
			} `xml:"Time"`
		} `xml:"Information_Source"`
	} `xml:"STIX_Header"`
	Indicators struct {
		Text      string `xml:",chardata"`
		Indicator []struct {
			Text      string `xml:",chardata"`
			ID        string `xml:"id,attr"`
			Timestamp string `xml:"timestamp,attr"`
			AttrType  string `xml:"type,attr"`
			Title     struct {
				Text string `xml:",chardata"` // Malicious IPv4 Indicator, Malicious IPv4 Indicator, Malicious IPv4 Indicator, Malicious IPv4 Indicat...
			} `xml:"Title"`
			Type struct {
				Text string `xml:",chardata"` // IP Watchlist, Compromised, Compromised, Compromised, Compromised, Compromised, Compromised, Compromi...
				Type string `xml:"type,attr"`
			} `xml:"Type"`
			Description struct {
				Text string `xml:",chardata"` // According to a trusted third party, between November 24, and November 30, 2016, Volgmer malware was ...
			} `xml:"Description"`
			Observable struct {
				Text   string `xml:",chardata"`
				ID     string `xml:"id,attr"`
				Object struct {
					Text       string `xml:",chardata"`
					ID         string `xml:"id,attr"`
					Properties struct {
						Text         string `xml:",chardata"`
						Type         string `xml:"type,attr"`
						Category     string `xml:"category,attr"`
						IsSpoofed    string `xml:"is_spoofed,attr"`
						AddressValue struct {
							Text      string `xml:",chardata"` // 199.68.196.125, 103.16.223.35, 113.28.244.194, 116.48.145.179, 186.116.9.20, 186.149.198.172, 195.28...
							Condition string `xml:"condition,attr"`
						} `xml:"Address_Value"`
						FileName struct {
							Text      string `xml:",chardata"` // S.exe, fnetsecenum.dll
							Condition string `xml:"condition,attr"`
						} `xml:"File_Name"`
						SizeInBytes struct {
							Text      string `xml:",chardata"` // 159748, 122880, 144020, 144000, 148000, 120000, 104500, 128000, 120000, 108000
							Condition string `xml:"condition,attr"`
						} `xml:"Size_In_Bytes"`
						Hashes struct {
							Text string `xml:",chardata"`
							Hash []struct {
								Text string `xml:",chardata"`
								Type struct {
									Text      string `xml:",chardata"` // MD5, SHA1, MD5, SHA1, SHA256, SSDEEP, MD5, SHA1, SHA256, SSDEEP
									Condition string `xml:"condition,attr"`
									Type      string `xml:"type,attr"`
								} `xml:"Type"`
								SimpleHashValue struct {
									Text      string `xml:",chardata"` // 2D2B88AE9F7E5B49B728AD7A1D220E84, 000270fd7f5d5a020ac05c87afe138f80acb120a, 9A5FA5C5F3915B2297A1C379...
									Condition string `xml:"condition,attr"`
								} `xml:"Simple_Hash_Value"`
								FuzzyHashValue struct {
									Text      string `xml:",chardata"` // 1536:oFzeDm/AzT/xmIaTyDfOyzCz1ETbzZunTdWuco6+TxRR+dh6kHJBtBf:oUYWTzaTTz14ZunJiNInwh6WJBtBf, 3072:Z1K...
									Condition string `xml:"condition,attr"`
								} `xml:"Fuzzy_Hash_Value"`
							} `xml:"Hash"`
						} `xml:"Hashes"`
					} `xml:"Properties"`
				} `xml:"Object"`
			} `xml:"Observable"`
			KillChainPhases struct {
				Text           string `xml:",chardata"`
				KillChainPhase struct {
					Text        string `xml:",chardata"`
					Ordinality  string `xml:"ordinality,attr"`
					Name        string `xml:"name,attr"`
					PhaseID     string `xml:"phase_id,attr"`
					KillChainID string `xml:"kill_chain_id,attr"`
				} `xml:"Kill_Chain_Phase"`
			} `xml:"Kill_Chain_Phases"`
			Handling struct {
				Text    string `xml:",chardata"`
				Marking struct {
					Text                string `xml:",chardata"`
					ControlledStructure struct {
						Text string `xml:",chardata"` // ../../../descendant-or-self::node() | ../../../descendant-or-self::node()/@*, ../../../descendant-or...
					} `xml:"Controlled_Structure"`
					MarkingStructure []struct {
						Text       string `xml:",chardata"`
						Type       string `xml:"type,attr"`
						Color      string `xml:"color,attr"`
						TermsOfUse struct {
							Text string `xml:",chardata"` // DISCLAIMER: This report is provided "as is" for informational purposes only. The Department of Homel...
						} `xml:"Terms_Of_Use"`
					} `xml:"Marking_Structure"`
				} `xml:"Marking"`
			} `xml:"Handling"`
		} `xml:"Indicator"`
	} `xml:"Indicators"`
	TTPs struct {
		Text       string `xml:",chardata"`
		KillChains struct {
			Text      string `xml:",chardata"`
			KillChain struct {
				Text           string `xml:",chardata"`
				Reference      string `xml:"reference,attr"`
				NumberOfPhases string `xml:"number_of_phases,attr"`
				ID             string `xml:"id,attr"`
				Definer        string `xml:"definer,attr"`
				Name           string `xml:"name,attr"`
				KillChainPhase []struct {
					Text       string `xml:",chardata"`
					Ordinality string `xml:"ordinality,attr"`
					Name       string `xml:"name,attr"`
					PhaseID    string `xml:"phase_id,attr"`
				} `xml:"Kill_Chain_Phase"`
			} `xml:"Kill_Chain"`
		} `xml:"Kill_Chains"`
	} `xml:"TTPs"`
}
