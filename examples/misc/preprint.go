package misc

import "encoding/xml"

// PreprintList was generated 2017-12-06 00:00:28 by tir on apollo.
type PreprintList struct {
	XMLName  xml.Name `xml:"PreprintList"`
	Text     string   `xml:",chardata"`
	Year     string   `xml:"Year,attr"`
	Preprint []struct {
		Text  string `xml:",chardata"`
		Num   string `xml:"num,attr"`
		Title struct {
			Text string `xml:",chardata"` // G2-structures and quantization of non-geometric M-theory backgrounds, Supersymmetric Yang-Mills Theo...
		} `xml:"title"`
		Author []struct {
			Text string `xml:",chardata"` // Vladislav Kupriyanov, Richard Szabo, Christian Sämann, Martin Wolf, Yasuhito Kaminaga, Gary T. Horo...
		} `xml:"author"`
		Eprint struct {
			Text string `xml:",chardata"` // 1701.02574, 1702.04160, 1703.06718, 1704.04071, 1705.09639, 1706.07383, 1709.03860, 1708.09660, 1711...
		} `xml:"eprint"`
	} `xml:"preprint"`
}
