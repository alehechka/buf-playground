package database

import (
	"context"

	inventory "github.com/alehechka/buf-playground/proto/gen/go/inventory/v1alpha1"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateItem(ctx context.Context, item *inventory.Item) (*inventory.Item, error) {
	createItem := newItem(item)

	res, err := itemCollection().InsertOne(ctx, createItem)
	if err != nil {
		return nil, err
	}

	createItem.ID = res.InsertedID.(primitive.ObjectID)
	return createItem.Item(), nil
}
