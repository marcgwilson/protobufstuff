SRC_DIR=rpc
DST_DIR=rpc

all: rpc test

.PHONY: rpc test

test:
	go test -v github.com/marcgwilson/protobufstuff/...

rpc:
		protoc --proto_path=/usr/local/include -I=$(SRC_DIR)/state --go_out=plugins=grpc:$(DST_DIR)/state $(SRC_DIR)/state/state.proto
		protoc --proto_path=/usr/local/include -I=$(SRC_DIR)/valuescache --go_out=plugins=grpc:$(DST_DIR)/valuescache $(SRC_DIR)/valuescache/valuescache.proto

default: test
