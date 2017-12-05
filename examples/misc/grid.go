package misc

import "encoding/xml"

// Metadata was generated 2017-12-05 23:35:35 by tir on apollo.
type Metadata struct {
	XMLName xml.Name `xml:"metadata"`
	Text    string   `xml:",chardata"`
	Idinfo  struct {
		Text     string `xml:",chardata"`
		Datsetid struct {
			Text string `xml:",chardata"` // gov.noaa.nws.ohd.hdsc: OHD-HDSC-NOAAAtlas14V8v2013...
		} `xml:"datsetid"`
		Citation struct {
			Text     string `xml:",chardata"`
			Citeinfo struct {
				Text   string `xml:",chardata"`
				Origin struct {
					Text string `xml:",chardata"` // NOAA/NWS/Office of Hydrologic Development, Hydrome...
				} `xml:"origin"`
				Pubdate struct {
					Text string `xml:",chardata"` // 20130415
				} `xml:"pubdate"`
				Title struct {
					Text string `xml:",chardata"` // Precipitation Frequency for Midwestern states, USA...
				} `xml:"title"`
				Geoform struct {
					Text string `xml:",chardata"` // raster digital data
				} `xml:"geoform"`
				Serinfo struct {
					Text    string `xml:",chardata"`
					Sername struct {
						Text string `xml:",chardata"` // NOAA Atlas 14
					} `xml:"sername"`
					Issue struct {
						Text string `xml:",chardata"` // Volume 8
					} `xml:"issue"`
				} `xml:"serinfo"`
				Pubinfo struct {
					Text     string `xml:",chardata"`
					Pubplace struct {
						Text string `xml:",chardata"` // Silver Spring, MD, USA
					} `xml:"pubplace"`
					Publish struct {
						Text string `xml:",chardata"` // NOAA/NWS/Office of Hydrologic Development, Hydrome...
					} `xml:"publish"`
				} `xml:"pubinfo"`
				Onlink []struct {
					Text string `xml:",chardata"` // http://hdsc.nws.noaa.gov/hdsc/pfds/pfds_gis.html, ...
				} `xml:"onlink"`
			} `xml:"citeinfo"`
		} `xml:"citation"`
		Descript struct {
			Text     string `xml:",chardata"`
			Abstract struct {
				Text string `xml:",chardata"` // This GIS grid atlas contains precipitation frequen...
			} `xml:"abstract"`
			Purpose struct {
				Text string `xml:",chardata"` // Precipitation frequency (PF) estimates are used as...
			} `xml:"purpose"`
			Supplinf struct {
				Text string `xml:",chardata"` // None
			} `xml:"supplinf"`
		} `xml:"descript"`
		Timeperd struct {
			Text     string `xml:",chardata"`
			Timeinfo struct {
				Text     string `xml:",chardata"`
				Rngdates struct {
					Text    string `xml:",chardata"`
					Begdate struct {
						Text string `xml:",chardata"` // 18361101
					} `xml:"begdate"`
					Enddate struct {
						Text string `xml:",chardata"` // 20130301
					} `xml:"enddate"`
				} `xml:"rngdates"`
			} `xml:"timeinfo"`
			Current struct {
				Text string `xml:",chardata"` // Climatological period from which the point precipi...
			} `xml:"current"`
		} `xml:"timeperd"`
		Status struct {
			Text     string `xml:",chardata"`
			Progress struct {
				Text string `xml:",chardata"` // Complete
			} `xml:"progress"`
			Update struct {
				Text string `xml:",chardata"` // As needed.
			} `xml:"update"`
		} `xml:"status"`
		Spdom struct {
			Text     string `xml:",chardata"`
			Bounding struct {
				Text   string `xml:",chardata"`
				Westbc struct {
					Text string `xml:",chardata"` // -109.3625
				} `xml:"westbc"`
				Eastbc struct {
					Text string `xml:",chardata"` // -82.371913
				} `xml:"eastbc"`
				Northbc struct {
					Text string `xml:",chardata"` // 49.428522
				} `xml:"northbc"`
				Southbc struct {
					Text string `xml:",chardata"` // 33.3125
				} `xml:"southbc"`
			} `xml:"bounding"`
		} `xml:"spdom"`
		Keywords struct {
			Text  string `xml:",chardata"`
			Theme []struct {
				Text    string `xml:",chardata"`
				Themekt struct {
					Text string `xml:",chardata"` // NASA GCMD Science Keywords, ISO 19115 Topic Catego...
				} `xml:"themekt"`
				Themekey struct {
					Text string `xml:",chardata"` // Atmosphere > Precipitation > Precipitation Frequen...
				} `xml:"themekey"`
			} `xml:"theme"`
			Place struct {
				Text    string `xml:",chardata"`
				Placekt struct {
					Text string `xml:",chardata"` // None
				} `xml:"placekt"`
				Placekey struct {
					Text string `xml:",chardata"` // Midwestern States
				} `xml:"placekey"`
			} `xml:"place"`
			Stratum struct {
				Text    string `xml:",chardata"`
				Stratkt struct {
					Text string `xml:",chardata"` // NASA GCMD Location Keywords
				} `xml:"stratkt"`
				Stratkey struct {
					Text string `xml:",chardata"` // Vertical Location > Ground Layer
				} `xml:"stratkey"`
			} `xml:"stratum"`
		} `xml:"keywords"`
		Thelayid struct {
			Text     string `xml:",chardata"`
			Numthlay struct {
				Text string `xml:",chardata"` // 1
			} `xml:"numthlay"`
			Layrname struct {
				Text  string `xml:",chardata"`
				Theme struct {
					Text    string `xml:",chardata"`
					Themekt struct {
						Text string `xml:",chardata"` // NASA GCMD Science Keywords
					} `xml:"themekt"`
					Themekey struct {
						Text string `xml:",chardata"` // Atmosphere > Precipitation > Precipitation Frequen...
					} `xml:"themekey"`
				} `xml:"theme"`
			} `xml:"layrname"`
		} `xml:"thelayid"`
		Accconst struct {
			Text string `xml:",chardata"` // n/a, no restrictions apply and data is public doma...
		} `xml:"accconst"`
		Useconst struct {
			Text string `xml:",chardata"` // Acknowledgement of NOAA/National Weather Service, ...
		} `xml:"useconst"`
		Ptcontac struct {
			Text    string `xml:",chardata"`
			Cntinfo struct {
				Text    string `xml:",chardata"`
				Cntorgp struct {
					Text   string `xml:",chardata"`
					Cntorg struct {
						Text string `xml:",chardata"` // NOAA/National Weather Service, Office of Hydrologi...
					} `xml:"cntorg"`
				} `xml:"cntorgp"`
				Cntaddr struct {
					Text     string `xml:",chardata"`
					Addrtype struct {
						Text string `xml:",chardata"` // mailing and physical address
					} `xml:"addrtype"`
					Address []struct {
						Text string `xml:",chardata"` // NOAA/NWS/Office of Hydrologic Development, SSMC2, ...
					} `xml:"address"`
					City struct {
						Text string `xml:",chardata"` // Silver Spring
					} `xml:"city"`
					State struct {
						Text string `xml:",chardata"` // MD
					} `xml:"state"`
					Postal struct {
						Text string `xml:",chardata"` // 20910
					} `xml:"postal"`
					Country struct {
						Text string `xml:",chardata"` // USA
					} `xml:"country"`
				} `xml:"cntaddr"`
				Cntemail struct {
					Text string `xml:",chardata"` // hdsc.questions@noaa.gov
				} `xml:"cntemail"`
				Hours struct {
					Text string `xml:",chardata"` // 9 a.m. to 5 p.m. Eastern, Monday-Friday
				} `xml:"hours"`
			} `xml:"cntinfo"`
		} `xml:"ptcontac"`
		Browse struct {
			Text    string `xml:",chardata"`
			Browsen struct {
				Text string `xml:",chardata"` // http://hdsc.nws.noaa.gov/hdsc/pfds/meta/images/mw1...
			} `xml:"browsen"`
			Browsed struct {
				Text string `xml:",chardata"` // MW 100-year 24-hour
			} `xml:"browsed"`
			Browset struct {
				Text string `xml:",chardata"` // PNG
			} `xml:"browset"`
		} `xml:"browse"`
		Secinfo struct {
			Text   string `xml:",chardata"`
			Secsys struct {
				Text string `xml:",chardata"` // None
			} `xml:"secsys"`
			Secclass struct {
				Text string `xml:",chardata"` // Unclassified
			} `xml:"secclass"`
			Sechandl struct {
				Text string `xml:",chardata"` // None
			} `xml:"sechandl"`
		} `xml:"secinfo"`
		Native struct {
			Text string `xml:",chardata"` // Microsoft Windows XP/7
		} `xml:"native"`
		Crossref []struct {
			Text     string `xml:",chardata"`
			Citeinfo struct {
				Text   string `xml:",chardata"`
				Origin struct {
					Text string `xml:",chardata"` // National Oceanic and Atmospheric Administration, N...
				} `xml:"origin"`
				Pubdate struct {
					Text string `xml:",chardata"` // 1977, 1961, 1973, 1964
				} `xml:"pubdate"`
				Title struct {
					Text string `xml:",chardata"` // Technical Memorandum NWS Hydro 35, Five to 60-minu...
				} `xml:"title"`
				Geoform struct {
					Text string `xml:",chardata"` // document, document, document, document
				} `xml:"geoform"`
				Serinfo struct {
					Text    string `xml:",chardata"`
					Sername struct {
						Text string `xml:",chardata"` // TM Hydro35, TP40, NOAA Atlas 2, TP49
					} `xml:"sername"`
					Issue struct {
						Text string `xml:",chardata"` // None, None, Vol 11, None
					} `xml:"issue"`
				} `xml:"serinfo"`
				Pubinfo struct {
					Text     string `xml:",chardata"`
					Pubplace struct {
						Text string `xml:",chardata"` // Silver Spring, MD, USA, Silver Spring, MD, USA, Si...
					} `xml:"pubplace"`
					Publish struct {
						Text string `xml:",chardata"` // National Weather Service, U.S. Weather Bureau, NOA...
					} `xml:"publish"`
				} `xml:"pubinfo"`
				Onlink struct {
					Text string `xml:",chardata"` // http://www.nws.noaa.gov/oh/hdsc/PF_documents/Techn...
				} `xml:"onlink"`
			} `xml:"citeinfo"`
		} `xml:"crossref"`
	} `xml:"idinfo"`
	Dataqual struct {
		Text    string `xml:",chardata"`
		Attracc struct {
			Text     string `xml:",chardata"`
			Attraccr struct {
				Text string `xml:",chardata"` // All attributes created were tested for internal co...
			} `xml:"attraccr"`
		} `xml:"attracc"`
		Logic struct {
			Text string `xml:",chardata"` // All data were based on precipitation-reporting sta...
		} `xml:"logic"`
		Complete struct {
			Text string `xml:",chardata"` // Observed point precipitation values originated fro...
		} `xml:"complete"`
		Lineage struct {
			Text    string `xml:",chardata"`
			Srcinfo []struct {
				Text    string `xml:",chardata"`
				Srccite struct {
					Text     string `xml:",chardata"`
					Citeinfo struct {
						Text   string `xml:",chardata"`
						Origin struct {
							Text string `xml:",chardata"` // National Climatic Data Center (NCDC), National Cli...
						} `xml:"origin"`
						Pubdate struct {
							Text string `xml:",chardata"` // 2011, 2011, 2011, Unknown, 2009, Unknown, Unknown,...
						} `xml:"pubdate"`
						Title struct {
							Text string `xml:",chardata"` // 15-minute Precipitation Data (DSI-3260), Hourly Pr...
						} `xml:"title"`
						Geoform struct {
							Text string `xml:",chardata"` // tabular digital data, tabular digital data, tabula...
						} `xml:"geoform"`
						Serinfo struct {
							Text    string `xml:",chardata"`
							Sername struct {
								Text string `xml:",chardata"` // DSI-3260, DSI-3240, DSI-3200, DSI-3202, DSI-3206, ...
							} `xml:"sername"`
							Issue struct {
								Text string `xml:",chardata"` // None, None, None, None, None, None, None, None, No...
							} `xml:"issue"`
						} `xml:"serinfo"`
						Pubinfo struct {
							Text     string `xml:",chardata"`
							Pubplace struct {
								Text string `xml:",chardata"` // Asheville, NC, Asheville, NC, Asheville, NC, City ...
							} `xml:"pubplace"`
							Publish struct {
								Text string `xml:",chardata"` // National Climatic Data Center (NCDC), National Cli...
							} `xml:"publish"`
						} `xml:"pubinfo"`
						Onlink struct {
							Text string `xml:",chardata"` // http://lwf.ncdc.noaa.gov/oa/ncdc.html, http://cdo....
						} `xml:"onlink"`
					} `xml:"citeinfo"`
				} `xml:"srccite"`
				Typesrc struct {
					Text string `xml:",chardata"` // online, online, online, online, http://agebb.misso...
				} `xml:"typesrc"`
				Srctime struct {
					Text     string `xml:",chardata"`
					Timeinfo struct {
						Text     string `xml:",chardata"`
						Rngdates struct {
							Text    string `xml:",chardata"`
							Begdate struct {
								Text string `xml:",chardata"` // 19700801, 19490701, 18400301, 19990901, 18900101, ...
							} `xml:"begdate"`
							Enddate struct {
								Text string `xml:",chardata"` // 20101231, 20101231, 20111031, 20081231, 20091031, ...
							} `xml:"enddate"`
						} `xml:"rngdates"`
					} `xml:"timeinfo"`
					Srccurr struct {
						Text string `xml:",chardata"` // ground condition, ground condition, ground conditi...
					} `xml:"srccurr"`
				} `xml:"srctime"`
				Srccitea struct {
					Text string `xml:",chardata"` // NOAA National Climatic Data Center (NCDC) ArcGIS S...
				} `xml:"srccitea"`
				Srccontr struct {
					Text string `xml:",chardata"` // 15-minute precipitation information, hourly precip...
				} `xml:"srccontr"`
			} `xml:"srcinfo"`
			Procstep struct {
				Text     string `xml:",chardata"`
				Procdesc struct {
					Text string `xml:",chardata"` // Datasets
				} `xml:"procdesc"`
				Procdate struct {
					Text string `xml:",chardata"` // Unknown
				} `xml:"procdate"`
			} `xml:"procstep"`
		} `xml:"lineage"`
	} `xml:"dataqual"`
	Spdoinfo struct {
		Text     string `xml:",chardata"`
		Indspref struct {
			Text string `xml:",chardata"` // This spatial dataset was prepared by spatially int...
		} `xml:"indspref"`
		Direct struct {
			Text string `xml:",chardata"` // Raster
		} `xml:"direct"`
		Rastinfo struct {
			Text     string `xml:",chardata"`
			Cvaltype struct {
				Text string `xml:",chardata"` // integer
			} `xml:"cvaltype"`
			Rasttype struct {
				Text string `xml:",chardata"` // Grid Cell
			} `xml:"rasttype"`
			Rowcount struct {
				Text string `xml:",chardata"` // 3239
			} `xml:"rowcount"`
			Colcount struct {
				Text string `xml:",chardata"` // 1935
			} `xml:"colcount"`
			Vrtcount struct {
				Text string `xml:",chardata"` // 1
			} `xml:"vrtcount"`
		} `xml:"rastinfo"`
	} `xml:"spdoinfo"`
	Spref struct {
		Text     string `xml:",chardata"`
		Horizsys struct {
			Text     string `xml:",chardata"`
			Geograph struct {
				Text   string `xml:",chardata"`
				Latres struct {
					Text string `xml:",chardata"` // 0.00833333
				} `xml:"latres"`
				Longres struct {
					Text string `xml:",chardata"` // 0.00833333
				} `xml:"longres"`
				Geogunit struct {
					Text string `xml:",chardata"` // Decimal degrees
				} `xml:"geogunit"`
			} `xml:"geograph"`
			Geodetic struct {
				Text    string `xml:",chardata"`
				Horizdn struct {
					Text string `xml:",chardata"` // North American Datum of 1983
				} `xml:"horizdn"`
				Ellips struct {
					Text string `xml:",chardata"` // Geodetic Reference System 80
				} `xml:"ellips"`
				Semiaxis struct {
					Text string `xml:",chardata"` // 6378137.000000
				} `xml:"semiaxis"`
				Denflat struct {
					Text string `xml:",chardata"` // 298.257222
				} `xml:"denflat"`
			} `xml:"geodetic"`
		} `xml:"horizsys"`
	} `xml:"spref"`
	Eainfo struct {
		Text     string `xml:",chardata"`
		Detailed struct {
			Text   string `xml:",chardata"`
			Enttyp struct {
				Text    string `xml:",chardata"`
				Enttypl struct {
					Text string `xml:",chardata"` // precipitation frequency grid cell value
				} `xml:"enttypl"`
				Enttypd struct {
					Text string `xml:",chardata"` // ASCII integer
				} `xml:"enttypd"`
				Enttypds struct {
					Text string `xml:",chardata"` // Self-evident
				} `xml:"enttypds"`
			} `xml:"enttyp"`
			Attr []struct {
				Text     string `xml:",chardata"`
				Attrlabl struct {
					Text string `xml:",chardata"` // precipitation frequency grid cell value, Value
				} `xml:"attrlabl"`
				Attrdef struct {
					Text string `xml:",chardata"` // gridded precipitation frequency estimates, Feature...
				} `xml:"attrdef"`
				Attrdefs struct {
					Text string `xml:",chardata"` // NWS/OHD/HDSC, HDSC
				} `xml:"attrdefs"`
				Attrdomv struct {
					Text string `xml:",chardata"`
					Udom struct {
						Text string `xml:",chardata"` // Sequential unique whole numbers that are automatic...
					} `xml:"udom"`
					Rdom struct {
						Text    string `xml:",chardata"`
						Rdommin struct {
							Text string `xml:",chardata"` // 0
						} `xml:"rdommin"`
						Rdommax struct {
							Text string `xml:",chardata"` // 9999999
						} `xml:"rdommax"`
						Attrunit struct {
							Text string `xml:",chardata"` // 1000ths of inches; -9 is missing
						} `xml:"attrunit"`
					} `xml:"rdom"`
				} `xml:"attrdomv"`
				Attrvai struct {
					Text   string `xml:",chardata"`
					Attrva struct {
						Text string `xml:",chardata"` // .01
					} `xml:"attrva"`
					Attrvae struct {
						Text string `xml:",chardata"` // measurement accuracy of precipitation gauges in in...
					} `xml:"attrvae"`
				} `xml:"attrvai"`
			} `xml:"attr"`
		} `xml:"detailed"`
	} `xml:"eainfo"`
	Distinfo struct {
		Text    string `xml:",chardata"`
		Distrib struct {
			Text    string `xml:",chardata"`
			Cntinfo struct {
				Text    string `xml:",chardata"`
				Cntorgp struct {
					Text   string `xml:",chardata"`
					Cntorg struct {
						Text string `xml:",chardata"` // NOAA/NWS/OHD/Hydrometeorological Design Studies Ce...
					} `xml:"cntorg"`
				} `xml:"cntorgp"`
				Cntaddr struct {
					Text     string `xml:",chardata"`
					Addrtype struct {
						Text string `xml:",chardata"` // mailing address
					} `xml:"addrtype"`
					Address []struct {
						Text string `xml:",chardata"` // NOAA/NWS/Office of Hydrologic Development, SSMC2, ...
					} `xml:"address"`
					City struct {
						Text string `xml:",chardata"` // Silver Spring
					} `xml:"city"`
					State struct {
						Text string `xml:",chardata"` // MD
					} `xml:"state"`
					Postal struct {
						Text string `xml:",chardata"` // 20910
					} `xml:"postal"`
					Country struct {
						Text string `xml:",chardata"` // USA
					} `xml:"country"`
				} `xml:"cntaddr"`
			} `xml:"cntinfo"`
		} `xml:"distrib"`
		Resdesc struct {
			Text string `xml:",chardata"` // Downloadable data.
		} `xml:"resdesc"`
		Distliab struct {
			Text string `xml:",chardata"` // Users must assume responsibility to determine the ...
		} `xml:"distliab"`
		Stdorder struct {
			Text    string `xml:",chardata"`
			Digform struct {
				Text     string `xml:",chardata"`
				Digtinfo struct {
					Text     string `xml:",chardata"`
					Formname struct {
						Text string `xml:",chardata"` // ASCII
					} `xml:"formname"`
					Transize struct {
						Text string `xml:",chardata"` // 74.0
					} `xml:"transize"`
				} `xml:"digtinfo"`
				Digtopt struct {
					Text     string `xml:",chardata"`
					Onlinopt struct {
						Text     string `xml:",chardata"`
						Computer struct {
							Text     string `xml:",chardata"`
							Networka struct {
								Text     string `xml:",chardata"`
								Networkr []struct {
									Text string `xml:",chardata"` // HDSC PFDS;, mwXyrYY (X = return period, YY = durat...
								} `xml:"networkr"`
							} `xml:"networka"`
						} `xml:"computer"`
						Accinstr struct {
							Text string `xml:",chardata"` // Web download
						} `xml:"accinstr"`
					} `xml:"onlinopt"`
				} `xml:"digtopt"`
			} `xml:"digform"`
			Fees struct {
				Text string `xml:",chardata"` // None
			} `xml:"fees"`
		} `xml:"stdorder"`
		Techpreq struct {
			Text string `xml:",chardata"` // A grid- or raster-based GIS, such as ArcInfo (GRID...
		} `xml:"techpreq"`
	} `xml:"distinfo"`
	Metainfo struct {
		Text string `xml:",chardata"`
		Metd struct {
			Text string `xml:",chardata"` // 20130415
		} `xml:"metd"`
		Metc struct {
			Text    string `xml:",chardata"`
			Cntinfo struct {
				Text    string `xml:",chardata"`
				Cntorgp struct {
					Text   string `xml:",chardata"`
					Cntorg struct {
						Text string `xml:",chardata"` // NOAA/NWS/OHD/Hydrometeorological Design Studies Ce...
					} `xml:"cntorg"`
				} `xml:"cntorgp"`
				Cntpos struct {
					Text string `xml:",chardata"` // N/A
				} `xml:"cntpos"`
				Cntaddr struct {
					Text     string `xml:",chardata"`
					Addrtype struct {
						Text string `xml:",chardata"` // mailing address
					} `xml:"addrtype"`
					Address []struct {
						Text string `xml:",chardata"` // NOAA/NWS/Office of Hydrologic Development, SSMC2, ...
					} `xml:"address"`
					City struct {
						Text string `xml:",chardata"` // Silver Spring
					} `xml:"city"`
					State struct {
						Text string `xml:",chardata"` // MD
					} `xml:"state"`
					Postal struct {
						Text string `xml:",chardata"` // 20910
					} `xml:"postal"`
					Country struct {
						Text string `xml:",chardata"` // USA
					} `xml:"country"`
				} `xml:"cntaddr"`
				Cntemail struct {
					Text string `xml:",chardata"` // hdsc.questions@noaa.gov
				} `xml:"cntemail"`
				Hours struct {
					Text string `xml:",chardata"` // 9 a.m. to 5 p.m. Eastern, Monday-Friday
				} `xml:"hours"`
			} `xml:"cntinfo"`
		} `xml:"metc"`
		Metstdn struct {
			Text string `xml:",chardata"` // FGDC Content Standard for Digital Geospatial Metad...
		} `xml:"metstdn"`
		Metstdv struct {
			Text string `xml:",chardata"` // FGDC-STD-012-2002
		} `xml:"metstdv"`
		Mettc struct {
			Text string `xml:",chardata"` // local time
		} `xml:"mettc"`
	} `xml:"metainfo"`
}
