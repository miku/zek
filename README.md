zek
===

Zek for XML acquisition.

Analyzing and transforming XML can be a burden (XSLT specification printout is about 900 pages).

Features
--------

* Readable Document summaries (tag, attribute utilization), also: machine-readable
* Go struct generation without schema
* Schema inference from data
* Fake data generation

Usage
-----

```shell
$ zek file.xml
```

```
$ zek -s file.xml > stub.go
```

Translation rules
-----------------

For each element:

* The struct name corresponds to the element name.
* The chardata is exposed as Text field. If a tag would result in a Text field,
  it is renamed: TextTag or similar.
* Attributes are expressed as fields with the same name. If there is a name
  clash, the attribute names are prefixed with Attr.

Precedence: Text is always chardata, Tag might be raenamed, if named text.
Attributes are least important.

* TODO: recursive structures. Detect Duplicate Subtrees. Extract the subtrees as
  separate structs and use them in the main struct.
* TODO: calculate some stats on the fields (cardinality, multiplicity, average values)
* TODO: correlate values between fields
* TODO: add examples as comments

Scope
-----

* Become a go to library for advanced XML related processing in Go
* https://kore-nordmann.de/talks/11_03_learning_xml_schema_definitions_from_xml_data.pdf
* Learn XSD, DTD or similar schema from XML instances
* Generate structs from XSD, DTD or plain XML
* Readable structs, with good annotations
