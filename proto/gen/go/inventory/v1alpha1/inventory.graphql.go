// Code generated by protoc-gen-graphql. DO NOT EDIT.
// versions:
// 	protoc-gen-graphql dev
// 	protoc             (unknown)
// source: inventory/v1alpha1/inventory.proto

package inventoryv1alpha1

import (
	"context"

	paymentv1alpha1 "github.com/alehechka/buf-playground/proto/gen/go/payment/v1alpha1"
	"github.com/alehechka/grpc-graphql-gateway/runtime"
	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	gql_ptypes_money "github.com/alehechka/grpc-graphql-gateway/ptypes/google/type/money"
)

var (
	gql__type_UpdateItemResponse    *graphql.Object      // message UpdateItemResponse in inventory/v1alpha1/inventory.proto
	gql__type_UpdateItemRequest     *graphql.Object      // message UpdateItemRequest in inventory/v1alpha1/inventory.proto
	gql__type_PurchaseItemResponse  *graphql.Object      // message PurchaseItemResponse in inventory/v1alpha1/inventory.proto
	gql__type_PurchaseItemRequest   *graphql.Object      // message PurchaseItemRequest in inventory/v1alpha1/inventory.proto
	gql__type_ListItemsResponse     *graphql.Object      // message ListItemsResponse in inventory/v1alpha1/inventory.proto
	gql__type_ListItemsRequest      *graphql.Object      // message ListItemsRequest in inventory/v1alpha1/inventory.proto
	gql__type_Item                  *graphql.Object      // message Item in inventory/v1alpha1/inventory.proto
	gql__type_GetItemResponse       *graphql.Object      // message GetItemResponse in inventory/v1alpha1/inventory.proto
	gql__type_GetItemRequest        *graphql.Object      // message GetItemRequest in inventory/v1alpha1/inventory.proto
	gql__type_DeleteItemResponse    *graphql.Object      // message DeleteItemResponse in inventory/v1alpha1/inventory.proto
	gql__type_DeleteItemRequest     *graphql.Object      // message DeleteItemRequest in inventory/v1alpha1/inventory.proto
	gql__type_CreateItemResponse    *graphql.Object      // message CreateItemResponse in inventory/v1alpha1/inventory.proto
	gql__type_CreateItemRequest     *graphql.Object      // message CreateItemRequest in inventory/v1alpha1/inventory.proto
	gql__input_UpdateItemResponse   *graphql.InputObject // message UpdateItemResponse in inventory/v1alpha1/inventory.proto
	gql__input_UpdateItemRequest    *graphql.InputObject // message UpdateItemRequest in inventory/v1alpha1/inventory.proto
	gql__input_PurchaseItemResponse *graphql.InputObject // message PurchaseItemResponse in inventory/v1alpha1/inventory.proto
	gql__input_PurchaseItemRequest  *graphql.InputObject // message PurchaseItemRequest in inventory/v1alpha1/inventory.proto
	gql__input_Money                *graphql.InputObject // message Money in google/type/money.proto
	gql__input_ListItemsResponse    *graphql.InputObject // message ListItemsResponse in inventory/v1alpha1/inventory.proto
	gql__input_ListItemsRequest     *graphql.InputObject // message ListItemsRequest in inventory/v1alpha1/inventory.proto
	gql__input_Item                 *graphql.InputObject // message Item in inventory/v1alpha1/inventory.proto
	gql__input_GetItemResponse      *graphql.InputObject // message GetItemResponse in inventory/v1alpha1/inventory.proto
	gql__input_GetItemRequest       *graphql.InputObject // message GetItemRequest in inventory/v1alpha1/inventory.proto
	gql__input_DeleteItemResponse   *graphql.InputObject // message DeleteItemResponse in inventory/v1alpha1/inventory.proto
	gql__input_DeleteItemRequest    *graphql.InputObject // message DeleteItemRequest in inventory/v1alpha1/inventory.proto
	gql__input_CreateItemResponse   *graphql.InputObject // message CreateItemResponse in inventory/v1alpha1/inventory.proto
	gql__input_CreateItemRequest    *graphql.InputObject // message CreateItemRequest in inventory/v1alpha1/inventory.proto
)

