BASE_PATH = $(shell pwd)
export PROJECT_ROOT := $(BASE_PATH)
export PATH := $(BASE_PATH)/bin:$(PATH)
export GOBIN := $(BASE_PATH)/bin

SHELL := env PATH=$(PATH) /bin/bash

# Commands
GOCMD=go
GORUN=$(GOCMD) run
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test
GOFMT=$(GOCMD) fmt
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGENERATE=$(GOCMD) generate

BINARY_NAME=gomon
BINARY_PATH=$(GOBIN)/$(BINARY_NAME)

# all the packages without vendor
ALL_PKGS = $(shell go list ./... | grep -v /vendor | grep -v /pkg/grpcapi)

# Colors
GREEN_COLOR   = "\033[0;32m"
PURPLE_COLOR  = "\033[0;35m"
DEFAULT_COLOR = "\033[m"

.PHONY: all help clean test lint fmt build

all: clean fmt lint build_test_apps build test

help:
	@echo 'Usage: make <TARGETS> ... <OPTIONS>'
	@echo ''
	@echo 'Available targets are:'
	@echo ''
	@echo '    help               Show this help screen.'
	@echo '    clean              Clean all the artifacts.'
	@echo '    test               Run unit tests.'
	@echo '    lint               Run all linters.'
	@echo '    fmt                Run gofmt on package sources.'
	@echo '    build              Compile packages and dependencies.'
	@echo '    generate           Generate step'
	@echo ''

clean:
	@echo -e [$(GREEN_COLOR)clean$(DEFAULT_COLOR)]
	@$(GOCLEAN)
	@rm -rf $(GOBIN)

test: build_test_apps
	@echo -e [$(GREEN_COLOR)test$(DEFAULT_COLOR)]
	@$(GOTEST) -v -race -count=1 ./...

lint:
	@echo -e [$(GREEN_COLOR)lint$(DEFAULT_COLOR)]
	@$(GOINSTALL) github.com/golangci/golangci-lint/cmd/golangci-lint
	@$(GOBIN)/golangci-lint run

fmt:
	@echo -e [$(GREEN_COLOR)format$(DEFAULT_COLOR)]
	@$(GOFMT) $(ALLPKGS)

build:
	@echo -e [$(GREEN_COLOR)build$(DEFAULT_COLOR)]
	@mkdir -p ./bin
	@$(GOBUILD) -o $(BINARY_PATH)

build_test_apps:
	@echo -e [$(GREEN_COLOR)build test apps$(DEFAULT_COLOR)]
	@mkdir -p ./bin
	@$(GOBUILD) -o $(GOBIN)/test-app-args ./cmd/test-app-args/
	@$(GOBUILD) -o $(GOBIN)/test-app-kill ./cmd/test-app-kill/

generate:
	@mkdir -p ./bin
	@echo -e $(PURPLE_COLOR)[building mockery]$(DEFAULT_COLOR)
	@$(GOINSTALL) github.com/vektra/mockery/cmd/mockery
	@echo -e $(PURPLE_COLOR)[mockery built]$(DEFAULT_COLOR)
	@echo -e [$(GREEN_COLOR)generate$(DEFAULT_COLOR)]
	@$(GOGENERATE) $(ALL_PKGS)
