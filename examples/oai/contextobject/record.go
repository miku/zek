package contextobject

import "encoding/xml"

type Record struct {
	XMLName xml.Name `xml:"Record"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text       string `xml:",chardata"`
		Status     string `xml:"status,attr"`
		Identifier struct {
			Text string `xml:",chardata"` // oai:generic.eprints.org:3...
		} `xml:"identifier"`
		Datestamp struct {
			Text string `xml:",chardata"` // 2010-01-02T04:42:29Z, 201...
		} `xml:"datestamp"`
		SetSpec []struct {
			Text string `xml:",chardata"` // 7374617475733D707562, 737...
		} `xml:"setSpec"`
	} `xml:"header"`
	Metadata struct {
		Text          string `xml:",chardata"`
		ContextObject struct {
			Text           string `xml:",chardata"`
			Xsi            string `xml:"xsi,attr"`
			Ctx            string `xml:"ctx,attr"`
			Timestamp      string `xml:"timestamp,attr"`
			SchemaLocation string `xml:"schemaLocation,attr"`
			Referent       struct {
				Text       string `xml:",chardata"`
				Identifier struct {
					Text string `xml:",chardata"` // info:oai:generic.eprints....
				} `xml:"identifier"`
				MetadataByVal struct {
					Text   string `xml:",chardata"`
					Format struct {
						Text string `xml:",chardata"` // info:ofi/fmt:xml:xsd:jour...
					} `xml:"format"`
					Metadata struct {
						Text    string `xml:",chardata"`
						Journal struct {
							Text           string `xml:",chardata"`
							Xsi            string `xml:"xsi,attr"`
							Jnl            string `xml:"jnl,attr"`
							SchemaLocation string `xml:"schemaLocation,attr"`
							Dis            string `xml:"dis,attr"`
							Authors        struct {
								Text   string `xml:",chardata"`
								Author []struct {
									Text   string `xml:",chardata"`
									Aulast struct {
										Text string `xml:",chardata"` // Sharp, Mohan, Maneschi, V...
									} `xml:"aulast"`
									Aufirst struct {
										Text string `xml:",chardata"` // P S, V, F, F, H R, J M, E...
									} `xml:"aufirst"`
									Au struct {
										Text string `xml:",chardata"` // Sharp, P S, Mohan, V, Man...
									} `xml:"au"`
								} `xml:"author"`
							} `xml:"authors"`
							Issue struct {
								Text string `xml:",chardata"` // 1, 1, 3, 2, 1, 2, 2, 6, 1...
							} `xml:"issue"`
							Volume struct {
								Text string `xml:",chardata"` // 36, 13, 13, 13, 24, 5, 19...
							} `xml:"volume"`
							Date struct {
								Text string `xml:",chardata"` // 1987, 1987, 1987, 1987, 1...
							} `xml:"date"`
							ISSN struct {
								Text string `xml:",chardata"` // 0026-0495, 0338-1684, 033...
							} `xml:"issn"`
							Atitle struct {
								Text string `xml:",chardata"` // Changes in plasma growth ...
							} `xml:"atitle"`
							Title struct {
								Text string `xml:",chardata"` // Metabolism: clinical and ...
							} `xml:"title"`
							Pages struct {
								Text string `xml:",chardata"` // 71-5, 27-30, 193-7, 140-1...
							} `xml:"pages"`
							Genre struct {
								Text string `xml:",chardata"` // article, article, article...
							} `xml:"genre"`
							Btitle struct {
								Text string `xml:",chardata"` // Cardiology update - 2008,...
							} `xml:"btitle"`
							ISBN struct {
								Text string `xml:",chardata"` // 81-901452-0-7, 81-901452-...
							} `xml:"isbn"`
							Degree struct {
								Text string `xml:",chardata"` // PhD
							} `xml:"degree"`
							Inst struct {
								Text string `xml:",chardata"` // The Tamilnadu Dr M G R Me...
							} `xml:"inst"`
							Tpages struct {
								Text string `xml:",chardata"` // 317, 13, 4
							} `xml:"tpages"`
						} `xml:"journal"`
					} `xml:"metadata"`
				} `xml:"metadata-by-val"`
			} `xml:"referent"`
		} `xml:"context-object"`
	} `xml:"metadata"`
	About struct {
		Text string `xml:",chardata"`
	} `xml:"about"`
}
