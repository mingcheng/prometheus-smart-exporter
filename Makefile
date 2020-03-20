.PHONY: all build clean test packr2

VERSION=0.0.1
BIN=./smart-exporter
DIR_SRC=./cmd/smart-exporter

GO_ENV=CGO_ENABLED=0
GO_FLAGS=-ldflags="-X main.version=$(VERSION) -X 'main.buildTime=`date`' -extldflags -static"
GO=$(GO_ENV) $(shell which go)
GOROOT=$(shell `which go` env GOROOT)
GOPATH=$(shell `which go` env GOPATH)

all: clean build

packr2:
	@go get -u github.com/gobuffalo/packr/v2/...
	@go get -u github.com/gobuffalo/packr/v2/packr2

build: packr2 $(DIR_SRC)
	@$(GOPATH)/bin/packr2 build
	@$(GO) build $(GO_FLAGS) -o $(BIN) $(DIR_SRC)

clean: packr2
	@packr2 clean
	@rm -f $(BIN)
	@$(GO) clean ./...

test:
	@$(GO) test ./...
