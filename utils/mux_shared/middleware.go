package mux_shared

import (
	"context"
	"net/http"

	"github.com/alehechka/buf-playground/utils/grpc_shared"
	"github.com/alehechka/go-utils/ginshared"
	graphql "github.com/alehechka/grpc-graphql-gateway/runtime"
	"github.com/gin-gonic/gin"
	rest "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type MuxRestHandler func(context.Context, grpc_shared.AuthToken) (*rest.ServeMux, error)

func RestMiddleware(handler MuxRestHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authToken := grpc_shared.AuthToken(ctx.GetHeader("Authorization"))

		mux, err := handler(ctx.Request.Context(), authToken)
		if ginshared.ShouldAbortWithError(ctx)(http.StatusInternalServerError, err) {
			return
		}

		mux.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

type MuxGraphqlHandler func(grpc_shared.AuthToken) (*graphql.ServeMux, error)

func GraphqlMiddleware(handler MuxGraphqlHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authToken := grpc_shared.AuthToken(ctx.GetHeader("Authorization"))

		mux, err := handler(authToken)
		if ginshared.ShouldAbortWithError(ctx)(http.StatusInternalServerError, err) {
			return
		}

		mux.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
