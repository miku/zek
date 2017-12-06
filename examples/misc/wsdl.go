package misc

import "encoding/xml"

// Definitions was generated 2017-12-06 00:50:11 by tir on apollo.
type Definitions struct {
	XMLName         xml.Name `xml:"definitions"`
	Text            string   `xml:",chardata"`
	Tm              string   `xml:"tm,attr"`
	Soapenc         string   `xml:"soapenc,attr"`
	Mime            string   `xml:"mime,attr"`
	Tns             string   `xml:"tns,attr"`
	Soap            string   `xml:"soap,attr"`
	S               string   `xml:"s,attr"`
	Soap12          string   `xml:"soap12,attr"`
	HTTP            string   `xml:"http,attr"`
	TargetNamespace string   `xml:"targetNamespace,attr"`
	Wsdl            string   `xml:"wsdl,attr"`
	Types           struct {
		Text   string `xml:",chardata"`
		Schema struct {
			Text               string `xml:",chardata"`
			ElementFormDefault string `xml:"elementFormDefault,attr"`
			TargetNamespace    string `xml:"targetNamespace,attr"`
			Element            []struct {
				Text        string `xml:",chardata"`
				Name        string `xml:"name,attr"`
				Nillable    string `xml:"nillable,attr"`
				Type        string `xml:"type,attr"`
				ComplexType struct {
					Text     string `xml:",chardata"`
					Sequence struct {
						Text    string `xml:",chardata"`
						Element []struct {
							Text      string `xml:",chardata"`
							MinOccurs string `xml:"minOccurs,attr"`
							MaxOccurs string `xml:"maxOccurs,attr"`
							Name      string `xml:"name,attr"`
							Type      string `xml:"type,attr"`
						} `xml:"element"`
					} `xml:"sequence"`
				} `xml:"complexType"`
			} `xml:"element"`
		} `xml:"schema"`
	} `xml:"types"`
	Message []struct {
		Text string `xml:",chardata"`
		Name string `xml:"name,attr"`
		Part []struct {
			Text    string `xml:",chardata"`
			Name    string `xml:"name,attr"`
			Element string `xml:"element,attr"`
			Type    string `xml:"type,attr"`
		} `xml:"part"`
	} `xml:"message"`
	PortType []struct {
		Text      string `xml:",chardata"`
		Name      string `xml:"name,attr"`
		Operation []struct {
			Text          string `xml:",chardata"`
			Name          string `xml:"name,attr"`
			Documentation struct {
				Text string `xml:",chardata"` // Get weather report for al...
				Wsdl string `xml:"wsdl,attr"`
			} `xml:"documentation"`
			Input struct {
				Text    string `xml:",chardata"`
				Message string `xml:"message,attr"`
			} `xml:"input"`
			Output struct {
				Text    string `xml:",chardata"`
				Message string `xml:"message,attr"`
			} `xml:"output"`
		} `xml:"operation"`
	} `xml:"portType"`
	Binding []struct {
		Text    string `xml:",chardata"`
		Name    string `xml:"name,attr"`
		Type    string `xml:"type,attr"`
		Binding struct {
			Text      string `xml:",chardata"`
			Transport string `xml:"transport,attr"`
			Verb      string `xml:"verb,attr"`
		} `xml:"binding"`
		Operation []struct {
			Text      string `xml:",chardata"`
			Name      string `xml:"name,attr"`
			Operation struct {
				Text       string `xml:",chardata"`
				SoapAction string `xml:"soapAction,attr"`
				Style      string `xml:"style,attr"`
				Location   string `xml:"location,attr"`
			} `xml:"operation"`
			Input struct {
				Text string `xml:",chardata"`
				Body struct {
					Text string `xml:",chardata"`
					Use  string `xml:"use,attr"`
				} `xml:"body"`
				UrlEncoded struct {
					Text string `xml:",chardata"`
				} `xml:"urlEncoded"`
				Content struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"content"`
			} `xml:"input"`
			Output struct {
				Text string `xml:",chardata"`
				Body struct {
					Text string `xml:",chardata"`
					Use  string `xml:"use,attr"`
				} `xml:"body"`
				MimeXml struct {
					Text string `xml:",chardata"`
					Part string `xml:"part,attr"`
				} `xml:"mimeXml"`
			} `xml:"output"`
		} `xml:"operation"`
	} `xml:"binding"`
	Service struct {
		Text string `xml:",chardata"`
		Name string `xml:"name,attr"`
		Port []struct {
			Text    string `xml:",chardata"`
			Name    string `xml:"name,attr"`
			Binding string `xml:"binding,attr"`
			Address struct {
				Text     string `xml:",chardata"`
				Location string `xml:"location,attr"`
			} `xml:"address"`
		} `xml:"port"`
	} `xml:"service"`
}
