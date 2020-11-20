PWD := $(shell pwd)
#GOPATH := $(shell go env GOPATH)

build:
	@GO111MODULE=on go build -mod=vendor -o $(PWD)/advent cmd/server/main.go

test:
	go test -mod=vendor . cmd/server

.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor
