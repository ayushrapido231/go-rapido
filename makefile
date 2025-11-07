.DEFAULT_GOAL := build

.PHONY: fmt vet build run
fmt:
	cd helloworld && make fmt

vet: fmt
	cd helloworld && make vet

build: vet
	cd helloworld && make build

run:
	cd helloworld && make run

