.PHONY: install dist clean build image

BUILD_VERSION ?= $(shell git describe --tags)
BUILD_FLAGS := -ldflags "-X main.Version=${BUILD_VERSION}"
DOCKER_BUILD_FLAGS := -a -installsuffix cgo
BUILDOUT ?= tallybox
IMAGE_NAME = "confio/tallybox"
IMAGE_VERSION = "$(IMAGE_NAME):${BUILD_VERSION}"

install:
	go install $(BUILD_FLAGS) .

dist: clean build image

clean:
	rm -f ${BUILDOUT}

build:
	GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build $(BUILD_FLAGS) $(DOCKER_BUILD_FLAGS) -o $(BUILDOUT) .

image:
	docker build --pull -t $(IMAGE_NAME) .
	docker tag $(IMAGE_NAME) $(IMAGE_VERSION)
	@echo "Built $(IMAGE_VERSION) as $(IMAGE_NAME):latest"
