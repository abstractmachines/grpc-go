# grpc-go

## Runbook

### Go Modules go mod init 
- See: blog.golang.org/using-go-modules

- `go mod init github.com/abstractmachines/grpc-go/instruments`
- `go list -m all`


## TODO 
import (
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)