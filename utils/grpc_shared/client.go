package grpc_shared

import (
	"github.com/alehechka/buf-playground/utils/otel"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

func ClientDialOptions(authToken AuthToken) []grpc.DialOption {
	return []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithPerRPCCredentials(authToken),
		unaryClientChain(),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	}
}

func unaryClientChain() grpc.DialOption {
	return grpc.WithUnaryInterceptor(
		grpc_middleware.ChainUnaryClient(
			otelgrpc.UnaryClientInterceptor(otelgrpc.WithTracerProvider(otel.OpenTelTracer)),
		),
	)
}
