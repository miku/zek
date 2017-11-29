package oailom

import "encoding/xml"

type Record struct {
	XMLName xml.Name `xml:"Record"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text       string `xml:",chardata"`
		Status     string `xml:"status,attr"`
		Identifier struct {
			Text string `xml:",chardata"` // oai:oka-pu.repo.nii.ac.jp...
		} `xml:"identifier"`
		Datestamp struct {
			Text string `xml:",chardata"` // 2015-09-28T05:17:04Z, 201...
		} `xml:"datestamp"`
		SetSpec struct {
			Text string `xml:",chardata"` // 00012:00013, 00069:00072,...
		} `xml:"setSpec"`
	} `xml:"header"`
	Metadata struct {
		Text string `xml:",chardata"`
		Lom  struct {
			Text           string `xml:",chardata"`
			SchemaLocation string `xml:"schemaLocation,attr"`
			Xmlns          string `xml:"xmlns,attr"`
			General        struct {
				Text       string `xml:",chardata"`
				Identifier []struct {
					Text    string `xml:",chardata"`
					Catalog struct {
						Text string `xml:",chardata"` // URI, identifier, identifi...
					} `xml:"catalog"`
					Entry struct {
						Text string `xml:",chardata"` // https://oka-pu.repo.nii.a...
					} `xml:"entry"`
				} `xml:"identifier"`
				Title struct {
					Text   string `xml:",chardata"`
					String struct {
						Text     string `xml:",chardata"` // åæåããã³éåæ...
						Language string `xml:"language,attr"`
					} `xml:"string"`
				} `xml:"title"`
				Language struct {
					Text string `xml:",chardata"` // ja, ja, ja, ja, ja, ja, e...
				} `xml:"language"`
				Description []struct {
					Text   string `xml:",chardata"`
					String struct {
						Text     string `xml:",chardata"` // 2013, å²¡å±±çç«å¤§å­¦å...
						Language string `xml:"language,attr"`
					} `xml:"string"`
				} `xml:"description"`
				Keyword []struct {
					Text   string `xml:",chardata"`
					String struct {
						Text string `xml:",chardata"` // èå¾³å¤ªå­, æè²å­¦, ...
					} `xml:"string"`
				} `xml:"keyword"`
			} `xml:"general"`
			LifeCycle struct {
				Text       string `xml:",chardata"`
				Contribute []struct {
					Text string `xml:",chardata"`
					Role struct {
						Text   string `xml:",chardata"`
						Source struct {
							Text string `xml:",chardata"` // LOMv1.0, LOMv1.0, LOMv1.0...
						} `xml:"source"`
						Value struct {
							Text string `xml:",chardata"` // author, author, publisher...
						} `xml:"value"`
					} `xml:"role"`
					Entity []struct {
						Text string `xml:",chardata"` // ç«é è²´ä¹, å®å± è...
					} `xml:"entity"`
				} `xml:"contribute"`
			} `xml:"lifeCycle"`
			Relation []struct {
				Text     string `xml:",chardata"`
				Resource struct {
					Text       string `xml:",chardata"`
					Identifier []struct {
						Text    string `xml:",chardata"`
						Catalog struct {
							Text string `xml:",chardata"` // identifier, identifier, i...
						} `xml:"catalog"`
						Entry struct {
							Text string `xml:",chardata"` // info:doi/10.4327/jsnfs.55...
						} `xml:"entry"`
					} `xml:"identifier"`
				} `xml:"resource"`
				Kind struct {
					Text   string `xml:",chardata"`
					Source struct {
						Text string `xml:",chardata"` // LOMv1.0, LOMv1.0, LOMv1.0...
					} `xml:"source"`
					Value struct {
						Text string `xml:",chardata"` // isbasedon, isbasedon, isb...
					} `xml:"value"`
				} `xml:"kind"`
			} `xml:"relation"`
			Rights struct {
				Text        string `xml:",chardata"`
				Description struct {
					Text   string `xml:",chardata"`
					String struct {
						Text string `xml:",chardata"` // Copyright Â© 2015 IEICE, ...
					} `xml:"string"`
				} `xml:"description"`
			} `xml:"rights"`
		} `xml:"lom"`
	} `xml:"metadata"`
	About struct {
		Text string `xml:",chardata"`
	} `xml:"about"`
}
