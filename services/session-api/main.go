package main

import (
	"fmt"
	"log"
	"net"

	session "github.com/alehechka/buf-playground/proto/gen/go/session/v1alpha1"
	sessionserver "github.com/alehechka/buf-playground/services/session-api/server"
	"google.golang.org/grpc"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	listenOn := "127.0.0.1:8081"
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", listenOn, err)
	}

	server := grpc.NewServer()
	session.RegisterSessionServiceServer(server, &sessionserver.SessionServiceServer{})
	log.Println("Listening on", listenOn)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}
