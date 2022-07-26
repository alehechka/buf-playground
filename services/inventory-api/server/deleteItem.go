package server

import (
	"context"

	inventory "github.com/alehechka/buf-playground/proto/gen/go/inventory/v1alpha1"
	"github.com/alehechka/buf-playground/services/inventory-api/database"
)

func (s *InventoryServiceServer) DeleteItem(ctx context.Context, req *inventory.DeleteItemRequest) (res *inventory.DeleteItemResponse, err error) {
	res = &inventory.DeleteItemResponse{}
	err = database.DeleteItem(ctx, req.GetItemId())
	return
}
