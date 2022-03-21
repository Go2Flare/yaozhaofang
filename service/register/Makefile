
GOPATH:=$(shell go env GOPATH)

ProtoOutput=./proto/register #相对的proto路径

.PHONY: proto
proto:

	protoc -I . --micro_out=$(ProtoOutput) --go_out=$(ProtoOutput) proto/register/register.proto

.PHONY: build
build: proto

	go build -o register-srv *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t register-srv:latest
