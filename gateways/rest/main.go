package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	inventorygw "github.com/alehechka/buf-playground/proto/gen/go/inventory/v1alpha1"
	sessiongw "github.com/alehechka/buf-playground/proto/gen/go/session/v1alpha1"
	"github.com/alehechka/buf-playground/utils"
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

	mux = runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	if err = inventorygw.RegisterInventoryServiceHandlerFromEndpoint(ctx, mux, utils.GRPCInventoryServerEndpoint, opts); err != nil {
		return nil, err
	}

	if err = sessiongw.RegisterSessionServiceHandlerFromEndpoint(ctx, mux, utils.GRPCSessionServerEndpoint, opts); err != nil {
		return nil, err
	}

	return
}
