package main

import (
	"fmt"
	"log"
	"net"
	"os"

	inventory "github.com/alehechka/buf-playground/proto/gen/go/inventory/v1alpha1"
	inventoryserver "github.com/alehechka/buf-playground/services/inventory-api/server"
	"github.com/alehechka/buf-playground/utils"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

func main() {
	shutdownTracer, err := utils.InitializeOpenTelTracer()
	utils.Check(err)
	defer shutdownTracer()

	disconnect, err := utils.InitializeMongoDB(os.Getenv("MONGODB_CONNECTION_STRING"), "inventory")
	utils.Check(err)
	defer disconnect()

	utils.Check(run())
}

func run() error {

	listener, err := net.Listen("tcp", utils.ListenOn)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", utils.ListenOn, err)
	}

	server := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
		grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
	)

	inventory.RegisterInventoryServiceServer(server, &inventoryserver.InventoryServiceServer{})
	log.Println("Listening on", utils.ListenOn)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}
