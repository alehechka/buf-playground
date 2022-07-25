package database

import (
	inventory "github.com/alehechka/buf-playground/proto/gen/go/inventory/v1alpha1"
	"github.com/alehechka/buf-playground/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/genproto/googleapis/type/money"
)

func itemCollection() *mongo.Collection {
	if utils.Database != nil {
		return utils.Database.Collection("items")
	}
	return nil
}

type item struct {
	id primitive.ObjectID `bson:"_id"`

	name     string       `bson:"n,omitempty"`
	weight   float32      `bson:"w,omitempty"`
	height   float32      `bson:"h,omitempty"`
	quantity int64        `bson:"q,omitempty"`
	price    *money.Money `bson:"p,omitempty"`
}

func newItem(i *inventory.Item) item {
	newI, _ := updateItem(i)
	return newI
}

func updateItem(i *inventory.Item) (item, error) {
	id, err := primitive.ObjectIDFromHex(i.GetItemId())

	return item{
		id:       id,
		name:     i.GetName(),
		weight:   i.GetWeight(),
		height:   i.GetHeight(),
		quantity: i.GetQuantity(),
		price:    i.GetPrice(),
	}, err
}

func (i *item) Item() *inventory.Item {
	return &inventory.Item{
		ItemId:   i.id.Hex(),
		Name:     i.name,
		Weight:   i.weight,
		Height:   i.height,
		Quantity: i.quantity,
		Price:    i.price,
	}
}
