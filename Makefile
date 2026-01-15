.PHONY: all test lint build clean example

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=ecash-sdk

all: test build

build:
	$(GOBUILD) -v ./...

test:
	$(GOTEST) -v ./...

lint:
	# Assuming golangci-lint is installed
	golangci-lint run

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

example:
	@echo "Running example..."
	@go run examples/simple_transfer/main.go

deps:
	$(GOGET) github.com/google/uuid
	$(GOGET) github.com/stretchr/testify
