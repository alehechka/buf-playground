package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	inventorygw "github.com/alehechka/buf-playground/proto/gen/go/inventory/v1alpha1"
	sessiongw "github.com/alehechka/buf-playground/proto/gen/go/session/v1alpha1"
	"github.com/alehechka/buf-playground/utils"
)

var (
	// gRPC server endpoint
	grpcInventoryServerEndpoint = utils.GetEnv("GRCP_INVENTORY_SERVER_ENDPOINT", "localhost:3001")
	grpcSessionServerEndpoint   = utils.GetEnv("GRCP_SESSION_SERVER_ENDPOINT", "localhost:3002")
)

func main() {
	mux, err := createRESTHandler()
	utils.Check(err)

	engine := gin.Default()
	engine.Any("/api/*path", gin.WrapH(mux))
	utils.Check(engine.Run())
}

func createRESTHandler() (mux *runtime.ServeMux, err error) {
	ctx := context.Background()

	fmt.Println("inventory: \t", grpcInventoryServerEndpoint)
	fmt.Println("session: \t", grpcSessionServerEndpoint)

	mux = runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	if err = inventorygw.RegisterInventoryServiceHandlerFromEndpoint(ctx, mux, grpcInventoryServerEndpoint, opts); err != nil {
		return nil, err
	}

	if err = sessiongw.RegisterSessionServiceHandlerFromEndpoint(ctx, mux, grpcSessionServerEndpoint, opts); err != nil {
		return nil, err
	}

	return
}
