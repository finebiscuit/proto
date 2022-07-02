.PHONY: all
all: biscuit biscuitdata

.PHONY: biscuit
biscuit:
	buf generate biscuit --template biscuit/buf.gen.yaml -o biscuit

.PHONY: biscuitdata
biscuitdata:
	buf generate biscuitdata --template biscuitdata/buf.gen.yaml -o biscuitdata

.PHONY: deps
deps:
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest

