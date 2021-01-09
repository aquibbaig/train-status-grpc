package main

import (
	"github.com/aquibbaig/train-status-grpc/enquiry"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Could not listen at PORT 9000 %v", err)
	}
	grpcServer := grpc.NewServer()

	// enquiry server struct.
	serv := enquiry.Server{}

	// Register the server.
	enquiry.RegisterTrainStatusServer(grpcServer, &serv)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error creating gRPC server %v", err)
	}
}
