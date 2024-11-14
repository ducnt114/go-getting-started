package main

import (
	"context"
	pb "go-getting-started/gen/go/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type server struct {
	//pb.UnimplementedProductServiceServer
	pb.PaymentServiceServer
}

func (s *server) Transfer(ctx context.Context, in *pb.TransferRequest) (*pb.TransferResponse, error) {
	return &pb.TransferResponse{
		TransactionId: "random-transaction-id",
		Timestamp:     time.Now().Unix(),
		Status:        "success",
		From:          in.From,
		To:            in.To,
		Amount:        in.Amount,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPaymentServiceServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
