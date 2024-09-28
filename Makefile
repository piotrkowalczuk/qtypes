VERSION?=$(shell git describe --tags --always --dirty)
PROTOC=/usr/local/bin/protoc

.PHONY: proto

gen:
	@${PROTOC} -I=/usr/local/include -I=. --go_out=${GOPATH}/src *.proto
	@${PROTOC} -I=/usr/local/include -I=. --swift_opt=Visibility=Public --swift_opt=FileNaming=DropPath --swift_out=. --grpc-swift_out=. *.proto
	@#python3 -m grpc_tools.protoc -I=/usr/local/include -I=. --python_out=./qtypes qtypes.proto
	@ls -al | grep "pb.go"
	@ls -al ./qtypes | grep "_pb2"
	@ls -al | grep "pb.swift"

version:
	@echo ${VERSION} > VERSION.txt
