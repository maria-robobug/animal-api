# Keep a `VERSION` file in the root so that anyone
# can clearly check what's the VERSION of `master` or any
# branch at any time by checking the `VERSION` in that git
# revision.
#
# Another benefit is that we can pass this file to our Docker
# build context and have the version set in the binary that ends
# up inside the Docker image too.
VERSION         :=      $(shell cat .VERSION)
IMAGE_NAME      :=      mariarobobug/animal-api

test:
	go test ./... -v

image:
	docker build -t $(IMAGE_NAME):$(VERSION) .
	docker tag ${IMAGE_NAME}:${VERSION} ${IMAGE_NAME}:latest

push_image:
	docker push ${IMAGE_NAME}:${VERSION}
	docker push ${IMAGE_NAME}:latest

release:
	make push_image
	./heroku_deploy.sh

.PHONY: test image release
