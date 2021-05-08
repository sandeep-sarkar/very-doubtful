package main

import (
	"log"
	"net"

	"github.com/very-doubtful/api"

	pb "github.com/very-doubtful/proto/calcstatisticsb"

	"google.golang.org/grpc"
)

const (
	port = ":50061"
)

func main() {
	log.Printf("Listening to port: %s", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	statServer := &api.Server{}
	pb.RegisterStatisticsCalculatorServer(grpcServer, statServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
