package misc

import "encoding/xml"

// Neuroml was generated 2017-12-06 01:04:54 by tir on apollo.
type Neuroml struct {
	XMLName        xml.Name `xml:"neuroml"`
	Text           string   `xml:",chardata"`
	Xmlns          string   `xml:"xmlns,attr"`
	Meta           string   `xml:"meta,attr"`
	Mml            string   `xml:"mml,attr"`
	Bio            string   `xml:"bio,attr"`
	Cml            string   `xml:"cml,attr"`
	Net            string   `xml:"net,attr"`
	Xsi            string   `xml:"xsi,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	LengthUnits    string   `xml:"lengthUnits,attr"`
	Cells          struct {
		Text string `xml:",chardata"`
		Cell struct {
			Text  string `xml:",chardata"`
			Name  string `xml:"name,attr"`
			Notes struct {
				Text string `xml:",chardata"` // An implementation using C...
			} `xml:"notes"`
			Segments struct {
				Text    string `xml:",chardata"`
				Segment struct {
					Text     string `xml:",chardata"`
					ID       string `xml:"id,attr"`
					Name     string `xml:"name,attr"`
					Cable    string `xml:"cable,attr"`
					Proximal struct {
						Text     string `xml:",chardata"`
						X        string `xml:"x,attr"`
						Y        string `xml:"y,attr"`
						Z        string `xml:"z,attr"`
						Diameter string `xml:"diameter,attr"`
					} `xml:"proximal"`
					Distal struct {
						Text     string `xml:",chardata"`
						X        string `xml:"x,attr"`
						Y        string `xml:"y,attr"`
						Z        string `xml:"z,attr"`
						Diameter string `xml:"diameter,attr"`
					} `xml:"distal"`
				} `xml:"segment"`
			} `xml:"segments"`
			Cables struct {
				Text  string `xml:",chardata"`
				Cable struct {
					Text  string `xml:",chardata"`
					ID    string `xml:"id,attr"`
					Name  string `xml:"name,attr"`
					Group []struct {
						Text string `xml:",chardata"` // all, soma_group
					} `xml:"group"`
				} `xml:"cable"`
				Cablegroup struct {
					Text  string `xml:",chardata"`
					Name  string `xml:"name,attr"`
					Cable struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
					} `xml:"cable"`
				} `xml:"cablegroup"`
			} `xml:"cables"`
			Biophysics struct {
				Text      string `xml:",chardata"`
				Units     string `xml:"units,attr"`
				Mechanism []struct {
					Text               string `xml:",chardata"`
					Name               string `xml:"name,attr"`
					Type               string `xml:"type,attr"`
					PassiveConductance string `xml:"passive_conductance,attr"`
					Parameter          []struct {
						Text  string `xml:",chardata"`
						Name  string `xml:"name,attr"`
						Value string `xml:"value,attr"`
						Group struct {
							Text string `xml:",chardata"` // all, all, all, soma_group...
						} `xml:"group"`
					} `xml:"parameter"`
				} `xml:"mechanism"`
				SpecCapacitance struct {
					Text      string `xml:",chardata"`
					Parameter struct {
						Text  string `xml:",chardata"`
						Value string `xml:"value,attr"`
						Group struct {
							Text string `xml:",chardata"` // all
						} `xml:"group"`
					} `xml:"parameter"`
				} `xml:"spec_capacitance"`
				SpecAxialResistance struct {
					Text      string `xml:",chardata"`
					Parameter struct {
						Text  string `xml:",chardata"`
						Value string `xml:"value,attr"`
						Group struct {
							Text string `xml:",chardata"` // all
						} `xml:"group"`
					} `xml:"parameter"`
				} `xml:"spec_axial_resistance"`
				InitMembPotential struct {
					Text      string `xml:",chardata"`
					Parameter struct {
						Text  string `xml:",chardata"`
						Value string `xml:"value,attr"`
						Group struct {
							Text string `xml:",chardata"` // all
						} `xml:"group"`
					} `xml:"parameter"`
				} `xml:"init_memb_potential"`
				IonProps []struct {
					Text      string `xml:",chardata"`
					Name      string `xml:"name,attr"`
					Parameter []struct {
						Text  string `xml:",chardata"`
						Name  string `xml:"name,attr"`
						Value string `xml:"value,attr"`
						Group struct {
							Text string `xml:",chardata"` // all, all, all, all, all
						} `xml:"group"`
					} `xml:"parameter"`
				} `xml:"ion_props"`
			} `xml:"biophysics"`
			Connectivity struct {
				Text            string `xml:",chardata"`
				PotentialSynLoc []struct {
					Text        string `xml:",chardata"`
					SynapseType string `xml:"synapse_type,attr"`
					Group       struct {
						Text string `xml:",chardata"` // all, all, all, all
					} `xml:"group"`
				} `xml:"potential_syn_loc"`
			} `xml:"connectivity"`
		} `xml:"cell"`
	} `xml:"cells"`
}
