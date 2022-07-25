package database

import (
	"context"

	inventory "github.com/alehechka/buf-playground/proto/gen/go/inventory/v1alpha1"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetItem attempts to retrieve a single item denoted the provided id.
func GetItem(id string) (*inventory.Item, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	res := itemCollection().FindOne(context.Background(), item{id: oid})
	if res.Err() != nil {
		return nil, res.Err()
	}

	internalItem := item{}
	err = res.Decode(&internalItem)
	return internalItem.Item(), err
}
