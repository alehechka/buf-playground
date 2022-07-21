package server

import (
	session "github.com/alehechka/buf-playground/proto/gen/go/session/v1alpha1"
)

// SessionServiceServer implements the SessionService API.
type SessionServiceServer struct {
	session.UnimplementedSessionServiceServer
}
