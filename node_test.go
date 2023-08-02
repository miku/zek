// Copyright 2021 by Leipzig University Library, http://ub.uni-leipzig.de
//                   The Finc Authors, http://finc.info
//                   Martin Czygan, <martin.czygan@uni-leipzig.de>
//
// This file is part of some open source application.
//
// Some open source application is free software: you can redistribute
// it and/or modify it under the terms of the GNU General Public
// License as published by the Free Software Foundation, either
// version 3 of the License, or (at your option) any later version.
//
// Some open source application is distributed in the hope that it will
// be useful, but WITHOUT ANY WARRANTY; without even the implied warranty
// of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Foobar.  If not, see <http://www.gnu.org/licenses/>.
//
// @license GPL-3.0+ <http://spdx.org/licenses/GPL-3.0+>

package zek

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestNodeMergeAttr(t *testing.T) {
	var cases = []struct {
		node   *Node
		attr   []xml.Attr
		result []xml.Attr
	}{
		{
			node:   &Node{Attr: []xml.Attr{}},
			attr:   []xml.Attr{},
			result: []xml.Attr{},
		},
		{
			node:   &Node{Attr: []xml.Attr{}},
			attr:   []xml.Attr{{Name: xml.Name{Local: "id"}}},
			result: []xml.Attr{{Name: xml.Name{Local: "id"}}},
		},
		{
			node:   &Node{Attr: []xml.Attr{{Name: xml.Name{Local: "id"}}}},
			attr:   []xml.Attr{{Name: xml.Name{Local: "id"}}},
			result: []xml.Attr{{Name: xml.Name{Local: "id"}}},
		},
	}
	for _, c := range cases {
		c.node.mergeAttr(c.attr)
		if !reflect.DeepEqual(c.node.Attr, c.result) {
			t.Errorf("got %v, want %v", c.node.Attr, c.result)
		}
	}
}

func deepEqualJSON(s, t string) (bool, error) {
	var a, b interface{}
	if err := json.Unmarshal([]byte(s), &a); err != nil {
		return false, err
	}
	if err := json.Unmarshal([]byte(t), &b); err != nil {
		return false, err
	}
	return reflect.DeepEqual(a, b), nil
}

