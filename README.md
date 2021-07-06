# grpc-go

## Runbook

> Generating code with protoc compiler

- Rather than using `$GOBIN` and `$GOPATH` for the binaries, 
install them locally in the project root / module you're working in,
so that you can ensure that the versions being used are correct:

```
go build -o bin/ google.golang.org/protobuf/cmd/protoc-gen-go
```

```
go build -o bin/ google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

```
PATH="${PATH}:${PWD}/bin" protoc --go_out=. --go_opt=paths=source_relative \
                                                           --go-grpc_out=. --go-grpc_opt=paths=source_relative \
                                                           instruments.proto
```

### Go Modules go mod init 
- See: blog.golang.org/using-go-modules

- `go mod init github.com/abstractmachines/grpc-go/instruments`
- `go list -m all`


## TODO 
import (
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
