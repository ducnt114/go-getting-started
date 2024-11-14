package main

import (
	"context"
	"fmt"
	pb "go-getting-started/gen/go/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewProductServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.Echo(ctx, &pb.StringMessage{Value: "Hello, World!"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Response: %s", r.GetValue())

	// test stream
	stream, err := client.ProductStream(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			if err := stream.Send(&pb.Product{
				Id:          time.Now().Unix(),
				Name:        "Product Name",
				Description: "Product Description",
				Price:       123.4,
			}); err != nil {
				log.Fatal(err)
			}
			time.Sleep(1 * time.Second)
		}
	}()
	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(reply.GetId())
		fmt.Println(reply.GetName())
		fmt.Println(reply.GetDescription())
		fmt.Println(reply.GetPrice())
	}

	time.Sleep(5 * time.Second)
}
