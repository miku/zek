package misc

import "encoding/xml"

// XMI was generated 2017-12-06 00:41:59 by tir on apollo.
type XMI struct {
	XMLName    xml.Name `xml:"XMI"`
	Text       string   `xml:",chardata"`
	XmiVersion string   `xml:"xmi.version,attr"`
	Model      string   `xml:"Model,attr"`
	Timestamp  string   `xml:"timestamp,attr"`
	XMIHeader  struct {
		Text             string `xml:",chardata"`
		XMIDocumentation struct {
			Text        string `xml:",chardata"`
			XMIExporter struct {
				Text string `xml:",chardata"` // Unisys.JCR.1
			} `xml:"XMI.exporter"`
			XMIExporterVersion struct {
				Text string `xml:",chardata"` // 1.3.2
			} `xml:"XMI.exporterVersion"`
		} `xml:"XMI.documentation"`
		XMIMetamodel struct {
			Text       string `xml:",chardata"`
			XmiName    string `xml:"xmi.name,attr"`
			XmiVersion string `xml:"xmi.version,attr"`
		} `xml:"XMI.metamodel"`
	} `xml:"XMI.header"`
	XMIContent struct {
		Text    string `xml:",chardata"`
		Package []struct {
			Text                   string `xml:",chardata"`
			XmiID                  string `xml:"xmi.id,attr"`
			Name                   string `xml:"name,attr"`
			IsRoot                 string `xml:"isRoot,attr"`
			IsLeaf                 string `xml:"isLeaf,attr"`
			IsAbstract             string `xml:"isAbstract,attr"`
			Visibility             string `xml:"visibility,attr"`
			Annotation             string `xml:"annotation,attr"`
			ModelElementAnnotation struct {
				Text string `xml:",chardata"` // The Core metamodel depend...
			} `xml:"ModelElement.annotation"`
			NamespaceContents struct {
				Text string `xml:",chardata"`
				Tag  []struct {
					Text       string `xml:",chardata"`
					XmiID      string `xml:"xmi.id,attr"`
					Name       string `xml:"name,attr"`
					Annotation string `xml:"annotation,attr"`
					TagId      string `xml:"tagId,attr"`
					Elements   string `xml:"elements,attr"`
					TagValues  struct {
						Text   string `xml:",chardata"`
						XMIAny struct {
							Text    string `xml:",chardata"` // org.omg.cwm.objectmodel, ...
							XmiType string `xml:"xmi.type,attr"`
							XmiName string `xml:"xmi.name,attr"`
						} `xml:"XMI.any"`
					} `xml:"Tag.values"`
				} `xml:"Tag"`
				Class []struct {
					Text                   string `xml:",chardata"`
					XmiID                  string `xml:"xmi.id,attr"`
					Name                   string `xml:"name,attr"`
					IsRoot                 string `xml:"isRoot,attr"`
					IsLeaf                 string `xml:"isLeaf,attr"`
					IsAbstract             string `xml:"isAbstract,attr"`
					Visibility             string `xml:"visibility,attr"`
					IsSingleton            string `xml:"isSingleton,attr"`
					Supertypes             string `xml:"supertypes,attr"`
					Annotation             string `xml:"annotation,attr"`
					ModelElementAnnotation struct {
						Text string `xml:",chardata"` // An element is an atomic c...
					} `xml:"ModelElement.annotation"`
					NamespaceContents struct {
						Text      string `xml:",chardata"`
						Attribute []struct {
							Text                          string `xml:",chardata"`
							XmiID                         string `xml:"xmi.id,attr"`
							Name                          string `xml:"name,attr"`
							Annotation                    string `xml:"annotation,attr"`
							Scope                         string `xml:"scope,attr"`
							Visibility                    string `xml:"visibility,attr"`
							IsChangeable                  string `xml:"isChangeable,attr"`
							IsDerived                     string `xml:"isDerived,attr"`
							Type                          string `xml:"type,attr"`
							StructuralFeatureMultiplicity struct {
								Text             string `xml:",chardata"`
								MultiplicityType struct {
									Text      string `xml:",chardata"`
									Lower     string `xml:"lower,attr"`
									Upper     string `xml:"upper,attr"`
									IsOrdered string `xml:"is_ordered,attr"`
									IsUnique  string `xml:"is_unique,attr"`
								} `xml:"MultiplicityType"`
							} `xml:"StructuralFeature.multiplicity"`
							ModelElementAnnotation struct {
								Text string `xml:",chardata"` // A textual representation ...
							} `xml:"ModelElement.annotation"`
						} `xml:"Attribute"`
						Reference []struct {
							Text                          string `xml:",chardata"`
							XmiID                         string `xml:"xmi.id,attr"`
							Name                          string `xml:"name,attr"`
							Annotation                    string `xml:"annotation,attr"`
							Scope                         string `xml:"scope,attr"`
							Visibility                    string `xml:"visibility,attr"`
							IsChangeable                  string `xml:"isChangeable,attr"`
							Type                          string `xml:"type,attr"`
							ReferencedEnd                 string `xml:"referencedEnd,attr"`
							StructuralFeatureMultiplicity struct {
								Text             string `xml:",chardata"`
								MultiplicityType struct {
									Text      string `xml:",chardata"`
									Lower     string `xml:"lower,attr"`
									Upper     string `xml:"upper,attr"`
									IsOrdered string `xml:"is_ordered,attr"`
									IsUnique  string `xml:"is_unique,attr"`
								} `xml:"MultiplicityType"`
							} `xml:"StructuralFeature.multiplicity"`
						} `xml:"Reference"`
					} `xml:"Namespace.contents"`
				} `xml:"Class"`
				DataType []struct {
					Text                   string `xml:",chardata"`
					XmiID                  string `xml:"xmi.id,attr"`
					Name                   string `xml:"name,attr"`
					IsRoot                 string `xml:"isRoot,attr"`
					IsLeaf                 string `xml:"isLeaf,attr"`
					IsAbstract             string `xml:"isAbstract,attr"`
					Visibility             string `xml:"visibility,attr"`
					Annotation             string `xml:"annotation,attr"`
					ModelElementAnnotation struct {
						Text string `xml:",chardata"` // The Any data type is used...
					} `xml:"ModelElement.annotation"`
					DataTypeTypeCode struct {
						Text             string `xml:",chardata"`
						XMICorbaTypeCode struct {
							Text          string `xml:",chardata"`
							XMICorbaTcAny struct {
								Text string `xml:",chardata"`
							} `xml:"XMI.CorbaTcAny"`
							XMICorbaTcBoolean struct {
								Text string `xml:",chardata"`
							} `xml:"XMI.CorbaTcBoolean"`
							XMICorbaTcFloat struct {
								Text string `xml:",chardata"`
							} `xml:"XMI.CorbaTcFloat"`
							XMICorbaTcLong struct {
								Text string `xml:",chardata"`
							} `xml:"XMI.CorbaTcLong"`
							XMICorbaTcAlias struct {
								Text             string `xml:",chardata"`
								XmiTcName        string `xml:"xmi.tcName,attr"`
								XMICorbaTypeCode struct {
									Text             string `xml:",chardata"`
									XMICorbaTcString struct {
										Text        string `xml:",chardata"`
										XmiTcLength string `xml:"xmi.tcLength,attr"`
									} `xml:"XMI.CorbaTcString"`
									XMICorbaTcLong struct {
										Text string `xml:",chardata"`
									} `xml:"XMI.CorbaTcLong"`
								} `xml:"XMI.CorbaTypeCode"`
							} `xml:"XMI.CorbaTcAlias"`
							XMICorbaTcString struct {
								Text        string `xml:",chardata"`
								XmiTcLength string `xml:"xmi.tcLength,attr"`
							} `xml:"XMI.CorbaTcString"`
							XMICorbaTcEnum struct {
								Text                string `xml:",chardata"`
								XmiTcName           string `xml:"xmi.tcName,attr"`
								XMICorbaTcEnumLabel []struct {
									Text      string `xml:",chardata"`
									XmiTcName string `xml:"xmi.tcName,attr"`
								} `xml:"XMI.CorbaTcEnumLabel"`
							} `xml:"XMI.CorbaTcEnum"`
						} `xml:"XMI.CorbaTypeCode"`
					} `xml:"DataType.typeCode"`
				} `xml:"DataType"`
				Association []struct {
					Text              string `xml:",chardata"`
					XmiID             string `xml:"xmi.id,attr"`
					Name              string `xml:"name,attr"`
					Annotation        string `xml:"annotation,attr"`
					IsRoot            string `xml:"isRoot,attr"`
					IsLeaf            string `xml:"isLeaf,attr"`
					IsAbstract        string `xml:"isAbstract,attr"`
					Visibility        string `xml:"visibility,attr"`
					IsDerived         string `xml:"isDerived,attr"`
					NamespaceContents struct {
						Text           string `xml:",chardata"`
						AssociationEnd []struct {
							Text                       string `xml:",chardata"`
							XmiID                      string `xml:"xmi.id,attr"`
							Name                       string `xml:"name,attr"`
							Annotation                 string `xml:"annotation,attr"`
							IsNavigable                string `xml:"isNavigable,attr"`
							Aggregation                string `xml:"aggregation,attr"`
							IsChangeable               string `xml:"isChangeable,attr"`
							Type                       string `xml:"type,attr"`
							AssociationEndMultiplicity struct {
								Text             string `xml:",chardata"`
								MultiplicityType struct {
									Text      string `xml:",chardata"`
									Lower     string `xml:"lower,attr"`
									Upper     string `xml:"upper,attr"`
									IsOrdered string `xml:"is_ordered,attr"`
									IsUnique  string `xml:"is_unique,attr"`
								} `xml:"MultiplicityType"`
							} `xml:"AssociationEnd.multiplicity"`
						} `xml:"AssociationEnd"`
					} `xml:"Namespace.contents"`
					ModelElementAnnotation struct {
						Text string `xml:",chardata"` // The StereotypeTaggedValue...
					} `xml:"ModelElement.annotation"`
				} `xml:"Association"`
				Import []struct {
					Text              string `xml:",chardata"`
					XmiID             string `xml:"xmi.id,attr"`
					Name              string `xml:"name,attr"`
					Annotation        string `xml:"annotation,attr"`
					Visibility        string `xml:"visibility,attr"`
					IsClustered       string `xml:"isClustered,attr"`
					ImportedNamespace string `xml:"importedNamespace,attr"`
				} `xml:"Import"`
				Package []struct {
					Text              string `xml:",chardata"`
					XmiID             string `xml:"xmi.id,attr"`
					Name              string `xml:"name,attr"`
					Annotation        string `xml:"annotation,attr"`
					IsRoot            string `xml:"isRoot,attr"`
					IsLeaf            string `xml:"isLeaf,attr"`
					IsAbstract        string `xml:"isAbstract,attr"`
					Visibility        string `xml:"visibility,attr"`
					NamespaceContents struct {
						Text     string `xml:",chardata"`
						DataType []struct {
							Text                   string `xml:",chardata"`
							XmiID                  string `xml:"xmi.id,attr"`
							Name                   string `xml:"name,attr"`
							IsRoot                 string `xml:"isRoot,attr"`
							IsLeaf                 string `xml:"isLeaf,attr"`
							IsAbstract             string `xml:"isAbstract,attr"`
							Visibility             string `xml:"visibility,attr"`
							Annotation             string `xml:"annotation,attr"`
							ModelElementAnnotation struct {
								Text string `xml:",chardata"` // Used in Trigger.  It indi...
							} `xml:"ModelElement.annotation"`
							DataTypeTypeCode struct {
								Text             string `xml:",chardata"`
								XMICorbaTypeCode struct {
									Text           string `xml:",chardata"`
									XMICorbaTcEnum struct {
										Text                string `xml:",chardata"`
										XmiTcName           string `xml:"xmi.tcName,attr"`
										XMICorbaTcEnumLabel []struct {
											Text      string `xml:",chardata"`
											XmiTcName string `xml:"xmi.tcName,attr"`
										} `xml:"XMI.CorbaTcEnumLabel"`
									} `xml:"XMI.CorbaTcEnum"`
									XMICorbaTcDouble struct {
										Text string `xml:",chardata"`
									} `xml:"XMI.CorbaTcDouble"`
								} `xml:"XMI.CorbaTypeCode"`
							} `xml:"DataType.typeCode"`
						} `xml:"DataType"`
						Class []struct {
							Text                   string `xml:",chardata"`
							XmiID                  string `xml:"xmi.id,attr"`
							Name                   string `xml:"name,attr"`
							IsRoot                 string `xml:"isRoot,attr"`
							IsLeaf                 string `xml:"isLeaf,attr"`
							IsAbstract             string `xml:"isAbstract,attr"`
							Visibility             string `xml:"visibility,attr"`
							IsSingleton            string `xml:"isSingleton,attr"`
							Supertypes             string `xml:"supertypes,attr"`
							Annotation             string `xml:"annotation,attr"`
							ModelElementAnnotation struct {
								Text string `xml:",chardata"` // This object provides a ma...
							} `xml:"ModelElement.annotation"`
							NamespaceContents struct {
								Text      string `xml:",chardata"`
								Reference []struct {
									Text                          string `xml:",chardata"`
									XmiID                         string `xml:"xmi.id,attr"`
									Name                          string `xml:"name,attr"`
									Annotation                    string `xml:"annotation,attr"`
									Scope                         string `xml:"scope,attr"`
									Visibility                    string `xml:"visibility,attr"`
									IsChangeable                  string `xml:"isChangeable,attr"`
									Type                          string `xml:"type,attr"`
									ReferencedEnd                 string `xml:"referencedEnd,attr"`
									StructuralFeatureMultiplicity struct {
										Text             string `xml:",chardata"`
										MultiplicityType struct {
											Text      string `xml:",chardata"`
											Lower     string `xml:"lower,attr"`
											Upper     string `xml:"upper,attr"`
											IsOrdered string `xml:"is_ordered,attr"`
											IsUnique  string `xml:"is_unique,attr"`
										} `xml:"MultiplicityType"`
									} `xml:"StructuralFeature.multiplicity"`
								} `xml:"Reference"`
								Attribute []struct {
									Text                   string `xml:",chardata"`
									XmiID                  string `xml:"xmi.id,attr"`
									Name                   string `xml:"name,attr"`
									Scope                  string `xml:"scope,attr"`
									Visibility             string `xml:"visibility,attr"`
									IsChangeable           string `xml:"isChangeable,attr"`
									IsDerived              string `xml:"isDerived,attr"`
									Type                   string `xml:"type,attr"`
									Annotation             string `xml:"annotation,attr"`
									ModelElementAnnotation struct {
										Text string `xml:",chardata"` // The usage attribute  indi...
									} `xml:"ModelElement.annotation"`
									StructuralFeatureMultiplicity struct {
										Text             string `xml:",chardata"`
										MultiplicityType struct {
											Text      string `xml:",chardata"`
											Lower     string `xml:"lower,attr"`
											Upper     string `xml:"upper,attr"`
											IsOrdered string `xml:"is_ordered,attr"`
											IsUnique  string `xml:"is_unique,attr"`
										} `xml:"MultiplicityType"`
									} `xml:"StructuralFeature.multiplicity"`
								} `xml:"Attribute"`
							} `xml:"Namespace.contents"`
						} `xml:"Class"`
						Tag []struct {
							Text       string `xml:",chardata"`
							XmiID      string `xml:"xmi.id,attr"`
							Name       string `xml:"name,attr"`
							Annotation string `xml:"annotation,attr"`
							TagId      string `xml:"tagId,attr"`
							Elements   string `xml:"elements,attr"`
							TagValues  struct {
								Text   string `xml:",chardata"`
								XMIAny struct {
									Text    string `xml:",chardata"` // umlAttribute, vsf_, mp_
									XmiType string `xml:"xmi.type,attr"`
									XmiName string `xml:"xmi.name,attr"`
								} `xml:"XMI.any"`
							} `xml:"Tag.values"`
						} `xml:"Tag"`
						Association []struct {
							Text              string `xml:",chardata"`
							XmiID             string `xml:"xmi.id,attr"`
							Name              string `xml:"name,attr"`
							Annotation        string `xml:"annotation,attr"`
							IsRoot            string `xml:"isRoot,attr"`
							IsLeaf            string `xml:"isLeaf,attr"`
							IsAbstract        string `xml:"isAbstract,attr"`
							Visibility        string `xml:"visibility,attr"`
							IsDerived         string `xml:"isDerived,attr"`
							NamespaceContents struct {
								Text           string `xml:",chardata"`
								AssociationEnd []struct {
									Text                       string `xml:",chardata"`
									XmiID                      string `xml:"xmi.id,attr"`
									Name                       string `xml:"name,attr"`
									Annotation                 string `xml:"annotation,attr"`
									IsNavigable                string `xml:"isNavigable,attr"`
									Aggregation                string `xml:"aggregation,attr"`
									IsChangeable               string `xml:"isChangeable,attr"`
									Type                       string `xml:"type,attr"`
									AssociationEndMultiplicity struct {
										Text             string `xml:",chardata"`
										MultiplicityType struct {
											Text      string `xml:",chardata"`
											Lower     string `xml:"lower,attr"`
											Upper     string `xml:"upper,attr"`
											IsOrdered string `xml:"is_ordered,attr"`
											IsUnique  string `xml:"is_unique,attr"`
										} `xml:"MultiplicityType"`
									} `xml:"AssociationEnd.multiplicity"`
								} `xml:"AssociationEnd"`
							} `xml:"Namespace.contents"`
						} `xml:"Association"`
					} `xml:"Namespace.contents"`
				} `xml:"Package"`
			} `xml:"Namespace.contents"`
		} `xml:"Package"`
	} `xml:"XMI.content"`
}
