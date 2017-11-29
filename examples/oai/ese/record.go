package ese

import "encoding/xml"

type Record struct {
	XMLName xml.Name `xml:"Record"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text       string `xml:",chardata"`
		Status     string `xml:"status,attr"`
		Identifier struct {
			Text string `xml:",chardata"` // oai:dsp.tecnalia.com:1155...
		} `xml:"identifier"`
		Datestamp struct {
			Text string `xml:",chardata"` // 2015-05-06T08:25:04Z, 201...
		} `xml:"datestamp"`
		SetSpec []struct {
			Text string `xml:",chardata"` // com_11556_1, col_11556_2,...
		} `xml:"setSpec"`
	} `xml:"header"`
	Metadata struct {
		Text   string `xml:",chardata"`
		Record struct {
			Text           string `xml:",chardata"`
			Europeana      string `xml:"europeana,attr"`
			Confman        string `xml:"confman,attr"`
			Doc            string `xml:"doc,attr"`
			Dcterms        string `xml:"dcterms,attr"`
			Xsi            string `xml:"xsi,attr"`
			Dc             string `xml:"dc,attr"`
			SchemaLocation string `xml:"schemaLocation,attr"`
			Title          []struct {
				Text string `xml:",chardata"` // Investigation of the Micr...
			} `xml:"title"`
			Creator []struct {
				Text string `xml:",chardata"` // Weintraub, Lior, Davidow,...
			} `xml:"creator"`
			Subject []struct {
				Text string `xml:",chardata"` // (PB,SN,GE)TE-BASED ALLOYS...
			} `xml:"subject"`
			Description []struct {
				Text string `xml:",chardata"` // In the frame of the curre...
			} `xml:"description"`
			Date []struct {
				Text string `xml:",chardata"` // 2015-05-05T15:10:18Z, 201...
			} `xml:"date"`
			Type struct {
				Text string `xml:",chardata"` // article, TEXT, conference...
			} `xml:"type"`
			Identifier []struct {
				Text string `xml:",chardata"` // Journal of Nanomaterials,...
			} `xml:"identifier"`
			Language []struct {
				Text string `xml:",chardata"` // eng, eng, spa, eng, eng, ...
			} `xml:"language"`
			Relation []struct {
				Text string `xml:",chardata"` // info:eu-repo/grantAgreeme...
			} `xml:"relation"`
			Rights []struct {
				Text string `xml:",chardata"` // openAccess, openAccess, o...
			} `xml:"rights"`
			Provider struct {
				Text string `xml:",chardata"` // Hispana, Hispana, Hispana...
			} `xml:"provider"`
			DataProvider struct {
				Text string `xml:",chardata"` // Universidad Rey Juan carl...
			} `xml:"dataProvider"`
			IsShownAt struct {
				Text string `xml:",chardata"` // http://hdl.handle.net/115...
			} `xml:"isShownAt"`
			Contributor []struct {
				Text string `xml:",chardata"` // Agote, Iñígo, San Juan,...
			} `xml:"contributor"`
			Publisher struct {
				Text string `xml:",chardata"` // UPV-EHU, ELSEVIER SCIENCE...
			} `xml:"publisher"`
			Object struct {
				Text string `xml:",chardata"` // http://dsp.tecnalia.com/b...
			} `xml:"object"`
		} `xml:"record"`
	} `xml:"metadata"`
	About struct {
		Text string `xml:",chardata"`
	} `xml:"about"`
}
