PROTOC=/usr/local/bin/protoc

.PHONY: proto

proto:
	@${PROTOC} -I=/usr/include -I=${GOPATH}/src -I=. --go_out=. qtypes.proto
	@ls -al | grep "pb.go"