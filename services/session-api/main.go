package main

import (
	"fmt"
	"log"
	"net"
	"os"

	session "github.com/alehechka/buf-playground/proto/gen/go/session/v1alpha1"
	sessionserver "github.com/alehechka/buf-playground/services/session-api/server"
	"github.com/alehechka/buf-playground/utils"
	"github.com/alehechka/buf-playground/utils/database"
	"github.com/alehechka/buf-playground/utils/grpc_shared"
	"github.com/alehechka/buf-playground/utils/otel"
)

func main() {
	shutdownTracer, err := otel.InitializeOpenTelTracer()
	utils.Check(err)
	defer shutdownTracer()

	disconnect, err := database.InitializeMongoDB(os.Getenv("MONGODB_CONNECTION_STRING"), "session")
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

	session.RegisterSessionServiceServer(server, &sessionserver.SessionServiceServer{})
	log.Println("Listening on", utils.ListenOn)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}