func Gql__type_UpdateItemResponse() *graphql.Object {
	if gql__type_UpdateItemResponse == nil {
		gql__type_UpdateItemResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Inventoryv1Alpha1_Type_UpdateItemResponse",
			Fields: graphql.Fields{
				"item": &graphql.Field{
					Type: graphql.NewNonNull(Gql__type_Item()),
				},
			},
		})
	}
	return gql__type_UpdateItemResponse
}

func Gql__type_UpdateItemRequest() *graphql.Object {
	if gql__type_UpdateItemRequest == nil {
		gql__type_UpdateItemRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Inventoryv1Alpha1_Type_UpdateItemRequest",
			Fields: graphql.Fields{
				"itemID": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"item": &graphql.Field{
					Type: graphql.NewNonNull(Gql__type_Item()),
				},
			},
		})
	}
	return gql__type_UpdateItemRequest
}

func Gql__type_PurchaseItemResponse() *graphql.Object {
	if gql__type_PurchaseItemResponse == nil {
		gql__type_PurchaseItemResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Inventoryv1Alpha1_Type_PurchaseItemResponse",
			Fields: graphql.Fields{
				"receipt": &graphql.Field{
					Type: paymentv1alpha1.Gql__type_Receipt(),
				},
			},
		})
	}
	return gql__type_PurchaseItemResponse
}

func Gql__type_PurchaseItemRequest() *graphql.Object {
	if gql__type_PurchaseItemRequest == nil {
		gql__type_PurchaseItemRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Inventoryv1Alpha1_Type_PurchaseItemRequest",
			Fields: graphql.Fields{
				"purchase": &graphql.Field{
					Type: paymentv1alpha1.Gql__type_Purchase(),
				},
			},
		})
	}
	return gql__type_PurchaseItemRequest
}

func Gql__type_ListItemsResponse() *graphql.Object {
	if gql__type_ListItemsResponse == nil {
		gql__type_ListItemsResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Inventoryv1Alpha1_Type_ListItemsResponse",
			Fields: graphql.Fields{
				"items": &graphql.Field{
					Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(Gql__type_Item()))),
				},
			},
		})
	}
	return gql__type_ListItemsResponse
}

func Gql__type_ListItemsRequest() *graphql.Object {
	if gql__type_ListItemsRequest == nil {
		gql__type_ListItemsRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Inventoryv1Alpha1_Type_ListItemsRequest",
			Fields: graphql.Fields{
				"_": &graphql.Field{
					Type: graphql.Boolean,
				},
			},
		})
	}
	return gql__type_ListItemsRequest
}

func Gql__type_Item() *graphql.Object {
	if gql__type_Item == nil {
		gql__type_Item = graphql.NewObject(graphql.ObjectConfig{
			Name: "Inventoryv1Alpha1_Type_Item",
			Fields: graphql.Fields{
				"itemID": &graphql.Field{
					Type: graphql.String,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"weight": &graphql.Field{
					Type: graphql.Float,
				},
				"height": &graphql.Field{
					Type: graphql.Float,
				},
				"quantity": &graphql.Field{
					Type: graphql.Float,
				},
				"price": &graphql.Field{
					Type: gql_ptypes_money.Gql__type_Money(),
				},
			},
		})
	}
	return gql__type_Item
}

func Gql__type_GetItemResponse() *graphql.Object {
	if gql__type_GetItemResponse == nil {
		gql__type_GetItemResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Inventoryv1Alpha1_Type_GetItemResponse",
			Fields: graphql.Fields{
				"item": &graphql.Field{
					Type: graphql.NewNonNull(Gql__type_Item()),
				},
			},
		})
	}
	return gql__type_GetItemResponse
}

func Gql__type_GetItemRequest() *graphql.Object {
	if gql__type_GetItemRequest == nil {
		gql__type_GetItemRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Inventoryv1Alpha1_Type_GetItemRequest",
			Fields: graphql.Fields{
				"itemID": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
		})
	}
	return gql__type_GetItemRequest
}

