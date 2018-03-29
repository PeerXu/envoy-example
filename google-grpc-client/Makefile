GOBIN := $(shell which go)

all:
	$(MAKE) -C protos all
	$(GOBIN) build -o bin/greet cmd/main.go
