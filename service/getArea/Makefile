
GOPATH:=$(shell go env GOPATH)
ProtoOutput=./proto/getArea #相对的proto路径

.PHONY: proto
proto:
	protoc -I . --micro_out=$(ProtoOutput) --go_out=$(ProtoOutput) proto/getArea/getArea.proto

.PHONY: build
build: proto

	go build -o getArea-srv *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t getArea-srv:latest