func Gql__type_DeleteItemResponse() *graphql.Object {
	if gql__type_DeleteItemResponse == nil {
		gql__type_DeleteItemResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Inventoryv1Alpha1_Type_DeleteItemResponse",
			Fields: graphql.Fields{
				"_": &graphql.Field{
					Type: graphql.Boolean,
				},
			},
		})
	}
	return gql__type_DeleteItemResponse
}

func Gql__type_DeleteItemRequest() *graphql.Object {
	if gql__type_DeleteItemRequest == nil {
		gql__type_DeleteItemRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Inventoryv1Alpha1_Type_DeleteItemRequest",
			Fields: graphql.Fields{
				"itemID": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
		})
	}
	return gql__type_DeleteItemRequest
}

func Gql__type_CreateItemResponse() *graphql.Object {
	if gql__type_CreateItemResponse == nil {
		gql__type_CreateItemResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Inventoryv1Alpha1_Type_CreateItemResponse",
			Fields: graphql.Fields{
				"item": &graphql.Field{
					Type: graphql.NewNonNull(Gql__type_Item()),
				},
			},
		})
	}
	return gql__type_CreateItemResponse
}

func Gql__type_CreateItemRequest() *graphql.Object {
	if gql__type_CreateItemRequest == nil {
		gql__type_CreateItemRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Inventoryv1Alpha1_Type_CreateItemRequest",
			Fields: graphql.Fields{
				"item": &graphql.Field{
					Type: graphql.NewNonNull(Gql__type_Item()),
				},
			},
		})
	}
	return gql__type_CreateItemRequest
}

func Gql__input_UpdateItemResponse() *graphql.InputObject {
	if gql__input_UpdateItemResponse == nil {
		gql__input_UpdateItemResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Inventoryv1Alpha1_Input_UpdateItemResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"item": &graphql.InputObjectFieldConfig{
					Type: graphql.NewNonNull(Gql__input_Item()),
				},
			},
		})
	}
	return gql__input_UpdateItemResponse
}

func Gql__input_UpdateItemRequest() *graphql.InputObject {
	if gql__input_UpdateItemRequest == nil {
		gql__input_UpdateItemRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Inventoryv1Alpha1_Input_UpdateItemRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"itemID": &graphql.InputObjectFieldConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"item": &graphql.InputObjectFieldConfig{
					Type: graphql.NewNonNull(Gql__input_Item()),
				},
			},
		})
	}
	return gql__input_UpdateItemRequest
}

func Gql__input_PurchaseItemResponse() *graphql.InputObject {
	if gql__input_PurchaseItemResponse == nil {
		gql__input_PurchaseItemResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Inventoryv1Alpha1_Input_PurchaseItemResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"receipt": &graphql.InputObjectFieldConfig{
					Type: paymentv1alpha1.Gql__input_Receipt(),
				},
			},
		})
	}
	return gql__input_PurchaseItemResponse
}

func Gql__input_PurchaseItemRequest() *graphql.InputObject {
	if gql__input_PurchaseItemRequest == nil {
		gql__input_PurchaseItemRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Inventoryv1Alpha1_Input_PurchaseItemRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"purchase": &graphql.InputObjectFieldConfig{
					Type: paymentv1alpha1.Gql__input_Purchase(),
				},
			},
		})
	}
	return gql__input_PurchaseItemRequest
}

func Gql__input_Money() *graphql.InputObject {
	if gql__input_Money == nil {
		gql__input_Money = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Inventoryv1Alpha1_Input_Money",
			Fields: graphql.InputObjectConfigFieldMap{
				"currencyCode": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"units": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"nanos": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__input_Money
}

func Gql__input_ListItemsResponse() *graphql.InputObject {
	if gql__input_ListItemsResponse == nil {
		gql__input_ListItemsResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Inventoryv1Alpha1_Input_ListItemsResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"items": &graphql.InputObjectFieldConfig{
					Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(Gql__input_Item()))),
				},
			},
		})
	}
	return gql__input_ListItemsResponse
}

