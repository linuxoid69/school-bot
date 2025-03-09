.PHONY: all build_linux build_darwin

VERSION ?= $(shell cat VERSION)
APP=school-bot
BUILD_CMD='GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -X 'main.Version=$(VERSION)'" -o $(APP)'
GROUP=linuxoid69
DOCKER_REGISTRY=ghcr.io
GOLANG_VERSION=$(shell cat go.mod | grep ^go | cut -d " " -f 2)
REGISTRY_USER=$(GROUP)

all:
	@echo 'DEFAULT:        '
	@echo 'make build_linux - build linux binary'
	@echo 'make test        - run tests'
	@echo 'make lint        - run lint'
	@echo 'make build_image - build docker image'
	@echo 'make push_image  - push docker image'

build_linux:
	$(shell echo $(BUILD_CMD))

test:
	go test -v ./...

lint:
	golangci-lint run

build_image:
	docker buildx build --no-cache --platform linux/amd64 \
						--build-arg BUILD_CMD=$(BUILD_CMD) \
						--build-arg GOLANG_VERSION=$(GOLANG_VERSION) \
						-t $(DOCKER_REGISTRY)/$(GROUP)/$(APP):$(VERSION) .

push_image:
ifeq ($(CI), true)
	docker login $(DOCKER_REGISTRY) -u $(REGISTRY_USER) -p $(REGISTRY_TOKEN)
endif
	docker push $(DOCKER_REGISTRY)/$(GROUP)/$(APP):$(VERSION)
