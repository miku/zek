SHELL := /bin/bash
TARGETS = zek

# http://docs.travis-ci.com/user/languages/go/#Default-Test-Script
test:
	go get -d && go test -v

imports:
	goimports -w .

fmt:
	go fmt ./...

all: fmt test
	go build

clean:
	go clean
	rm -f $(TARGETS)

zek:
	go build -o zek cmd/zek/main.go