func Gql__input_ListItemsRequest() *graphql.InputObject {
	if gql__input_ListItemsRequest == nil {
		gql__input_ListItemsRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Inventoryv1Alpha1_Input_ListItemsRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"_": &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
			},
		})
	}
	return gql__input_ListItemsRequest
}

func Gql__input_Item() *graphql.InputObject {
	if gql__input_Item == nil {
		gql__input_Item = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Inventoryv1Alpha1_Input_Item",
			Fields: graphql.InputObjectConfigFieldMap{
				"itemID": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"name": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"weight": &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				"height": &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				"quantity": &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				"price": &graphql.InputObjectFieldConfig{
					Type: gql_ptypes_money.Gql__input_Money(),
				},
			},
		})
	}
	return gql__input_Item
}

func Gql__input_GetItemResponse() *graphql.InputObject {
	if gql__input_GetItemResponse == nil {
		gql__input_GetItemResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Inventoryv1Alpha1_Input_GetItemResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"item": &graphql.InputObjectFieldConfig{
					Type: graphql.NewNonNull(Gql__input_Item()),
				},
			},
		})
	}
	return gql__input_GetItemResponse
}

func Gql__input_GetItemRequest() *graphql.InputObject {
	if gql__input_GetItemRequest == nil {
		gql__input_GetItemRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Inventoryv1Alpha1_Input_GetItemRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"itemID": &graphql.InputObjectFieldConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
		})
	}
	return gql__input_GetItemRequest
}

func Gql__input_DeleteItemResponse() *graphql.InputObject {
	if gql__input_DeleteItemResponse == nil {
		gql__input_DeleteItemResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Inventoryv1Alpha1_Input_DeleteItemResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"_": &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
			},
		})
	}
	return gql__input_DeleteItemResponse
}

func Gql__input_DeleteItemRequest() *graphql.InputObject {
	if gql__input_DeleteItemRequest == nil {
		gql__input_DeleteItemRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Inventoryv1Alpha1_Input_DeleteItemRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"itemID": &graphql.InputObjectFieldConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
		})
	}
	return gql__input_DeleteItemRequest
}

func Gql__input_CreateItemResponse() *graphql.InputObject {
	if gql__input_CreateItemResponse == nil {
		gql__input_CreateItemResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Inventoryv1Alpha1_Input_CreateItemResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"item": &graphql.InputObjectFieldConfig{
					Type: graphql.NewNonNull(Gql__input_Item()),
				},
			},
		})
	}
	return gql__input_CreateItemResponse
}

func Gql__input_CreateItemRequest() *graphql.InputObject {
	if gql__input_CreateItemRequest == nil {
		gql__input_CreateItemRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Inventoryv1Alpha1_Input_CreateItemRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"item": &graphql.InputObjectFieldConfig{
					Type: graphql.NewNonNull(Gql__input_Item()),
				},
			},
		})
	}
	return gql__input_CreateItemRequest
}

// graphql__resolver_InventoryService is a struct for making query, mutation and resolve fields.
// This struct must be implemented runtime.SchemaBuilder interface.
type graphql__resolver_InventoryService struct {

	// Automatic connection host
	host string

	// grpc dial options
	dialOptions []grpc.DialOption

	// grpc client connection.
	// this connection may be provided by user
	conn *grpc.ClientConn
}

// new_graphql_resolver_InventoryService creates pointer of service struct
func new_graphql_resolver_InventoryService(conn *grpc.ClientConn, host string, opts []grpc.DialOption) *graphql__resolver_InventoryService {
	return &graphql__resolver_InventoryService{
		conn:        conn,
		host:        host,
		dialOptions: opts,
	}
}

