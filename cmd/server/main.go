package main

// The entry point to the ForwardBidServer.
// This will start and run continiously

import (
	"fmt"
	pb "github.com/A-Chidalu/forwardbiddingservice/api/proto"
	"github.com/A-Chidalu/forwardbiddingservice/internal/database"
	"golang.org/x/net/context"
)

type ForwardBidServer struct {
	pb.UnimplementedForwardBidServer
}

func (s *ForwardBidServer) PlaceBid(ctx context.Context, req *pb.ForwardBidRequest) {

}

func main() {

	//Open a connection to the database
	err := database.ConnectDB()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connected to the dbConn successfully!!")

		request := pb.ForwardBidRequest{
			UserId:        1,
			ItemId:        2,
			Amount:        78.0,
			IsTerminating: false,
		}

		bid, err := database.SaveForwardBid(&request)
		if err != nil {
			fmt.Println("THERE WAS AN ERROR", err)
		} else {
			fmt.Println("The new bid that was created is ", bid.ID)
		}

		bids, err := database.GetAllBids()

		if err != nil {
			fmt.Println("THERE WAS AN ERROR", err)
		} else {
			for _, bid := range bids {
				fmt.Printf("Bid with amount %g \n", bid.Amount)
			}
		}
	}

	defer database.CloseDB()

}
