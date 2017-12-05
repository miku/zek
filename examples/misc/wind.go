package misc

import "encoding/xml"

// Layer was generated 2017-12-05 23:34:45 by tir on apollo.
type Layer struct {
	XMLName     xml.Name `xml:"Layer"`
	Text        string   `xml:",chardata"`
	Version     string   `xml:"version,attr"`
	LayerType   string   `xml:"layerType,attr"`
	DisplayName struct {
		Text string `xml:",chardata"` // Political Boundaries
	} `xml:"DisplayName"`
	Service struct {
		Text               string `xml:",chardata"`
		ServiceName        string `xml:"serviceName,attr"`
		Version            string `xml:"version,attr"`
		GetCapabilitiesURL struct {
			Text string `xml:",chardata"` // http://worldwind22.arc.na...
		} `xml:"GetCapabilitiesURL"`
		GetMapURL struct {
			Text string `xml:",chardata"` // http://worldwind22.arc.na...
		} `xml:"GetMapURL"`
		LayerNames struct {
			Text string `xml:",chardata"` // topp:cia
		} `xml:"LayerNames"`
		StyleNames struct {
			Text string `xml:",chardata"` // countryboundaries
		} `xml:"StyleNames"`
	} `xml:"Service"`
	LastUpdate struct {
		Text string `xml:",chardata"` // 25 11 2009 22:21:47 GMT
	} `xml:"LastUpdate"`
	DataCacheName struct {
		Text string `xml:",chardata"` // Earth/PoliticalBoundaries...
	} `xml:"DataCacheName"`
	ImageFormat struct {
		Text string `xml:",chardata"` // image/png
	} `xml:"ImageFormat"`
	AvailableImageFormats struct {
		Text        string `xml:",chardata"`
		ImageFormat struct {
			Text string `xml:",chardata"` // image/png
		} `xml:"ImageFormat"`
	} `xml:"AvailableImageFormats"`
	FormatSuffix struct {
		Text string `xml:",chardata"` // .dds
	} `xml:"FormatSuffix"`
	NumLevels struct {
		Text     string `xml:",chardata"`
		Count    string `xml:"count,attr"`
		NumEmpty string `xml:"numEmpty,attr"`
	} `xml:"NumLevels"`
	TileOrigin struct {
		Text   string `xml:",chardata"`
		LatLon struct {
			Text      string `xml:",chardata"`
			Units     string `xml:"units,attr"`
			Latitude  string `xml:"latitude,attr"`
			Longitude string `xml:"longitude,attr"`
		} `xml:"LatLon"`
	} `xml:"TileOrigin"`
	LevelZeroTileDelta struct {
		Text   string `xml:",chardata"`
		LatLon struct {
			Text      string `xml:",chardata"`
			Units     string `xml:"units,attr"`
			Latitude  string `xml:"latitude,attr"`
			Longitude string `xml:"longitude,attr"`
		} `xml:"LatLon"`
	} `xml:"LevelZeroTileDelta"`
	TileSize struct {
		Text      string `xml:",chardata"`
		Dimension struct {
			Text   string `xml:",chardata"`
			Width  string `xml:"width,attr"`
			Height string `xml:"height,attr"`
		} `xml:"Dimension"`
	} `xml:"TileSize"`
	Sector struct {
		Text      string `xml:",chardata"`
		SouthWest struct {
			Text   string `xml:",chardata"`
			LatLon struct {
				Text      string `xml:",chardata"`
				Units     string `xml:"units,attr"`
				Latitude  string `xml:"latitude,attr"`
				Longitude string `xml:"longitude,attr"`
			} `xml:"LatLon"`
		} `xml:"SouthWest"`
		NorthEast struct {
			Text   string `xml:",chardata"`
			LatLon struct {
				Text      string `xml:",chardata"`
				Units     string `xml:"units,attr"`
				Latitude  string `xml:"latitude,attr"`
				Longitude string `xml:"longitude,attr"`
			} `xml:"LatLon"`
		} `xml:"NorthEast"`
	} `xml:"Sector"`
	UseTransparentTextures struct {
		Text string `xml:",chardata"` // true
	} `xml:"UseTransparentTextures"`
}
