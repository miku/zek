package misc
// Fuentes was generated 2017-12-05 23:56:07 by tir on apollo.
type Fuentes struct {
	XMLName xml.Name `xml:"fuentes"`
	Text    string   `xml:",chardata"`
	Fuente  []struct {
		Text string `xml:",chardata"`
		ID   struct {
			Text string `xml:",chardata"` // 3, 4, 7
		} `xml:"id"`
		Descripcion struct {
			Text string `xml:",chardata"` // Observatorio Demográfico...
		} `xml:"descripcion"`
		URLPublicacion struct {
			Text string `xml:",chardata"` // http://www.eclac.cl/cgi-b...
		} `xml:"url_publicacion"`
		SiglaOrganismo struct {
			Text string `xml:",chardata"` // CELADE, CELADE, CEPAL
		} `xml:"sigla_organismo"`
		NombreOrganismo struct {
			Text string `xml:",chardata"` // Centro Latinoamericano y ...
		} `xml:"nombre_organismo"`
		URLOrganizacion struct {
			Text string `xml:",chardata"` // http://www.eclac.cl/celad...
		} `xml:"url_organizacion"`
	} `xml:"fuente"`
} 

