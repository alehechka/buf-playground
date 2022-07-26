package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteItem(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	res := itemCollection().FindOneAndDelete(ctx, item{ID: oid})

	return res.Err()
}
