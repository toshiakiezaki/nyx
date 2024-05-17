PWD=$(shell pwd -L)
GO_CMD=go
TZ=UTC

ifneq (,$(wildcard .env))
    include .env
    export
endif

.PHONY: all

all: help

build: build-app build-docker                  ## Build all resources

build-app: fmt                                 ## Build application
	@echo "Building application of this project..."
	@mkdir -p bin
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO_CMD) build -v -a -tags netgo -ldflags "-X main.gitVersion=$$(git rev-parse HEAD ) -w -extldflags "-static""  -o bin/nyx main.go

build-docker: fmt                              ## Build docker image
	@echo "Building docker image of this project..."
	@$(DOCKER_CMD) build -t eidolons/nyx:development -f Dockerfile .

tidy:                                          ## Download the dependencies needed to build this project
	@echo "Downloading Go dependencies needed to build this project..."
	@$(GO_CMD) mod tidy

fmt: tidy                                      ## Format code files
	@echo "Formatting code using Go code style..."
	@$(GO_CMD) fmt ./...

help:                                          ## Display help screen
	@echo "Usage:"
	@echo "	 make [COMMAND]"
	@echo "	 make help \n"
	@echo "Commands: \n"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
