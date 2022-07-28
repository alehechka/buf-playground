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
	ID primitive.ObjectID `bson:"_id"`

	Name     string       `bson:"n,omitempty"`
	Weight   float32      `bson:"w,omitempty"`
	Height   float32      `bson:"h,omitempty"`
	Quantity float32      `bson:"q,omitempty"`
	Price    *money.Money `bson:"p,omitempty"`
}

func newItem(i *inventory.Item) item {
	return updateItem(i, primitive.NewObjectID())
}

func updateItem(i *inventory.Item, id primitive.ObjectID) item {
	return item{
		ID:       id,
		Name:     i.GetName(),
		Weight:   i.GetWeight(),
		Height:   i.GetHeight(),
		Quantity: i.GetQuantity(),
		Price:    i.GetPrice(),
	}
}

func (i *item) Item() *inventory.Item {
	return &inventory.Item{
		ItemId:   i.ID.Hex(),
		Name:     i.Name,
		Weight:   i.Weight,
		Height:   i.Height,
		Quantity: i.Quantity,
		Price:    i.Price,
	}
}
