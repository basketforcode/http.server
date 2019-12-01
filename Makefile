.PHONY: build
build:
	go build -v ./cmd/api

.PHONY: test
test:
	go test -v -race -timeout 30s ./tests

.DEFAULT_GOAL := build