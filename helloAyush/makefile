.DEFAULT_GOAL := build

.PHONY: fmt vet build run deps
fmt:
	go fmt ./...

vet: fmt
	go vet ./...

deps:
	go mod download
	go mod tidy

build: deps vet
	go build -o helloayush

run: build
	./helloayush


