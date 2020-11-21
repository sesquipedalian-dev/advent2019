PWD := $(shell pwd)
#GOPATH := $(shell go env GOPATH)

build:
	go build -o $(PWD)/advent ./cmd/server

test:
	go test . cmd/server

.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor

run: build
	./advent