func TestNodeReadFrom(t *testing.T) {
	var cases = []struct {
		input            string // Raw XML input string.
		serializedResult string // Compare to JSON serialization.
		err              error
	}{
		{
			input: ``,
			serializedResult: `
			{
					"name": {
							"Space": "",
							"Local": ""
					}
			}`,
			err: nil,
		},
		{
			input: `<?xml version="1.0"?><a></a>`,
			serializedResult: `
			{
					"name": {
							"Space": "",
							"Local": "a"
					}
			}`,
			err: nil,
		},
		{
			input: `<a></a>`,
			serializedResult: `
			{
					"name": {
							"Space": "",
							"Local": "a"
					}
			}`,
			err: nil,
		},
		{
			input: `<a><b></b></a>`,
			serializedResult: `
			{
					"name": {
							"Space": "",
							"Local": "a"
					},
					"children": [
							{
									"name": {
											"Space": "",
											"Local": "b"
									}
							}
					]
			}`,
			err: nil,
		},
		{
			input: `<a><b></b></a>`,
			serializedResult: `
			{
					"name": {
							"Space": "",
							"Local": "a"
					},
					"children": [
							{
									"name": {
											"Space": "",
											"Local": "b"
									}
							}
					]
			}`,
			err: nil,
		},
		{
			input: `<a><b></b><b></b></a>`,
			serializedResult: `
			{
					"name": {
							"Space": "",
							"Local": "a"
					},
					"children": [
							{
									"name": {
											"Space": "",
											"Local": "b"
									}
							}
					]
			}`,
			err: nil,
		},
		{
			input: `<a><b><c></c></b></a>`,
			serializedResult: `
			{
					"name": {
							"Space": "",
							"Local": "a"
					},
					"children": [
							{
									"name": {
											"Space": "",
											"Local": "b"
									},
									"children": [
											{
													"name": {
															"Space": "",
															"Local": "c"
													}
											}
									]
							}
					]
			}`,
			err: nil,
		},
		{
			input: `<a id="1"><b><c></c></b></a>`,
			serializedResult: `
			{
					"name": {
							"Space": "",
							"Local": "a"
					},
					"attr": [
							{
									"Name": {
											"Space": "",
											"Local": "id"
									},
									"Value": "1"
							}
					],
					"children": [
							{
									"name": {
											"Space": "",
											"Local": "b"
									},
									"children": [
											{
													"name": {
															"Space": "",
															"Local": "c"
													}
											}
									]
							}
					]
			}`,
			err: nil,
		},
		{
			input: `<a id="2"><b><c></c></b></a><a id="1"></a>`,
			serializedResult: `
			{
					"name": {
							"Space": "",
							"Local": "a"
					},
					"attr": [
							{
									"Name": {
											"Space": "",
											"Local": "id"
									},
									"Value": "2"
							}
					],
					"children": [
							{
									"name": {
											"Space": "",
											"Local": "b"
									},
									"children": [
											{
													"name": {
															"Space": "",
															"Local": "c"
													}
											}
									]
							}
					]
			}`,
			err: nil,
		},
		{
			input: `<x><a id="2"><b><c></c></b></a><a id="1"><b><d></d></b></a></x>`,
			serializedResult: `
			{
					"name": {
							"Space": "",
							"Local": "x"
					},
					"children": [
							{
									"name": {
											"Space": "",
											"Local": "a"
									},
									"attr": [
											{
													"Name": {
															"Space": "",
															"Local": "id"
													},
													"Value": "2"
											}
									],
									"children": [
											{
													"name": {
															"Space": "",
															"Local": "b"
													},
													"children": [
															{
																	"name": {
																			"Space": "",
																			"Local": "c"
																	}
															},
															{
																	"name": {
																			"Space": "",
																			"Local": "d"
																	}
															}
													]
											}
									]
							}
					]
			}`,
			err: nil,
		},
		{
			input: `<x><a id="2"><b><c></c></b></a><a name="A"><b><d></d></b></a></x>`,
			serializedResult: `
			{
					"name": {
							"Space": "",
							"Local": "x"
					},
					"children": [
							{
									"name": {
											"Space": "",
											"Local": "a"
									},
									"attr": [
											{
													"Name": {
															"Space": "",
															"Local": "id"
													},
													"Value": "2"
											},
											{
												"Name": {
														"Space": "",
														"Local": "name"
												},
												"Value": "A"
										}
									],
									"children": [
											{
													"name": {
															"Space": "",
															"Local": "b"
													},
													"children": [
															{
																	"name": {
																			"Space": "",
																			"Local": "c"
																	}
															},
															{
																	"name": {
																			"Space": "",
																			"Local": "d"
																	}
															}
													]
											}
									]
							}
					]
			}`,
			err: nil,
		},
		{
			input: `<x><a><b><c></c><c></c></b></a><a><b><d></d><d></d></b></a></x>`,
			serializedResult: `
			{
					"name": {
							"Space": "",
							"Local": "x"
					},
					"children": [
							{
									"name": {
											"Space": "",
											"Local": "a"
									},
									"children": [
											{
													"name": {
															"Space": "",
															"Local": "b"
													},
													"children": [
															{
																	"name": {
																			"Space": "",
																			"Local": "c"
																	}
															},
															{
																	"name": {
																			"Space": "",
																			"Local": "d"
																	}
															}
													]
											}
									]
							}
					]
			}`,
			err: nil,
		},
	}

	for _, c := range cases {
		r := strings.NewReader(c.input)
		node := new(Node)
		_, err := node.ReadFrom(r, &ReadOpts{MaxExamples: 10})
		if err != c.err {
			t.Errorf("got %v, want %v", err, c.err)
		}
		b, err := json.Marshal(node)
		if err != nil {
			t.Errorf("unexpected error during serialization: %v", err)
		}
		s := string(b)
		if ok, err := deepEqualJSON(s, c.serializedResult); !ok {
			if err != nil {
				t.Errorf("documents differ: %v", err)
			}
			t.Errorf("got %v, want %v", s, c.serializedResult)
		}
	}
}

