package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"

	gw "go-getting-started/gen/go/proto"
)

var (
	grpcProductServerEndpoint = "localhost:8081"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := gw.RegisterProductServiceHandlerFromEndpoint(ctx, mux, grpcProductServerEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Println("server listening at :8080")
	return http.ListenAndServe(":8080", mux)
}

func main() {

	if err := run(); err != nil {
		grpclog.Fatal(err)
	}
}
