package main

// The entry point to the ForwardBidServer.
// This will start and run continiously

import (
	"flag"
	"fmt"
	pb "github.com/A-Chidalu/forwardbiddingservice/api/proto"
	"github.com/A-Chidalu/forwardbiddingservice/internal/database"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

type ForwardBidServer struct {
	pb.UnimplementedForwardBidServer
}

// PlaceBid This is a function the ForwardBidServer implements
func (s *ForwardBidServer) PlaceBid(ctx context.Context, req *pb.ForwardBidRequest) (*pb.ForwardBidResponse, error) {
	bid, err := database.SaveForwardBid(req)

	if err != nil {
		log.Fatalf("There was an error in saving the request %v with error %v", req, err)
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Could not place bid. Reason: %v", err))
	}

	response := &pb.ForwardBidResponse{BidId: bid.ID}

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
	pb.RegisterForwardBidServer(grpcServer, &ForwardBidServer{})
	grpcServer.Serve(lis)

	defer database.CloseDB()

}
