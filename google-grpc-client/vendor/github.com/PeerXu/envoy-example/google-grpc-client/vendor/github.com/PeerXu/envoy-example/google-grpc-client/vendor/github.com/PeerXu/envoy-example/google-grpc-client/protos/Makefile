GRPC_FLAGS := -I. -I/usr/local/include -I$(GOPATH)/src -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis

PROTO_FILES := greet.proto

all:
	protoc $(GRPC_FLAGS) \
	--go_out=plugins=grpc:. \
	$(PROTO_FILES)

	protoc $(GRPC_FLAGS) \
	--grpc-gateway_out=logtostderr=true:. \
	$(PROTO_FILES)
