package main

import (
	"fmt"
	"log"
	"net"
	"os"

	session "github.com/alehechka/buf-playground/proto/gen/go/session/v1alpha1"
	sessionserver "github.com/alehechka/buf-playground/services/session-api/server"
	"github.com/alehechka/buf-playground/utils"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

func main() {
	shutdownTracer, err := utils.InitializeOpenTelTracer()
	utils.Check(err)
	defer shutdownTracer()

	disconnect, err := utils.InitializeMongoDB(os.Getenv("MONGODB_CONNECTION_STRING"), "session")
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

	session.RegisterSessionServiceServer(server, &sessionserver.SessionServiceServer{})
	log.Println("Listening on", utils.ListenOn)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}
