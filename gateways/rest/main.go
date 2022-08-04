package main

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	inventory "github.com/alehechka/buf-playground/proto/gen/go/inventory/v1alpha1"
	session "github.com/alehechka/buf-playground/proto/gen/go/session/v1alpha1"
	"github.com/alehechka/buf-playground/utils"
	"github.com/alehechka/buf-playground/utils/grpc_shared"
	"github.com/alehechka/buf-playground/utils/mux_shared"
	"github.com/alehechka/buf-playground/utils/otel"
)

func main() {
	shutdownTracer, err := otel.InitializeOpenTelTracer()
	utils.Check(err)
	defer shutdownTracer()

	engine := gin.Default()
	engine.Use(otelgin.Middleware(os.Getenv("OTEL_SERVICE_NAME"), otelgin.WithTracerProvider(otel.OpenTelTracer)))
	engine.Any("/api/*path", mux_shared.RestMiddleware(createHandler))
	utils.Check(engine.Run())
}

func createHandler(ctx context.Context, authToken grpc_shared.AuthToken) (mux *runtime.ServeMux, err error) {
	mux = runtime.NewServeMux()
	opts := grpc_shared.ClientDialOptions(authToken)

	if err = inventory.RegisterInventoryServiceHandlerFromEndpoint(ctx, mux, utils.GRPCInventoryServerEndpoint, opts); err != nil {
		return nil, err
	}

	if err = session.RegisterSessionServiceHandlerFromEndpoint(ctx, mux, utils.GRPCSessionServerEndpoint, opts); err != nil {
		return nil, err
	}

	return
}
