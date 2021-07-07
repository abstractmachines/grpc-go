// Package main implements a client for Instrument service.
package main

import (
	"context"
	"log"
	// "os"
	// "strconv"
	"time"

	"google.golang.org/grpc"
	pb "github.com/abstractmachines/grpc-go/instruments"
)

const (
	address     = "localhost:50051"
	defaultId int32 = 123
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
	// if len(os.Args) > 1 { // later we'll care about argv.
	// 	// id = os.Args[1]
	// 	idStr := os.Args[1]
	// 	idInt, err = strconv.Atoi(idStr)
	// 	id = int32(idInt)
	// 	// TODO if err, bad, gib feedback
	// }
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetInstrument(ctx, &pb.InstrumentRequest{Id: id})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("henlo world. and the id: %s", r.GetId())
}
