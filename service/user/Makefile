
GOPATH:=$(shell go env GOPATH)

ProtoOutput=./proto/user #相对的proto路径

.PHONY: proto
proto:

	protoc -I . --micro_out=$(ProtoOutput) --go_out=$(ProtoOutput) proto/user/user.proto

.PHONY: buildrun
build: proto

	go build -o user-srv.exe *.go
	./user-srv.exe

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t user-srv:latest
