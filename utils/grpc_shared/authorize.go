package grpc_shared

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const authorizationKey string = "authorization"

func authorize(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx, status.Errorf(codes.PermissionDenied, "retrieving metadata failed")
	}

	fmt.Println("Meta: ", md)

	auth, ok := md[authorizationKey]
	if !ok {
		return ctx, status.Errorf(codes.PermissionDenied, "no auth details supplied")
	}

	if len(auth) == 0 {
		return ctx, status.Errorf(codes.PermissionDenied, "invalid authorization provided")
	}

	fmt.Println("Authorized: ", auth)

	return ctx, nil
}

type AuthToken string

func (a AuthToken) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		authorizationKey: string(a),
	}, nil
}

func (a AuthToken) RequireTransportSecurity() bool {
	return false
}
