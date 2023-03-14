package main

import (
	"context"
	pb "github.com/A-Chidalu/biddingservice/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	client := pb.NewBidClient(conn)

	response, err := client.PlaceBid(context.Background(), &pb.BidRequest{
		UserId:        2,
		ItemId:        5,
		Amount:        20201,
		IsTerminating: false,
		BidType:       pb.BID_TYPE_FORWARD,
	})

	log.Printf("The response is %v", response.BidId)

	defer conn.Close()

}
