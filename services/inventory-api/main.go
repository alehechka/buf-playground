package main

import (
	"fmt"
	"log"
	"net"
	"os"

	inventory "github.com/alehechka/buf-playground/proto/gen/go/inventory/v1alpha1"
	inventoryserver "github.com/alehechka/buf-playground/services/inventory-api/server"
	"github.com/alehechka/buf-playground/utils"
	"github.com/alehechka/buf-playground/utils/database"
	"github.com/alehechka/buf-playground/utils/grpc_shared"
	"github.com/alehechka/buf-playground/utils/otel"
)

func main() {
	shutdownTracer, err := otel.InitializeOpenTelTracer()
	utils.Check(err)
	defer shutdownTracer()

	disconnect, err := database.InitializeMongoDB(os.Getenv("MONGODB_CONNECTION_STRING"), "inventory")
	utils.Check(err)
	defer disconnect()

	utils.Check(run())
}

func run() error {

	listener, err := net.Listen("tcp", utils.ListenOn)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", utils.ListenOn, err)
	}

	server := grpc_shared.Server()

	inventory.RegisterInventoryServiceServer(server, &inventoryserver.InventoryServiceServer{})

	log.Println("Listening on", utils.ListenOn)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}
