package main

import "encoding/xml"

// KOMBI was generated 2017-12-05 17:42:05 by tir on apollo.
type KOMBI struct {
	XMLName xml.Name `xml:"KOMBI"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	SDG     struct {
		Text  string `xml:",chardata"`
		DDKEY struct {
			Text string `xml:",chardata"`
		} `xml:"D_DKEY"`
		DSKEY struct {
			Text string `xml:",chardata"`
		} `xml:"D_SKEY"`
		ARRFA struct {
			Text string `xml:",chardata"`
			ABBR struct {
				Text string `xml:",chardata"`
			} `xml:"ABBR"`
		} `xml:"AR_RFA"`
		PNR struct {
			Text string `xml:",chardata"`
		} `xml:"PNR"`
		POOL struct {
			Text    string `xml:",chardata"`
			Abbrev  string `xml:"abbrev,attr"`
			Code    string `xml:"code,attr"`
			STATION struct {
				Text string `xml:",chardata"`
			} `xml:"STATION"`
		} `xml:"POOL"`
		PE struct {
			Text string `xml:",chardata"`
			ABBR struct {
				Text string `xml:",chardata"`
			} `xml:"ABBR"`
		} `xml:"PE"`
		BSTAT struct {
			Text string `xml:",chardata"`
			ABBR struct {
				Text string `xml:",chardata"`
			} `xml:"ABBR"`
		} `xml:"BSTAT"`
		KST struct {
			Text string `xml:",chardata"`
		} `xml:"KST"`
		KASS struct {
			Text string `xml:",chardata"`
		} `xml:"KASS"`
		DEEDAT struct {
			Text string `xml:",chardata"`
		} `xml:"DE_EDAT"`
		DEKDAT struct {
			Text string `xml:",chardata"`
		} `xml:"DE_KDAT"`
		ERF struct {
			Text string `xml:",chardata"`
		} `xml:"ERF"`
		TIT struct {
			Text string `xml:",chardata"`
			SHTI []struct {
				Text   string `xml:",chardata"`
				Descr  string `xml:"descr,attr"`
				TITTXT struct {
					Text string `xml:",chardata"`
				} `xml:"TIT_TXT"`
				TITANM struct {
					Text string `xml:",chardata"`
				} `xml:"TIT_ANM"`
				TITTEIL struct {
					Text string `xml:",chardata"`
				} `xml:"TIT_TEIL"`
			} `xml:"SHTI"`
			SRTI struct {
				Text   string `xml:",chardata"`
				Descr  string `xml:"descr,attr"`
				TITTXT struct {
					Text string `xml:",chardata"`
				} `xml:"TIT_TXT"`
				TITTEIL struct {
					Text string `xml:",chardata"`
				} `xml:"TIT_TEIL"`
			} `xml:"SRTI"`
			UNTI struct {
				Text   string `xml:",chardata"`
				Descr  string `xml:"descr,attr"`
				TITTXT struct {
					Text string `xml:",chardata"`
				} `xml:"TIT_TXT"`
			} `xml:"UNTI"`
			SUTI struct {
				Text   string `xml:",chardata"`
				Descr  string `xml:"descr,attr"`
				TITTXT struct {
					Text string `xml:",chardata"`
				} `xml:"TIT_TXT"`
			} `xml:"SUTI"`
		} `xml:"TIT"`
		UPM struct {
			Text string `xml:",chardata"`
			SSV  []struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
				VNAME struct {
					Text string `xml:",chardata"`
				} `xml:"VNAME"`
				NNAME struct {
					Text string `xml:",chardata"`
				} `xml:"NNAME"`
			} `xml:"SSV"`
			MOD []struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
				VNAME struct {
					Text string `xml:",chardata"`
				} `xml:"VNAME"`
				NNAME struct {
					Text string `xml:",chardata"`
				} `xml:"NNAME"`
				UFZUS struct {
					Text string `xml:",chardata"`
				} `xml:"U_FZUS"`
			} `xml:"MOD"`
			RDN []struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
				ORGA  struct {
					Text string `xml:",chardata"`
				} `xml:"ORGA"`
			} `xml:"RDN"`
			VOK []struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
				VNAME struct {
					Text string `xml:",chardata"`
				} `xml:"VNAME"`
				NNAME struct {
					Text string `xml:",chardata"`
				} `xml:"NNAME"`
			} `xml:"VOK"`
			REG []struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
				VNAME struct {
					Text string `xml:",chardata"`
				} `xml:"VNAME"`
				NNAME struct {
					Text string `xml:",chardata"`
				} `xml:"NNAME"`
			} `xml:"REG"`
			SSP []struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
				VNAME struct {
					Text string `xml:",chardata"`
				} `xml:"VNAME"`
				NNAME struct {
					Text string `xml:",chardata"`
				} `xml:"NNAME"`
			} `xml:"SSP"`
			AUB struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
				VNAME struct {
					Text string `xml:",chardata"`
				} `xml:"VNAME"`
				NNAME struct {
					Text string `xml:",chardata"`
				} `xml:"NNAME"`
			} `xml:"AUB"`
			ENSMV struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
				ORGA  struct {
					Text string `xml:",chardata"`
				} `xml:"ORGA"`
			} `xml:"ENSMV"`
			CUT []struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
				VNAME struct {
					Text string `xml:",chardata"`
				} `xml:"VNAME"`
				NNAME struct {
					Text string `xml:",chardata"`
				} `xml:"NNAME"`
			} `xml:"CUT"`
			KAM []struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
				VNAME struct {
					Text string `xml:",chardata"`
				} `xml:"VNAME"`
				NNAME struct {
					Text string `xml:",chardata"`
				} `xml:"NNAME"`
			} `xml:"KAM"`
			AUT []struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
				VNAME struct {
					Text string `xml:",chardata"`
				} `xml:"VNAME"`
				NNAME struct {
					Text string `xml:",chardata"`
				} `xml:"NNAME"`
				UFZUS struct {
					Text string `xml:",chardata"`
				} `xml:"U_FZUS"`
			} `xml:"AUT"`
			RED struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
				VNAME struct {
					Text string `xml:",chardata"`
				} `xml:"VNAME"`
				NNAME struct {
					Text string `xml:",chardata"`
				} `xml:"NNAME"`
			} `xml:"RED"`
			ENS []struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
				ORGA  struct {
					Text string `xml:",chardata"`
				} `xml:"ORGA"`
			} `xml:"ENS"`
			SPR []struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
				VNAME struct {
					Text string `xml:",chardata"`
				} `xml:"VNAME"`
				NNAME struct {
					Text string `xml:",chardata"`
				} `xml:"NNAME"`
				UFZUS struct {
					Text string `xml:",chardata"`
				} `xml:"U_FZUS"`
			} `xml:"SPR"`
			INT struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
				VNAME struct {
					Text string `xml:",chardata"`
				} `xml:"VNAME"`
				NNAME struct {
					Text string `xml:",chardata"`
				} `xml:"NNAME"`
			} `xml:"INT"`
			BER struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
				VNAME struct {
					Text string `xml:",chardata"`
				} `xml:"VNAME"`
				NNAME struct {
					Text string `xml:",chardata"`
				} `xml:"NNAME"`
			} `xml:"BER"`
			MITS []struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
				VNAME struct {
					Text string `xml:",chardata"`
				} `xml:"VNAME"`
				NNAME struct {
					Text string `xml:",chardata"`
				} `xml:"NNAME"`
				UFZUS struct {
					Text string `xml:",chardata"`
				} `xml:"U_FZUS"`
			} `xml:"MITS"`
			KAB []struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
				VNAME struct {
					Text string `xml:",chardata"`
				} `xml:"VNAME"`
				NNAME struct {
					Text string `xml:",chardata"`
				} `xml:"NNAME"`
			} `xml:"KAB"`
			DIL struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
				VNAME struct {
					Text string `xml:",chardata"`
				} `xml:"VNAME"`
				NNAME struct {
					Text string `xml:",chardata"`
				} `xml:"NNAME"`
			} `xml:"DIL"`
		} `xml:"UPM"`
		SDA struct {
			Text string `xml:",chardata"`
			ESD  []struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
				SRFA  struct {
					Text string `xml:",chardata"`
					ABBR struct {
						Text string `xml:",chardata"`
					} `xml:"ABBR"`
				} `xml:"S_RFA"`
				SDAT struct {
					Text string `xml:",chardata"`
				} `xml:"SDAT"`
				SDAU struct {
					Text string `xml:",chardata"`
				} `xml:"SDAU"`
				SKST struct {
					Text string `xml:",chardata"`
				} `xml:"S_KST"`
				SEDAT struct {
					Text string `xml:",chardata"`
				} `xml:"S_EDAT"`
				SKDAT struct {
					Text string `xml:",chardata"`
				} `xml:"S_KDAT"`
				VG struct {
					Text string `xml:",chardata"`
					ABBR struct {
						Text string `xml:",chardata"`
					} `xml:"ABBR"`
				} `xml:"VG"`
				SPNR struct {
					Text string `xml:",chardata"`
				} `xml:"S_PNR"`
				SBEG struct {
					Text string `xml:",chardata"`
				} `xml:"SBEG"`
			} `xml:"ESD"`
			WSD []struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
				SRFA  struct {
					Text string `xml:",chardata"`
					ABBR struct {
						Text string `xml:",chardata"`
					} `xml:"ABBR"`
				} `xml:"S_RFA"`
				VG struct {
					Text string `xml:",chardata"`
					ABBR struct {
						Text string `xml:",chardata"`
					} `xml:"ABBR"`
				} `xml:"VG"`
				SDAT struct {
					Text string `xml:",chardata"`
				} `xml:"SDAT"`
				SDAU struct {
					Text string `xml:",chardata"`
				} `xml:"SDAU"`
				SEDAT struct {
					Text string `xml:",chardata"`
				} `xml:"S_EDAT"`
				SKDAT struct {
					Text string `xml:",chardata"`
				} `xml:"S_KDAT"`
				SBEG struct {
					Text string `xml:",chardata"`
				} `xml:"SBEG"`
				UPM struct {
					Text string `xml:",chardata"`
					RED  struct {
						Text  string `xml:",chardata"`
						Descr string `xml:"descr,attr"`
						VNAME struct {
							Text string `xml:",chardata"`
						} `xml:"VNAME"`
						NNAME struct {
							Text string `xml:",chardata"`
						} `xml:"NNAME"`
						UFZUS struct {
							Text string `xml:",chardata"`
						} `xml:"U_FZUS"`
					} `xml:"RED"`
					RDN struct {
						Text  string `xml:",chardata"`
						Descr string `xml:"descr,attr"`
						ORGA  struct {
							Text string `xml:",chardata"`
						} `xml:"ORGA"`
					} `xml:"RDN"`
				} `xml:"UPM"`
				SKST struct {
					Text string `xml:",chardata"`
				} `xml:"S_KST"`
				SPNR struct {
					Text string `xml:",chardata"`
				} `xml:"S_PNR"`
				WHTI struct {
					Text string `xml:",chardata"`
				} `xml:"WHTI"`
			} `xml:"WSD"`
			APDB struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
				SRFA  struct {
					Text string `xml:",chardata"`
					ABBR struct {
						Text string `xml:",chardata"`
					} `xml:"ABBR"`
				} `xml:"S_RFA"`
				SDAT struct {
					Text string `xml:",chardata"`
				} `xml:"SDAT"`
				SEDAT struct {
					Text string `xml:",chardata"`
				} `xml:"S_EDAT"`
				SKDAT struct {
					Text string `xml:",chardata"`
				} `xml:"S_KDAT"`
				VG struct {
					Text string `xml:",chardata"`
					ABBR struct {
						Text string `xml:",chardata"`
					} `xml:"ABBR"`
				} `xml:"VG"`
				SDAU struct {
					Text string `xml:",chardata"`
				} `xml:"SDAU"`
			} `xml:"APDB"`
			ADAT struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
				SRFA  struct {
					Text string `xml:",chardata"`
					ABBR struct {
						Text string `xml:",chardata"`
					} `xml:"ABBR"`
				} `xml:"S_RFA"`
				VG struct {
					Text string `xml:",chardata"`
					ABBR struct {
						Text string `xml:",chardata"`
					} `xml:"ABBR"`
				} `xml:"VG"`
				SDAT struct {
					Text string `xml:",chardata"`
				} `xml:"SDAT"`
				SDAU struct {
					Text string `xml:",chardata"`
				} `xml:"SDAU"`
				SEDAT struct {
					Text string `xml:",chardata"`
				} `xml:"S_EDAT"`
				SKDAT struct {
					Text string `xml:",chardata"`
				} `xml:"S_KDAT"`
			} `xml:"ADAT"`
		} `xml:"SDA"`
		HD struct {
			Text string `xml:",chardata"`
			STD  []struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
			} `xml:"STD"`
			SACH []struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
			} `xml:"SACH"`
			GEO []struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
			} `xml:"GEO"`
			PERS []struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
			} `xml:"PERS"`
			INST []struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
			} `xml:"INST"`
			WERK struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
			} `xml:"WERK"`
		} `xml:"HD"`
		INH struct {
			Text string `xml:",chardata"`
			T    struct {
				Text   string `xml:",chardata"`
				Descr  string `xml:"descr,attr"`
				Format string `xml:"format,attr"`
			} `xml:"T"`
			P struct {
				Text   string `xml:",chardata"`
				Descr  string `xml:"descr,attr"`
				Format string `xml:"format,attr"`
			} `xml:"P"`
		} `xml:"INH"`
		BTT []struct {
			Text             string `xml:",chardata"`
			Digi             string `xml:"digi,attr"`
			DigiAbbr         string `xml:"digi_abbr,attr"`
			DigiDescr        string `xml:"digi_descr,attr"`
			HasKeyframes     string `xml:"hasKeyframes,attr"`
			IsDefaultPreview string `xml:"isDefaultPreview,attr"`
			Key              string `xml:"key,attr"`
			LoanValue        string `xml:"loanValue,attr"`
			Manzeige         string `xml:"manzeige,attr"`
			Mauslinfo        string `xml:"mauslinfo,attr"`
			Mpreview         string `xml:"mpreview,attr"`
			Mvormerk         string `xml:"mvormerk,attr"`
			Station          string `xml:"station,attr"`
			BERFA            struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BE_RFA"`
			ANK struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"ANK"`
			MANR struct {
				Text string `xml:",chardata"`
			} `xml:"M_ANR"`
			BEH struct {
				Text string `xml:",chardata"`
			} `xml:"BEH"`
			ROL struct {
				Text string `xml:",chardata"`
			} `xml:"ROL"`
			LGO struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
				SUPP struct {
					Text string `xml:",chardata"`
				} `xml:"SUPP"`
			} `xml:"LGO"`
			BTTSPR struct {
				Text string `xml:",chardata"`
				STAT struct {
					Text string `xml:",chardata"`
				} `xml:"STAT"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BTT_SPR"`
			BTTDS struct {
				Text string `xml:",chardata"`
			} `xml:"BTT_DS"`
			BTTPFR struct {
				Text string `xml:",chardata"`
			} `xml:"BTT_PFR"`
			BTTDAU struct {
				Text string `xml:",chardata"`
			} `xml:"BTT_DAU"`
			ARDAT struct {
				Text string `xml:",chardata"`
			} `xml:"ARDAT"`
			BTTEDAT struct {
				Text string `xml:",chardata"`
			} `xml:"BTT_EDAT"`
			BTTTA struct {
				Text string `xml:",chardata"`
				STAT struct {
					Text string `xml:",chardata"`
				} `xml:"STAT"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
				TEXT struct {
					Text string `xml:",chardata"`
				} `xml:"TEXT"`
				DAT struct {
					Text string `xml:",chardata"`
				} `xml:"DAT"`
			} `xml:"BTT_TA"`
			BTTPOS struct {
				Text string `xml:",chardata"`
			} `xml:"BTT_POS"`
			BTTTC struct {
				Text     string `xml:",chardata"`
				Accuracy string `xml:"accuracy,attr"`
			} `xml:"BTT_TC"`
			BTTLZ struct {
				Text string `xml:",chardata"`
			} `xml:"BTT_LZ"`
			BTTBPR struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BTT_BPR"`
			AZDAT struct {
				Text string `xml:",chardata"`
			} `xml:"AZDAT"`
			BTTINQU struct {
				Text string `xml:",chardata"`
			} `xml:"BTT_INQU"`
			BTTID struct {
				Text  string `xml:",chardata"`
				IDBCD struct {
					Text  string `xml:",chardata"`
					Descr string `xml:"descr,attr"`
				} `xml:"ID_BCD"`
				IDESS struct {
					Text  string `xml:",chardata"`
					Descr string `xml:"descr,attr"`
				} `xml:"ID_ESS"`
			} `xml:"BTT_ID"`
			BTTCONF struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BTT_CONF"`
			BTTRV struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BTT_RV"`
			BTTRVH struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BTT_RVH"`
			BTTVAR struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BTT_VAR"`
			BTTVQ struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BTT_VQ"`
			BTTAAF struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BTT_AAF"`
			BTTAQ struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BTT_AQ"`
			BTTFSS struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BTT_FSS"`
			TSP []struct {
				Text   string `xml:",chardata"`
				TRACK  string `xml:"TRACK,attr"`
				TSWERT struct {
					Text string `xml:",chardata"`
					ABBR struct {
						Text string `xml:",chardata"`
					} `xml:"ABBR"`
				} `xml:"TS_WERT"`
				TSRDA struct {
					Text string `xml:",chardata"`
					ABBR struct {
						Text string `xml:",chardata"`
					} `xml:"ABBR"`
				} `xml:"TS_RDA"`
				TSRUN struct {
					Text string `xml:",chardata"`
					ABBR struct {
						Text string `xml:",chardata"`
					} `xml:"ABBR"`
				} `xml:"TS_RUN"`
				TSANM struct {
					Text string `xml:",chardata"`
				} `xml:"TS_ANM"`
			} `xml:"TSP"`
			RNG struct {
				Text string `xml:",chardata"`
			} `xml:"RNG"`
			BTTLOE struct {
				Text string `xml:",chardata"`
			} `xml:"BTT_LOE"`
			BTTSYS struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BTT_SYS"`
			BTTFAB struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BTT_FAB"`
			BTTKDAT struct {
				Text string `xml:",chardata"`
			} `xml:"BTT_KDAT"`
			BTTMANM struct {
				Text string `xml:",chardata"`
			} `xml:"BTT_MANM"`
			BTTFRB struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BTT_FRB"`
			BTTRST struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BTT_RST"`
			BTTPST struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BTT_PST"`
			BTTKAP struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BTT_KAP"`
			BTTFTN struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BTT_FTN"`
			BTTHVM struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BTT_HVM"`
			BTTANM struct {
				Text string `xml:",chardata"`
			} `xml:"BTT_ANM"`
			BTTBAK struct {
				Text string `xml:",chardata"`
			} `xml:"BTT_BAK"`
			BTTFRNR struct {
				Text string `xml:",chardata"`
			} `xml:"BTT_FRNR"`
			BTTPTL struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BTT_PTL"`
			BTTTIT struct {
				Text string `xml:",chardata"`
			} `xml:"BTT_TIT"`
			BTTFR struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BTT_FR"`
			BTTRDS struct {
				Text string `xml:",chardata"`
				Abbr struct {
					Text string `xml:",chardata"`
				} `xml:"abbr"`
			} `xml:"BTT_RDS"`
			BTTZSD struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BTT_ZSD"`
			BTTDR struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BTT_DR"`
		} `xml:"BTT"`
		SEQ []struct {
			Text  string `xml:",chardata"`
			ZUORD []struct {
				Text  string `xml:",chardata"`
				RTAKE struct {
					Text string `xml:",chardata"`
				} `xml:"R_TAKE"`
				RIDX struct {
					Text string `xml:",chardata"`
				} `xml:"R_IDX"`
				RTC struct {
					Text     string `xml:",chardata"`
					Accuracy string `xml:"accuracy,attr"`
				} `xml:"R_TC"`
				RDAU struct {
					Text string `xml:",chardata"`
				} `xml:"R_DAU"`
				EDOK struct {
					Text  string `xml:",chardata"`
					RDKEY struct {
						Text string `xml:",chardata"`
					} `xml:"R_DKEY"`
					RSKEY struct {
						Text string `xml:",chardata"`
					} `xml:"R_SKEY"`
					RTIT struct {
						Text string `xml:",chardata"`
					} `xml:"R_TIT"`
					RTAKE struct {
						Text string `xml:",chardata"`
					} `xml:"R_TAKE"`
					RTC struct {
						Text     string `xml:",chardata"`
						Accuracy string `xml:"accuracy,attr"`
					} `xml:"R_TC"`
					RDAU struct {
						Text string `xml:",chardata"`
					} `xml:"R_DAU"`
				} `xml:"EDOK"`
			} `xml:"ZUORD"`
			DDKEY struct {
				Text string `xml:",chardata"`
			} `xml:"D_DKEY"`
			DSKEY struct {
				Text string `xml:",chardata"`
			} `xml:"D_SKEY"`
			ARRFA struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"AR_RFA"`
			POOL struct {
				Text    string `xml:",chardata"`
				Abbrev  string `xml:"abbrev,attr"`
				Code    string `xml:"code,attr"`
				STATION struct {
					Text string `xml:",chardata"`
				} `xml:"STATION"`
			} `xml:"POOL"`
			KASS struct {
				Text string `xml:",chardata"`
			} `xml:"KASS"`
			DEEDAT struct {
				Text string `xml:",chardata"`
			} `xml:"DE_EDAT"`
			ERF struct {
				Text string `xml:",chardata"`
			} `xml:"ERF"`
			INH struct {
				Text string `xml:",chardata"`
				B    struct {
					Text   string `xml:",chardata"`
					Descr  string `xml:"descr,attr"`
					Format string `xml:"format,attr"`
				} `xml:"B"`
				O struct {
					Text   string `xml:",chardata"`
					Descr  string `xml:"descr,attr"`
					Format string `xml:"format,attr"`
				} `xml:"O"`
			} `xml:"INH"`
			DV struct {
				Text string `xml:",chardata"`
			} `xml:"DV"`
			DEKDAT struct {
				Text string `xml:",chardata"`
			} `xml:"DE_KDAT"`
			DRA struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"DRA"`
		} `xml:"SEQ"`
		DV struct {
			Text string `xml:",chardata"`
		} `xml:"DV"`
		KAT struct {
			Text string `xml:",chardata"`
			KAP  []struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
			} `xml:"KAP"`
			KGG struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
			} `xml:"KGG"`
			KAI []struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
			} `xml:"KAI"`
			SPA []struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
			} `xml:"SPA"`
			KAZ struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
			} `xml:"KAZ"`
		} `xml:"KAT"`
		DEID struct {
			Text  string `xml:",chardata"`
			IDLIC struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
			} `xml:"ID_LIC"`
		} `xml:"DE_ID"`
		ANM struct {
			Text  string `xml:",chardata"`
			APABR struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
			} `xml:"A_PABR"`
		} `xml:"ANM"`
		PORT struct {
			Text string `xml:",chardata"`
		} `xml:"PORT"`
		ZGRB struct {
			Text string `xml:",chardata"`
			SP   struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
				RTEXT struct {
					Text string `xml:",chardata"`
				} `xml:"R_TEXT"`
			} `xml:"SP"`
			VWB struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
				RTEXT struct {
					Text string `xml:",chardata"`
				} `xml:"R_TEXT"`
				RVDAT struct {
					Text string `xml:",chardata"`
				} `xml:"R_VDAT"`
			} `xml:"VWB"`
			RH struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
				RTEXT struct {
					Text string `xml:",chardata"`
				} `xml:"R_TEXT"`
			} `xml:"RH"`
		} `xml:"ZGRB"`
		VERW struct {
			Text string `xml:",chardata"`
			BMAT struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
			} `xml:"BMAT"`
		} `xml:"VERW"`
		DRA struct {
			Text string `xml:",chardata"`
			ABBR struct {
				Text string `xml:",chardata"`
			} `xml:"ABBR"`
		} `xml:"DRA"`
		INFO struct {
			Text string `xml:",chardata"`
			IDIV struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
				Key   string `xml:"key,attr"`
				INAME struct {
					Text string `xml:",chardata"`
				} `xml:"I_NAME"`
				IMIME struct {
					Text string `xml:",chardata"`
					ABBR struct {
						Text string `xml:",chardata"`
					} `xml:"ABBR"`
					SUPP struct {
						Text string `xml:",chardata"`
					} `xml:"SUPP"`
				} `xml:"I_MIME"`
			} `xml:"I_DIV"`
		} `xml:"INFO"`
		HER struct {
			Text string `xml:",chardata"`
		} `xml:"HER"`
	} `xml:"SDG"`
	PREVIEW struct {
		Text  string `xml:",chardata"`
		VIDEO struct {
			Text     string `xml:",chardata"`
			VIDEOURL struct {
				Text string `xml:",chardata"`
			} `xml:"VIDEOURL"`
			TCIN struct {
				Text string `xml:",chardata"`
			} `xml:"TCIN"`
		} `xml:"VIDEO"`
	} `xml:"PREVIEW"`
	BTR struct {
		Text  string `xml:",chardata"`
		ZUORD []struct {
			Text  string `xml:",chardata"`
			RTAKE struct {
				Text string `xml:",chardata"`
			} `xml:"R_TAKE"`
			RIDX struct {
				Text string `xml:",chardata"`
			} `xml:"R_IDX"`
			RTC struct {
				Text     string `xml:",chardata"`
				Accuracy string `xml:"accuracy,attr"`
			} `xml:"R_TC"`
			RDAU struct {
				Text string `xml:",chardata"`
			} `xml:"R_DAU"`
			EDOK struct {
				Text  string `xml:",chardata"`
				RDKEY struct {
					Text string `xml:",chardata"`
				} `xml:"R_DKEY"`
				RSKEY struct {
					Text string `xml:",chardata"`
				} `xml:"R_SKEY"`
				RTIT struct {
					Text string `xml:",chardata"`
				} `xml:"R_TIT"`
				RTAKE struct {
					Text string `xml:",chardata"`
				} `xml:"R_TAKE"`
				RTC struct {
					Text     string `xml:",chardata"`
					Accuracy string `xml:"accuracy,attr"`
				} `xml:"R_TC"`
				RDAU struct {
					Text string `xml:",chardata"`
				} `xml:"R_DAU"`
			} `xml:"EDOK"`
		} `xml:"ZUORD"`
		DDKEY struct {
			Text string `xml:",chardata"`
		} `xml:"D_DKEY"`
		DSKEY struct {
			Text string `xml:",chardata"`
		} `xml:"D_SKEY"`
		ARRFA struct {
			Text string `xml:",chardata"`
			ABBR struct {
				Text string `xml:",chardata"`
			} `xml:"ABBR"`
		} `xml:"AR_RFA"`
		PNR struct {
			Text string `xml:",chardata"`
		} `xml:"PNR"`
		POOL struct {
			Text    string `xml:",chardata"`
			Abbrev  string `xml:"abbrev,attr"`
			Code    string `xml:"code,attr"`
			STATION struct {
				Text string `xml:",chardata"`
			} `xml:"STATION"`
		} `xml:"POOL"`
		PE struct {
			Text string `xml:",chardata"`
			ABBR struct {
				Text string `xml:",chardata"`
			} `xml:"ABBR"`
		} `xml:"PE"`
		BSTAT struct {
			Text string `xml:",chardata"`
			ABBR struct {
				Text string `xml:",chardata"`
			} `xml:"ABBR"`
		} `xml:"BSTAT"`
		KST struct {
			Text string `xml:",chardata"`
		} `xml:"KST"`
		KASS struct {
			Text string `xml:",chardata"`
		} `xml:"KASS"`
		DEEDAT struct {
			Text string `xml:",chardata"`
		} `xml:"DE_EDAT"`
		DEKDAT struct {
			Text string `xml:",chardata"`
		} `xml:"DE_KDAT"`
		ERF struct {
			Text string `xml:",chardata"`
		} `xml:"ERF"`
		TIT struct {
			Text string `xml:",chardata"`
			BETI struct {
				Text   string `xml:",chardata"`
				Descr  string `xml:"descr,attr"`
				TITTXT struct {
					Text string `xml:",chardata"`
				} `xml:"TIT_TXT"`
			} `xml:"BETI"`
		} `xml:"TIT"`
		UPM struct {
			Text string `xml:",chardata"`
			RDN  struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
				ORGA  struct {
					Text string `xml:",chardata"`
				} `xml:"ORGA"`
			} `xml:"RDN"`
			AUT []struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
				VNAME struct {
					Text string `xml:",chardata"`
				} `xml:"VNAME"`
				NNAME struct {
					Text string `xml:",chardata"`
				} `xml:"NNAME"`
			} `xml:"AUT"`
			INT struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
				VNAME struct {
					Text string `xml:",chardata"`
				} `xml:"VNAME"`
				NNAME struct {
					Text string `xml:",chardata"`
				} `xml:"NNAME"`
			} `xml:"INT"`
		} `xml:"UPM"`
		SDA struct {
			Text string `xml:",chardata"`
			ESD  []struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
				SRFA  struct {
					Text string `xml:",chardata"`
					ABBR struct {
						Text string `xml:",chardata"`
					} `xml:"ABBR"`
				} `xml:"S_RFA"`
				VG struct {
					Text string `xml:",chardata"`
					ABBR struct {
						Text string `xml:",chardata"`
					} `xml:"ABBR"`
				} `xml:"VG"`
				SDAT struct {
					Text string `xml:",chardata"`
				} `xml:"SDAT"`
				SBEG struct {
					Text string `xml:",chardata"`
				} `xml:"SBEG"`
				SDAU struct {
					Text string `xml:",chardata"`
				} `xml:"SDAU"`
				SKST struct {
					Text string `xml:",chardata"`
				} `xml:"S_KST"`
				SEDAT struct {
					Text string `xml:",chardata"`
				} `xml:"S_EDAT"`
				SPNR struct {
					Text string `xml:",chardata"`
				} `xml:"S_PNR"`
				SKDAT struct {
					Text string `xml:",chardata"`
				} `xml:"S_KDAT"`
			} `xml:"ESD"`
			WSD []struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
				SRFA  struct {
					Text string `xml:",chardata"`
					ABBR struct {
						Text string `xml:",chardata"`
					} `xml:"ABBR"`
				} `xml:"S_RFA"`
				VG struct {
					Text string `xml:",chardata"`
					ABBR struct {
						Text string `xml:",chardata"`
					} `xml:"ABBR"`
				} `xml:"VG"`
				SDAT struct {
					Text string `xml:",chardata"`
				} `xml:"SDAT"`
				SBEG struct {
					Text string `xml:",chardata"`
				} `xml:"SBEG"`
				SDAU struct {
					Text string `xml:",chardata"`
				} `xml:"SDAU"`
				SPNR struct {
					Text string `xml:",chardata"`
				} `xml:"S_PNR"`
				SKST struct {
					Text string `xml:",chardata"`
				} `xml:"S_KST"`
				SEDAT struct {
					Text string `xml:",chardata"`
				} `xml:"S_EDAT"`
			} `xml:"WSD"`
		} `xml:"SDA"`
		HD struct {
			Text string `xml:",chardata"`
			GEO  []struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
			} `xml:"GEO"`
			SACH []struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
			} `xml:"SACH"`
			STD []struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
			} `xml:"STD"`
		} `xml:"HD"`
		INH struct {
			Text string `xml:",chardata"`
			T    struct {
				Text   string `xml:",chardata"`
				Descr  string `xml:"descr,attr"`
				Format string `xml:"format,attr"`
			} `xml:"T"`
		} `xml:"INH"`
		BTT []struct {
			Text             string `xml:",chardata"`
			Digi             string `xml:"digi,attr"`
			DigiAbbr         string `xml:"digi_abbr,attr"`
			DigiDescr        string `xml:"digi_descr,attr"`
			HasKeyframes     string `xml:"hasKeyframes,attr"`
			IsDefaultPreview string `xml:"isDefaultPreview,attr"`
			Key              string `xml:"key,attr"`
			LoanValue        string `xml:"loanValue,attr"`
			Manzeige         string `xml:"manzeige,attr"`
			Mauslinfo        string `xml:"mauslinfo,attr"`
			Mpreview         string `xml:"mpreview,attr"`
			Mvormerk         string `xml:"mvormerk,attr"`
			Station          string `xml:"station,attr"`
			BERFA            struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BE_RFA"`
			ANK struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"ANK"`
			MANR struct {
				Text string `xml:",chardata"`
			} `xml:"M_ANR"`
			BEH struct {
				Text string `xml:",chardata"`
			} `xml:"BEH"`
			ROL struct {
				Text string `xml:",chardata"`
			} `xml:"ROL"`
			LGO struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"LGO"`
			BTTSPR struct {
				Text string `xml:",chardata"`
			} `xml:"BTT_SPR"`
			BTTDS struct {
				Text string `xml:",chardata"`
			} `xml:"BTT_DS"`
			BTTPFR struct {
				Text string `xml:",chardata"`
			} `xml:"BTT_PFR"`
			BTTDAU struct {
				Text string `xml:",chardata"`
			} `xml:"BTT_DAU"`
			ARDAT struct {
				Text string `xml:",chardata"`
			} `xml:"ARDAT"`
			BTTEDAT struct {
				Text string `xml:",chardata"`
			} `xml:"BTT_EDAT"`
			BTTTA struct {
				Text string `xml:",chardata"`
			} `xml:"BTT_TA"`
			BTTTC struct {
				Text     string `xml:",chardata"`
				Accuracy string `xml:"accuracy,attr"`
			} `xml:"BTT_TC"`
			BTTLZ struct {
				Text string `xml:",chardata"`
			} `xml:"BTT_LZ"`
			AZDAT struct {
				Text string `xml:",chardata"`
			} `xml:"AZDAT"`
			BTTINQU struct {
				Text string `xml:",chardata"`
			} `xml:"BTT_INQU"`
			BTTID struct {
				Text  string `xml:",chardata"`
				IDBCD struct {
					Text  string `xml:",chardata"`
					Descr string `xml:"descr,attr"`
				} `xml:"ID_BCD"`
				IDESS struct {
					Text  string `xml:",chardata"`
					Descr string `xml:"descr,attr"`
				} `xml:"ID_ESS"`
			} `xml:"BTT_ID"`
			BTTCONF struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BTT_CONF"`
			BTTRV struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BTT_RV"`
			BTTRVH struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BTT_RVH"`
			BTTVAR struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BTT_VAR"`
			BTTVQ struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BTT_VQ"`
			BTTAAF struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BTT_AAF"`
			BTTAQ struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BTT_AQ"`
			BTTBPR struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BTT_BPR"`
			BTTFSS struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BTT_FSS"`
			BTTHVM struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BTT_HVM"`
			BTTANM struct {
				Text string `xml:",chardata"`
			} `xml:"BTT_ANM"`
			TSP []struct {
				Text   string `xml:",chardata"`
				TRACK  string `xml:"TRACK,attr"`
				TSWERT struct {
					Text string `xml:",chardata"`
					ABBR struct {
						Text string `xml:",chardata"`
					} `xml:"ABBR"`
				} `xml:"TS_WERT"`
				TSRDA struct {
					Text string `xml:",chardata"`
					ABBR struct {
						Text string `xml:",chardata"`
					} `xml:"ABBR"`
				} `xml:"TS_RDA"`
			} `xml:"TSP"`
			RNG struct {
				Text string `xml:",chardata"`
			} `xml:"RNG"`
			BTTSYS struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BTT_SYS"`
			BTTKDAT struct {
				Text string `xml:",chardata"`
			} `xml:"BTT_KDAT"`
			BTTFRB struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BTT_FRB"`
			BTTPST struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BTT_PST"`
			BTTPOS struct {
				Text string `xml:",chardata"`
			} `xml:"BTT_POS"`
			BTTRST struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"BTT_RST"`
		} `xml:"BTT"`
		SEQ []struct {
			Text  string `xml:",chardata"`
			ZUORD []struct {
				Text  string `xml:",chardata"`
				RTAKE struct {
					Text string `xml:",chardata"`
				} `xml:"R_TAKE"`
				RIDX struct {
					Text string `xml:",chardata"`
				} `xml:"R_IDX"`
				RTC struct {
					Text     string `xml:",chardata"`
					Accuracy string `xml:"accuracy,attr"`
				} `xml:"R_TC"`
				RDAU struct {
					Text string `xml:",chardata"`
				} `xml:"R_DAU"`
				EDOK struct {
					Text  string `xml:",chardata"`
					RDKEY struct {
						Text string `xml:",chardata"`
					} `xml:"R_DKEY"`
					RSKEY struct {
						Text string `xml:",chardata"`
					} `xml:"R_SKEY"`
					RTIT struct {
						Text string `xml:",chardata"`
					} `xml:"R_TIT"`
					RTAKE struct {
						Text string `xml:",chardata"`
					} `xml:"R_TAKE"`
					RTC struct {
						Text     string `xml:",chardata"`
						Accuracy string `xml:"accuracy,attr"`
					} `xml:"R_TC"`
					RDAU struct {
						Text string `xml:",chardata"`
					} `xml:"R_DAU"`
				} `xml:"EDOK"`
			} `xml:"ZUORD"`
			DDKEY struct {
				Text string `xml:",chardata"`
			} `xml:"D_DKEY"`
			DSKEY struct {
				Text string `xml:",chardata"`
			} `xml:"D_SKEY"`
			ARRFA struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"AR_RFA"`
			POOL struct {
				Text    string `xml:",chardata"`
				Abbrev  string `xml:"abbrev,attr"`
				Code    string `xml:"code,attr"`
				STATION struct {
					Text string `xml:",chardata"`
				} `xml:"STATION"`
			} `xml:"POOL"`
			KASS struct {
				Text string `xml:",chardata"`
			} `xml:"KASS"`
			DEEDAT struct {
				Text string `xml:",chardata"`
			} `xml:"DE_EDAT"`
			ERF struct {
				Text string `xml:",chardata"`
			} `xml:"ERF"`
			INH struct {
				Text string `xml:",chardata"`
				B    struct {
					Text   string `xml:",chardata"`
					Descr  string `xml:"descr,attr"`
					Format string `xml:"format,attr"`
				} `xml:"B"`
				O struct {
					Text   string `xml:",chardata"`
					Descr  string `xml:"descr,attr"`
					Format string `xml:"format,attr"`
				} `xml:"O"`
			} `xml:"INH"`
			DV struct {
				Text string `xml:",chardata"`
			} `xml:"DV"`
			DRA struct {
				Text string `xml:",chardata"`
				ABBR struct {
					Text string `xml:",chardata"`
				} `xml:"ABBR"`
			} `xml:"DRA"`
		} `xml:"SEQ"`
		DV struct {
			Text string `xml:",chardata"`
		} `xml:"DV"`
		DRA struct {
			Text string `xml:",chardata"`
			ABBR struct {
				Text string `xml:",chardata"`
			} `xml:"ABBR"`
		} `xml:"DRA"`
		ANM struct {
			Text  string `xml:",chardata"`
			APABR struct {
				Text  string `xml:",chardata"`
				Descr string `xml:"descr,attr"`
			} `xml:"A_PABR"`
		} `xml:"ANM"`
	} `xml:"BTR"`
}
