package main

import (
	"fmt"
	"log"
	"net"

	session "github.com/alehechka/buf-playground/proto/gen/go/session/v1alpha1"
	sessionserver "github.com/alehechka/buf-playground/services/session-api/server"
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
	session.RegisterSessionServiceServer(server, &sessionserver.SessionServiceServer{})
	log.Println("Listening on", utils.ListenOn)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}
