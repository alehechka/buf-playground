package server

import (
	"context"

	inventory "github.com/alehechka/buf-playground/proto/gen/go/inventory/v1alpha1"
	"github.com/alehechka/buf-playground/services/inventory-api/database"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *InventoryServiceServer) CreateItem(ctx context.Context, req *inventory.CreateItemRequest) (*inventory.CreateItemResponse, error) {
	item, err := database.CreateItem(ctx, req.GetItem())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &inventory.CreateItemResponse{Item: item}, nil
}
