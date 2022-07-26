package database

import (
	"context"

	inventory "github.com/alehechka/buf-playground/proto/gen/go/inventory/v1alpha1"
	"go.mongodb.org/mongo-driver/bson"
)

// ListItems retrieves a list of Items
func ListItems(ctx context.Context) (items []*inventory.Item, err error) {
	cursor, err := itemCollection().Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		internalItem := item{}
		err = cursor.Decode(&internalItem)
		if err != nil {
			return
		}

		items = append(items, internalItem.Item())
	}
	return
}
