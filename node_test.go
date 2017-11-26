package zek

import (
	"encoding/json"
	"encoding/xml"
	"reflect"
	"strings"
	"testing"
)

func TestAttrListContains(t *testing.T) {
	var cases = []struct {
		attrs  []xml.Attr
		attr   xml.Attr
		result bool
	}{
		{
			attrs:  []xml.Attr{xml.Attr{}},
			attr:   xml.Attr{},
			result: true,
		},
		{
			attrs:  []xml.Attr{xml.Attr{}},
			attr:   xml.Attr{Name: xml.Name{Local: "id"}},
			result: false,
		},
		{
			attrs:  []xml.Attr{xml.Attr{Name: xml.Name{Local: "id"}}},
			attr:   xml.Attr{Name: xml.Name{Local: "id"}},
			result: true,
		},
		{
			attrs:  []xml.Attr{xml.Attr{Name: xml.Name{Local: "id"}, Value: "yyy"}},
			attr:   xml.Attr{Name: xml.Name{Local: "id"}, Value: "xxx"},
			result: true,
		},
	}

	for _, c := range cases {
		got := attrListContains(c.attrs, c.attr)
		if c.result != got {
			t.Errorf("got %v, want %v", got, c.result)
		}
	}
}

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
			attr:   []xml.Attr{xml.Attr{Name: xml.Name{Local: "id"}}},
			result: []xml.Attr{xml.Attr{Name: xml.Name{Local: "id"}}},
		},
		{
			node:   &Node{Attr: []xml.Attr{xml.Attr{Name: xml.Name{Local: "id"}}}},
			attr:   []xml.Attr{xml.Attr{Name: xml.Name{Local: "id"}}},
			result: []xml.Attr{xml.Attr{Name: xml.Name{Local: "id"}}},
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
		_, err := node.ReadFrom(r)
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
		_, err := node.ReadFrom(r)
		if err != nil {
			t.Errorf("failed to parse tree: %s", err)
		}
		if node.Height() != c.height {
			t.Errorf("got %v, want %v", node.Height(), c.height)
		}
	}
}
