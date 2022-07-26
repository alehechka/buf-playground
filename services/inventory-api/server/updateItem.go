package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	inventory "github.com/alehechka/buf-playground/proto/gen/go/inventory/v1alpha1"
	"github.com/alehechka/buf-playground/services/inventory-api/database"
)

func (s *InventoryServiceServer) UpdateItem(ctx context.Context, req *inventory.UpdateItemRequest) (*inventory.UpdateItemResponse, error) {
	item, err := database.UpdateItem(ctx, req.GetItemId(), req.GetItem())
	if err != nil {
		return nil, status.Error(codes.Unavailable, err.Error())
	}

	return &inventory.UpdateItemResponse{Item: item}, nil
}
