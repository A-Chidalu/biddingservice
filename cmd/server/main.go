package main

// The entry point to the  BidServer.
// This will start and run continiously

import (
	"flag"
	"fmt"
	pb "github.com/A-Chidalu/biddingservice/api/proto"
	"github.com/A-Chidalu/biddingservice/internal/database"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

type BidServer struct {
	pb.UnimplementedBidServer
}

// PlaceBid This is a function the BidServer implements
func (s *BidServer) PlaceBid(ctx context.Context, req *pb.BidRequest) (*pb.BidResponse, error) {
	bid, err := database.SaveBid(req)

	if err != nil {
		log.Fatalf("There was an error in saving the request %v with error %v", req, err)
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Could not place bid. Reason: %v", err))
	}

	response := &pb.BidResponse{BidId: bid.ID}

	return response, nil

}

func (s *BidServer) GetWinningBidder(ctx context.Context, req *pb.BidWinnerRequest) (*pb.BidWinnerResponse, error) {
	bid, err := database.GetLatestBidForItem(req.GetItemId())

	if err != nil {
		log.Fatalf("There was an error getting the winning bidder for the request: %v %v", req, err)
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Could not get bid winner. Reason: %v", err))
	}

	response := &pb.BidWinnerResponse{UserId: bid.UserID, ItemId: bid.ItemID}

	return response, nil

}

func main() {

	//Open a connection to the database
	err := database.ConnectDB()
	if err != nil {
		panic("Could not connect to the database")
	}

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("grpcServer:8080"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterBidServer(grpcServer, &BidServer{})
	grpcServer.Serve(lis)

	defer database.CloseDB()

}
