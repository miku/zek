package misc

import "encoding/xml"

// RFC was generated 2017-12-05 23:37:07 by tir on apollo.
// curl -s https://tools.ietf.org/id/draft-templin-6man-ndrioiana-00.xml | zek -e > rfc.go
type RFC struct {
	XMLName  xml.Name `xml:"rfc"`
	Text     string   `xml:",chardata"`
	Category string   `xml:"category,attr"`
	DocName  string   `xml:"docName,attr"`
	Ipr      string   `xml:"ipr,attr"`
	Updates  string   `xml:"updates,attr"`
	Front    struct {
		Text  string `xml:",chardata"`
		Title struct {
			Text   string `xml:",chardata"` // IPv6 ND RIO Flags IANA   ...
			Abbrev string `xml:"abbrev,attr"`
		} `xml:"title"`
		Author struct {
			Text         string `xml:",chardata"`
			Fullname     string `xml:"fullname,attr"`
			Initials     string `xml:"initials,attr"`
			Role         string `xml:"role,attr"`
			Surname      string `xml:"surname,attr"`
			Organization struct {
				Text string `xml:",chardata"` // Boeing Research & Technol...
			} `xml:"organization"`
			Address struct {
				Text   string `xml:",chardata"`
				Postal struct {
					Text   string `xml:",chardata"`
					Street struct {
						Text string `xml:",chardata"` // P.O. Box 3707
					} `xml:"street"`
					City struct {
						Text string `xml:",chardata"` // Seattle
					} `xml:"city"`
					Region struct {
						Text string `xml:",chardata"` // WA
					} `xml:"region"`
					Code struct {
						Text string `xml:",chardata"` // 98124
					} `xml:"code"`
					Country struct {
						Text string `xml:",chardata"` // USA
					} `xml:"country"`
				} `xml:"postal"`
				Email struct {
					Text string `xml:",chardata"` // fltemplin@acm.org
				} `xml:"email"`
			} `xml:"address"`
		} `xml:"author"`
		Date struct {
			Text  string `xml:",chardata"`
			Day   string `xml:"day,attr"`
			Month string `xml:"month,attr"`
			Year  string `xml:"year,attr"`
		} `xml:"date"`
		Keyword []struct {
			Text string `xml:",chardata"` // I-D, Internet-Draft
		} `xml:"keyword"`
		Abstract struct {
			Text string `xml:",chardata"`
			T    struct {
				Text string `xml:",chardata"` // The IPv6 Neighbor Discove...
			} `xml:"t"`
		} `xml:"abstract"`
	} `xml:"front"`
	Middle struct {
		Text    string `xml:",chardata"`
		Section []struct {
			Text   string `xml:",chardata"`
			Anchor string `xml:"anchor,attr"`
			Title  string `xml:"title,attr"`
			T      []struct {
				Text string `xml:",chardata"` // The Route Information Opt...
				Xref []struct {
					Text   string `xml:",chardata"`
					Target string `xml:"target,attr"`
				} `xml:"xref"`
				Figure struct {
					Text    string `xml:",chardata"`
					Artwork struct {
						Text string `xml:",chardata"` // 0 1 2 3 4 5 6 7       +-+...
					} `xml:"artwork"`
				} `xml:"figure"`
			} `xml:"t"`
		} `xml:"section"`
	} `xml:"middle"`
	Back struct {
		Text       string `xml:",chardata"`
		References []struct {
			Text  string `xml:",chardata"`
			Title string `xml:"title,attr"`
		} `xml:"references"`
	} `xml:"back"`
}
