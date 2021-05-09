package main

import (
	"log"
	"net"
	"net/http"

	"github.com/very-doubtful/api"
	health "github.com/very-doubtful/health"
	api_health "google.golang.org/grpc/health/grpc_health_v1"

	pb "github.com/very-doubtful/proto/calcstatisticsb"

	"google.golang.org/grpc"
)

const (
	port = ":50061"
)

func fileServer(port string) {
	log.Printf("Starting file server on port %s", port)
	err := http.ListenAndServe(port, http.FileServer(http.Dir("result")))
	if err != nil {
		log.Fatalf("Error handling request :%v", err)
	}
}

func main() {
	log.Printf("Listening to port: %s", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	go fileServer(":50071")

	grpcServer := grpc.NewServer()

	//health server
	healthServ := health.NewHealthCheckService()
	api_health.RegisterHealthServer(grpcServer, healthServ)

	statServer := &api.Server{}
	pb.RegisterStatisticsCalculatorServer(grpcServer, statServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
