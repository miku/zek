SHELL = /bin/bash

TARGETS = zek
PKGNAME = zek
ARCH = amd64
VERSION = 0.1.14

.PHONY: all
all: $(TARGETS)

$(TARGETS): %: cmd/%/main.go
	go get -v ./...
	go build -ldflags="-s -w" -v -o $@ $<

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

.PHONY: deb
deb: $(TARGETS) docs/$(PKGNAME).1
	mkdir -p packaging/deb/$(PKGNAME)/usr/sbin
	mkdir -p packaging/deb/$(PKGNAME)/usr/share/man/man1
	cp docs/$(PKGNAME).1 packaging/deb/$(PKGNAME)/usr/share/man/man1
	find packaging/deb/$(PKGNAME)/usr -type d -exec chmod 0755 {} \;
	find packaging/deb/$(PKGNAME)/usr -type f -exec chmod 0644 {} \;
	mkdir -p packaging/deb/$(PKGNAME)/DEBIAN/
	cp $(TARGETS) packaging/deb/$(PKGNAME)/usr/sbin
	cp packaging/deb/control.$(ARCH) packaging/deb/$(PKGNAME)/DEBIAN/control
	cd packaging/deb && fakeroot dpkg-deb --build $(PKGNAME) .
	mv packaging/deb/$(PKGNAME)_*.deb .

.PHONY: rpm
rpm: $(TARGETS) docs/$(PKGNAME).1
	mkdir -p $(HOME)/rpmbuild/{BUILD,SOURCES,SPECS,RPMS}
	cp ./packaging/rpm/$(PKGNAME).spec $(HOME)/rpmbuild/SPECS
	cp $(TARGETS) $(HOME)/rpmbuild/BUILD
	cp docs/$(PKGNAME).1 $(HOME)/rpmbuild/BUILD
	./packaging/rpm/buildrpm.sh $(PKGNAME)
	cp $(HOME)/rpmbuild/RPMS/x86_64/$(PKGNAME)*.rpm .

.PHONY: update-version
update-version:
	sed -i -e 's@^const Version =.*@const Version = "$(VERSION)"@' version.go
	sed -i -e 's@^Version:.*@Version: $(VERSION)@' packaging/deb/zek/DEBIAN/control
	sed -i -e 's@^Version:.*@Version: $(VERSION)@' packaging/deb/control.amd64
	sed -i -e 's@^Version:.*@Version: $(VERSION)@' packaging/deb/control.any
	sed -i -e 's@^Version:.*@Version: $(VERSION)@' packaging/deb/control.amd64
	sed -i -e 's@^Version:.*@Version: $(VERSION)@' packaging/rpm/zek.spec

