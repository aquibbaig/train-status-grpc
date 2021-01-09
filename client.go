package main

import (
	"github.com/aquibbaig/train-status-grpc/enquiry"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var (
		conn *grpc.ClientConn
	)
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not establish client connection %v", err)
	}
	defer conn.Close()

	cl := enquiry.NewTrainStatusClient(conn)
	message := enquiry.Message{
		TrainNumber: "242",
	}
	response, err := cl.ReturnTrainDetails(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error while fetching response from grpc server: %v", err)
	}

	log.Printf("Response: %v", response)
}
