SHELL := /bin/bash

BIN_DIR := $(CURDIR)/bin
AIR := $(BIN_DIR)/air
COVER_PROFILE := coverage.out

.PHONY: tidy test cover lint dev run

tidy:
	go mod tidy

test:
	go test ./... -coverprofile=$(COVER_PROFILE) -covermode=atomic

cover: test
	go tool cover -func=$(COVER_PROFILE)

lint:
	golangci-lint run

run:
	go run ./cmd/api

dev: $(AIR)
	$(AIR)

$(AIR):
	@mkdir -p $(BIN_DIR)
	@echo "Installing air..."
	@GOBIN=$(BIN_DIR) go install github.com/air-verse/air@latest
