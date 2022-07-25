package main

import (
	"fmt"
	"log"
	"net"

	inventory "github.com/alehechka/buf-playground/proto/gen/go/inventory/v1alpha1"
	inventoryserver "github.com/alehechka/buf-playground/services/inventory-api/server"
	"github.com/alehechka/buf-playground/utils"
	"google.golang.org/grpc"
)

func main() {
	utils.Check(run())
}

func run() error {
	listener, err := net.Listen("tcp", utils.ListenOn)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", utils.ListenOn, err)
	}

	server := grpc.NewServer()
	inventory.RegisterInventoryServiceServer(server, &inventoryserver.InventoryServiceServer{})
	log.Println("Listening on", utils.ListenOn)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}
