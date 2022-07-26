package main

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"

	inventory "github.com/alehechka/buf-playground/proto/gen/go/inventory/v1alpha1"
	session "github.com/alehechka/buf-playground/proto/gen/go/session/v1alpha1"
	"github.com/alehechka/buf-playground/utils"
)

func main() {
	shutdownTracer, err := utils.InitializeOpenTelTracer()
	utils.Check(err)
	defer shutdownTracer()

	mux, err := createRESTHandler()
	utils.Check(err)

	engine := gin.Default()
	engine.Use(otelgin.Middleware(os.Getenv("OTEL_SERVICE_NAME"), otelgin.WithTracerProvider(utils.OpenTelTracer)))
	engine.Any("/api/*path", gin.WrapH(mux))
	utils.Check(engine.Run())
}

func createRESTHandler() (mux *runtime.ServeMux, err error) {
	ctx := context.Background()

	mux = runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	}

	if err = inventory.RegisterInventoryServiceHandlerFromEndpoint(ctx, mux, utils.GRPCInventoryServerEndpoint, opts); err != nil {
		return nil, err
	}

	if err = session.RegisterSessionServiceHandlerFromEndpoint(ctx, mux, utils.GRPCSessionServerEndpoint, opts); err != nil {
		return nil, err
	}

	return
}
