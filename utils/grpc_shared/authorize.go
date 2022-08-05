package grpc_shared

import (
	"context"

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

	auth, ok := md[authorizationKey]
	if !ok {
		return ctx, status.Errorf(codes.PermissionDenied, "no auth details supplied")
	}

	if len(auth) == 0 {
		return ctx, status.Errorf(codes.PermissionDenied, "invalid authorization provided")
	}

	return ctx, nil
}
