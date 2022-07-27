package main

import (
	"os"

	inventory "github.com/alehechka/buf-playground/proto/gen/go/inventory/v1alpha1"
	session "github.com/alehechka/buf-playground/proto/gen/go/session/v1alpha1"
	"github.com/alehechka/buf-playground/utils"
	"github.com/alehechka/grpc-graphql-gateway/runtime"
	"github.com/friendsofgo/graphiql"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

func main() {
	shutdownTracer, err := utils.InitializeOpenTelTracer()
	utils.Check(err)
	defer shutdownTracer()

	mux, err := createGraphQLHandler()
	utils.Check(err)

	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	utils.Check(err)

	engine := gin.Default()
	engine.Use(otelgin.Middleware(os.Getenv("OTEL_SERVICE_NAME"), otelgin.WithTracerProvider(utils.OpenTelTracer)))
	engine.POST("/graphql", gin.WrapH(mux))
	engine.GET("/graphiql", gin.WrapH(graphiqlHandler))
	utils.Check(engine.Run())
}

func createGraphQLHandler() (mux *runtime.ServeMux, err error) {
	mux = runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	}

	if err = inventory.RegisterInventoryServiceGraphql(mux, utils.GRPCInventoryServerEndpoint, opts...); err != nil {
		return nil, err
	}

	if err = session.RegisterSessionServiceGraphql(mux, utils.GRPCSessionServerEndpoint, opts...); err != nil {
		return nil, err
	}

	return
}
