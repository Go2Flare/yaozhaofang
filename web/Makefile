
GOPATH:=$(shell go env GOPATH)

.PHONY: build docker


build:
	go build -o yzf-web main.go


docker:
	docker build . -t yzf-web:latest

run: build docker