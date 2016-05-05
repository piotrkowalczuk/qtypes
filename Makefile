PROTOC=/usr/local/bin/protoc

.PHONY: proto

proto:
	@${PROTOC} --proto_path=${GOPATH}/src --proto_path=. --go_out=. qtypes.proto
	@ls -al | grep "pb.go"