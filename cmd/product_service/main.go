package main

import (
	"context"
	pb "go-getting-started/gen/go/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	//pb.UnimplementedProductServiceServer
	pb.ProductServiceServer
}

func (s *server) Echo(ctx context.Context, in *pb.StringMessage) (*pb.StringMessage, error) {
	return &pb.StringMessage{Value: in.Value}, nil
}

func (s *server) GetProductDetail(ctx context.Context, in *pb.GetProductParam) (*pb.Product, error) {
	return &pb.Product{
		Id:          in.Id,
		Name:        "Product Name",
		Description: "Product Description",
		Price:       123.4,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterProductServiceServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
