.DEFAULT_GOAL := build

.PHONY: fmt vet build run
fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build

run:
	go run main.go

