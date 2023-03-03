package main

import (
	"context"
	pb "github.com/A-Chidalu/forwardbiddingservice/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	client := pb.NewForwardBidClient(conn)

	response, err := client.PlaceBid(context.Background(), &pb.ForwardBidRequest{
		UserId:        2,
		ItemId:        3,
		Amount:        1000.1,
		IsTerminating: false,
	})

	log.Printf("The response is %v", response.BidId)

	defer conn.Close()

}
