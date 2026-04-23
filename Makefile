SHELL = /bin/bash

TARGETS = zek
PKGNAME = zek
ARCH = amd64
VERSION = 0.1.31

.PHONY: all
all: $(TARGETS)

$(TARGETS): %: cmd/%/main.go
	go get -v ./...
	go build -ldflags="-s -w -X github.com/miku/zek.Version=$(VERSION)" -v -o $@ $<

.PHONY: test
test:
	go test -cover -v ./...

.PHONY: clean
clean:
	rm -f $(TARGETS)
	rm -f $(PKGNAME)*.deb
	rm -f $(PKGNAME)-*.rpm
	rm -rf packaging/deb/$(PKGNAME)/usr
	rm -f docs/$(PKGNAME).1
	rm -rf dist

docs/$(PKGNAME).1: docs/$(PKGNAME).md
	pandoc -s -t man $< > $@

# nfpm-based packaging.
SEMVER := $(shell echo $(VERSION) | sed 's/^v//')

.PHONY: deb
deb: $(TARGETS) docs/$(PKGNAME).1
	SEMVER=$(SEMVER) GOARCH=$(ARCH) nfpm package -p deb -f nfpm.yaml

.PHONY: rpm
rpm: $(TARGETS) docs/$(PKGNAME).1
	SEMVER=$(SEMVER) GOARCH=$(ARCH) nfpm package -p rpm -f nfpm.yaml

.PHONY: release
release:
	@echo export GITHUB_TOKEN="..."
	@echo go tag $(VERSION)
	@echo goreleaser release --rm-dist



