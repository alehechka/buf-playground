package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	inventory "github.com/alehechka/buf-playground/proto/gen/go/inventory/v1alpha1"
	"github.com/alehechka/buf-playground/services/inventory-api/database"
)

// GetItem retrieves a random item from the InventoryService.
func (s *InventoryServiceServer) GetItem(ctx context.Context, req *inventory.GetItemRequest) (*inventory.GetItemResponse, error) {
	item, err := database.GetItem(ctx, req.GetItemId())
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &inventory.GetItemResponse{Item: item}, nil
}
