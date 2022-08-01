// Code generated by protoc-gen-graphql. DO NOT EDIT.
// versions:
// 	protoc-gen-graphql dev
// 	protoc             (unknown)
// source: payment/v1alpha1/payment.proto

package paymentv1alpha1

import (
	gql_ptypes_timestamppb "github.com/alehechka/grpc-graphql-gateway/ptypes/google/protobuf/timestamp"
	"github.com/graphql-go/graphql"

	gql_ptypes_money "github.com/alehechka/grpc-graphql-gateway/ptypes/google/type/money"
)

var (
	gql__type_Receipt   *graphql.Object      // message Receipt in payment/v1alpha1/payment.proto
	gql__type_Purchase  *graphql.Object      // message Purchase in payment/v1alpha1/payment.proto
	gql__input_Receipt  *graphql.InputObject // message Receipt in payment/v1alpha1/payment.proto
	gql__input_Purchase *graphql.InputObject // message Purchase in payment/v1alpha1/payment.proto
)

func Gql__type_Receipt() *graphql.Object {
	if gql__type_Receipt == nil {
		gql__type_Receipt = graphql.NewObject(graphql.ObjectConfig{
			Name: "Paymentv1Alpha1_Type_Receipt",
			Fields: graphql.Fields{
				"purchase": &graphql.Field{
					Type: Gql__type_Purchase(),
				},
				"receiptID": &graphql.Field{
					Type: graphql.String,
				},
				"createdAt": &graphql.Field{
					Type: gql_ptypes_timestamppb.Gql__type_Timestamp(),
				},
			},
		})
	}
	return gql__type_Receipt
}

func Gql__type_Purchase() *graphql.Object {
	if gql__type_Purchase == nil {
		gql__type_Purchase = graphql.NewObject(graphql.ObjectConfig{
			Name: "Paymentv1Alpha1_Type_Purchase",
			Fields: graphql.Fields{
				"itemID": &graphql.Field{
					Type: graphql.String,
				},
				"quantity": &graphql.Field{
					Type: graphql.Float,
				},
				"userID": &graphql.Field{
					Type: graphql.String,
				},
				"price": &graphql.Field{
					Type: gql_ptypes_money.Gql__type_Money(),
				},
			},
		})
	}
	return gql__type_Purchase
}

func Gql__input_Receipt() *graphql.InputObject {
	if gql__input_Receipt == nil {
		gql__input_Receipt = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Paymentv1Alpha1_Input_Receipt",
			Fields: graphql.InputObjectConfigFieldMap{
				"purchase": &graphql.InputObjectFieldConfig{
					Type: Gql__input_Purchase(),
				},
				"receiptID": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"createdAt": &graphql.InputObjectFieldConfig{
					Type: gql_ptypes_timestamppb.Gql__input_Timestamp(),
				},
			},
		})
	}
	return gql__input_Receipt
}

func Gql__input_Purchase() *graphql.InputObject {
	if gql__input_Purchase == nil {
		gql__input_Purchase = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Paymentv1Alpha1_Input_Purchase",
			Fields: graphql.InputObjectConfigFieldMap{
				"itemID": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"quantity": &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				"userID": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"price": &graphql.InputObjectFieldConfig{
					Type: gql_ptypes_money.Gql__input_Money(),
				},
			},
		})
	}
	return gql__input_Purchase
}
