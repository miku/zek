package misc

import "encoding/xml"

// StorageUnit was generated 2017-12-06 01:02:45 by tir on apollo.
type StorageUnit struct {
	XMLName            xml.Name `xml:"_StorageUnit"`
	Text               string   `xml:",chardata"`
	Release            string   `xml:"release,attr"`
	Originator         string   `xml:"originator,attr"`
	PackageGuid        string   `xml:"packageGuid,attr"`
	Time               string   `xml:"time,attr"`
	CHEMNonStdChemComp struct {
		Text                string `xml:",chardata"`
		ID                  string `xml:"_ID,attr"`
		CcpCode             string `xml:"ccpCode,attr"`
		CreatedBy           string `xml:"createdBy,attr"`
		Guid                string `xml:"guid,attr"`
		MolType             string `xml:"molType,attr"`
		CHEMChemCompDetails struct {
			Text     string `xml:",chardata"`
			IMPLText struct {
				Text string `xml:",chardata"` // Acetic acid with acetate ...
			} `xml:"IMPL.Text"`
		} `xml:"CHEM.ChemComp.details"`
		CHEMChemCompName struct {
			Text     string `xml:",chardata"`
			IMPLText struct {
				Text string `xml:",chardata"` // Acetic acid
			} `xml:"IMPL.Text"`
		} `xml:"CHEM.ChemComp.name"`
		CHEMChemCompChemAtomSets struct {
			Text            string `xml:",chardata"`
			CHEMChemAtomSet struct {
				Text                     string `xml:",chardata"`
				ID                       string `xml:"_ID,attr"`
				IsEquivalent             string `xml:"isEquivalent,attr"`
				IsProchiral              string `xml:"isProchiral,attr"`
				Name                     string `xml:"name,attr"`
				CHEMChemAtomSetChemAtoms struct {
					Text string `xml:",chardata"` // _3 _4 _5
				} `xml:"CHEM.ChemAtomSet.chemAtoms"`
			} `xml:"CHEM.ChemAtomSet"`
		} `xml:"CHEM.ChemComp.chemAtomSets"`
		CHEMChemCompChemAtoms struct {
			Text         string `xml:",chardata"`
			CHEMChemAtom []struct {
				Text                      string `xml:",chardata"`
				ID                        string `xml:"_ID,attr"`
				ElementSymbol             string `xml:"elementSymbol,attr"`
				Name                      string `xml:"name,attr"`
				WaterExchangeable         string `xml:"waterExchangeable,attr"`
				CHEMChemAtomBoundLinkEnds struct {
					Text string `xml:",chardata"` // _7
				} `xml:"CHEM.ChemAtom.boundLinkEnds"`
				CHEMAbstractChemAtomChemBonds struct {
					Text string `xml:",chardata"` // _8 _9 _10 _11, _16 _9 _17...
				} `xml:"CHEM.AbstractChemAtom.chemBonds"`
				CHEMAbstractChemAtomChemCompVars struct {
					Text string `xml:",chardata"` // _12 _13 _14, _12 _13 _14,...
				} `xml:"CHEM.AbstractChemAtom.chemCompVars"`
			} `xml:"CHEM.ChemAtom"`
			CHEMLinkAtom struct {
				Text                     string `xml:",chardata"`
				ID                       string `xml:"_ID,attr"`
				Name                     string `xml:"name,attr"`
				CHEMLinkAtomBoundLinkEnd struct {
					Text string `xml:",chardata"` // _7
				} `xml:"CHEM.LinkAtom.boundLinkEnd"`
				CHEMAbstractChemAtomChemBonds struct {
					Text string `xml:",chardata"` // _8
				} `xml:"CHEM.AbstractChemAtom.chemBonds"`
				CHEMAbstractChemAtomChemCompVars struct {
					Text string `xml:",chardata"` // _12
				} `xml:"CHEM.AbstractChemAtom.chemCompVars"`
			} `xml:"CHEM.LinkAtom"`
		} `xml:"CHEM.ChemComp.chemAtoms"`
		CHEMChemCompChemBonds struct {
			Text         string `xml:",chardata"`
			CHEMChemBond []struct {
				Text                  string `xml:",chardata"`
				ID                    string `xml:"_ID,attr"`
				BondType              string `xml:"bondType,attr"`
				CHEMChemBondChemAtoms struct {
					Text string `xml:",chardata"` // _6 _15, _22 _19, _6 _22, ...
				} `xml:"CHEM.ChemBond.chemAtoms"`
			} `xml:"CHEM.ChemBond"`
		} `xml:"CHEM.ChemComp.chemBonds"`
		CHEMChemCompChemCompVars struct {
			Text            string `xml:",chardata"`
			CHEMChemCompVar []struct {
				Text                      string `xml:",chardata"`
				ID                        string `xml:"_ID,attr"`
				FormalCharge              string `xml:"formalCharge,attr"`
				IsAromatic                string `xml:"isAromatic,attr"`
				IsDefaultVar              string `xml:"isDefaultVar,attr"`
				IsParamagnetic            string `xml:"isParamagnetic,attr"`
				Linking                   string `xml:"linking,attr"`
				CHEMChemCompVarDescriptor struct {
					Text     string `xml:",chardata"`
					IMPLLine struct {
						Text string `xml:",chardata"` // deprot:HO1, prot:HO1, neu...
					} `xml:"IMPL.Line"`
				} `xml:"CHEM.ChemCompVar.descriptor"`
				CHEMChemCompVarChemAtoms struct {
					Text string `xml:",chardata"` // _6 _22 _4 _21 _3 _15 _5, ...
				} `xml:"CHEM.ChemCompVar.chemAtoms"`
			} `xml:"CHEM.ChemCompVar"`
		} `xml:"CHEM.ChemComp.chemCompVars"`
		CHEMChemCompLinkEnds struct {
			Text        string `xml:",chardata"`
			CHEMLinkEnd struct {
				Text                     string `xml:",chardata"`
				ID                       string `xml:"_ID,attr"`
				LinkCode                 string `xml:"linkCode,attr"`
				CHEMLinkEndBoundLinkAtom struct {
					Text string `xml:",chardata"` // _23
				} `xml:"CHEM.LinkEnd.boundLinkAtom"`
			} `xml:"CHEM.LinkEnd"`
		} `xml:"CHEM.ChemComp.linkEnds"`
	} `xml:"CHEM.NonStdChemComp"`
}