func TestHeight(t *testing.T) {
	var cases = []struct {
		input  string
		height int
	}{
		{
			input: "", height: 0,
		},
		{
			input: "<a></a>", height: 1,
		},
		{
			input: "<a><b></b></a>", height: 2,
		},
		{
			input: "<a><b></b><c><d></d></c></a>", height: 3,
		},
		{
			input: "<a><b></b><b></b><b></b><b></b></a>", height: 2,
		},
	}

	for _, c := range cases {
		r := strings.NewReader(c.input)
		node := new(Node)
		_, err := node.ReadFrom(r, &ReadOpts{MaxExamples: 10})
		if err != nil {
			t.Errorf("failed to parse tree: %s", err)
		}
		if node.Height() != c.height {
			t.Errorf("got %v, want %v", node.Height(), c.height)
		}
	}
}

func TestByName(t *testing.T) {
	r := strings.NewReader(`<a><b><c></c></b></a>`)
	root := new(Node)
	if _, err := root.ReadFrom(r, &ReadOpts{MaxExamples: 10}); err != nil {
		t.Errorf("got %v, want nil", err)
	}

	var cases = []struct {
		name   string
		result *Node
	}{
		{
			name: "", result: root,
		},
		{
			name: "a", result: root,
		},
		{
			name: "b", result: root.Children[0],
		},
		{
			name: "c", result: root.Children[0].Children[0],
		},
		{
			name: "d", result: nil,
		},
	}
	for _, c := range cases {
		result := root.ByName(c.name)
		if result != c.result {
			t.Errorf("got %v, want %v", result, c.result)
		}
	}

	// Test case, where node is nil.
	var nn *Node
	if got := nn.ByName("anything"); got != nil {
		t.Errorf("got %v, want nil", got)
	}
}

