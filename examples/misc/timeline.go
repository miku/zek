package misc

import "encoding/xml"

// Timeline was generated 2017-12-05 23:44:27 by tir on apollo.
type Timeline struct {
	XMLName xml.Name `xml:"timeline"`
	Text    string   `xml:",chardata"`
	Decade  []struct {
		Text  string `xml:",chardata"`
		DID   string `xml:"d_id,attr"`
		DYear []struct {
			Text     string `xml:",chardata"`
			YID      string `xml:"y_id,attr"`
			Gorenews struct {
				Text      string `xml:",chardata"`
				GBodytext struct {
					Text string `xml:",chardata"` // <p>January 1: Wilbert L. (&quot;Bill&quot;) and Genevieve (&quot;Vieve&quot;) Gore launch W. L. Gore...
				} `xml:"g_bodytext"`
				GImages struct {
					Text   string `xml:",chardata"`
					GImage struct {
						Text string `xml:",chardata"` // <p>Co-founders of the Enterprise, Bill and Vieve Gore</p>, <p>Early brochure for TETRA-ETCH® etchan...
						GSrc string `xml:"g_src,attr"`
					} `xml:"g_image"`
				} `xml:"g_images"`
			} `xml:"gorenews"`
			Worldnews struct {
				Text      string `xml:",chardata"`
				WBodytext struct {
					Text string `xml:",chardata"` // <p>British Airways launches the first transatlantic passenger-jet service.<br><br>U.S. Congress pass...
				} `xml:"w_bodytext"`
				WImages struct {
					Text   string `xml:",chardata"`
					WImage struct {
						Text string `xml:",chardata"` // Pelé competes for Brazil in the 1958 World Cup., Nixon and Khruschev at the kitchen display of the ...
						WSrc string `xml:"w_src,attr"`
					} `xml:"w_image"`
				} `xml:"w_images"`
			} `xml:"worldnews"`
		} `xml:"d_year"`
	} `xml:"decade"`
}
