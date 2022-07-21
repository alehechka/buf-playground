package server

import (
	inventory "github.com/alehechka/buf-playground/proto/gen/go/inventory/v1alpha1"
)

// InventoryServiceServer implements the InventoryService API.
type InventoryServiceServer struct {
	inventory.UnimplementedInventoryServiceServer
}
