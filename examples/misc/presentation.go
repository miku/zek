package misc

import "encoding/xml"

// Presentation was generated 2017-12-05 23:54:02 by tir on apollo.
type Presentation struct {
	XMLName           xml.Name `xml:"presentation"`
	Text              string   `xml:",chardata"`
	Width             string   `xml:"width,attr"`
	Height            string   `xml:"height,attr"`
	Start             string   `xml:"start,attr"`
	SlideTitleFormat  string   `xml:"slideTitleFormat,attr"`
	BuildNum          string   `xml:"buildNum,attr"`
	PresentationTitle struct {
		Text string `xml:",chardata"` // GIPA Trademarks Spanish
	} `xml:"presentationTitle"`
	Speakers struct {
		Text    string `xml:",chardata"`
		Speaker struct {
			Text string `xml:",chardata"`
			Name struct {
				Text string `xml:",chardata"` // Giselle Agosto
			} `xml:"name"`
		} `xml:"speaker"`
	} `xml:"speakers"`
	Backgrounds struct {
		Text       string `xml:",chardata"`
		Background []struct {
			Text string `xml:",chardata"`
			URL  string `xml:"url,attr"`
		} `xml:"background"`
	} `xml:"backgrounds"`
	Search struct {
		Text string `xml:",chardata"`
		URL  string `xml:"url,attr"`
	} `xml:"search"`
	Slides struct {
		Text  string `xml:",chardata"`
		Slide []struct {
			Text        string `xml:",chardata"`
			ID          string `xml:"id,attr"`
			FrameRate   string `xml:"frameRate,attr"`
			TotalFrames string `xml:"totalFrames,attr"`
			Advance     string `xml:"advance,attr"`
			Type        string `xml:"type,attr"`
			Content     struct {
				Text string `xml:",chardata"`
				URL  string `xml:"url,attr"`
			} `xml:"content"`
			SlideBackground struct {
				Text           string `xml:",chardata"`
				Index          string `xml:"index,attr"`
				FrameToPlay    string `xml:"frameToPlay,attr"`
				Duration       string `xml:"duration,attr"`
				ShowAfterFrame string `xml:"showAfterFrame,attr"`
			} `xml:"slideBackground"`
			SlideTitle struct {
				Text string `xml:",chardata"` // INFORMACIÓN GENERAL DE L...
			} `xml:"slideTitle"`
			Notes struct {
				Text   string `xml:",chardata"` // Hola,&#x20;mi&#x20;nombre...
				IsHTML string `xml:"isHTML,attr"`
			} `xml:"notes"`
			Video struct {
				Text string `xml:",chardata"`
				URL  string `xml:"url,attr"`
			} `xml:"video"`
			SlideThumbnail struct {
				Text  string `xml:",chardata"`
				URL   string `xml:"url,attr"`
				Frame string `xml:"frame,attr"`
			} `xml:"slideThumbnail"`
			SlideSpeaker struct {
				Text  string `xml:",chardata"`
				Index string `xml:"index,attr"`
			} `xml:"slideSpeaker"`
			Audio struct {
				Text string `xml:",chardata"`
				URL  string `xml:"url,attr"`
			} `xml:"audio"`
		} `xml:"slide"`
	} `xml:"slides"`
	Attachments struct {
		Text string `xml:",chardata"`
	} `xml:"attachments"`
}
