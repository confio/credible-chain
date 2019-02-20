.PHONY: all build test image tf protoc clean dist

BUILD_VERSION ?= $(shell git describe --tags)
BUILD_FLAGS := -ldflags "-X main.Version=${BUILD_VERSION}"
DOCKER_BUILD_FLAGS := -a -installsuffix cgo
BUILDOUT ?= credible-chain
IMAGE_NAME = "confio/credible-chain"
IMAGE_VERSION = "confio/credible-chain:${BUILD_VERSION}"

### Basic

all: deps test install

install:
	go install $(BUILD_FLAGS) .

# test never caches and checks for race conditons
test:
	go test -count=1 -race ./...

# tf runs the quick version of the tests
tf:
	go test -short ./...

### Docker

dist: clean test build image

build:
	GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build $(BUILD_FLAGS) $(DOCKER_BUILD_FLAGS) -o $(BUILDOUT) .

clean:
	rm -f ${BUILDOUT}

image:
	@echo "Only builds docker image if all changes have been committed\n"
	@git diff-index --quiet HEAD
	docker build --pull -t $(IMAGE_NAME) .
	docker tag $(IMAGE_NAME) $(IMAGE_VERSION)
	@echo "Built $(IMAGE_VERSION) as $(IMAGE_NAME):latest"

### Tools

deps: tools
	@rm -rf vendor/
	dep ensure -vendor-only

tools:
	@go get github.com/golang/dep/cmd/dep

protoc: prototools
	protoc --gogofaster_out=. -I=. -I=./vendor -I=$(GOPATH)/src app/*.proto
	protoc --gogofaster_out=. -I=. -I=./vendor -I=$(GOPATH)/src x/votes/*.proto

### cross-platform check for installing protoc ###

MYOS := $(shell uname -s)

ifeq ($(MYOS),Darwin)  # Mac OS X
	ZIP := protoc-3.4.0-osx-x86_64.zip
endif
ifeq ($(MYOS),Linux)
	ZIP := protoc-3.4.0-linux-x86_64.zip
endif

/usr/local/bin/protoc:
	@ curl -L https://github.com/google/protobuf/releases/download/v3.4.0/$(ZIP) > $(ZIP)
	@ unzip -q $(ZIP) -d protoc3
	@ rm $(ZIP)
	sudo mv protoc3/bin/protoc /usr/local/bin/
	@ sudo mv protoc3/include/* /usr/local/include/
	@ sudo chown `whoami` /usr/local/bin/protoc
	@ sudo chown -R `whoami` /usr/local/include/google
	@ rm -rf protoc3

prototools: /usr/local/bin/protoc deps
	# install all tools from our vendored dependencies
	@go install ./vendor/github.com/gogo/protobuf/proto
	@go install ./vendor/github.com/gogo/protobuf/gogoproto
	@go install ./vendor/github.com/gogo/protobuf/protoc-gen-gogofaster


