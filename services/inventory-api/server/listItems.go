package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	inventory "github.com/alehechka/buf-playground/proto/gen/go/inventory/v1alpha1"
	"github.com/alehechka/buf-playground/services/inventory-api/database"
)

// ListItems retrieves a list of Items
func (s *InventoryServiceServer) ListItems(ctx context.Context, req *inventory.ListItemsRequest) (*inventory.ListItemsResponse, error) {

	items, err := database.ListItems(ctx)
	if err != nil {
		return nil, status.Error(codes.Unavailable, err.Error())
	}

	return &inventory.ListItemsResponse{Items: items}, nil
}