// CreateConnection() returns grpc connection which user specified or newly connected and closing function
func (x *graphql__resolver_InventoryService) CreateConnection(ctx context.Context) (*grpc.ClientConn, func(), error) {
	// If x.conn is not nil, user injected their own connection
	if x.conn != nil {
		return x.conn, func() {}, nil
	}

	// Otherwise, this handler opens connection with specified host
	conn, err := grpc.DialContext(ctx, x.host, x.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	return conn, func() { conn.Close() }, nil
}

// GetQueries returns acceptable graphql.Fields for Query.
func (x *graphql__resolver_InventoryService) GetQueries(conn *grpc.ClientConn) graphql.Fields {
	return graphql.Fields{
		"getItem": &graphql.Field{
			Type: Gql__type_GetItemResponse(),
			Args: graphql.FieldConfigArgument{
				"itemID": &graphql.ArgumentConfig{
					Type:         graphql.NewNonNull(graphql.String),
					DefaultValue: "",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req GetItemRequest
				if err := runtime.MarshalRequest(p.Args, &req); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for getItem")
				}
				client := NewInventoryServiceClient(conn)
				resp, err := client.GetItem(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC GetItem")
				}
				return runtime.MarshalResponse(resp), nil
			},
		},
		"listItems": &graphql.Field{
			Type: Gql__type_ListItemsResponse(),
			Args: graphql.FieldConfigArgument{},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req ListItemsRequest
				if err := runtime.MarshalRequest(p.Args, &req); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for listItems")
				}
				client := NewInventoryServiceClient(conn)
				resp, err := client.ListItems(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC ListItems")
				}
				return runtime.MarshalResponse(resp), nil
			},
		},
	}
}

// GetMutations returns acceptable graphql.Fields for Mutation.
func (x *graphql__resolver_InventoryService) GetMutations(conn *grpc.ClientConn) graphql.Fields {
	return graphql.Fields{
		"createItem": &graphql.Field{
			Type: Gql__type_CreateItemResponse(),
			Args: graphql.FieldConfigArgument{
				"item": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(Gql__input_Item()),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req CreateItemRequest
				if err := runtime.MarshalRequest(p.Args, &req); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for createItem")
				}
				client := NewInventoryServiceClient(conn)
				resp, err := client.CreateItem(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC CreateItem")
				}
				return runtime.MarshalResponse(resp), nil
			},
		},

		"updateItem": &graphql.Field{
			Type: Gql__type_UpdateItemResponse(),
			Args: graphql.FieldConfigArgument{
				"itemID": &graphql.ArgumentConfig{
					Type:         graphql.NewNonNull(graphql.String),
					DefaultValue: "",
				},
				"item": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(Gql__input_Item()),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req UpdateItemRequest
				if err := runtime.MarshalRequest(p.Args, &req); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for updateItem")
				}
				client := NewInventoryServiceClient(conn)
				resp, err := client.UpdateItem(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC UpdateItem")
				}
				return runtime.MarshalResponse(resp), nil
			},
		},

		"deleteItem": &graphql.Field{
			Type: Gql__type_DeleteItemResponse(),
			Args: graphql.FieldConfigArgument{
				"itemID": &graphql.ArgumentConfig{
					Type:         graphql.NewNonNull(graphql.String),
					DefaultValue: "",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req DeleteItemRequest
				if err := runtime.MarshalRequest(p.Args, &req); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for deleteItem")
				}
				client := NewInventoryServiceClient(conn)
				resp, err := client.DeleteItem(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC DeleteItem")
				}
				return runtime.MarshalResponse(resp), nil
			},
		},
	}
}

// Register package divided graphql handler "without" *grpc.ClientConn,
// therefore gRPC connection will be opened and closed automatically.
// Occasionally you may worry about open/close performance for each handling graphql request,
// then you can call RegisterInventoryServiceGraphqlHandler with *grpc.ClientConn manually.
func RegisterInventoryServiceGraphql(mux *runtime.ServeMux, host string, opts ...grpc.DialOption) error {
	return RegisterInventoryServiceGraphqlHandler(mux, nil, host, opts...)
}

// Register package divided graphql handler "with" *grpc.ClientConn.
// this function accepts your defined grpc connection, so that we reuse that and never close connection inside.
// You need to close it manually when application will terminate.
func RegisterInventoryServiceGraphqlHandler(mux *runtime.ServeMux, conn *grpc.ClientConn, host string, opts ...grpc.DialOption) error {
	return mux.AddHandler(new_graphql_resolver_InventoryService(conn, host, opts))
}
