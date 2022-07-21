package main

import (
	"fmt"
	"log"
	"net"

	inventory "github.com/alehechka/buf-playground/proto/gen/go/inventory/v1alpha1"
	inventoryserver "github.com/alehechka/buf-playground/services/inventory-api/server"
	"google.golang.org/grpc"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	listenOn := "127.0.0.1:8080"
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", listenOn, err)
	}

	server := grpc.NewServer()
	inventory.RegisterInventoryServiceServer(server, &inventoryserver.InventoryServiceServer{})
	log.Println("Listening on", listenOn)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}
