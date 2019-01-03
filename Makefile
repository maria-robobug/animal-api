export GO111MODULE=on
BINARY_NAME=dogfacts

all: deps build
install:
	go install cmd/dogfacts/*.go
build:
	go build -o dogfacts cmd/dogfacts/*.go
test:
	go test -v ./...
clean:
	rm -f ./$(BINARY_NAME)
deps:
	go build -v ./...
upgrade:
	go get -u
