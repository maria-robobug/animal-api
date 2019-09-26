# Keep a `VERSION` file in the root so that anyone
# can clearly check what's the VERSION of `master` or any
# branch at any time by checking the `VERSION` in that git
# revision.
VERSION      :=     $(shell cat .VERSION)
BINARY_NAME  :=     animal-api-$(VERSION)
REGISTRY     :=     registry.heroku.com/go-animal-api/web

# Build for mac
build-darwin:
	go build -o $(BINARY_NAME) cmd/animal-api/main.go 
test:
	go test ./... -v

run:
	docker run -p 9000:8080 --rm -it -v $(PWD)/.env:/.env $(REGISTRY):$(VERSION)

# Docker
image:
	docker build --build-arg VERSION_ARG=$(shell cat .VERSION) -t $(REGISTRY):$(VERSION) .

push_image:
	docker push $(REGISTRY):$(VERSION)

release:
	make image
	make push_image

.PHONY: test image
