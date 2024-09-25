.PHONY: all build_linux build_darwin

VERSION ?= $(shell cat VERSION)

all:
	@echo 'DEFAULT:                                                               '
	@echo '   make build                                                    '

build_linux:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -X 'main.Version=$(VERSION)'"  -o school main.go

build_darwin:
	CGO_ENABLED=0  GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w -X 'main.Version=$(VERSION)'"  -o school main.go

test:
	go test -v ./...

lint:
	golangci-lint run

build_image:
	docker buildx build --no-cache --platform linux/amd64 -t git.my-itclub.ru/bots/school:$(VERSION) .

push_image:
	docker push git.my-itclub.ru/bots/school:$(VERSION)
