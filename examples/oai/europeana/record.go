package europeana

import "encoding/xml"

type Record struct {
	XMLName xml.Name `xml:"Record"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text       string `xml:",chardata"`
		Status     string `xml:"status,attr"`
		Identifier struct {
			Text string `xml:",chardata"` // urn:nbn:de:0220-gd-804646...
		} `xml:"identifier"`
		Datestamp struct {
			Text string `xml:",chardata"` // 2014-12-08T14:01:21Z, 201...
		} `xml:"datestamp"`
		SetSpec struct {
			Text string `xml:",chardata"` // kaiserreichgeschichtsschu...
		} `xml:"setSpec"`
	} `xml:"header"`
	Metadata struct {
		Text   string `xml:",chardata"`
		Record struct {
			Text       string `xml:",chardata"`
			Identifier struct {
				Text string `xml:",chardata"` // http://gei-digital.gei.de...
			} `xml:"identifier"`
			Title struct {
				Text string `xml:",chardata"` // Lehrbuch der Geschichte f...
			} `xml:"title"`
			Date struct {
				Text string `xml:",chardata"` // 1895, 1900, 1915, 1902, 1...
			} `xml:"date"`
			Creator []struct {
				Text string `xml:",chardata"` // Wessel, Paul, Stein, Rich...
			} `xml:"creator"`
			Publisher struct {
				Text string `xml:",chardata"` // Perthes, Schöningh, Buch...
			} `xml:"publisher"`
			Type struct {
				Text string `xml:",chardata"` // Monograph, TEXT, Volume, ...
			} `xml:"type"`
			Format []struct {
				Text string `xml:",chardata"` // image/jpeg, application/p...
			} `xml:"format"`
			Source struct {
				Text string `xml:",chardata"` // Wessel, Paul: Lehrbuch de...
			} `xml:"source"`
			Provider struct {
				Text string `xml:",chardata"` // Georg-Eckert-Institut - L...
			} `xml:"provider"`
			Rights struct {
				Text string `xml:",chardata"` // http://creativecommons.or...
			} `xml:"rights"`
			DataProvider struct {
				Text string `xml:",chardata"` // Georg-Eckert-Institut - L...
			} `xml:"dataProvider"`
			IsShownAt struct {
				Text string `xml:",chardata"` // http://gei-digital.gei.de...
			} `xml:"isShownAt"`
			Language struct {
				Text string `xml:",chardata"` // ger, ger, ger, ger, ger, ...
			} `xml:"language"`
		} `xml:"record"`
	} `xml:"metadata"`
	About struct {
		Text string `xml:",chardata"`
	} `xml:"about"`
}
