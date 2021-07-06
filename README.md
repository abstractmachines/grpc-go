# grpc-go

gRPC, protobufs, Go

## References:
- [Google Developers Protobuf Tutorial] (https://developers.google.com/protocol-buffers/docs/gotutorial)
- [gRPC quickstart](https://grpc.io/docs/languages/go/quickstart/)

## Working with Protobufs (high level)

1. Define data structures in a .proto file.
2. Compile that file to generate server methods and related client stubs for those same methods.
3. Profit.

## Motivation

> Distributed Systems and Services

With gRPC, a client can "directly" call a method on a server, as if it were a local object. _The client has "stubs" of the server's methods._

This makes creating distributed applications and services a lot easier than thinking of clients and servers/services as having separate "methods."

> Performance

When we use the `protoc` compiler, that generates data access methods (getters, setters et al). 

_The protobuf compiler encodes into an efficient binary format._ 

_The generated code takes care of the underlying details of reading/writing to/from services for us._

## Runbook

> Generating code with protoc compiler

1. Install protoc binaries for go and go grpc:

	- Rather than using `$GOBIN` and `$GOPATH` for the binaries, 
	install them locally in the project root / module you're working in,
	so that you can ensure that the versions being used are correct:

	```
	go build -o bin/ google.golang.org/protobuf/cmd/protoc-gen-go
	```

	```
	go build -o bin/ google.golang.org/grpc/cmd/protoc-gen-go-grpc
	```

	> Expected result: you now have a new `/bin` dir in your project 
	> with the binaries `protoc-gen-go*      protoc-gen-go-grpc*`

2. Compile the protoc file, generating Go code:

	```
	PATH="${PATH}:${PWD}/bin" protoc --go_out=. --go_opt=paths=source_relative \
															--go-grpc_out=. --go-grpc_opt=paths=source_relative \
															instruments.proto
	```
	> Expected result: You'll have a `*_grpc.pb.go` file and `*.pb.go` file

### Go Modules go mod init 
- See: blog.golang.org/using-go-modules

- `go mod init github.com/abstractmachines/grpc-go/instruments`
- `go list -m all`


## TODO 
import (
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
