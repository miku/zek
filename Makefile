SHELL = /bin/bash

TARGETS = zek
PKGNAME = zek
ARCH = amd64 # $$(dpkg --print-architecture)

all: $(TARGETS)

$(TARGETS): %: cmd/%/main.go
	go get -v ./...
	go build -ldflags="-s -w" -v -o $@ $<

clean:
	rm -f $(TARGETS)
	rm -f $(PKGNAME)*.deb
	rm -f $(PKGNAME)-*.rpm
	rm -rf packaging/deb/$(PKGNAME)/usr

deb: $(TARGETS)
	mkdir -p packaging/deb/$(PKGNAME)/usr/sbin
	cp $(TARGETS) packaging/deb/$(PKGNAME)/usr/sbin
	# md2man-roff $(PKGNAME).md | gzip -n -9 -c > $(PKGNAME).1.gz
	# mkdir -p packaging/deb/$(PKGNAME)/usr/share/man/man1
	# cp docs/$(PKGNAME).1.gz packaging/deb/$(PKGNAME)/usr/share/man/man1
	find packaging/deb/$(PKGNAME)/usr -type d -exec chmod 0755 {} \;
	find packaging/deb/$(PKGNAME)/usr -type f -exec chmod 0644 {} \;
	mkdir -p packaging/deb/$(PKGNAME)/DEBIAN/
	cp packaging/deb/control.$(ARCH) packaging/deb/$(PKGNAME)/DEBIAN/control
	cd packaging/deb && fakeroot dpkg-deb --build $(PKGNAME) .
	mv packaging/deb/$(PKGNAME)_*.deb .

rpm: $(TARGETS)
	mkdir -p $(HOME)/rpmbuild/{BUILD,SOURCES,SPECS,RPMS}
	cp ./packaging/rpm/$(PKGNAME).spec $(HOME)/rpmbuild/SPECS
	cp $(TARGETS) $(HOME)/rpmbuild/BUILD
	# cp docs/$(PKGNAME).1.gz $(HOME)/rpmbuild/BUILD
	./packaging/rpm/buildrpm.sh $(PKGNAME)
	cp $(HOME)/rpmbuild/RPMS/x86_64/$(PKGNAME)*.rpm .