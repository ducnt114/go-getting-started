package main

import (
	"context"
	pb "go-getting-started/gen/go/proto"
	"google.golang.org/grpc"
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
}
