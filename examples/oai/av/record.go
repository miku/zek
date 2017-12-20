package av

import "encoding/xml"

// Record was generated 2017-12-20 15:02:45 by tir on hayiti.
type Record struct {
	XMLName xml.Name `xml:"Record"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text       string `xml:",chardata"`
		Status     string `xml:"status,attr"`
		Identifier struct {
			Text string `xml:",chardata"` // oai::a8450::ton|70900174|...
		} `xml:"identifier"`
		Datestamp struct {
			Text string `xml:",chardata"` // 2015-09-24T07:44:53Z, 201...
		} `xml:"datestamp"`
		SetSpec struct {
			Text string `xml:",chardata"` // audio, audio, audio, audi...
		} `xml:"setSpec"`
	} `xml:"header"`
	Metadata struct {
		Text string `xml:",chardata"`
		Dc   struct {
			Text           string `xml:",chardata"`
			Edp            string `xml:"edp,attr"`
			Dc             string `xml:"dc,attr"`
			Xsi            string `xml:"xsi,attr"`
			OaiDc          string `xml:"oai_dc,attr"`
			Europeana      string `xml:"europeana,attr"`
			Edpm           string `xml:"edpm,attr"`
			SchemaLocation string `xml:"schemaLocation,attr"`
			Contributor    []struct {
				Text string `xml:",chardata"` // Hindemith, Paul, Hindemit...
			} `xml:"contributor"`
			Coverage []struct {
				Text string `xml:",chardata"` // Sächsische Landesbibliot...
			} `xml:"coverage"`
			Title []struct {
				Text string `xml:",chardata"` // Trios, Vl Va Vc, Nr. 2 / ...
			} `xml:"title"`
			Subject []struct {
				Text string `xml:",chardata"` // Schellackplatte, Schellac...
			} `xml:"subject"`
			Date []struct {
				Text string `xml:",chardata"` // 1935, 1935, 2007, 2007, 1...
			} `xml:"date"`
			Description struct {
				Text string `xml:",chardata"` // Zweites Streichtrio (1933...
			} `xml:"description"`
			Format struct {
				Text string `xml:",chardata"` // sound/mp3, sound/mp3, sou...
			} `xml:"format"`
			Source struct {
				Text string `xml:",chardata"` // SLUB Dresden, SLUB Dresde...
			} `xml:"source"`
			Type struct {
				Text string `xml:",chardata"` // sound, SOUND, sound, SOUN...
			} `xml:"type"`
			Provider struct {
				Text string `xml:",chardata"` // SLUB Dresden / Mediathek,...
			} `xml:"provider"`
			Identifier struct {
				Text string `xml:",chardata"` // http://mediathek.slub-dre...
			} `xml:"identifier"`
			IsShownAt struct {
				Text string `xml:",chardata"` // http://mediathek.slub-dre...
			} `xml:"isShownAt"`
			IsShownBy struct {
				Text string `xml:",chardata"` // http://media.slub-dresden...
			} `xml:"isShownBy"`
			Object struct {
				Text string `xml:",chardata"` // http://media.slub-dresden...
			} `xml:"object"`
			Language struct {
				Text string `xml:",chardata"` // de-DE, de, de-DE, de, de-...
			} `xml:"language"`
			Rights []struct {
				Text string `xml:",chardata"` // SLUB Dresden / Mediathek,...
			} `xml:"rights"`
			Country struct {
				Text string `xml:",chardata"` // DE, DE, DE, DE, DE, DE, D...
			} `xml:"country"`
			Creator []struct {
				Text string `xml:",chardata"` // Benndorf, Hirsch, Ernst, ...
			} `xml:"creator"`
		} `xml:"dc"`
	} `xml:"metadata"`
	About struct {
		Text string `xml:",chardata"`
	} `xml:"about"`
}
