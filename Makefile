GOPATH ?= $(shell go env GOPATH)
BUILD_DIR ?= build
OUTPUT ?= $(BUILD_DIR)/bcrypt-me
.PHONY: build install test lint clean
.DEFAULT_GOAL := build

build:
	GO111MODULE=on go build -o $(OUTPUT) ./cmd/bcrypt-me

install:
	GO111MODULE=on go install ./cmd/bcrypt-me

test:
	GO111MODULE=on go test -cover -race ./...

$(GOPATH)/bin/golangci-lint:
	GO111MODULE=off go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

lint: $(GOPATH)/bin/golangci-lint
	GO111MODULE=on $(GOPATH)/bin/golangci-lint run ./...

clean:
	rm -rf $(BUILD_DIR)
