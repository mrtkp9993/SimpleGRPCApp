GOCMD=go
GOBUILD=$(GOCMD) build
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod download

OS := $(shell uname -s | awk '{print tolower($$0)}')
TAG = $$(git rev-parse --short HEAD)

BINARY_SERVER = bin/server
BINARY_CLIENT = bin/client

.PHONY: deps
deps:
		$(GOMOD)

.PHONY: bin
bin:
		$(GOBUILD) -o ${BINARY_SERVER}-$(OS)-${TAG} -v cmd/server/main.go
		$(GOBUILD) -o ${BINARY_CLIENT}-$(OS)-${TAG} -v cmd/client/main.go

.PHONY: clean
clean:
		rm -f $(BINARY_SERVER)
		rm -f $(BINARY_CLIENT)
