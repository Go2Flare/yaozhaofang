
GOPATH:=$(shell go env GOPATH)
ProtoOutput=./proto/house #相对的proto路径

.PHONY: proto
proto:
	protoc -I . --micro_out=$(ProtoOutput) --go_out=$(ProtoOutput) proto/house/house.proto

.PHONY: build
build: proto

	go build -o house-srv *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t house-srv:latest
