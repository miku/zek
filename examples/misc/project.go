package misc

import "encoding/xml"

// Project was generated 2017-12-06 01:15:18 by tir on apollo.
type Project struct {
	XMLName   xml.Name `xml:"project"`
	Text      string   `xml:",chardata"`
	Version   string   `xml:"version,attr"`
	Component []struct {
		Text                   string `xml:",chardata"`
		Name                   string `xml:"name,attr"`
		AttrDetectionDone      string `xml:"detection-done,attr"`
		AttrSorting            string `xml:"sorting,attr"`
		ExtendedState          string `xml:"extendedState,attr"`
		SettingsEditedManually string `xml:"settingsEditedManually,attr"`
		ShowRecycled           string `xml:"show_recycled,attr"`
		List                   struct {
			Text    string `xml:",chardata"`
			Default string `xml:"default,attr"`
			ID      string `xml:"id,attr"`
			Name    string `xml:"name,attr"`
			Comment string `xml:"comment,attr"`
			Change  []struct {
				Text       string `xml:",chardata"`
				BeforePath string `xml:"beforePath,attr"`
				AfterPath  string `xml:"afterPath,attr"`
			} `xml:"change"`
		} `xml:"list"`
		Option []struct {
			Text  string `xml:",chardata"`
			Name  string `xml:"name,attr"`
			Value string `xml:"value,attr"`
			List  struct {
				Text   string `xml:",chardata"`
				Option []struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"option"`
				RuleState []struct {
					Text   string `xml:",chardata"`
					Option struct {
						Text  string `xml:",chardata"`
						Name  string `xml:"name,attr"`
						Value string `xml:"value,attr"`
					} `xml:"option"`
				} `xml:"RuleState"`
			} `xml:"list"`
		} `xml:"option"`
		Leaf struct {
			Text                 string `xml:",chardata"`
			SIDETABSSIZELIMITKEY string `xml:"SIDE_TABS_SIZE_LIMIT_KEY,attr"`
			File                 []struct {
				Text         string `xml:",chardata"`
				LeafFileName string `xml:"leaf-file-name,attr"`
				Pinned       string `xml:"pinned,attr"`
				CurrentInTab string `xml:"current-in-tab,attr"`
				Entry        struct {
					Text     string `xml:",chardata"`
					File     string `xml:"file,attr"`
					Provider struct {
						Text         string `xml:",chardata"`
						Selected     string `xml:"selected,attr"`
						EditorTypeID string `xml:"editor-type-id,attr"`
						State        struct {
							Text                  string `xml:",chardata"`
							RelativeCaretPosition string `xml:"relative-caret-position,attr"`
							Caret                 struct {
								Text                 string `xml:",chardata"`
								Line                 string `xml:"line,attr"`
								Column               string `xml:"column,attr"`
								LeanForward          string `xml:"lean-forward,attr"`
								SelectionStartLine   string `xml:"selection-start-line,attr"`
								SelectionStartColumn string `xml:"selection-start-column,attr"`
								SelectionEndLine     string `xml:"selection-end-line,attr"`
								SelectionEndColumn   string `xml:"selection-end-column,attr"`
							} `xml:"caret"`
							Folding struct {
								Text    string `xml:",chardata"`
								Element struct {
									Text      string `xml:",chardata"`
									Signature string `xml:"signature,attr"`
									Expanded  string `xml:"expanded,attr"`
								} `xml:"element"`
							} `xml:"folding"`
						} `xml:"state"`
					} `xml:"provider"`
				} `xml:"entry"`
			} `xml:"file"`
		} `xml:"leaf"`
		DetectionDone struct {
			Text string `xml:",chardata"` // true
		} `xml:"detection-done"`
		Sorting struct {
			Text string `xml:",chardata"` // DEFINITION_ORDER
		} `xml:"sorting"`
		Navigator struct {
			Text            string `xml:",chardata"`
			CurrentView     string `xml:"currentView,attr"`
			Proportions     string `xml:"proportions,attr"`
			Version         string `xml:"version,attr"`
			FlattenPackages struct {
				Text string `xml:",chardata"`
			} `xml:"flattenPackages"`
			ShowMembers struct {
				Text string `xml:",chardata"`
			} `xml:"showMembers"`
			ShowModules struct {
				Text string `xml:",chardata"`
			} `xml:"showModules"`
			ShowLibraryContents struct {
				Text string `xml:",chardata"`
			} `xml:"showLibraryContents"`
			HideEmptyPackages struct {
				Text string `xml:",chardata"`
			} `xml:"hideEmptyPackages"`
			AbbreviatePackageNames struct {
				Text string `xml:",chardata"`
			} `xml:"abbreviatePackageNames"`
			AutoscrollToSource struct {
				Text string `xml:",chardata"`
			} `xml:"autoscrollToSource"`
			AutoscrollFromSource struct {
				Text string `xml:",chardata"`
			} `xml:"autoscrollFromSource"`
			SortByType struct {
				Text string `xml:",chardata"`
			} `xml:"sortByType"`
			ManualOrder struct {
				Text string `xml:",chardata"`
			} `xml:"manualOrder"`
			FoldersAlwaysOnTop struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"foldersAlwaysOnTop"`
		} `xml:"navigator"`
		Panes struct {
			Text string `xml:",chardata"`
			Pane []struct {
				Text    string `xml:",chardata"`
				ID      string `xml:"id,attr"`
				SubPane struct {
					Text   string `xml:",chardata"`
					Expand struct {
						Text string `xml:",chardata"`
						Path []struct {
							Text string `xml:",chardata"`
							Item []struct {
								Text string `xml:",chardata"`
								Name string `xml:"name,attr"`
								Type string `xml:"type,attr"`
							} `xml:"item"`
						} `xml:"path"`
					} `xml:"expand"`
					Select struct {
						Text string `xml:",chardata"`
					} `xml:"select"`
				} `xml:"subPane"`
			} `xml:"pane"`
		} `xml:"panes"`
		Property []struct {
			Text  string `xml:",chardata"`
			Name  string `xml:"name,attr"`
			Value string `xml:"value,attr"`
		} `xml:"property"`
		Key []struct {
			Text   string `xml:",chardata"`
			Name   string `xml:"name,attr"`
			Recent []struct {
				Text string `xml:",chardata"`
				Name string `xml:"name,attr"`
			} `xml:"recent"`
		} `xml:"key"`
		Configuration struct {
			Text        string `xml:",chardata"`
			Name        string `xml:"name,attr"`
			Type        string `xml:"type,attr"`
			FactoryName string `xml:"factoryName,attr"`
			Option      []struct {
				Text  string `xml:",chardata"`
				Name  string `xml:"name,attr"`
				Value string `xml:"value,attr"`
			} `xml:"option"`
			Envs struct {
				Text string `xml:",chardata"`
				Env  []struct {
					Text  string `xml:",chardata"`
					Name  string `xml:"name,attr"`
					Value string `xml:"value,attr"`
				} `xml:"env"`
			} `xml:"envs"`
			Module struct {
				Text string `xml:",chardata"`
				Name string `xml:"name,attr"`
			} `xml:"module"`
		} `xml:"configuration"`
		Task struct {
			Text       string `xml:",chardata"`
			Active     string `xml:"active,attr"`
			ID         string `xml:"id,attr"`
			Summary    string `xml:"summary,attr"`
			Changelist struct {
				Text    string `xml:",chardata"`
				ID      string `xml:"id,attr"`
				Name    string `xml:"name,attr"`
				Comment string `xml:"comment,attr"`
			} `xml:"changelist"`
			Created struct {
				Text string `xml:",chardata"` // 1511979393981
			} `xml:"created"`
			Option []struct {
				Text  string `xml:",chardata"`
				Name  string `xml:"name,attr"`
				Value string `xml:"value,attr"`
			} `xml:"option"`
			Updated struct {
				Text string `xml:",chardata"` // 1511979393981
			} `xml:"updated"`
		} `xml:"task"`
		Servers struct {
			Text string `xml:",chardata"`
		} `xml:"servers"`
		Frame struct {
			Text          string `xml:",chardata"`
			X             string `xml:"x,attr"`
			Y             string `xml:"y,attr"`
			Width         string `xml:"width,attr"`
			Height        string `xml:"height,attr"`
			ExtendedState string `xml:"extended-state,attr"`
		} `xml:"frame"`
		Editor struct {
			Text   string `xml:",chardata"`
			Active string `xml:"active,attr"`
		} `xml:"editor"`
		Layout struct {
			Text       string `xml:",chardata"`
			WindowInfo []struct {
				Text             string `xml:",chardata"`
				ID               string `xml:"id,attr"`
				Active           string `xml:"active,attr"`
				Anchor           string `xml:"anchor,attr"`
				AutoHide         string `xml:"auto_hide,attr"`
				InternalType     string `xml:"internal_type,attr"`
				Type             string `xml:"type,attr"`
				Visible          string `xml:"visible,attr"`
				ShowStripeButton string `xml:"show_stripe_button,attr"`
				Weight           string `xml:"weight,attr"`
				SideWeight       string `xml:"sideWeight,attr"`
				Order            string `xml:"order,attr"`
				SideTool         string `xml:"side_tool,attr"`
				ContentUi        string `xml:"content_ui,attr"`
			} `xml:"window_info"`
		} `xml:"layout"`
		BreakpointManager struct {
			Text string `xml:",chardata"`
		} `xml:"breakpoint-manager"`
		WatchesManager struct {
			Text string `xml:",chardata"`
		} `xml:"watches-manager"`
		Entry []struct {
			Text     string `xml:",chardata"`
			File     string `xml:"file,attr"`
			Provider struct {
				Text         string `xml:",chardata"`
				Selected     string `xml:"selected,attr"`
				EditorTypeID string `xml:"editor-type-id,attr"`
				State        struct {
					Text                  string `xml:",chardata"`
					RelativeCaretPosition string `xml:"relative-caret-position,attr"`
					Caret                 struct {
						Text                 string `xml:",chardata"`
						Line                 string `xml:"line,attr"`
						Column               string `xml:"column,attr"`
						LeanForward          string `xml:"lean-forward,attr"`
						SelectionStartLine   string `xml:"selection-start-line,attr"`
						SelectionStartColumn string `xml:"selection-start-column,attr"`
						SelectionEndLine     string `xml:"selection-end-line,attr"`
						SelectionEndColumn   string `xml:"selection-end-column,attr"`
					} `xml:"caret"`
					Folding struct {
						Text    string `xml:",chardata"`
						Element struct {
							Text      string `xml:",chardata"`
							Signature string `xml:"signature,attr"`
							Expanded  string `xml:"expanded,attr"`
						} `xml:"element"`
					} `xml:"folding"`
				} `xml:"state"`
			} `xml:"provider"`
		} `xml:"entry"`
	} `xml:"component"`
}
