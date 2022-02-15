GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install
BINARY_NAME=square
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build
build:
		$(GOBUILD) -o bin/$(BINARY_NAME) -v ./cmd/square
test:
		$(GOTEST) ./...
clean:
		$(GOCLEAN)
		rm -f bin/$(BINARY_NAME)
		rm -f bin/$(BINARY_UNIX)
install:
		$(GOINSTALL)
