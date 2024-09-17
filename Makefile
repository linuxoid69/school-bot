.PHONY: all build_linux build_darwin

all:
	@echo 'DEFAULT:                                                               '
	@echo '   make build                                                    '

build_linux:
	CGO_ENABLED=0  GOOS=linux GOARCH=amd64 go build -ldflags="-s -w"  -o school main.go

build_darwin:
	CGO_ENABLED=0  GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w"  -o school main.go

test:
	go test -v ./...

lint:
	golangci-lint run
