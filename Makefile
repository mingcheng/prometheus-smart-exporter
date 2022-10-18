.PHONY: all build clean test pkger install

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
	@go install github.com/markbates/pkger/cmd/pkger@latest

build: pkger
	@$(GOPATH)/bin/pkger -include /scripts
	@$(GO) build $(GO_FLAGS) -o $(BIN) $(DIR_SRC)

clean: pkger
	@rm -f $(BIN) ./pkged.go
	@$(GO) clean ./...

install: $(BIN)
	@install -m 0755 $(BIN) /usr/local/bin/
	@if [ -d "/etc/systemd/system" ]; then install -m 0644 ./systemd/prometheus-smart-exporter.service /etc/systemd/system; fi
	@if [ -d "/etc/supervisor.d/" ]; then install -m 0644 ./supervisor/prometheus-smart-exporter.ini /etc/supervisor.d; fi

test: build
	@$(GO) test ./...
	@./scripts/test.sh
