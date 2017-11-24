SHELL := /bin/bash
TARGETS = zek

zek:
	go build -o zek cmd/zek/main.go

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

