.PHONY: all build_linux build_darwin

VERSION ?= $(shell cat VERSION)

all:
	@echo 'DEFAULT:                                                               '
	@echo '   make build                                                    '

build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -X 'main.Version=$(VERSION)'"  -o school main.go

build_darwin:
	CGO_ENABLED=0  GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w -X 'main.Version=$(VERSION)'"  -o school main.go

test:
	go test -v ./...

lint:
	golangci-lint run

build_image:
	docker buildx build --no-cache --platform linux/amd64 -t ghcr.io/linuxoid69/school-bot:$(VERSION) .

push_image:
	docker push ghcr.io/linuxoid69/school-bot:$(VERSION)
