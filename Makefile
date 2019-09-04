# Keep a `VERSION` file in the root so that anyone
# can clearly check what's the VERSION of `master` or any
# branch at any time by checking the `VERSION` in that git
# revision.
VERSION         :=      $(shell cat .VERSION)
REGISTRY				:= 			registry.heroku.com/go-animal-api/web

test:
	go test ./... -v

image:
	docker build -t $(REGISTRY):$(VERSION) .

push_image:
	docker push $(REGISTRY):$(VERSION)

release:
	make image
	make push_image

.PHONY: test image
