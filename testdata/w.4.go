type Records struct {
	XMLName xml.Name `xml:"Records"`
	Text    string   `xml:",chardata"`
	Xsi     string   `xml:"xsi,attr"`
	Record  []struct {
		XMLName xml.Name `xml:"Record"`
		Text    string   `xml:",chardata"`
		Header  struct {
			XMLName    xml.Name `xml:"header"`
			Text       string   `xml:",chardata"`
			Status     string   `xml:"status,attr"`
			Identifier struct {
				XMLName xml.Name `xml:"identifier"`
				Text    string   `xml:",chardata"`
			}
			Datestamp struct {
				XMLName xml.Name `xml:"datestamp"`
				Text    string   `xml:",chardata"`
			}
			SetSpec struct {
				XMLName xml.Name `xml:"setSpec"`
				Text    string   `xml:",chardata"`
			}
		}
		Metadata struct {
			XMLName xml.Name `xml:"metadata"`
			Text    string   `xml:",chardata"`
			Rfc1807 struct {
				XMLName        xml.Name `xml:"rfc1807"`
				Text           string   `xml:",chardata"`
				Xmlns          string   `xml:"xmlns,attr"`
				Xsi            string   `xml:"xsi,attr"`
				SchemaLocation string   `xml:"schemaLocation,attr"`
				BibVersion     struct {
					XMLName xml.Name `xml:"bib-version"`
					Text    string   `xml:",chardata"`
				}
				ID struct {
					XMLName xml.Name `xml:"id"`
					Text    string   `xml:",chardata"`
				}
				Entry struct {
					XMLName xml.Name `xml:"entry"`
					Text    string   `xml:",chardata"`
				}
				Organization []struct {
					XMLName xml.Name `xml:"organization"`
					Text    string   `xml:",chardata"`
				}
				Title struct {
					XMLName xml.Name `xml:"title"`
					Text    string   `xml:",chardata"`
				}
				Type struct {
					XMLName xml.Name `xml:"type"`
					Text    string   `xml:",chardata"`
				}
				Author []struct {
					XMLName xml.Name `xml:"author"`
					Text    string   `xml:",chardata"`
				}
				Copyright struct {
					XMLName xml.Name `xml:"copyright"`
					Text    string   `xml:",chardata"`
				}
				OtherAccess struct {
					XMLName xml.Name `xml:"other_access"`
					Text    string   `xml:",chardata"`
				}
				Keyword struct {
					XMLName xml.Name `xml:"keyword"`
					Text    string   `xml:",chardata"`
				}
				Period []struct {
					XMLName xml.Name `xml:"period"`
					Text    string   `xml:",chardata"`
				}
				Monitoring struct {
					XMLName xml.Name `xml:"monitoring"`
					Text    string   `xml:",chardata"`
				}
				Language struct {
					XMLName xml.Name `xml:"language"`
					Text    string   `xml:",chardata"`
				}
				Abstract struct {
					XMLName xml.Name `xml:"abstract"`
					Text    string   `xml:",chardata"`
				}
				Date struct {
					XMLName xml.Name `xml:"date"`
					Text    string   `xml:",chardata"`
				}
			}
		}
		About struct {
			XMLName xml.Name `xml:"about"`
			Text    string   `xml:",chardata"`
		}
	}
}
