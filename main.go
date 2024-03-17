package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/almaraz333/finance-tracker-proto-files/expense"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedExpenseServer
}

func (s *server) CreateExpense(ctx context.Context, in *pb.CreateExenseRequest) (*pb.CreateExpenseResponse, error) {
	log.Printf("Received: %v, %v, %v", in.GetAmount(), in.GetCategory(), in.GetCreatedAt())

	return &pb.CreateExpenseResponse{
			Amount: in.GetAmount(),
		},
		nil
}

func StartServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", 50051))

	if err != nil {
		log.Fatalf("Could not lister on port: %v, with error: %v", 50051, err)
	}

	s := grpc.NewServer()
	pb.RegisterExpenseServer(s, &server{})
	log.Printf("Server listening on port %v ...", 50051)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}

func main() {
	StartServer()
}
