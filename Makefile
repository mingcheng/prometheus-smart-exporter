.PHONY: all build clean test pkger

VERSION=0.0.2
BIN=./smart-exporter
DIR_SRC=./cmd/smart-exporter

GO_ENV=CGO_ENABLED=0
GO_FLAGS=-ldflags="-X main.version=$(VERSION) -X 'main.buildTime=`date`' -extldflags -static"
GO=$(GO_ENV) $(shell which go)
GOROOT=$(shell `which go` env GOROOT)
GOPATH=$(shell `which go` env GOPATH)

all: clean build

pkger:
	@go get github.com/markbates/pkger/cmd/pkger

build: pkger
	@$(GOPATH)/bin/pkger -include /scripts
	@$(GO) build $(GO_FLAGS) -o $(BIN) $(DIR_SRC)

clean: pkger
	@rm -f $(BIN) ./pkged.go
	@$(GO) clean ./...

test: build
	@$(GO) test ./...
	@./scripts/test.sh
