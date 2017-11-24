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

