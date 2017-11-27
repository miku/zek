package formats

import (
	"encoding/xml"
)

// RssItem from https://geoscan.nrcan.gc.ca/rss/newpub_e.rss.
type RssItem struct {
	Text  string `xml:",chardata"`
	Title struct {
		Text string `xml:",chardata"` // Surficial geology, Aberde...
	} `xml:"title"`
	Link struct {
		Text string `xml:",chardata"` // https://geoscan.nrcan.gc....
	} `xml:"link"`
	Description struct {
		Text string `xml:",chardata"` // Geological Survey of Cana...
	} `xml:"description"`
	Guid struct {
		Text        string `xml:",chardata"` // 304279, 306212, 306175, 3...
		IsPermaLink string `xml:"isPermaLink,attr"`
	} `xml:"guid"`
	PubDate struct {
		Text string `xml:",chardata"` // Fri, 24 Nov 2017 00:00:00...
	} `xml:"pubDate"`
	Polygon []struct {
		Text string `xml:",chardata"` // 64.0000 -98.0000 64.0000 ...
	} `xml:"polygon"`
	Download struct {
		Text string `xml:",chardata"` // https://geoscan.nrcan.gc....
	} `xml:"download"`
	License struct {
		Text string `xml:",chardata"` // http://data.gc.ca/eng/ope...
	} `xml:"license"`
	Author struct {
		Text string `xml:",chardata"` // Geological Survey of Cana...
	} `xml:"author"`
	Source struct {
		Text string `xml:",chardata"` // Geological Survey of Cana...
	} `xml:"source"`
	SndSeries struct {
		Text string `xml:",chardata"` // Bedford Institute of Ocea...
	} `xml:"SndSeries"`
	Publisher struct {
		Text string `xml:",chardata"` // Natural Resources Canada,...
	} `xml:"publisher"`
	Edition struct {
		Text string `xml:",chardata"` // prelim., surficial data m...
	} `xml:"edition"`
	Meeting struct {
		Text string `xml:",chardata"` // Geological Association of...
	} `xml:"meeting"`
	Documenttype struct {
		Text string `xml:",chardata"` // serial, open file, serial...
	} `xml:"documenttype"`
	Language struct {
		Text string `xml:",chardata"` // English, English, English...
	} `xml:"language"`
	Maps struct {
		Text string `xml:",chardata"` // 1 map, 5 maps, Publicatio...
	} `xml:"maps"`
	Mapinfo struct {
		Text string `xml:",chardata"` // surficial geology, surfic...
	} `xml:"mapinfo"`
	Medium struct {
		Text string `xml:",chardata"` // on-line; digital, digital...
	} `xml:"medium"`
	Province struct {
		Text string `xml:",chardata"` // Nunavut, Northwest Territ...
	} `xml:"province"`
	Nts struct {
		Text string `xml:",chardata"` // 066B, 095J; 095N; 095O; 0...
	} `xml:"nts"`
	Area struct {
		Text string `xml:",chardata"` // Aberdeen Lake, Mackenzie ...
	} `xml:"area"`
	Subjects struct {
		Text string `xml:",chardata"`
	} `xml:"subjects"`
	Program struct {
		Text string `xml:",chardata"` // GEM2: Geo-mapping for Ene...
	} `xml:"program"`
	Project struct {
		Text string `xml:",chardata"` // Rae Province Project Mana...
	} `xml:"project"`
	Projectnumber struct {
		Text string `xml:",chardata"` // 340521, 343202, 340557, 3...
	} `xml:"projectnumber"`
	Abstract struct {
		Text string `xml:",chardata"` // This new surficial geolog...
	} `xml:"abstract"`
	Links struct {
		Text string `xml:",chardata"` // Online - En ligne (PDF, 9...
	} `xml:"links"`
	Readme struct {
		Text string `xml:",chardata"` // readme | https://geoscan....
	} `xml:"readme"`
	PPIid struct {
		Text string `xml:",chardata"` // 34532, 35096, 35438, 2563...
	} `xml:"PPIid"`
}

// Rss from https://geoscan.nrcan.gc.ca/rss/newpub_e.rss.
type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Rdf     string   `xml:"rdf,attr"`
	Dc      string   `xml:"dc,attr"`
	Geoscan string   `xml:"geoscan,attr"`
	Media   string   `xml:"media,attr"`
	Gml     string   `xml:"gml,attr"`
	Taxo    string   `xml:"taxo,attr"`
	Georss  string   `xml:"georss,attr"`
	Content string   `xml:"content,attr"`
	Geo     string   `xml:"geo,attr"`
	Version string   `xml:"version,attr"`
	Channel struct {
		Text  string `xml:",chardata"`
		Title struct {
			Text string `xml:",chardata"` // ESS New Releases (Display...
		} `xml:"title"`
		Link struct {
			Text string `xml:",chardata"` // http://tinyurl.com/ESSNew...
		} `xml:"link"`
		Description struct {
			Text string `xml:",chardata"` // New releases from the Ear...
		} `xml:"description"`
		LastBuildDate struct {
			Text string `xml:",chardata"` // Mon, 27 Nov 2017 00:06:35...
		} `xml:"lastBuildDate"`
		Items []Item `xml:"item"`
	} `xml:"channel"`
}
