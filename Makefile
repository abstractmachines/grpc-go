build_tmp=bin/

.PHONY: list
list:
	@grep '^[^#[:space:]].*:' Makefile

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: build
build:
	go build -o bin/ google.golang.org/protobuf/cmd/protoc-gen-go
	go build -o bin/ google.golang.org/grpc/cmd/protoc-gen-go-grpc

.PHONY: clean
clean:
	rm -rf ${build_tmp}

.PHONY: start-client
start-client:
	$(MAKE) -C cmd/client run

.PHONY: start-server
start-server:
	$(MAKE) proto-gen
	$(MAKE) server

.PHONY: server
server:
	$(MAKE) -C cmd/server run

.PHONY: proto-gen
proto-gen:	
	PATH="${PATH}:${PWD}/bin" protoc --go_out=. --go_opt=paths=source_relative \
	       						   --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    							   instruments.proto
