package main

import (
	"context"
	pb "github.com/A-Chidalu/biddingservice/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:5003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	client := pb.NewBidClient(conn)

	response, err := client.PlaceBid(context.Background(), &pb.BidRequest{
		UserId:  1,
		ItemId:  "89487gbiuaybs1872g817r",
		Amount:  200,
		BidType: pb.BID_TYPE_FORWARD,
	})

	client.PlaceBid(context.Background(), &pb.BidRequest{
		UserId:  1,
		ItemId:  "89487gbiuaybs1872g817r",
		Amount:  201,
		BidType: pb.BID_TYPE_FORWARD,
	})

	client.PlaceBid(context.Background(), &pb.BidRequest{
		UserId:  1,
		ItemId:  "89487gbiuaybs1872g817r",
		Amount:  202,
		BidType: pb.BID_TYPE_FORWARD,
	})

	client.PlaceBid(context.Background(), &pb.BidRequest{
		UserId:  1,
		ItemId:  "89487gbiuaybs1872g817r",
		Amount:  203,
		BidType: pb.BID_TYPE_FORWARD,
	})

	client.PlaceBid(context.Background(), &pb.BidRequest{
		UserId:  5,
		ItemId:  "89487gbiuaybs1872g817r",
		Amount:  204,
		BidType: pb.BID_TYPE_FORWARD,
	})

	log.Printf("The response is %v", response.BidId)

	response1, err := client.GetWinningBidder(context.Background(), &pb.BidWinnerRequest{
		ItemId: "89487gbiuaybs1872g817r",
	})

	log.Printf("The winning Bidder is is %v", response1.UserId)

	defer conn.Close()

}
