package database

import (
	"context"

	inventory "github.com/alehechka/buf-playground/proto/gen/go/inventory/v1alpha1"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateItem(ctx context.Context, id string, i *inventory.Item) (*inventory.Item, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	updateItem := updateItem(i, oid)

	res := itemCollection().FindOneAndUpdate(ctx, item{ID: oid}, bson.M{"$set": updateItem}, options.FindOneAndUpdate().SetReturnDocument(1))
	if res.Err() != nil {
		return nil, res.Err()
	}

	internalItem := item{}
	err = res.Decode(&internalItem)
	return internalItem.Item(), err
}
