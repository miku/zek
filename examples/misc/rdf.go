package misc

import "encoding/xml"

// RDF was generated 2018-07-09 16:09:13 by tir on sol.
type RDF struct {
	XMLName           xml.Name `xml:"RDF"`
	Text              string   `xml:",chardata"`
	Rdf               string   `xml:"rdf,attr"`
	Dc                string   `xml:"dc,attr"`
	Functx            string   `xml:"functx,attr"`
	Gn                string   `xml:"gn,attr"`
	Szd               string   `xml:"szd,attr"`
	Skos              string   `xml:"skos,attr"`
	T                 string   `xml:"t,attr"`
	Dcterms           string   `xml:"dcterms,attr"`
	Gnd               string   `xml:"gnd,attr"`
	Gndo              string   `xml:"gndo,attr"`
	Owl               string   `xml:"owl,attr"`
	Rdfs              string   `xml:"rdfs,attr"`
	Foaf              string   `xml:"foaf,attr"`
	Gams              string   `xml:"gams,attr"`
	BiographicalEvent []struct {
		Text           string `xml:",chardata"`
		About          string `xml:"about,attr"`
		TextualContent struct {
			Text string `xml:",chardata"` // Wien,  28. November 1881 ...
		} `xml:"textualContent"`
		When struct {
			Text string `xml:",chardata"` // 1881-11-28, 1892-04, 1895...
		} `xml:"when"`
		Head struct {
			Text string `xml:",chardata"` // Wien, 28. November 1881, ...
		} `xml:"head"`
		Content struct {
			Text string `xml:",chardata"` // Stefan Zweig wird als zwe...
		} `xml:"content"`
		RelationTo []struct {
			Text     string `xml:",chardata"`
			Resource string `xml:"resource,attr"`
		} `xml:"relationTo"`
		IsMemberOfCollection struct {
			Text     string `xml:",chardata"`
			Resource string `xml:"resource,attr"`
		} `xml:"isMemberOfCollection"`
	} `xml:"BiographicalEvent"`
}
