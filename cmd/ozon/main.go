package main

import (
	"fmt"
	"log"
	"net"
	"ozon/internal/api/grpc/gen"
	linkshorter "ozon/internal/api/grpc/link_shorter"

	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	server := &linkshorter.GRPCServer{}

	gen.RegisterLinkShorterServer(grpcServer, server)

	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(fmt.Errorf("error occurred on listening tcp:8000: %w", err))
	}

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(fmt.Errorf("error occurred on serving gRPC server: %w", err))
	}
}
