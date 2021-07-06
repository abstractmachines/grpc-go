// Package main implements a client for Instrument service.
package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	// TODO fix
	pb "github.com/abstractmachines/grpc-go"
)

const (
	address     = "localhost:50051"
	defaultId = 123
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewInstrumentClient(conn)

	// Contact the server and print out its response.
	id := defaultId
	if len(os.Args) > 1 {
		id = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetInstrument(ctx, &pb.HelloRequest{Id: id})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
