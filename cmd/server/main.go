// Package main implements a server for Instrument service.
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/abstractmachines/grpc-go/instruments"
)

const (
	port = ":50051"
)

// server is used to implement InstrumentServer
type server struct {
	pb.UnimplementedInstrumentServer
}

// GetInstrument implements InstrumentServer
func (s *server) GetInstrument(ctx context.Context, in *pb.InstrumentRequest) (*pb.InstrumentResponse, error) {
	log.Printf("Received: %v", in.GetId())
	return &pb.InstrumentResponse{Id: 456 + in.GetId()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterInstrumentServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
