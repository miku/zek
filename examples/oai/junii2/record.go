package junii2

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
		Text   string `xml:",chardata"`
		Junii2 struct {
			Text           string `xml:",chardata"`
			SchemaLocation string `xml:"schemaLocation,attr"`
			Title          struct {
				Text string `xml:",chardata"` // åæåããã³éåæ...
				Lang string `xml:"lang,attr"`
			} `xml:"title"`
			Language struct {
				Text string `xml:",chardata"` // jpn, jpn, jpn, jpn, jpn, ...
			} `xml:"language"`
			URI struct {
				Text string `xml:",chardata"` // http://id.nii.ac.jp/1212/...
			} `xml:"URI"`
			SelfDOI struct {
				Text string `xml:",chardata"` // info:doi/10.15009/0000069...
				Ra   string `xml:"ra,attr"`
			} `xml:"selfDOI"`
			NIItype struct {
				Text string `xml:",chardata"` // Thesis or Dissertation, D...
			} `xml:"NIItype"`
			Alternative []struct {
				Text string `xml:",chardata"` // ãã¦ã­ã¬ã¿ ãªã¨ã...
				Lang string `xml:"lang,attr"`
			} `xml:"alternative"`
			Creator []struct {
				Text string `xml:",chardata"` // ç«é, è²´ä¹, å®å±, è...
				Lang string `xml:"lang,attr"`
				ID   string `xml:"id,attr"`
			} `xml:"creator"`
			Description []struct {
				Text string `xml:",chardata"` // 2013, 2014, 2014, 2014, 2...
			} `xml:"description"`
			Grantid struct {
				Text string `xml:",chardata"` // 25301A92, 25301A93, 25301...
			} `xml:"grantid"`
			Dateofgranted struct {
				Text string `xml:",chardata"` // 2013-09-23, 2014-03-24, 2...
			} `xml:"dateofgranted"`
			Degreename struct {
				Text string `xml:",chardata"` // åå£«ï¼å·¥å­¦ï¼, åå£...
			} `xml:"degreename"`
			Grantor struct {
				Text string `xml:",chardata"` // å²¡å±±çç«å¤§å­¦å¤§å­¦é...
			} `xml:"grantor"`
			Textversion struct {
				Text string `xml:",chardata"` // ETD, publisher, publisher...
			} `xml:"textversion"`
			FullTextURL []struct {
				Text string `xml:",chardata"` // https://oka-pu.repo.nii.a...
			} `xml:"fullTextURL"`
			Date struct {
				Text string `xml:",chardata"` // 2014-08-13, 2015-01-20, 2...
			} `xml:"date"`
			Subject []struct {
				Text string `xml:",chardata"` // èå¾³å¤ªå­, æè²å­¦, ...
			} `xml:"subject"`
			Type struct {
				Text string `xml:",chardata"` // P(è«æ), P(è«æ), P(è...
			} `xml:"type"`
			Contributor []struct {
				Text string `xml:",chardata"` // å²¡å±±çæ é¤ç­æå¤§å...
				Lang string `xml:"lang,attr"`
			} `xml:"contributor"`
			Publisher []struct {
				Text string `xml:",chardata"` // å²¡å±±çæ é¤ç­æå¤§å...
			} `xml:"publisher"`
			Jtitle struct {
				Text string `xml:",chardata"` // å²¡å±±çæ é¤ç­æå¤§å...
			} `xml:"jtitle"`
			Issue struct {
				Text string `xml:",chardata"` // 1, 1, 1, 1, 1, 1, 2, 2, 2...
			} `xml:"issue"`
			Spage struct {
				Text string `xml:",chardata"` // ä¸, åã, äºã, 1, 4...
			} `xml:"spage"`
			Epage struct {
				Text string `xml:",chardata"` // ä¸ä¹, åä¹, äºå­, ä...
			} `xml:"epage"`
			Dateofissued struct {
				Text string `xml:",chardata"` // 1957, 1957, 1957, 1957, 1...
			} `xml:"dateofissued"`
			ISSN struct {
				Text string `xml:",chardata"` // 02870983, 02870983, 02870...
			} `xml:"issn"`
			NCID struct {
				Text string `xml:",chardata"` // AN00032660, AN00032660, A...
			} `xml:"NCID"`
			Identifier struct {
				Text string `xml:",chardata"` // 40000320142, 40000320140,...
			} `xml:"identifier"`
			Rights []struct {
				Text string `xml:",chardata"` // CC BY-NC-ND, CC BY-NC-ND,...
			} `xml:"rights"`
			Volume struct {
				Text string `xml:",chardata"` // 37, 37, 32, 32, 32, 32, 3...
			} `xml:"volume"`
			Doi struct {
				Text string `xml:",chardata"` // info:doi/10.4327/jsnfs.55...
			} `xml:"doi"`
			Source []struct {
				Text string `xml:",chardata"` // https://www.jstage.jst.go...
			} `xml:"source"`
			NAID struct {
				Text string `xml:",chardata"` // http://ci.nii.ac.jp/naid/...
			} `xml:"NAID"`
		} `xml:"junii2"`
	} `xml:"metadata"`
	About struct {
		Text string `xml:",chardata"`
	} `xml:"about"`
}
