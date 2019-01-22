export
GO111MODULE=on
BINARY_NAME=animal-api
ARCH=amd64

HEROKU_APP_NAME=animal-facts-api
BIN_LINUX=$(BINARY_NAME)-linux-$(ARCH)
BIN_DARWIN=$(BINARY_NAME)-darwin-$(ARCH)

SOURCE=cmd/$(BINARY_NAME)/main.go

VERSION=$(shell ./version.sh)

.PHONY: test clean all

all: upgrade deps

build-linux:
	GOARCH=$(ARCH) GOOS=linux CGO_ENABLED=0 go build -o bin/$(BIN_LINUX) $(SOURCE)

build-darwin:
	GOARCH=$(ARCH) GOOS=darwin go build -o bin/$(BIN_DARWIN) $(SOURCE)

test:
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out -o coverage.html
	rm -f coverage.out

tag:
	git tag -a $(VERSION) -m "new release candidate"
	git push origin $(VERSION)

clean:
	rm -fr bin
	rm -f coverage.html

deps:
	go build -v ./...

upgrade:
	go get -u
