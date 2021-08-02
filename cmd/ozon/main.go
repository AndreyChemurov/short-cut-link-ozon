package main

import (
	"fmt"
	"log"
	"net"
	"ozon/internal/api/grpc/gen"
	linkshorter "ozon/internal/api/grpc/link_shorter"
	"ozon/internal/database"

	"google.golang.org/grpc"
)

func main() {
	// Create database table
	if err := database.CreateTableAndIndex(); err != nil {
		log.Fatal(fmt.Errorf("error occured on creating database: %w", err))
	}

	// Create gRPC server
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
