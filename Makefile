export
GO111MODULE=on
BINARY_NAME=animal-api
ARCH=amd64

BIN_LINUX=$(BINARY_NAME)-linux-$(ARCH)
BIN_DARWIN=$(BINARY_NAME)-darwin-$(ARCH)

SOURCE=cmd/$(BINARY_NAME)/main.go

.PHONY: test clean all

all: upgrade deps build-linux

build-linux:
	GOARCH=$(ARCH) GOOS=linux CGO_ENABLED=0 go build -o bin/$(BIN_LINUX) $(SOURCE)

build-darwin:
	GOARCH=$(ARCH) GOOS=darwin go build -o bin/$(BIN_DARWIN) $(SOURCE)

test:
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

clean:
	rm -f $(BIN_DARWIN)
	rm -f $(BIN_LINUX)
	rm -f coverage.out

deps:
	go build -v ./...

upgrade:
	go get -u
