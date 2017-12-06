package misc

import "encoding/xml"

// Site was generated 2017-12-06 00:36:21 by tir on apollo.
// time curl -sL http://www.ins.cwi.nl/projects/xmark/Assets/standard.gz | gunzip -c | zek -e > xmark.go
// real 0m30.917s
// user 0m14.661s
// sys  0m2.095s
type Site struct {
	XMLName xml.Name `xml:"site"`
	Text    string   `xml:",chardata"`
	Regions struct {
		Text   string `xml:",chardata"`
		Africa struct {
			Text string `xml:",chardata"`
			Item []struct {
				Text     string `xml:",chardata"`
				ID       string `xml:"id,attr"`
				Featured string `xml:"featured,attr"`
				Location struct {
					Text string `xml:",chardata"` // United States, Moldova, R...
				} `xml:"location"`
				Quantity struct {
					Text string `xml:",chardata"` // 1, 1, 1, 1, 1, 1, 1, 1, 1...
				} `xml:"quantity"`
				Name struct {
					Text string `xml:",chardata"` // duteous nine eighteen, co...
				} `xml:"name"`
				Payment struct {
					Text string `xml:",chardata"` // Creditcard, Money order, ...
				} `xml:"payment"`
				Description struct {
					Chardata string `xml:",chardata"`
					Parlist  struct {
						Text     string `xml:",chardata"`
						Listitem []struct {
							Chardata string `xml:",chardata"`
							Text     struct {
								Text    string `xml:",chardata"` // page rous lady idle autho...
								Keyword []struct {
									Text string `xml:",chardata"` // officer embrace such fear...
									Emph []struct {
										Text string `xml:",chardata"` // person whate lily pilgrim...
									} `xml:"emph"`
									Bold struct {
										Text string `xml:",chardata"` // despite flourish musical ...
									} `xml:"bold"`
								} `xml:"keyword"`
								Bold []struct {
									Text string `xml:",chardata"` // naturally sanctuary riddl...
									Emph struct {
										Text string `xml:",chardata"` // sir hor shape suspicious ...
									} `xml:"emph"`
									Keyword []struct {
										Text string `xml:",chardata"` // petter boist viewed that ...
									} `xml:"keyword"`
								} `xml:"bold"`
								Emph []struct {
									Text    string `xml:",chardata"` // preventions spurr edmund ...
									Keyword []struct {
										Text string `xml:",chardata"` // cleomenes thrust, inn, pr...
									} `xml:"keyword"`
									Bold struct {
										Text string `xml:",chardata"` // conveyance voice strives ...
									} `xml:"bold"`
								} `xml:"emph"`
							} `xml:"text"`
							Parlist struct {
								Text     string `xml:",chardata"`
								Listitem []struct {
									Chardata string `xml:",chardata"`
									Text     struct {
										Text string `xml:",chardata"` // will little haunt reasons...
										Emph []struct {
											Text    string `xml:",chardata"` // were mounts true rounds h...
											Keyword struct {
												Text string `xml:",chardata"` // arrived sees wrestled mul...
											} `xml:"keyword"`
											Bold struct {
												Text string `xml:",chardata"` // undiscover heretic nobler...
											} `xml:"bold"`
										} `xml:"emph"`
										Bold []struct {
											Text string `xml:",chardata"` // sold, coventry entreat mu...
											Emph struct {
												Text string `xml:",chardata"` // dost lost toward corners ...
											} `xml:"emph"`
											Keyword struct {
												Text string `xml:",chardata"` // despite laud treasons min...
											} `xml:"keyword"`
										} `xml:"bold"`
										Keyword []struct {
											Text string `xml:",chardata"` // cimber paper admittance t...
											Bold []struct {
												Text string `xml:",chardata"` // note, cell rivers flesh l...
											} `xml:"bold"`
											Emph struct {
												Text string `xml:",chardata"` // antony commonwealth, abro...
											} `xml:"emph"`
										} `xml:"keyword"`
									} `xml:"text"`
								} `xml:"listitem"`
							} `xml:"parlist"`
						} `xml:"listitem"`
					} `xml:"parlist"`
					Text struct {
						Text string `xml:",chardata"` // gold promotions despair f...
						Bold []struct {
							Text    string `xml:",chardata"` // laws prey egg, brabantio,...
							Keyword []struct {
								Text string `xml:",chardata"` // colouring preventions pre...
							} `xml:"keyword"`
							Emph struct {
								Text string `xml:",chardata"` // trade pleas milk, mystery...
							} `xml:"emph"`
						} `xml:"bold"`
						Emph []struct {
							Text    string `xml:",chardata"` // woman, wanton pregnant re...
							Keyword struct {
								Text string `xml:",chardata"` // alexas madness countrymen...
							} `xml:"keyword"`
							Bold []struct {
								Text string `xml:",chardata"` // scarce, troths castle wat...
							} `xml:"bold"`
						} `xml:"emph"`
						Keyword []struct {
							Text string `xml:",chardata"` // twinn, bull bora senate k...
							Emph struct {
								Text string `xml:",chardata"` // behalf suspend crystal al...
							} `xml:"emph"`
							Bold []struct {
								Text string `xml:",chardata"` // york, hurt knits imprison...
							} `xml:"bold"`
						} `xml:"keyword"`
					} `xml:"text"`
				} `xml:"description"`
				Shipping struct {
					Text string `xml:",chardata"` // Will ship internationally...
				} `xml:"shipping"`
				Incategory []struct {
					Text     string `xml:",chardata"`
					Category string `xml:"category,attr"`
				} `xml:"incategory"`
				Mailbox struct {
					Text string `xml:",chardata"`
					Mail []struct {
						Chardata string `xml:",chardata"`
						From     struct {
							Text string `xml:",chardata"` // Libero Rive mailto:Rive@h...
						} `xml:"from"`
						To struct {
							Text string `xml:",chardata"` // Benedikte Glew mailto:Gle...
						} `xml:"to"`
						Date struct {
							Text string `xml:",chardata"` // 08/05/1999, 11/28/1999, 0...
						} `xml:"date"`
						Text struct {
							Text    string `xml:",chardata"` // virgin preventions half l...
							Keyword []struct {
								Text string `xml:",chardata"` // girdles deserts flood geo...
								Bold []struct {
									Text string `xml:",chardata"` // jar drabbing broken troth...
								} `xml:"bold"`
								Emph []struct {
									Text string `xml:",chardata"` // stranger, provost corrupt...
								} `xml:"emph"`
							} `xml:"keyword"`
							Emph []struct {
								Text string `xml:",chardata"` // armed, dotage discipline ...
								Bold []struct {
									Text string `xml:",chardata"` // tongue brains frenchman, ...
								} `xml:"bold"`
								Keyword []struct {
									Text string `xml:",chardata"` // whoever load drop, poise ...
								} `xml:"keyword"`
							} `xml:"emph"`
							Bold []struct {
								Text    string `xml:",chardata"` // childish carrions imagina...
								Keyword struct {
									Text string `xml:",chardata"` // displeasure sow law fooli...
								} `xml:"keyword"`
								Emph []struct {
									Text string `xml:",chardata"` // gar knife intemperate pai...
								} `xml:"emph"`
							} `xml:"bold"`
						} `xml:"text"`
					} `xml:"mail"`
				} `xml:"mailbox"`
			} `xml:"item"`
		} `xml:"africa"`
		Asia struct {
			Text string `xml:",chardata"`
			Item []struct {
				Text     string `xml:",chardata"`
				ID       string `xml:"id,attr"`
				Featured string `xml:"featured,attr"`
				Location struct {
					Text string `xml:",chardata"` // United States, United Sta...
				} `xml:"location"`
				Quantity struct {
					Text string `xml:",chardata"` // 1, 1, 1, 1, 1, 1, 1, 1, 1...
				} `xml:"quantity"`
				Name struct {
					Text string `xml:",chardata"` // stubborn ben huswife raim...
				} `xml:"name"`
				Payment struct {
					Text string `xml:",chardata"` // Money order, Creditcard, ...
				} `xml:"payment"`
				Description struct {
					Chardata string `xml:",chardata"`
					Text     struct {
						Text string `xml:",chardata"` // skin hisses valour fortin...
						Emph []struct {
							Text string `xml:",chardata"` // drove nemean sauce, porti...
							Bold []struct {
								Text string `xml:",chardata"` // fly, meditations clay lan...
							} `xml:"bold"`
							Keyword []struct {
								Text string `xml:",chardata"` // midwife fifty, marian, th...
							} `xml:"keyword"`
						} `xml:"emph"`
						Bold []struct {
							Text    string `xml:",chardata"` // flattering, satire make i...
							Keyword []struct {
								Text string `xml:",chardata"` // gear, unfortunate loathso...
							} `xml:"keyword"`
							Emph []struct {
								Text string `xml:",chardata"` // weeds shepherdess, pays, ...
							} `xml:"emph"`
						} `xml:"bold"`
						Keyword []struct {
							Text string `xml:",chardata"` // sentences knights note ki...
							Bold []struct {
								Text string `xml:",chardata"` // two thereby apothecary da...
							} `xml:"bold"`
							Emph []struct {
								Text string `xml:",chardata"` // knowing main monument bet...
							} `xml:"emph"`
						} `xml:"keyword"`
					} `xml:"text"`
					Parlist struct {
						Text     string `xml:",chardata"`
						Listitem []struct {
							Chardata string `xml:",chardata"`
							Text     struct {
								Text string `xml:",chardata"` // damnation flattery melted...
								Emph []struct {
									Text    string `xml:",chardata"` // look otherwise loyalty th...
									Keyword []struct {
										Text string `xml:",chardata"` // cure preventions, busy wo...
									} `xml:"keyword"`
									Bold []struct {
										Text string `xml:",chardata"` // islanders vows grows, str...
									} `xml:"bold"`
								} `xml:"emph"`
								Keyword []struct {
									Text string `xml:",chardata"` // altogether broad, enough ...
									Bold []struct {
										Text string `xml:",chardata"` // lear would, wipe, draw su...
									} `xml:"bold"`
									Emph []struct {
										Text string `xml:",chardata"` // bounds, word brings crave...
									} `xml:"emph"`
								} `xml:"keyword"`
								Bold []struct {
									Text    string `xml:",chardata"` // hope third yawn seek, dig...
									Keyword []struct {
										Text string `xml:",chardata"` // phebe ask lewd, shun, bes...
									} `xml:"keyword"`
									Emph []struct {
										Text string `xml:",chardata"` // holiday, preventions like...
									} `xml:"emph"`
								} `xml:"bold"`
							} `xml:"text"`
							Parlist struct {
								Text     string `xml:",chardata"`
								Listitem []struct {
									Chardata string `xml:",chardata"`
									Text     struct {
										Text string `xml:",chardata"` // parish pish sunset worse ...
										Emph []struct {
											Text string `xml:",chardata"` // slumber basket verges hap...
											Bold []struct {
												Text string `xml:",chardata"` // falstaff effect, standing...
											} `xml:"bold"`
											Keyword []struct {
												Text string `xml:",chardata"` // othello kisses witness fa...
											} `xml:"keyword"`
										} `xml:"emph"`
										Keyword []struct {
											Text string `xml:",chardata"` // romans cedius split wak c...
											Bold []struct {
												Text string `xml:",chardata"` // here makes remember, scru...
											} `xml:"bold"`
											Emph struct {
												Text string `xml:",chardata"` // honey horses led asunder ...
											} `xml:"emph"`
										} `xml:"keyword"`
										Bold []struct {
											Text    string `xml:",chardata"` // kingly cooling pitch, maj...
											Keyword []struct {
												Text string `xml:",chardata"` // laws overheard hides sir ...
											} `xml:"keyword"`
											Emph []struct {
												Text string `xml:",chardata"` // cooling stone ophelia, en...
											} `xml:"emph"`
										} `xml:"bold"`
									} `xml:"text"`
								} `xml:"listitem"`
							} `xml:"parlist"`
						} `xml:"listitem"`
					} `xml:"parlist"`
				} `xml:"description"`
				Shipping struct {
					Text string `xml:",chardata"` // Will ship internationally...
				} `xml:"shipping"`
				Incategory []struct {
					Text     string `xml:",chardata"`
					Category string `xml:"category,attr"`
				} `xml:"incategory"`
				Mailbox struct {
					Text string `xml:",chardata"`
					Mail []struct {
						Chardata string `xml:",chardata"`
						From     struct {
							Text string `xml:",chardata"` // Tamio Gomide mailto:Gomid...
						} `xml:"from"`
						To struct {
							Text string `xml:",chardata"` // Rudiger Grunberger mailto...
						} `xml:"to"`
						Date struct {
							Text string `xml:",chardata"` // 06/02/2000, 06/03/1998, 0...
						} `xml:"date"`
						Text struct {
							Text string `xml:",chardata"` // propos assaulted minute l...
							Emph []struct {
								Text string `xml:",chardata"` // exacting curiosity outfro...
								Bold []struct {
									Text string `xml:",chardata"` // particulars, apparel, lad...
								} `xml:"bold"`
								Keyword []struct {
									Text string `xml:",chardata"` // shin courtier cradle thei...
								} `xml:"keyword"`
							} `xml:"emph"`
							Bold []struct {
								Text string `xml:",chardata"` // hypocrite violence ratcli...
								Emph []struct {
									Text string `xml:",chardata"` // fame, volumes will disobe...
								} `xml:"emph"`
								Keyword []struct {
									Text string `xml:",chardata"` // saucy tarry drowning slav...
								} `xml:"keyword"`
							} `xml:"bold"`
							Keyword []struct {
								Text string `xml:",chardata"` // regard figures ventur ber...
								Emph []struct {
									Text string `xml:",chardata"` // richer badge vienna monum...
								} `xml:"emph"`
								Bold []struct {
									Text string `xml:",chardata"` // exits acted enquire came ...
								} `xml:"bold"`
							} `xml:"keyword"`
						} `xml:"text"`
					} `xml:"mail"`
				} `xml:"mailbox"`
			} `xml:"item"`
		} `xml:"asia"`
		Australia struct {
			Text string `xml:",chardata"`
			Item []struct {
				Text     string `xml:",chardata"`
				ID       string `xml:"id,attr"`
				Featured string `xml:"featured,attr"`
				Location struct {
					Text string `xml:",chardata"` // United States, United Sta...
				} `xml:"location"`
				Quantity struct {
					Text string `xml:",chardata"` // 1, 1, 1, 1, 1, 1, 1, 1, 1...
				} `xml:"quantity"`
				Name struct {
					Text string `xml:",chardata"` // fought hadst grieving, be...
				} `xml:"name"`
				Payment struct {
					Text string `xml:",chardata"` // Personal Check, Cash, Mon...
				} `xml:"payment"`
				Description struct {
					Chardata string `xml:",chardata"`
					Parlist  struct {
						Text     string `xml:",chardata"`
						Listitem []struct {
							Chardata string `xml:",chardata"`
							Text     struct {
								Text    string `xml:",chardata"` // little master paltry tomb...
								Keyword []struct {
									Text string `xml:",chardata"` // marks, faithful care says...
									Bold []struct {
										Text string `xml:",chardata"` // procurator frail bear, st...
									} `xml:"bold"`
									Emph struct {
										Text string `xml:",chardata"` // flay, spots captain, mask...
									} `xml:"emph"`
								} `xml:"keyword"`
								Emph []struct {
									Text    string `xml:",chardata"` // precisely, fashion sup jo...
									Keyword []struct {
										Text string `xml:",chardata"` // collatine, raise, revenge...
									} `xml:"keyword"`
									Bold []struct {
										Text string `xml:",chardata"` // left men being dry dram l...
									} `xml:"bold"`
								} `xml:"emph"`
								Bold []struct {
									Text    string `xml:",chardata"` // line weather legs almost ...
									Keyword []struct {
										Text string `xml:",chardata"` // doleful, happily though l...
									} `xml:"keyword"`
									Emph []struct {
										Text string `xml:",chardata"` // beseech seditious left ru...
									} `xml:"emph"`
								} `xml:"bold"`
							} `xml:"text"`
							Parlist struct {
								Text     string `xml:",chardata"`
								Listitem []struct {
									Chardata string `xml:",chardata"`
									Text     struct {
										Text string `xml:",chardata"` // honesty rascal may fairy ...
										Bold []struct {
											Text    string `xml:",chardata"` // possession liege fair lan...
											Keyword []struct {
												Text string `xml:",chardata"` // philip northumberland rep...
											} `xml:"keyword"`
											Emph []struct {
												Text string `xml:",chardata"` // waits thinks, lepidus wit...
											} `xml:"emph"`
										} `xml:"bold"`
										Keyword []struct {
											Text string `xml:",chardata"` // preventions, lord, meant ...
											Bold []struct {
												Text string `xml:",chardata"` // redeem tape blind royal m...
											} `xml:"bold"`
											Emph []struct {
												Text string `xml:",chardata"` // himself, enters, wrote ap...
											} `xml:"emph"`
										} `xml:"keyword"`
										Emph []struct {
											Text    string `xml:",chardata"` // envy, eve, lighted commen...
											Keyword struct {
												Text string `xml:",chardata"` // calpurnia stir hollow hap...
											} `xml:"keyword"`
											Bold struct {
												Text string `xml:",chardata"` // directly, gets living, so...
											} `xml:"bold"`
										} `xml:"emph"`
									} `xml:"text"`
								} `xml:"listitem"`
							} `xml:"parlist"`
						} `xml:"listitem"`
					} `xml:"parlist"`
					Text struct {
						Text string `xml:",chardata"` // gurney company well grove...
						Emph []struct {
							Text    string `xml:",chardata"` // whip, paris gallant down ...
							Keyword []struct {
								Text string `xml:",chardata"` // sage lend tapers pots ren...
							} `xml:"keyword"`
							Bold []struct {
								Text string `xml:",chardata"` // enough, first steeled ass...
							} `xml:"bold"`
						} `xml:"emph"`
						Keyword []struct {
							Text string `xml:",chardata"` // abroad, creature, country...
							Emph []struct {
								Text string `xml:",chardata"` // puissance ent, treble der...
							} `xml:"emph"`
							Bold []struct {
								Text string `xml:",chardata"` // ominous unmatched thorns ...
							} `xml:"bold"`
						} `xml:"keyword"`
						Bold []struct {
							Text    string `xml:",chardata"` // conveyance understand kee...
							Keyword []struct {
								Text string `xml:",chardata"` // trees notice miles, heroi...
							} `xml:"keyword"`
							Emph []struct {
								Text string `xml:",chardata"` // nice, thither, pasture tr...
							} `xml:"emph"`
						} `xml:"bold"`
					} `xml:"text"`
				} `xml:"description"`
				Shipping struct {
					Text string `xml:",chardata"` // Will ship only within cou...
				} `xml:"shipping"`
				Incategory []struct {
					Text     string `xml:",chardata"`
					Category string `xml:"category,attr"`
				} `xml:"incategory"`
				Mailbox struct {
					Text string `xml:",chardata"`
					Mail []struct {
						Chardata string `xml:",chardata"`
						From     struct {
							Text string `xml:",chardata"` // Sumali Promhouse mailto:P...
						} `xml:"from"`
						To struct {
							Text string `xml:",chardata"` // Mehrdad Nollmann mailto:N...
						} `xml:"to"`
						Date struct {
							Text string `xml:",chardata"` // 10/26/1999, 06/19/2000, 1...
						} `xml:"date"`
						Text struct {
							Text string `xml:",chardata"` // tokens rock linen shield ...
							Emph []struct {
								Text    string `xml:",chardata"` // wombs hug captain, duty s...
								Keyword []struct {
									Text string `xml:",chardata"` // whoreson orlando disdain ...
								} `xml:"keyword"`
								Bold []struct {
									Text string `xml:",chardata"` // paid phrases tartness gol...
								} `xml:"bold"`
							} `xml:"emph"`
							Keyword []struct {
								Text string `xml:",chardata"` // lady boot jest semblable ...
								Bold []struct {
									Text string `xml:",chardata"` // unthrifty quondam, palate...
								} `xml:"bold"`
								Emph []struct {
									Text string `xml:",chardata"` // porter compact chiding ni...
								} `xml:"emph"`
							} `xml:"keyword"`
							Bold []struct {
								Text string `xml:",chardata"` // ros moon misprizing profi...
								Emph []struct {
									Text string `xml:",chardata"` // christians rise tripping ...
								} `xml:"emph"`
								Keyword []struct {
									Text string `xml:",chardata"` // badge mark entreats frame...
								} `xml:"keyword"`
							} `xml:"bold"`
						} `xml:"text"`
					} `xml:"mail"`
				} `xml:"mailbox"`
			} `xml:"item"`
		} `xml:"australia"`
		Europe struct {
			Text string `xml:",chardata"`
			Item []struct {
				Text     string `xml:",chardata"`
				ID       string `xml:"id,attr"`
				Featured string `xml:"featured,attr"`
				Location struct {
					Text string `xml:",chardata"` // Liberia, United States, S...
				} `xml:"location"`
				Quantity struct {
					Text string `xml:",chardata"` // 1, 1, 2, 1, 1, 1, 1, 1, 1...
				} `xml:"quantity"`
				Name struct {
					Text string `xml:",chardata"` // insolence, prize hearts l...
				} `xml:"name"`
				Payment struct {
					Text string `xml:",chardata"` // Money order, Creditcard, ...
				} `xml:"payment"`
				Description struct {
					Chardata string `xml:",chardata"`
					Text     struct {
						Text    string `xml:",chardata"` // mus instrument pupil suck...
						Keyword []struct {
							Text string `xml:",chardata"` // posterity chanced grudge,...
							Emph []struct {
								Text string `xml:",chardata"` // behalf burgundy through w...
							} `xml:"emph"`
							Bold []struct {
								Text string `xml:",chardata"` // sociable clay fast will c...
							} `xml:"bold"`
						} `xml:"keyword"`
						Bold []struct {
							Text string `xml:",chardata"` // letter arm, subdued, leng...
							Emph []struct {
								Text string `xml:",chardata"` // appear gravity rive pines...
							} `xml:"emph"`
							Keyword []struct {
								Text string `xml:",chardata"` // navarre imprisonment couc...
							} `xml:"keyword"`
						} `xml:"bold"`
						Emph []struct {
							Text string `xml:",chardata"` // tardy stanley subjects st...
							Bold []struct {
								Text string `xml:",chardata"` // pronounce unadvis knowing...
							} `xml:"bold"`
							Keyword []struct {
								Text string `xml:",chardata"` // deliver bullets since, ea...
							} `xml:"keyword"`
						} `xml:"emph"`
					} `xml:"text"`
					Parlist struct {
						Text     string `xml:",chardata"`
						Listitem []struct {
							Chardata string `xml:",chardata"`
							Text     struct {
								Text    string `xml:",chardata"` // breath rosalinde, unmatch...
								Keyword []struct {
									Text string `xml:",chardata"` // solely endure, studies he...
									Bold []struct {
										Text string `xml:",chardata"` // sin perverted fain, does,...
									} `xml:"bold"`
									Emph []struct {
										Text string `xml:",chardata"` // weep engine, kingdom nut ...
									} `xml:"emph"`
								} `xml:"keyword"`
								Emph []struct {
									Text string `xml:",chardata"` // power grandam air sobbing...
									Bold []struct {
										Text string `xml:",chardata"` // ophelia live, contrive wa...
									} `xml:"bold"`
									Keyword []struct {
										Text string `xml:",chardata"` // money infer bloody pilgri...
									} `xml:"keyword"`
								} `xml:"emph"`
								Bold []struct {
									Text string `xml:",chardata"` // grapes nurse save public ...
									Emph []struct {
										Text string `xml:",chardata"` // sights countess antony ha...
									} `xml:"emph"`
									Keyword []struct {
										Text string `xml:",chardata"` // untimely servingman ear m...
									} `xml:"keyword"`
								} `xml:"bold"`
							} `xml:"text"`
							Parlist struct {
								Text     string `xml:",chardata"`
								Listitem []struct {
									Chardata string `xml:",chardata"`
									Text     struct {
										Text string `xml:",chardata"` // rome stranger level pace,...
										Emph []struct {
											Text    string `xml:",chardata"` // factious own violent legs...
											Keyword []struct {
												Text string `xml:",chardata"` // athwart tapers modest bon...
											} `xml:"keyword"`
											Bold []struct {
												Text string `xml:",chardata"` // jealous steeled assuredly...
											} `xml:"bold"`
										} `xml:"emph"`
										Bold []struct {
											Text    string `xml:",chardata"` // know, mounting demand fig...
											Keyword []struct {
												Text string `xml:",chardata"` // desires notest surges map...
											} `xml:"keyword"`
											Emph []struct {
												Text string `xml:",chardata"` // offended folded breathe, ...
											} `xml:"emph"`
										} `xml:"bold"`
										Keyword []struct {
											Text string `xml:",chardata"` // proportion castaways nece...
											Emph []struct {
												Text string `xml:",chardata"` // live practise branches, m...
											} `xml:"emph"`
											Bold []struct {
												Text string `xml:",chardata"` // worser, rack talent, gath...
											} `xml:"bold"`
										} `xml:"keyword"`
									} `xml:"text"`
								} `xml:"listitem"`
							} `xml:"parlist"`
						} `xml:"listitem"`
					} `xml:"parlist"`
				} `xml:"description"`
				Shipping struct {
					Text string `xml:",chardata"` // Buyer pays fixed shipping...
				} `xml:"shipping"`
				Incategory []struct {
					Text     string `xml:",chardata"`
					Category string `xml:"category,attr"`
				} `xml:"incategory"`
				Mailbox struct {
					Text string `xml:",chardata"`
					Mail []struct {
						Chardata string `xml:",chardata"`
						From     struct {
							Text string `xml:",chardata"` // King Gruenwald mailto:Gru...
						} `xml:"from"`
						To struct {
							Text string `xml:",chardata"` // Limsson Maamir mailto:Maa...
						} `xml:"to"`
						Date struct {
							Text string `xml:",chardata"` // 06/11/1998, 10/06/2000, 0...
						} `xml:"date"`
						Text struct {
							Text string `xml:",chardata"` // cerberus corner bills ran...
							Bold []struct {
								Text string `xml:",chardata"` // subdue trunk, suppose cus...
								Emph []struct {
									Text string `xml:",chardata"` // greet send stead, strain ...
								} `xml:"emph"`
								Keyword []struct {
									Text string `xml:",chardata"` // likeness air vices med, w...
								} `xml:"keyword"`
							} `xml:"bold"`
							Emph []struct {
								Text    string `xml:",chardata"` // steward, unruly lands bus...
								Keyword []struct {
									Text string `xml:",chardata"` // deeds hubert drops logger...
								} `xml:"keyword"`
								Bold []struct {
									Text string `xml:",chardata"` // manner ignoble, sold geff...
								} `xml:"bold"`
							} `xml:"emph"`
							Keyword []struct {
								Text string `xml:",chardata"` // weather second lower appr...
								Emph []struct {
									Text string `xml:",chardata"` // inconstant carved grown w...
								} `xml:"emph"`
								Bold []struct {
									Text string `xml:",chardata"` // home babe, percy rosalind...
								} `xml:"bold"`
							} `xml:"keyword"`
						} `xml:"text"`
					} `xml:"mail"`
				} `xml:"mailbox"`
			} `xml:"item"`
		} `xml:"europe"`
		Namerica struct {
			Text string `xml:",chardata"`
			Item []struct {
				Text     string `xml:",chardata"`
				ID       string `xml:"id,attr"`
				Featured string `xml:"featured,attr"`
				Location struct {
					Text string `xml:",chardata"` // United States, United Sta...
				} `xml:"location"`
				Quantity struct {
					Text string `xml:",chardata"` // 1, 1, 2, 1, 1, 1, 1, 1, 1...
				} `xml:"quantity"`
				Name struct {
					Text string `xml:",chardata"` // creeping spot, attach, ch...
				} `xml:"name"`
				Payment struct {
					Text string `xml:",chardata"` // Money order, Creditcard, ...
				} `xml:"payment"`
				Description struct {
					Chardata string `xml:",chardata"`
					Parlist  struct {
						Text     string `xml:",chardata"`
						Listitem []struct {
							Chardata string `xml:",chardata"`
							Parlist  struct {
								Text     string `xml:",chardata"`
								Listitem []struct {
									Chardata string `xml:",chardata"`
									Text     struct {
										Text    string `xml:",chardata"` // steep britaine prevention...
										Keyword []struct {
											Text string `xml:",chardata"` // lag audrey ding stare sin...
											Bold []struct {
												Text string `xml:",chardata"` // windsor, prizes push jet ...
											} `xml:"bold"`
											Emph []struct {
												Text string `xml:",chardata"` // person whate lily pilgrim...
											} `xml:"emph"`
										} `xml:"keyword"`
										Emph []struct {
											Text    string `xml:",chardata"` // daylight, hereford, there...
											Keyword []struct {
												Text string `xml:",chardata"` // looks trumpets firmly dam...
											} `xml:"keyword"`
											Bold []struct {
												Text string `xml:",chardata"` // return name earth, calms ...
											} `xml:"bold"`
										} `xml:"emph"`
										Bold []struct {
											Text    string `xml:",chardata"` // calchas extend play servi...
											Keyword []struct {
												Text string `xml:",chardata"` // inclining lose revenge ch...
											} `xml:"keyword"`
											Emph []struct {
												Text string `xml:",chardata"` // sir hor shape laboring ta...
											} `xml:"emph"`
										} `xml:"bold"`
									} `xml:"text"`
								} `xml:"listitem"`
							} `xml:"parlist"`
							Text struct {
								Text    string `xml:",chardata"` // ham show andromache madam...
								Keyword []struct {
									Text string `xml:",chardata"` // mouse, greet, character, ...
									Emph []struct {
										Text string `xml:",chardata"` // feature, collected, stand...
									} `xml:"emph"`
									Bold []struct {
										Text string `xml:",chardata"` // thine, bestow extended un...
									} `xml:"bold"`
								} `xml:"keyword"`
								Emph []struct {
									Text string `xml:",chardata"` // tax seriously fat, verifi...
									Bold []struct {
										Text string `xml:",chardata"` // stewardship ulysses where...
									} `xml:"bold"`
									Keyword []struct {
										Text string `xml:",chardata"` // threat conjur costly proc...
									} `xml:"keyword"`
								} `xml:"emph"`
								Bold []struct {
									Text    string `xml:",chardata"` // town seemed, load knotted...
									Keyword []struct {
										Text string `xml:",chardata"` // laughs, pence yea gem tog...
									} `xml:"keyword"`
									Emph []struct {
										Text string `xml:",chardata"` // light mariana horatio gir...
									} `xml:"emph"`
								} `xml:"bold"`
							} `xml:"text"`
						} `xml:"listitem"`
					} `xml:"parlist"`
					Text struct {
						Text    string `xml:",chardata"` // iron notes edmund vanish ...
						Keyword []struct {
							Text string `xml:",chardata"` // henceforward, pen intent,...
							Emph []struct {
								Text string `xml:",chardata"` // maintained methinks aye, ...
							} `xml:"emph"`
							Bold []struct {
								Text string `xml:",chardata"` // abuse dishonour numbers k...
							} `xml:"bold"`
						} `xml:"keyword"`
						Emph []struct {
							Text string `xml:",chardata"` // meditating arise behalf i...
							Bold []struct {
								Text string `xml:",chardata"` // cure infects nobody, ass,...
							} `xml:"bold"`
							Keyword []struct {
								Text string `xml:",chardata"` // vain comparison, scruple ...
							} `xml:"keyword"`
						} `xml:"emph"`
						Bold []struct {
							Text string `xml:",chardata"` // cannon, bathe warranted a...
							Emph []struct {
								Text string `xml:",chardata"` // brutus form ninth garment...
							} `xml:"emph"`
							Keyword []struct {
								Text string `xml:",chardata"` // others uses esteem glouce...
							} `xml:"keyword"`
						} `xml:"bold"`
					} `xml:"text"`
				} `xml:"description"`
				Shipping struct {
					Text string `xml:",chardata"` // Will ship only within cou...
				} `xml:"shipping"`
				Incategory []struct {
					Text     string `xml:",chardata"`
					Category string `xml:"category,attr"`
				} `xml:"incategory"`
				Mailbox struct {
					Text string `xml:",chardata"`
					Mail []struct {
						Chardata string `xml:",chardata"`
						From     struct {
							Text string `xml:",chardata"` // Sujeet Rustenburg mailto:...
						} `xml:"from"`
						To struct {
							Text string `xml:",chardata"` // Sitki Kryukova mailto:Kry...
						} `xml:"to"`
						Date struct {
							Text string `xml:",chardata"` // 11/24/1999, 02/20/2000, 0...
						} `xml:"date"`
						Text struct {
							Text    string `xml:",chardata"` // drift others halberd nigh...
							Keyword []struct {
								Text string `xml:",chardata"` // measure, push souls, gold...
								Bold []struct {
									Text string `xml:",chardata"` // emilia could sea friar om...
								} `xml:"bold"`
								Emph []struct {
									Text string `xml:",chardata"` // temper knock unkind humph...
								} `xml:"emph"`
							} `xml:"keyword"`
							Bold []struct {
								Text    string `xml:",chardata"` // special prompt exterior m...
								Keyword []struct {
									Text string `xml:",chardata"` // lustre creep tremblingly ...
								} `xml:"keyword"`
								Emph []struct {
									Text string `xml:",chardata"` // malady woos, tarry, depar...
								} `xml:"emph"`
							} `xml:"bold"`
							Emph []struct {
								Text    string `xml:",chardata"` // distemper quick tenths ca...
								Keyword []struct {
									Text string `xml:",chardata"` // honours tribune account t...
								} `xml:"keyword"`
								Bold []struct {
									Text string `xml:",chardata"` // flow greasy clamours wide...
								} `xml:"bold"`
							} `xml:"emph"`
						} `xml:"text"`
					} `xml:"mail"`
				} `xml:"mailbox"`
			} `xml:"item"`
		} `xml:"namerica"`
		Samerica struct {
			Text string `xml:",chardata"`
			Item []struct {
				Text     string `xml:",chardata"`
				ID       string `xml:"id,attr"`
				Featured string `xml:"featured,attr"`
				Location struct {
					Text string `xml:",chardata"` // United States, United Sta...
				} `xml:"location"`
				Quantity struct {
					Text string `xml:",chardata"` // 1, 1, 1, 1, 1, 1, 1, 1, 1...
				} `xml:"quantity"`
				Name struct {
					Text string `xml:",chardata"` // greetings, commends, fift...
				} `xml:"name"`
				Payment struct {
					Text string `xml:",chardata"` // Money order, Personal Che...
				} `xml:"payment"`
				Description struct {
					Chardata string `xml:",chardata"`
					Text     struct {
						Text string `xml:",chardata"` // aliena pleases shrunk hon...
						Bold []struct {
							Text    string `xml:",chardata"` // try, princes, beseech dea...
							Keyword []struct {
								Text string `xml:",chardata"` // alight observance, faires...
							} `xml:"keyword"`
							Emph []struct {
								Text string `xml:",chardata"` // consideration, preserver ...
							} `xml:"emph"`
						} `xml:"bold"`
						Emph []struct {
							Text string `xml:",chardata"` // messengers heave, afterwa...
							Bold []struct {
								Text string `xml:",chardata"` // methinks, rode suitors mu...
							} `xml:"bold"`
							Keyword []struct {
								Text string `xml:",chardata"` // rewards apart, prevention...
							} `xml:"keyword"`
						} `xml:"emph"`
						Keyword []struct {
							Text string `xml:",chardata"` // abuses beget reports rose...
							Emph struct {
								Text string `xml:",chardata"` // fear, leg companion, meet...
							} `xml:"emph"`
							Bold struct {
								Text string `xml:",chardata"` // eleven newer thrive, leon...
							} `xml:"bold"`
						} `xml:"keyword"`
					} `xml:"text"`
					Parlist struct {
						Text     string `xml:",chardata"`
						Listitem []struct {
							Chardata string `xml:",chardata"`
							Text     struct {
								Text string `xml:",chardata"` // maine success stains good...
								Emph []struct {
									Text string `xml:",chardata"` // sails shent face, forced ...
									Bold struct {
										Text string `xml:",chardata"` // salt preventions disguise...
									} `xml:"bold"`
									Keyword struct {
										Text string `xml:",chardata"` // banished, safety stray of...
									} `xml:"keyword"`
								} `xml:"emph"`
								Bold []struct {
									Text    string `xml:",chardata"` // swing, may, holds evermor...
									Keyword []struct {
										Text string `xml:",chardata"` // marble stand bold yoursel...
									} `xml:"keyword"`
									Emph []struct {
										Text string `xml:",chardata"` // embodied teach express, d...
									} `xml:"emph"`
								} `xml:"bold"`
								Keyword []struct {
									Text string `xml:",chardata"` // blow villainous mouse our...
									Bold []struct {
										Text string `xml:",chardata"` // poison didst rascally bat...
									} `xml:"bold"`
									Emph []struct {
										Text string `xml:",chardata"` // fifty, malefactions valle...
									} `xml:"emph"`
								} `xml:"keyword"`
							} `xml:"text"`
							Parlist struct {
								Text     string `xml:",chardata"`
								Listitem []struct {
									Chardata string `xml:",chardata"`
									Text     struct {
										Text    string `xml:",chardata"` // froth brown beam commotio...
										Keyword []struct {
											Text string `xml:",chardata"` // behind, foes majestas, pr...
											Bold []struct {
												Text string `xml:",chardata"` // fault scruples amazement ...
											} `xml:"bold"`
											Emph []struct {
												Text string `xml:",chardata"` // minutes, england, resolut...
											} `xml:"emph"`
										} `xml:"keyword"`
										Bold []struct {
											Text string `xml:",chardata"` // troops unfold glad, liken...
											Emph struct {
												Text string `xml:",chardata"` // backward injure, too cate...
											} `xml:"emph"`
											Keyword []struct {
												Text string `xml:",chardata"` // petition shalt sugar, mot...
											} `xml:"keyword"`
										} `xml:"bold"`
										Emph []struct {
											Text    string `xml:",chardata"` // orb pale forward impatien...
											Keyword struct {
												Text string `xml:",chardata"` // complete danger envy main...
											} `xml:"keyword"`
											Bold struct {
												Text string `xml:",chardata"` // bright wert, thought norw...
											} `xml:"bold"`
										} `xml:"emph"`
									} `xml:"text"`
								} `xml:"listitem"`
							} `xml:"parlist"`
						} `xml:"listitem"`
					} `xml:"parlist"`
				} `xml:"description"`
				Shipping struct {
					Text string `xml:",chardata"` // Buyer pays fixed shipping...
				} `xml:"shipping"`
				Incategory []struct {
					Text     string `xml:",chardata"`
					Category string `xml:"category,attr"`
				} `xml:"incategory"`
				Mailbox struct {
					Text string `xml:",chardata"`
					Mail []struct {
						Chardata string `xml:",chardata"`
						From     struct {
							Text string `xml:",chardata"` // Taekyoung Litzler mailto:...
						} `xml:"from"`
						To struct {
							Text string `xml:",chardata"` // Saumya Makinen mailto:Mak...
						} `xml:"to"`
						Date struct {
							Text string `xml:",chardata"` // 04/22/2000, 12/13/2000, 1...
						} `xml:"date"`
						Text struct {
							Text string `xml:",chardata"` // heirs griefs sea antony d...
							Emph []struct {
								Text    string `xml:",chardata"` // packthread outside, three...
								Keyword []struct {
									Text string `xml:",chardata"` // delivered unmoan weariest...
								} `xml:"keyword"`
								Bold []struct {
									Text string `xml:",chardata"` // night drums, money beauti...
								} `xml:"bold"`
							} `xml:"emph"`
							Bold []struct {
								Text string `xml:",chardata"` // gowns implore sullies sem...
								Emph []struct {
									Text string `xml:",chardata"` // armado spends, cheerful w...
								} `xml:"emph"`
								Keyword []struct {
									Text string `xml:",chardata"` // strong witty bitterly sud...
								} `xml:"keyword"`
							} `xml:"bold"`
							Keyword []struct {
								Text string `xml:",chardata"` // suffocate foolish hamlet ...
								Emph []struct {
									Text string `xml:",chardata"` // infectious, silk break wh...
								} `xml:"emph"`
								Bold []struct {
									Text string `xml:",chardata"` // conduct thrown ulysses, p...
								} `xml:"bold"`
							} `xml:"keyword"`
						} `xml:"text"`
					} `xml:"mail"`
				} `xml:"mailbox"`
			} `xml:"item"`
		} `xml:"samerica"`
	} `xml:"regions"`
	Categories struct {
		Text     string `xml:",chardata"`
		Category []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Name struct {
				Text string `xml:",chardata"` // arm solemn weak finds, co...
			} `xml:"name"`
			Description struct {
				Chardata string `xml:",chardata"`
				Text     struct {
					Text    string `xml:",chardata"` // loyalty merit wicked plag...
					Keyword []struct {
						Text string `xml:",chardata"` // flowers, scornful heal le...
						Bold []struct {
							Text string `xml:",chardata"` // peace hymen, mouth spare ...
						} `xml:"bold"`
						Emph []struct {
							Text string `xml:",chardata"` // pol warrant dexterity, ke...
						} `xml:"emph"`
					} `xml:"keyword"`
					Emph []struct {
						Text    string `xml:",chardata"` // infinite peerless very st...
						Keyword []struct {
							Text string `xml:",chardata"` // poise qualifying die hors...
						} `xml:"keyword"`
						Bold []struct {
							Text string `xml:",chardata"` // beyond says clock poisons...
						} `xml:"bold"`
					} `xml:"emph"`
					Bold []struct {
						Text    string `xml:",chardata"` // stupid iras hides despise...
						Keyword []struct {
							Text string `xml:",chardata"` // doleful, rom egypt warble...
						} `xml:"keyword"`
						Emph []struct {
							Text string `xml:",chardata"` // party cried appearing wan...
						} `xml:"emph"`
					} `xml:"bold"`
				} `xml:"text"`
				Parlist struct {
					Text     string `xml:",chardata"`
					Listitem []struct {
						Chardata string `xml:",chardata"`
						Parlist  struct {
							Text     string `xml:",chardata"`
							Listitem []struct {
								Chardata string `xml:",chardata"`
								Text     struct {
									Text    string `xml:",chardata"` // greece griev bids smites ...
									Keyword []struct {
										Text string `xml:",chardata"` // township, maidens rebelli...
										Bold []struct {
											Text string `xml:",chardata"` // pronounc, come, type sly ...
										} `xml:"bold"`
										Emph struct {
											Text string `xml:",chardata"` // notice, vinaigre benedick...
										} `xml:"emph"`
									} `xml:"keyword"`
									Emph []struct {
										Text    string `xml:",chardata"` // green sententious nay pas...
										Keyword []struct {
											Text string `xml:",chardata"` // before, becomes, fools in...
										} `xml:"keyword"`
										Bold []struct {
											Text string `xml:",chardata"` // troths castle watchman mi...
										} `xml:"bold"`
									} `xml:"emph"`
									Bold []struct {
										Text    string `xml:",chardata"` // fantastical writer coward...
										Keyword struct {
											Text string `xml:",chardata"` // planetary, heap spleen do...
										} `xml:"keyword"`
										Emph struct {
											Text string `xml:",chardata"` // city underprop forty odd ...
										} `xml:"emph"`
									} `xml:"bold"`
								} `xml:"text"`
							} `xml:"listitem"`
						} `xml:"parlist"`
						Text struct {
							Text string `xml:",chardata"` // mad cost spirit heavily b...
							Bold []struct {
								Text    string `xml:",chardata"` // florence drawing constanc...
								Keyword []struct {
									Text string `xml:",chardata"` // book transgressing, preci...
								} `xml:"keyword"`
								Emph []struct {
									Text string `xml:",chardata"` // curfew rogue courtly arch...
								} `xml:"emph"`
							} `xml:"bold"`
							Emph []struct {
								Text string `xml:",chardata"` // only, lain worm, going om...
								Bold []struct {
									Text string `xml:",chardata"` // made foining knit inaudib...
								} `xml:"bold"`
								Keyword []struct {
									Text string `xml:",chardata"` // dull mettle apparel east ...
								} `xml:"keyword"`
							} `xml:"emph"`
							Keyword []struct {
								Text string `xml:",chardata"` // herein, defend heard phil...
								Bold []struct {
									Text string `xml:",chardata"` // redeem tape blind royal m...
								} `xml:"bold"`
								Emph struct {
									Text string `xml:",chardata"` // proclamation presages, fi...
								} `xml:"emph"`
							} `xml:"keyword"`
						} `xml:"text"`
					} `xml:"listitem"`
				} `xml:"parlist"`
			} `xml:"description"`
		} `xml:"category"`
	} `xml:"categories"`
	Catgraph struct {
		Text string `xml:",chardata"`
		Edge []struct {
			Text string `xml:",chardata"`
			From string `xml:"from,attr"`
			To   string `xml:"to,attr"`
		} `xml:"edge"`
	} `xml:"catgraph"`
	People struct {
		Text   string `xml:",chardata"`
		Person []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Name struct {
				Text string `xml:",chardata"` // Kasidit Treweek, Harpreet...
			} `xml:"name"`
			Emailaddress struct {
				Text string `xml:",chardata"` // mailto:Treweek@cohera.com...
			} `xml:"emailaddress"`
			Phone struct {
				Text string `xml:",chardata"` // +44 (645) 96317722, +27 (...
			} `xml:"phone"`
			Homepage struct {
				Text string `xml:",chardata"` // http://www.cohera.com/~Tr...
			} `xml:"homepage"`
			Creditcard struct {
				Text string `xml:",chardata"` // 9941 9701 2489 4716, 8299...
			} `xml:"creditcard"`
			Profile struct {
				Text     string `xml:",chardata"`
				Income   string `xml:"income,attr"`
				Interest []struct {
					Text     string `xml:",chardata"`
					Category string `xml:"category,attr"`
				} `xml:"interest"`
				Education struct {
					Text string `xml:",chardata"` // Graduate School, Other, G...
				} `xml:"education"`
				Business struct {
					Text string `xml:",chardata"` // No, Yes, No, Yes, No, Yes...
				} `xml:"business"`
				Gender struct {
					Text string `xml:",chardata"` // male, male, female, femal...
				} `xml:"gender"`
				Age struct {
					Text string `xml:",chardata"` // 18, 18, 37, 29, 18, 41, 1...
				} `xml:"age"`
			} `xml:"profile"`
			Watches struct {
				Text  string `xml:",chardata"`
				Watch []struct {
					Text        string `xml:",chardata"`
					OpenAuction string `xml:"open_auction,attr"`
				} `xml:"watch"`
			} `xml:"watches"`
			Address struct {
				Text   string `xml:",chardata"`
				Street struct {
					Text string `xml:",chardata"` // 99 Lazaro St, 18 Ali St, ...
				} `xml:"street"`
				City struct {
					Text string `xml:",chardata"` // El, Cody, Fortaleza, Pana...
				} `xml:"city"`
				Country struct {
					Text string `xml:",chardata"` // Svalbard and Jan Mayen Is...
				} `xml:"country"`
				Province struct {
					Text string `xml:",chardata"` // Basler, Qizhong, Pavani, ...
				} `xml:"province"`
				Zipcode struct {
					Text string `xml:",chardata"` // 27, 15, 13, 12, 38, 26, 8...
				} `xml:"zipcode"`
			} `xml:"address"`
		} `xml:"person"`
	} `xml:"people"`
	OpenAuctions struct {
		Text        string `xml:",chardata"`
		OpenAuction []struct {
			Text    string `xml:",chardata"`
			ID      string `xml:"id,attr"`
			Initial struct {
				Text string `xml:",chardata"` // 168.02, 16.92, 86.67, 36....
			} `xml:"initial"`
			Reserve struct {
				Text string `xml:",chardata"` // 2410.24, 61.41, 85.79, 57...
			} `xml:"reserve"`
			Bidder []struct {
				Text string `xml:",chardata"`
				Date struct {
					Text string `xml:",chardata"` // 05/04/1999, 08/12/2000, 0...
				} `xml:"date"`
				Time struct {
					Text string `xml:",chardata"` // 14:15:28, 18:03:05, 23:55...
				} `xml:"time"`
				Personref struct {
					Text   string `xml:",chardata"`
					Person string `xml:"person,attr"`
				} `xml:"personref"`
				Increase struct {
					Text string `xml:",chardata"` // 7.50, 12.00, 1.50, 1.50, ...
				} `xml:"increase"`
			} `xml:"bidder"`
			Current struct {
				Text string `xml:",chardata"` // 175.52, 16.92, 100.17, 12...
			} `xml:"current"`
			Itemref struct {
				Text string `xml:",chardata"`
				Item string `xml:"item,attr"`
			} `xml:"itemref"`
			Seller struct {
				Text   string `xml:",chardata"`
				Person string `xml:"person,attr"`
			} `xml:"seller"`
			Annotation struct {
				Text   string `xml:",chardata"`
				Author struct {
					Text   string `xml:",chardata"`
					Person string `xml:"person,attr"`
				} `xml:"author"`
				Description struct {
					Chardata string `xml:",chardata"`
					Text     struct {
						Text string `xml:",chardata"` // find rays assemble loaths...
						Emph []struct {
							Text    string `xml:",chardata"` // welcome shame burdens fad...
							Keyword []struct {
								Text string `xml:",chardata"` // resolv gaging, sold othel...
							} `xml:"keyword"`
							Bold []struct {
								Text string `xml:",chardata"` // aquitaine fifty commit sa...
							} `xml:"bold"`
						} `xml:"emph"`
						Keyword []struct {
							Text string `xml:",chardata"` // told gnat visible wouldst...
							Emph []struct {
								Text string `xml:",chardata"` // clap, crave oswald, brain...
							} `xml:"emph"`
							Bold []struct {
								Text string `xml:",chardata"` // match triumphant philip t...
							} `xml:"bold"`
						} `xml:"keyword"`
						Bold []struct {
							Text string `xml:",chardata"` // presume newness putting p...
							Emph []struct {
								Text string `xml:",chardata"` // gawds absent, glou forbid...
							} `xml:"emph"`
							Keyword []struct {
								Text string `xml:",chardata"` // forbid, pierce commander ...
							} `xml:"keyword"`
						} `xml:"bold"`
					} `xml:"text"`
					Parlist struct {
						Text     string `xml:",chardata"`
						Listitem []struct {
							Chardata string `xml:",chardata"`
							Text     struct {
								Text string `xml:",chardata"` // dainty unworthy right sti...
								Emph []struct {
									Text string `xml:",chardata"` // study bad ruder denial ro...
									Bold []struct {
										Text string `xml:",chardata"` // just offenders kingdom sa...
									} `xml:"bold"`
									Keyword []struct {
										Text string `xml:",chardata"` // tonight foam lisp more un...
									} `xml:"keyword"`
								} `xml:"emph"`
								Bold []struct {
									Text string `xml:",chardata"` // consent heel bad requite ...
									Emph []struct {
										Text string `xml:",chardata"` // worse hap lays first comm...
									} `xml:"emph"`
									Keyword []struct {
										Text string `xml:",chardata"` // signify reasonable denies...
									} `xml:"keyword"`
								} `xml:"bold"`
								Keyword []struct {
									Text string `xml:",chardata"` // return meant, rush slash,...
									Emph []struct {
										Text string `xml:",chardata"` // ulysses barefoot, protest...
									} `xml:"emph"`
									Bold []struct {
										Text string `xml:",chardata"` // forbear, eyes capacity mu...
									} `xml:"bold"`
								} `xml:"keyword"`
							} `xml:"text"`
							Parlist struct {
								Text     string `xml:",chardata"`
								Listitem []struct {
									Chardata string `xml:",chardata"`
									Text     struct {
										Text string `xml:",chardata"` // foot length confess monta...
										Emph []struct {
											Text string `xml:",chardata"` // falling sign elder cliffo...
											Bold []struct {
												Text string `xml:",chardata"` // ground stand disasters he...
											} `xml:"bold"`
											Keyword []struct {
												Text string `xml:",chardata"` // yields wits thrown, scum ...
											} `xml:"keyword"`
										} `xml:"emph"`
										Keyword []struct {
											Text string `xml:",chardata"` // epicure rich throwing vis...
											Bold []struct {
												Text string `xml:",chardata"` // feast garment famish fast...
											} `xml:"bold"`
											Emph []struct {
												Text string `xml:",chardata"` // encourage, weep engine, h...
											} `xml:"emph"`
										} `xml:"keyword"`
										Bold []struct {
											Text    string `xml:",chardata"` // winds, age robed tie repe...
											Keyword []struct {
												Text string `xml:",chardata"` // host, shamefully, trust a...
											} `xml:"keyword"`
											Emph []struct {
												Text string `xml:",chardata"` // presume musing giddy law ...
											} `xml:"emph"`
										} `xml:"bold"`
									} `xml:"text"`
								} `xml:"listitem"`
							} `xml:"parlist"`
						} `xml:"listitem"`
					} `xml:"parlist"`
				} `xml:"description"`
				Happiness struct {
					Text string `xml:",chardata"` // 4, 8, 3, 7, 8, 8, 3, 7, 9...
				} `xml:"happiness"`
			} `xml:"annotation"`
			Quantity struct {
				Text string `xml:",chardata"` // 1, 1, 1, 1, 1, 1, 1, 1, 1...
			} `xml:"quantity"`
			Type struct {
				Text string `xml:",chardata"` // Regular, Featured, Featur...
			} `xml:"type"`
			Interval struct {
				Text  string `xml:",chardata"`
				Start struct {
					Text string `xml:",chardata"` // 01/03/1998, 08/19/1999, 0...
				} `xml:"start"`
				End struct {
					Text string `xml:",chardata"` // 02/11/2000, 06/07/1999, 0...
				} `xml:"end"`
			} `xml:"interval"`
			Privacy struct {
				Text string `xml:",chardata"` // Yes, Yes, Yes, No, Yes, Y...
			} `xml:"privacy"`
		} `xml:"open_auction"`
	} `xml:"open_auctions"`
	ClosedAuctions struct {
		Text          string `xml:",chardata"`
		ClosedAuction []struct {
			Text   string `xml:",chardata"`
			Seller struct {
				Text   string `xml:",chardata"`
				Person string `xml:"person,attr"`
			} `xml:"seller"`
			Buyer struct {
				Text   string `xml:",chardata"`
				Person string `xml:"person,attr"`
			} `xml:"buyer"`
			Itemref struct {
				Text string `xml:",chardata"`
				Item string `xml:"item,attr"`
			} `xml:"itemref"`
			Price struct {
				Text string `xml:",chardata"` // 50.03, 42.56, 4.94, 52.15...
			} `xml:"price"`
			Date struct {
				Text string `xml:",chardata"` // 12/17/2000, 05/11/1999, 0...
			} `xml:"date"`
			Quantity struct {
				Text string `xml:",chardata"` // 1, 1, 1, 1, 1, 1, 1, 2, 1...
			} `xml:"quantity"`
			Type struct {
				Text string `xml:",chardata"` // Featured, Regular, Regula...
			} `xml:"type"`
			Annotation struct {
				Text   string `xml:",chardata"`
				Author struct {
					Text   string `xml:",chardata"`
					Person string `xml:"person,attr"`
				} `xml:"author"`
				Description struct {
					Chardata string `xml:",chardata"`
					Text     struct {
						Text    string `xml:",chardata"` // whereon writes university...
						Keyword []struct {
							Text string `xml:",chardata"` // allegiance letter stoop, ...
							Emph []struct {
								Text string `xml:",chardata"` // passes, pay moment shoote...
							} `xml:"emph"`
							Bold []struct {
								Text string `xml:",chardata"` // equally excursions, sacre...
							} `xml:"bold"`
						} `xml:"keyword"`
						Emph []struct {
							Text    string `xml:",chardata"` // brawl, pluck fleet, gaunt...
							Keyword []struct {
								Text string `xml:",chardata"` // alexas madness countrymen...
							} `xml:"keyword"`
							Bold []struct {
								Text string `xml:",chardata"` // afeard, helen sold shallo...
							} `xml:"bold"`
						} `xml:"emph"`
						Bold []struct {
							Text    string `xml:",chardata"` // deal heat comes jupiter, ...
							Keyword []struct {
								Text string `xml:",chardata"` // afresh, forgive whore, bo...
							} `xml:"keyword"`
							Emph []struct {
								Text string `xml:",chardata"` // trade pleas milk, curfew ...
							} `xml:"emph"`
						} `xml:"bold"`
					} `xml:"text"`
					Parlist struct {
						Text     string `xml:",chardata"`
						Listitem []struct {
							Chardata string `xml:",chardata"`
							Text     struct {
								Text string `xml:",chardata"` // clouds dry allhallowmas a...
								Emph []struct {
									Text string `xml:",chardata"` // rests, wept, solicit cock...
									Bold []struct {
										Text string `xml:",chardata"` // health imaginary height c...
									} `xml:"bold"`
									Keyword []struct {
										Text string `xml:",chardata"` // whoever load drop, preven...
									} `xml:"keyword"`
								} `xml:"emph"`
								Keyword []struct {
									Text string `xml:",chardata"` // players creditors tent, f...
									Emph []struct {
										Text string `xml:",chardata"` // behalf suspend crystal al...
									} `xml:"emph"`
									Bold []struct {
										Text string `xml:",chardata"` // worthy, mouth spare laid ...
									} `xml:"bold"`
								} `xml:"keyword"`
								Bold []struct {
									Text    string `xml:",chardata"` // glanced infected honest s...
									Keyword []struct {
										Text string `xml:",chardata"` // petter boist viewed that ...
									} `xml:"keyword"`
									Emph []struct {
										Text string `xml:",chardata"` // watch joys lash, wonder, ...
									} `xml:"emph"`
								} `xml:"bold"`
							} `xml:"text"`
							Parlist struct {
								Text     string `xml:",chardata"`
								Listitem []struct {
									Chardata string `xml:",chardata"`
									Text     struct {
										Text    string `xml:",chardata"` // commanders merrily politi...
										Keyword []struct {
											Text string `xml:",chardata"` // expectation soul happier ...
											Bold []struct {
												Text string `xml:",chardata"` // despite flourish musical ...
											} `xml:"bold"`
											Emph []struct {
												Text string `xml:",chardata"` // crack hourly which born d...
											} `xml:"emph"`
										} `xml:"keyword"`
										Emph []struct {
											Text string `xml:",chardata"` // millions, too, spell air,...
											Bold []struct {
												Text string `xml:",chardata"` // riddle wander chin poloni...
											} `xml:"bold"`
											Keyword []struct {
												Text string `xml:",chardata"` // west avail griefs took al...
											} `xml:"keyword"`
										} `xml:"emph"`
										Bold []struct {
											Text string `xml:",chardata"` // garland comfort eat think...
											Emph []struct {
												Text string `xml:",chardata"` // priest hill, forester utt...
											} `xml:"emph"`
											Keyword []struct {
												Text string `xml:",chardata"` // subornation, scratch trav...
											} `xml:"keyword"`
										} `xml:"bold"`
									} `xml:"text"`
								} `xml:"listitem"`
							} `xml:"parlist"`
						} `xml:"listitem"`
					} `xml:"parlist"`
				} `xml:"description"`
				Happiness struct {
					Text string `xml:",chardata"` // 5, 6, 8, 9, 9, 7, 5, 2, 1...
				} `xml:"happiness"`
			} `xml:"annotation"`
		} `xml:"closed_auction"`
	} `xml:"closed_auctions"`
}
