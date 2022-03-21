
GOPATH:=$(shell go env GOPATH)

ProtoOutput=./proto/userOrder #相对的proto路径

.PHONY: proto
proto:

	protoc -I . --micro_out=$(ProtoOutput) --go_out=$(ProtoOutput) proto/userOrder/userOrder.proto

.PHONY: build
build: proto

	go build -o userOrder-srv *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t userOrder-srv:latest
