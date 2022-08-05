package grpc_shared

import (
	"github.com/alehechka/buf-playground/utils/otel"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

// Server creates and returns a gRPC server instance with all ServerOptions baked in
func Server() *grpc.Server {
	return grpc.NewServer(
		unaryServerChain(),
		streamingServerChain(),
	)
}

func unaryServerChain() grpc.ServerOption {
	return grpc.UnaryInterceptor(
		grpc_middleware.ChainUnaryServer(
			otelgrpc.UnaryServerInterceptor(otelgrpc.WithTracerProvider(otel.OpenTelTracer)),
			grpc_auth.UnaryServerInterceptor(authorize),
		),
	)
}

func streamingServerChain() grpc.ServerOption {
	return grpc.StreamInterceptor(
		grpc_middleware.ChainStreamServer(
			otelgrpc.StreamServerInterceptor(otelgrpc.WithTracerProvider(otel.OpenTelTracer)),
			grpc_auth.StreamServerInterceptor(authorize),
		),
	)
}
