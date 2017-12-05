package misc

import "encoding/xml"

// IatiActivities was generated 2017-12-05 23:48:45 by tir on apollo.
type IatiActivities struct {
	XMLName           xml.Name `xml:"iati-activities"`
	Text              string   `xml:",chardata"`
	GeneratedDatetime string   `xml:"generated-datetime,attr"`
	Version           string   `xml:"version,attr"`
	Usg               string   `xml:"usg,attr"`
	IatiActivity      []struct {
		Text                string `xml:",chardata"`
		DefaultCurrency     string `xml:"default-currency,attr"`
		Hierarchy           string `xml:"hierarchy,attr"`
		LastUpdatedDatetime string `xml:"last-updated-datetime,attr"`
		Lang                string `xml:"lang,attr"`
		IatiIdentifier      struct {
			Text string `xml:",chardata"` // US-GOV-1-NI-0004A3515A, US-GOV-1-NI-006856901623, ...
		} `xml:"iati-identifier"`
		ReportingOrg struct {
			Text      string `xml:",chardata"`
			Ref       string `xml:"ref,attr"`
			Type      string `xml:"type,attr"`
			Narrative struct {
				Text string `xml:",chardata"` // U.S. Agency for International Development, U.S. Ag...
			} `xml:"narrative"`
		} `xml:"reporting-org"`
		Title struct {
			Text      string `xml:",chardata"`
			Narrative struct {
				Text string `xml:",chardata"` // Capacity Building, Preparedness, and Planning, Adm...
			} `xml:"narrative"`
		} `xml:"title"`
		Description struct {
			Text      string `xml:",chardata"`
			Narrative struct {
				Text string `xml:",chardata"` // Improve the ability of the USG, host countries and...
			} `xml:"narrative"`
		} `xml:"description"`
		ParticipatingOrg []struct {
			Text      string `xml:",chardata"`
			Ref       string `xml:"ref,attr"`
			Role      string `xml:"role,attr"`
			Type      string `xml:"type,attr"`
			Narrative struct {
				Text string `xml:",chardata"` // U.S. Agency for International Development, U.S. Ag...
			} `xml:"narrative"`
		} `xml:"participating-org"`
		ActivityStatus struct {
			Text string `xml:",chardata"`
			Code string `xml:"code,attr"`
		} `xml:"activity-status"`
		ActivityDate []struct {
			Text    string `xml:",chardata"`
			IsoDate string `xml:"iso-date,attr"`
			Type    string `xml:"type,attr"`
		} `xml:"activity-date"`
		ContactInfo struct {
			Text         string `xml:",chardata"`
			Type         string `xml:"type,attr"`
			Organisation struct {
				Text      string `xml:",chardata"`
				Narrative struct {
					Text string `xml:",chardata"` // U.S. Agency for International Development, U.S. Ag...
				} `xml:"narrative"`
			} `xml:"organisation"`
			PersonName struct {
				Text      string `xml:",chardata"`
				Narrative struct {
					Text string `xml:",chardata"` // Deidra Winston, Deidra Winston, Deidra Winston, De...
				} `xml:"narrative"`
			} `xml:"person-name"`
			Telephone struct {
				Text string `xml:",chardata"` // +1 (202) 712 4810, +1 (202) 712 4810, +1 (202) 712...
			} `xml:"telephone"`
			Email struct {
				Text string `xml:",chardata"` // USAID_IATI_Data@usaid.gov, dwinston@usaid.gov, dwi...
			} `xml:"email"`
			Website struct {
				Text string `xml:",chardata"` // https://www.usaid.gov/who-we-are/organization/bure...
			} `xml:"website"`
			MailingAddress struct {
				Text      string `xml:",chardata"`
				Narrative struct {
					Text string `xml:",chardata"` // 1300 Pennsylvania Ave NW, Washington DC 20004, 130...
				} `xml:"narrative"`
			} `xml:"mailing-address"`
		} `xml:"contact-info"`
		ActivityScope struct {
			Text string `xml:",chardata"`
			Code string `xml:"code,attr"`
		} `xml:"activity-scope"`
		RecipientCountry struct {
			Text       string `xml:",chardata"`
			Code       string `xml:"code,attr"`
			Percentage string `xml:"percentage,attr"`
		} `xml:"recipient-country"`
		CollaborationType struct {
			Text string `xml:",chardata"`
			Code string `xml:"code,attr"`
		} `xml:"collaboration-type"`
		DefaultFlowType struct {
			Text string `xml:",chardata"`
			Code string `xml:"code,attr"`
		} `xml:"default-flow-type"`
		DefaultFinanceType struct {
			Text string `xml:",chardata"`
			Code string `xml:"code,attr"`
		} `xml:"default-finance-type"`
		DefaultAidType struct {
			Text string `xml:",chardata"`
			Code string `xml:"code,attr"`
		} `xml:"default-aid-type"`
		DefaultTiedStatus struct {
			Text string `xml:",chardata"`
			Code string `xml:"code,attr"`
		} `xml:"default-tied-status"`
		Transaction []struct {
			Text            string `xml:",chardata"`
			Humanitarian    string `xml:"humanitarian,attr"`
			TransactionType struct {
				Text string `xml:",chardata"`
				Code string `xml:"code,attr"`
			} `xml:"transaction-type"`
			TransactionDate struct {
				Text    string `xml:",chardata"`
				IsoDate string `xml:"iso-date,attr"`
			} `xml:"transaction-date"`
			Value struct {
				Text      string `xml:",chardata"` // 2356.00, 1301.00, -1055.00, 206269.20, 206269.20, ...
				ValueDate string `xml:"value-date,attr"`
			} `xml:"value"`
			Description struct {
				Text      string `xml:",chardata"`
				Narrative struct {
					Text string `xml:",chardata"` // Improve the ability of the USG, host countries and...
				} `xml:"narrative"`
			} `xml:"description"`
			DisbursementChannel struct {
				Text string `xml:",chardata"`
				Code string `xml:"code,attr"`
			} `xml:"disbursement-channel"`
			Sector []struct {
				Text       string `xml:",chardata"`
				Code       string `xml:"code,attr"`
				Vocabulary string `xml:"vocabulary,attr"`
				Narrative  struct {
					Text string `xml:",chardata"` // Disaster prevention and preparedness, Disaster Rea...
				} `xml:"narrative"`
			} `xml:"sector"`
			TreasuryAccount struct {
				Text           string `xml:",chardata"`
				RegularAccount struct {
					Text string `xml:",chardata"`
					Code string `xml:"code,attr"`
				} `xml:"regular-account"`
				MainAccount struct {
					Text string `xml:",chardata"` // International Disaster and Famine Assistance, Inte...
					Code string `xml:"code,attr"`
				} `xml:"main-account"`
				FiscalFundingYear struct {
					Text  string `xml:",chardata"`
					Begin string `xml:"begin,attr"`
					End   string `xml:"end,attr"`
				} `xml:"fiscal-funding-year"`
			} `xml:"treasury-account"`
		} `xml:"transaction"`
		RelatedActivity struct {
			Text string `xml:",chardata"`
			Ref  string `xml:"ref,attr"`
			Type string `xml:"type,attr"`
		} `xml:"related-activity"`
		Conditions struct {
			Text     string `xml:",chardata"`
			Attached string `xml:"attached,attr"`
		} `xml:"conditions"`
		MechanismSigningDate struct {
			Text    string `xml:",chardata"`
			IsoDate string `xml:"iso-date,attr"`
		} `xml:"mechanism-signing-date"`
		DunsNumber struct {
			Text      string `xml:",chardata"`
			Narrative struct {
				Text string `xml:",chardata"` // 790238638, 4868105, 123456787, 927923250, 41799197...
			} `xml:"narrative"`
		} `xml:"duns-number"`
		StateLocation struct {
			Text      string `xml:",chardata"`
			Narrative struct {
				Text string `xml:",chardata"` // Nicaragua, Nicaragua, Nicaragua, Nicaragua, Nicara...
			} `xml:"narrative"`
		} `xml:"state-location"`
		DocumentLink []struct {
			Text   string `xml:",chardata"`
			Format string `xml:"format,attr"`
			URL    string `xml:"url,attr"`
			Title  struct {
				Text      string `xml:",chardata"`
				Narrative struct {
					Text string `xml:",chardata"` // ADS Chapter 302 USAID Direct Contracting, ADS Chap...
				} `xml:"narrative"`
			} `xml:"title"`
			Category struct {
				Text string `xml:",chardata"`
				Code string `xml:"code,attr"`
			} `xml:"category"`
			Language struct {
				Text string `xml:",chardata"`
				Code string `xml:"code,attr"`
			} `xml:"language"`
			DocumentDate struct {
				Text    string `xml:",chardata"`
				IsoDate string `xml:"iso-date,attr"`
			} `xml:"document-date"`
		} `xml:"document-link"`
		Location struct {
			Text          string `xml:",chardata"`
			LocationReach struct {
				Text string `xml:",chardata"`
				Code string `xml:"code,attr"`
			} `xml:"location-reach"`
			Name struct {
				Text      string `xml:",chardata"`
				Narrative struct {
					Text string `xml:",chardata"` // Managua, Managua, Managua, Managua, Managua, Manag...
				} `xml:"narrative"`
			} `xml:"name"`
			Point struct {
				Text    string `xml:",chardata"`
				SrsName string `xml:"srsName,attr"`
				Pos     struct {
					Text string `xml:",chardata"` // 12.1333 -86.25, 12.1333 -86.25, 12.1333 -86.25, 12...
				} `xml:"pos"`
			} `xml:"point"`
			Exactness struct {
				Text string `xml:",chardata"`
				Code string `xml:"code,attr"`
			} `xml:"exactness"`
			LocationClass struct {
				Text string `xml:",chardata"`
				Code string `xml:"code,attr"`
			} `xml:"location-class"`
		} `xml:"location"`
		Tec1 struct {
			Text      string `xml:",chardata"`
			Narrative struct {
				Text string `xml:",chardata"` // 133076.86, 2562387.45, 24000000.00, 41685.00, 2599...
			} `xml:"narrative"`
		} `xml:"tec1"`
		Budget struct {
			Text        string `xml:",chardata"`
			Status      string `xml:"status,attr"`
			PeriodStart struct {
				Text    string `xml:",chardata"`
				IsoDate string `xml:"iso-date,attr"`
			} `xml:"period-start"`
			PeriodEnd struct {
				Text    string `xml:",chardata"`
				IsoDate string `xml:"iso-date,attr"`
			} `xml:"period-end"`
			Value struct {
				Text      string `xml:",chardata"` // 700000.00, 920000.00, 5000.00, 150000.00, 30741.00...
				Currency  string `xml:"currency,attr"`
				ValueDate string `xml:"value-date,attr"`
			} `xml:"value"`
		} `xml:"budget"`
	} `xml:"iati-activity"`
}
