package misc

import "encoding/xml"

// Sbml was generated 2017-12-06 00:01:15 by tir on apollo.
type Sbml struct {
	XMLName xml.Name `xml:"sbml"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Metaid  string   `xml:"metaid,attr"`
	Level   string   `xml:"level,attr"`
	Version string   `xml:"version,attr"`
	Model   struct {
		Text   string `xml:",chardata"`
		Metaid string `xml:"metaid,attr"`
		ID     string `xml:"id,attr"`
		Name   string `xml:"name,attr"`
		Notes  struct {
			Text string `xml:",chardata"`
			Body struct {
				Text  string `xml:",chardata"`
				Xmlns string `xml:"xmlns,attr"`
				P     struct {
					Text string `xml:",chardata"` // This model originates fro...
					Br   []struct {
						Text string `xml:",chardata"`
					} `xml:"br"`
					A []struct {
						Text   string `xml:",chardata"` // terms of use, Li C, Doniz...
						Href   string `xml:"href,attr"`
						Target string `xml:"target,attr"`
					} `xml:"a"`
				} `xml:"p"`
			} `xml:"body"`
		} `xml:"notes"`
		Annotation struct {
			Text string `xml:",chardata"`
			RDF  struct {
				Text        string `xml:",chardata"`
				Rdf         string `xml:"rdf,attr"`
				Dc          string `xml:"dc,attr"`
				Dcterms     string `xml:"dcterms,attr"`
				VCard       string `xml:"vCard,attr"`
				Bqbiol      string `xml:"bqbiol,attr"`
				Bqmodel     string `xml:"bqmodel,attr"`
				Description struct {
					Text    string `xml:",chardata"`
					About   string `xml:"about,attr"`
					Creator struct {
						Text      string `xml:",chardata"`
						ParseType string `xml:"parseType,attr"`
						Bag       struct {
							Text string `xml:",chardata"`
							Li   []struct {
								Text      string `xml:",chardata"`
								ParseType string `xml:"parseType,attr"`
								N         struct {
									Text      string `xml:",chardata"`
									ParseType string `xml:"parseType,attr"`
									Family    struct {
										Text string `xml:",chardata"` // Proctor, Chelliah, Gray
									} `xml:"Family"`
									Given struct {
										Text string `xml:",chardata"` // Carole, Vijayalakshmi, Do...
									} `xml:"Given"`
								} `xml:"N"`
								EMAIL struct {
									Text string `xml:",chardata"` // c.j.proctor@ncl.ac.uk, vi...
								} `xml:"EMAIL"`
								ORG struct {
									Text    string `xml:",chardata"`
									Orgname struct {
										Text string `xml:",chardata"` // University of Newcastle, ...
									} `xml:"Orgname"`
								} `xml:"ORG"`
							} `xml:"li"`
						} `xml:"Bag"`
					} `xml:"creator"`
					Created struct {
						Text      string `xml:",chardata"`
						ParseType string `xml:"parseType,attr"`
						W3CDTF    struct {
							Text string `xml:",chardata"` // 2008-09-05T13:47:15Z
						} `xml:"W3CDTF"`
					} `xml:"created"`
					Modified struct {
						Text      string `xml:",chardata"`
						ParseType string `xml:"parseType,attr"`
						W3CDTF    struct {
							Text string `xml:",chardata"` // 2010-02-02T13:35:39Z
						} `xml:"W3CDTF"`
					} `xml:"modified"`
					Is []struct {
						Text string `xml:",chardata"`
						Bag  struct {
							Text string `xml:",chardata"`
							Li   struct {
								Text     string `xml:",chardata"`
								Resource string `xml:"resource,attr"`
							} `xml:"li"`
						} `xml:"Bag"`
					} `xml:"is"`
					IsDescribedBy struct {
						Text string `xml:",chardata"`
						Bag  struct {
							Text string `xml:",chardata"`
							Li   struct {
								Text     string `xml:",chardata"`
								Resource string `xml:"resource,attr"`
							} `xml:"li"`
						} `xml:"Bag"`
					} `xml:"isDescribedBy"`
					OccursIn struct {
						Text string `xml:",chardata"`
						Bag  struct {
							Text string `xml:",chardata"`
							Li   struct {
								Text     string `xml:",chardata"`
								Resource string `xml:"resource,attr"`
							} `xml:"li"`
						} `xml:"Bag"`
					} `xml:"occursIn"`
					IsPartOf struct {
						Text string `xml:",chardata"`
						Bag  struct {
							Text string `xml:",chardata"`
							Li   []struct {
								Text     string `xml:",chardata"`
								Resource string `xml:"resource,attr"`
							} `xml:"li"`
						} `xml:"Bag"`
					} `xml:"isPartOf"`
					IsVersionOf struct {
						Text string `xml:",chardata"`
						Bag  struct {
							Text string `xml:",chardata"`
							Li   struct {
								Text     string `xml:",chardata"`
								Resource string `xml:"resource,attr"`
							} `xml:"li"`
						} `xml:"Bag"`
					} `xml:"isVersionOf"`
				} `xml:"Description"`
			} `xml:"RDF"`
		} `xml:"annotation"`
		ListOfUnitDefinitions struct {
			Text           string `xml:",chardata"`
			UnitDefinition []struct {
				Text        string `xml:",chardata"`
				Metaid      string `xml:"metaid,attr"`
				ID          string `xml:"id,attr"`
				Name        string `xml:"name,attr"`
				ListOfUnits struct {
					Text string `xml:",chardata"`
					Unit []struct {
						Text       string `xml:",chardata"`
						Kind       string `xml:"kind,attr"`
						Exponent   string `xml:"exponent,attr"`
						Multiplier string `xml:"multiplier,attr"`
					} `xml:"unit"`
				} `xml:"listOfUnits"`
			} `xml:"unitDefinition"`
		} `xml:"listOfUnitDefinitions"`
		ListOfCompartments struct {
			Text        string `xml:",chardata"`
			Compartment struct {
				Text       string `xml:",chardata"`
				Metaid     string `xml:"metaid,attr"`
				ID         string `xml:"id,attr"`
				Size       string `xml:"size,attr"`
				Annotation struct {
					Text string `xml:",chardata"`
					RDF  struct {
						Text        string `xml:",chardata"`
						Rdf         string `xml:"rdf,attr"`
						Dc          string `xml:"dc,attr"`
						Dcterms     string `xml:"dcterms,attr"`
						VCard       string `xml:"vCard,attr"`
						Bqbiol      string `xml:"bqbiol,attr"`
						Bqmodel     string `xml:"bqmodel,attr"`
						Description struct {
							Text  string `xml:",chardata"`
							About string `xml:"about,attr"`
							Is    struct {
								Text string `xml:",chardata"`
								Bag  struct {
									Text string `xml:",chardata"`
									Li   struct {
										Text     string `xml:",chardata"`
										Resource string `xml:"resource,attr"`
									} `xml:"li"`
								} `xml:"Bag"`
							} `xml:"is"`
						} `xml:"Description"`
					} `xml:"RDF"`
				} `xml:"annotation"`
			} `xml:"compartment"`
		} `xml:"listOfCompartments"`
		ListOfSpecies struct {
			Text    string `xml:",chardata"`
			Species []struct {
				Text                  string `xml:",chardata"`
				Metaid                string `xml:"metaid,attr"`
				ID                    string `xml:"id,attr"`
				Compartment           string `xml:"compartment,attr"`
				InitialAmount         string `xml:"initialAmount,attr"`
				HasOnlySubstanceUnits string `xml:"hasOnlySubstanceUnits,attr"`
				SboTerm               string `xml:"sboTerm,attr"`
				BoundaryCondition     string `xml:"boundaryCondition,attr"`
				Constant              string `xml:"constant,attr"`
				Annotation            struct {
					Text string `xml:",chardata"`
					RDF  struct {
						Text        string `xml:",chardata"`
						Rdf         string `xml:"rdf,attr"`
						Dc          string `xml:"dc,attr"`
						Dcterms     string `xml:"dcterms,attr"`
						VCard       string `xml:"vCard,attr"`
						Bqbiol      string `xml:"bqbiol,attr"`
						Bqmodel     string `xml:"bqmodel,attr"`
						Description struct {
							Text  string `xml:",chardata"`
							About string `xml:"about,attr"`
							Is    struct {
								Text string `xml:",chardata"`
								Bag  struct {
									Text string `xml:",chardata"`
									Li   struct {
										Text     string `xml:",chardata"`
										Resource string `xml:"resource,attr"`
									} `xml:"li"`
								} `xml:"Bag"`
							} `xml:"is"`
							IsEncodedBy struct {
								Text string `xml:",chardata"`
								Bag  struct {
									Text string `xml:",chardata"`
									Li   struct {
										Text     string `xml:",chardata"`
										Resource string `xml:"resource,attr"`
									} `xml:"li"`
								} `xml:"Bag"`
							} `xml:"isEncodedBy"`
							IsVersionOf struct {
								Text string `xml:",chardata"`
								Bag  struct {
									Text string `xml:",chardata"`
									Li   []struct {
										Text     string `xml:",chardata"`
										Resource string `xml:"resource,attr"`
									} `xml:"li"`
								} `xml:"Bag"`
							} `xml:"isVersionOf"`
							HasPart struct {
								Text string `xml:",chardata"`
								Bag  struct {
									Text string `xml:",chardata"`
									Li   []struct {
										Text     string `xml:",chardata"`
										Resource string `xml:"resource,attr"`
									} `xml:"li"`
								} `xml:"Bag"`
							} `xml:"hasPart"`
						} `xml:"Description"`
					} `xml:"RDF"`
				} `xml:"annotation"`
			} `xml:"species"`
		} `xml:"listOfSpecies"`
		ListOfParameters struct {
			Text      string `xml:",chardata"`
			Parameter []struct {
				Text     string `xml:",chardata"`
				Metaid   string `xml:"metaid,attr"`
				ID       string `xml:"id,attr"`
				Value    string `xml:"value,attr"`
				Units    string `xml:"units,attr"`
				Constant string `xml:"constant,attr"`
			} `xml:"parameter"`
		} `xml:"listOfParameters"`
		ListOfRules struct {
			Text           string `xml:",chardata"`
			AssignmentRule []struct {
				Text     string `xml:",chardata"`
				Metaid   string `xml:"metaid,attr"`
				Variable string `xml:"variable,attr"`
				Math     struct {
					Text  string `xml:",chardata"`
					Xmlns string `xml:"xmlns,attr"`
					Apply struct {
						Text string `xml:",chardata"`
						Plus struct {
							Text string `xml:",chardata"`
						} `xml:"plus"`
						Ci []struct {
							Text string `xml:",chardata"` // p53, Mdm2_p53, Mdm2, Mdm2...
						} `xml:"ci"`
					} `xml:"apply"`
				} `xml:"math"`
			} `xml:"assignmentRule"`
		} `xml:"listOfRules"`
		ListOfReactions struct {
			Text     string `xml:",chardata"`
			Reaction []struct {
				Text       string `xml:",chardata"`
				Metaid     string `xml:"metaid,attr"`
				ID         string `xml:"id,attr"`
				Reversible string `xml:"reversible,attr"`
				Annotation struct {
					Text string `xml:",chardata"`
					RDF  struct {
						Text        string `xml:",chardata"`
						Rdf         string `xml:"rdf,attr"`
						Dc          string `xml:"dc,attr"`
						Dcterms     string `xml:"dcterms,attr"`
						VCard       string `xml:"vCard,attr"`
						Bqbiol      string `xml:"bqbiol,attr"`
						Bqmodel     string `xml:"bqmodel,attr"`
						Description struct {
							Text  string `xml:",chardata"`
							About string `xml:"about,attr"`
							Is    struct {
								Text string `xml:",chardata"`
								Bag  struct {
									Text string `xml:",chardata"`
									Li   struct {
										Text     string `xml:",chardata"`
										Resource string `xml:"resource,attr"`
									} `xml:"li"`
								} `xml:"Bag"`
							} `xml:"is"`
							IsVersionOf struct {
								Text string `xml:",chardata"`
								Bag  struct {
									Text string `xml:",chardata"`
									Li   struct {
										Text     string `xml:",chardata"`
										Resource string `xml:"resource,attr"`
									} `xml:"li"`
								} `xml:"Bag"`
							} `xml:"isVersionOf"`
						} `xml:"Description"`
					} `xml:"RDF"`
				} `xml:"annotation"`
				ListOfReactants struct {
					Text             string `xml:",chardata"`
					SpeciesReference []struct {
						Text    string `xml:",chardata"`
						Species string `xml:"species,attr"`
					} `xml:"speciesReference"`
				} `xml:"listOfReactants"`
				ListOfProducts struct {
					Text             string `xml:",chardata"`
					SpeciesReference []struct {
						Text    string `xml:",chardata"`
						Species string `xml:"species,attr"`
					} `xml:"speciesReference"`
				} `xml:"listOfProducts"`
				KineticLaw struct {
					Text string `xml:",chardata"`
					Math struct {
						Text  string `xml:",chardata"`
						Xmlns string `xml:"xmlns,attr"`
						Apply struct {
							Text  string `xml:",chardata"`
							Times struct {
								Text string `xml:",chardata"`
							} `xml:"times"`
							Ci []struct {
								Text string `xml:",chardata"` // ksynMdm2, Mdm2_mRNA, ksyn...
							} `xml:"ci"`
						} `xml:"apply"`
					} `xml:"math"`
				} `xml:"kineticLaw"`
			} `xml:"reaction"`
		} `xml:"listOfReactions"`
		ListOfEvents struct {
			Text  string `xml:",chardata"`
			Event []struct {
				Text       string `xml:",chardata"`
				Metaid     string `xml:"metaid,attr"`
				ID         string `xml:"id,attr"`
				Annotation struct {
					Text string `xml:",chardata"`
					RDF  struct {
						Text        string `xml:",chardata"`
						Rdf         string `xml:"rdf,attr"`
						Dc          string `xml:"dc,attr"`
						Dcterms     string `xml:"dcterms,attr"`
						VCard       string `xml:"vCard,attr"`
						Bqbiol      string `xml:"bqbiol,attr"`
						Bqmodel     string `xml:"bqmodel,attr"`
						Description struct {
							Text  string `xml:",chardata"`
							About string `xml:"about,attr"`
							Is    struct {
								Text string `xml:",chardata"`
								Bag  struct {
									Text string `xml:",chardata"`
									Li   struct {
										Text     string `xml:",chardata"`
										Resource string `xml:"resource,attr"`
									} `xml:"li"`
								} `xml:"Bag"`
							} `xml:"is"`
						} `xml:"Description"`
					} `xml:"RDF"`
				} `xml:"annotation"`
				Trigger struct {
					Text string `xml:",chardata"`
					Math struct {
						Text  string `xml:",chardata"`
						Xmlns string `xml:"xmlns,attr"`
						Apply struct {
							Text string `xml:",chardata"`
							Geq  struct {
								Text string `xml:",chardata"`
							} `xml:"geq"`
							Csymbol struct {
								Text          string `xml:",chardata"` // t, t
								Encoding      string `xml:"encoding,attr"`
								DefinitionURL string `xml:"definitionURL,attr"`
							} `xml:"csymbol"`
							Cn struct {
								Text string `xml:",chardata"` // 3600, 3660
							} `xml:"cn"`
						} `xml:"apply"`
					} `xml:"math"`
				} `xml:"trigger"`
				ListOfEventAssignments struct {
					Text            string `xml:",chardata"`
					EventAssignment struct {
						Text     string `xml:",chardata"`
						Variable string `xml:"variable,attr"`
						Math     struct {
							Text  string `xml:",chardata"`
							Xmlns string `xml:"xmlns,attr"`
							Cn    struct {
								Text string `xml:",chardata"` // 25, 0
							} `xml:"cn"`
						} `xml:"math"`
					} `xml:"eventAssignment"`
				} `xml:"listOfEventAssignments"`
			} `xml:"event"`
		} `xml:"listOfEvents"`
	} `xml:"model"`
}