func TestNodeReadFromAll(t *testing.T) {
	var cases = []struct {
		inputs           []string // Raw XML input strings.
		serializedResult string   // Compare to JSON serialization.
		err              error
	}{
		{
			inputs: []string{``, ``},
			serializedResult: `
			{
					"name": {
							"Space": "",
							"Local": ""
					}
			}`,
			err: nil,
		},
		{
			inputs: []string{`<?xml version="1.0"?><a></a>`},
			serializedResult: `
			{
					"name": {
							"Space": "",
							"Local": "a"
					}
			}`,
			err: nil,
		},
		{
			inputs: []string{`<a><b></b></a>`, `<a><c></c></a>`},
			serializedResult: `
			{
				"name": {
					"Space": "",
					"Local": "a"
				},
				"children": [
					{
						"name": {
							"Space": "",
							"Local": "b"
						}
					},
					{
						"name": {
							"Space": "",
							"Local": "c"
						}
					}
				]
			}
			`,
			err: nil,
		},
		{
			inputs: []string{`<a></a>`, ``},
			serializedResult: `
			{
					"name": {
							"Space": "",
							"Local": "a"
					}
			}`,
			err: nil,
		},
		{
			inputs: []string{`<a><b></b></a>`, `<a></a>`},
			serializedResult: `
			{
					"name": {
							"Space": "",
							"Local": "a"
					},
					"children": [
							{
									"name": {
											"Space": "",
											"Local": "b"
									}
							}
					]
			}`,
			err: nil,
		},
		{
			inputs: []string{`<a></a>`, `<a><b></b></a>`},
			serializedResult: `
			{
					"name": {
							"Space": "",
							"Local": "a"
					},
					"children": [
							{
									"name": {
											"Space": "",
											"Local": "b"
									}
							}
					]
			}`,
			err: nil,
		},
		{
			inputs: []string{`<a><b></b><b></b></a>`, `<a><b></b></a>`},
			serializedResult: `
			{
					"name": {
							"Space": "",
							"Local": "a"
					},
					"children": [
							{
									"name": {
											"Space": "",
											"Local": "b"
									}
							}
					]
			}`,
			err: nil,
		},
		{
			inputs: []string{`<a><b><c></c></b></a>`, ``},
			serializedResult: `
			{
					"name": {
							"Space": "",
							"Local": "a"
					},
					"children": [
							{
									"name": {
											"Space": "",
											"Local": "b"
									},
									"children": [
											{
													"name": {
															"Space": "",
															"Local": "c"
													}
											}
									]
							}
					]
			}`,
			err: nil,
		},
		{
			inputs: []string{`<a id="1"><b><c></c></b></a>`, ``},
			serializedResult: `
			{
					"name": {
							"Space": "",
							"Local": "a"
					},
					"attr": [
							{
									"Name": {
											"Space": "",
											"Local": "id"
									},
									"Value": "1"
							}
					],
					"children": [
							{
									"name": {
											"Space": "",
											"Local": "b"
									},
									"children": [
											{
													"name": {
															"Space": "",
															"Local": "c"
													}
											}
									]
							}
					]
			}`,
			err: nil,
		},
		{
			inputs: []string{`<a id="2"><b><c></c></b></a><a id="1"></a>`, ``, ``},
			serializedResult: `
			{
					"name": {
							"Space": "",
							"Local": "a"
					},
					"attr": [
							{
									"Name": {
											"Space": "",
											"Local": "id"
									},
									"Value": "2"
							}
					],
					"children": [
							{
									"name": {
											"Space": "",
											"Local": "b"
									},
									"children": [
											{
													"name": {
															"Space": "",
															"Local": "c"
													}
											}
									]
							}
					]
			}`,
			err: nil,
		},
		{
			inputs: []string{`<x><a id="2"><b><c></c></b></a><a id="1"><b><d></d></b></a></x>`, ``},
			serializedResult: `
			{
					"name": {
							"Space": "",
							"Local": "x"
					},
					"children": [
							{
									"name": {
											"Space": "",
											"Local": "a"
									},
									"attr": [
											{
													"Name": {
															"Space": "",
															"Local": "id"
													},
													"Value": "2"
											}
									],
									"children": [
											{
													"name": {
															"Space": "",
															"Local": "b"
													},
													"children": [
															{
																	"name": {
																			"Space": "",
																			"Local": "c"
																	}
															},
															{
																	"name": {
																			"Space": "",
																			"Local": "d"
																	}
															}
													]
											}
									]
							}
					]
			}`,
			err: nil,
		},
		{
			inputs: []string{`<x><a id="2"><b><c></c></b></a><a name="A"><b><d></d></b></a></x>`},
			serializedResult: `
			{
					"name": {
							"Space": "",
							"Local": "x"
					},
					"children": [
							{
									"name": {
											"Space": "",
											"Local": "a"
									},
									"attr": [
											{
													"Name": {
															"Space": "",
															"Local": "id"
													},
													"Value": "2"
											},
											{
												"Name": {
														"Space": "",
														"Local": "name"
												},
												"Value": "A"
										}
									],
									"children": [
											{
													"name": {
															"Space": "",
															"Local": "b"
													},
													"children": [
															{
																	"name": {
																			"Space": "",
																			"Local": "c"
																	}
															},
															{
																	"name": {
																			"Space": "",
																			"Local": "d"
																	}
															}
													]
											}
									]
							}
					]
			}`,
			err: nil,
		},
		{
			inputs: []string{`<x><a><b><c></c><c></c></b></a><a><b><d></d><d></d></b></a></x>`, ``},
			serializedResult: `
			{
					"name": {
							"Space": "",
							"Local": "x"
					},
					"children": [
							{
									"name": {
											"Space": "",
											"Local": "a"
									},
									"children": [
											{
													"name": {
															"Space": "",
															"Local": "b"
													},
													"children": [
															{
																	"name": {
																			"Space": "",
																			"Local": "c"
																	}
															},
															{
																	"name": {
																			"Space": "",
																			"Local": "d"
																	}
															}
													]
											}
									]
							}
					]
			}`,
			err: nil,
		},
	}

	for _, c := range cases {
		var readers []io.Reader

		for _, input := range c.inputs {
			readers = append(readers, strings.NewReader(input))
		}

		node := new(Node)
		_, err := node.ReadFrom(io.MultiReader(readers...), &ReadOpts{MaxExamples: 10})
		if err != c.err {
			t.Errorf("got %v, want %v", err, c.err)
		}
		b, err := json.Marshal(node)
		if err != nil {
			t.Errorf("unexpected error during serialization: %v", err)
		}
		s := string(b)
		if ok, err := deepEqualJSON(s, c.serializedResult); !ok {
			if err != nil {
				t.Errorf("documents differ: %v", err)
			}
			t.Errorf("got %v, want %v", s, c.serializedResult)
		}
	}
}
