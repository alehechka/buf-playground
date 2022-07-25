package main

import (
	inventorygw "github.com/alehechka/buf-playground/proto/gen/go/inventory/v1alpha1"
	sessiongw "github.com/alehechka/buf-playground/proto/gen/go/session/v1alpha1"
	"github.com/alehechka/buf-playground/utils"
	"github.com/friendsofgo/graphiql"
	"github.com/gin-gonic/gin"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"google.golang.org/grpc"
)

func main() {
	mux, err := createGraphQLHandler()
	utils.Check(err)

	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	utils.Check(err)

	engine := gin.Default()
	engine.POST("/graphql", gin.WrapH(mux))
	engine.GET("/graphiql", gin.WrapH(graphiqlHandler))
	utils.Check(engine.Run())
}

func createGraphQLHandler() (mux *runtime.ServeMux, err error) {
	mux = runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err = inventorygw.RegisterInventoryServiceGraphql(mux, utils.GRPCInventoryServerEndpoint, opts...); err != nil {
		return nil, err
	}

	if err = sessiongw.RegisterSessionServiceGraphql(mux, utils.GRPCSessionServerEndpoint, opts...); err != nil {
		return nil, err
	}

	return
}
