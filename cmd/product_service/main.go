package main

import (
	"context"
	pb "go-getting-started/gen/go/proto"
	"google.golang.org/grpc"
	"io"
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

func (s *server) ProductStream(streamServer pb.ProductService_ProductStreamServer) error {
	for {
		args, err := streamServer.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		reply := &pb.Product{
			Id:          args.GetId(),
			Name:        "server_response: " + args.GetName(),
			Description: "server_response: " + args.GetDescription(),
			Price:       args.GetPrice(),
		}

		err = streamServer.Send(reply)
		if err != nil {
			return err
		}
	}
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
