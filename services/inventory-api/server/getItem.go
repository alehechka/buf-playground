package server

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	inventory "github.com/alehechka/buf-playground/proto/gen/go/inventory/v1alpha1"
	"github.com/alehechka/buf-playground/services/inventory-api/database"
)

// GetItem retrieves a random item from the InventoryService.
func (s *InventoryServiceServer) GetItem(ctx context.Context, req *inventory.GetItemRequest) (*inventory.GetItemResponse, error) {
	itemID := req.GetItemId()
	if len(itemID) == 0 {
		return nil, status.Error(codes.InvalidArgument, "no item_id provided")
	}

	log.Println("Got a request to retrieve item with ID:", itemID)

	item, err := database.GetItem(itemID)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &inventory.GetItemResponse{Item: item}, nil
}
