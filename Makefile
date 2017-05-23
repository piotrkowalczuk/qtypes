VERSION?=$(shell git describe --tags --always --dirty)
PROTOC=/usr/local/bin/protoc

.PHONY: proto

proto:
	@${PROTOC} -I=/usr/include -I=. --go_out=. qtypes.proto
	@python -m grpc_tools.protoc -I=/usr/include -I=. --python_out=. qtypes.proto
	@ls -al | grep "pb.go"
	@ls -al | grep "_pb2"

version:
	@echo ${VERSION} > VERSION.txt