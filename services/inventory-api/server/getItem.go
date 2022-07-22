package server

import (
	"context"
	"log"

	"google.golang.org/genproto/googleapis/type/money"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	inventory "github.com/alehechka/buf-playground/proto/gen/go/inventory/v1alpha1"
)

// GetItem retrieves a random item from the InventoryService.
func (s *InventoryServiceServer) GetItem(ctx context.Context, req *inventory.GetItemRequest) (*inventory.GetItemResponse, error) {
	itemID := req.GetItemId()
	if len(itemID) == 0 {
		return nil, status.Error(codes.InvalidArgument, "no item_id provided")
	}

	log.Println("Got a request to retrieve item with ID:", itemID)

	return &inventory.GetItemResponse{Item: &inventory.Item{
		ItemId:   itemID,
		Name:     "basketball",
		Weight:   2.45,
		Height:   1.3,
		Quantity: 32,
		Price: &money.Money{
			CurrencyCode: "USD",
			Units:        19,
			Nanos:        190000000,
		},
	}}, nil
}
