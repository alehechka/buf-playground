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

var (
	listenOn = ":" + utils.GetEnv("PORT", "80")
)

func main() {
	utils.Check(run())
}

func run() error {
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
