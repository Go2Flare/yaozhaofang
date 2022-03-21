
GOPATH:=$(shell go env GOPATH)

ProtoOutput=./proto/getCaptcha #相对的proto路径

.PHONY: proto
proto:
	protoc -I . --micro_out=$(ProtoOutput) --go_out=$(ProtoOutput) proto/getCaptcha/getCaptcha.proto

.PHONY: build
build: proto

	go build -o getCaptcha-srv *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t getCaptcha-srv:latest
