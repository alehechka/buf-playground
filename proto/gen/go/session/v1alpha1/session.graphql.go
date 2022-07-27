// Code generated by proroc-gen-graphql, DO NOT EDIT.
package sessionv1alpha1

import (
	"context"

	"github.com/alehechka/grpc-graphql-gateway/runtime"
	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

var (
	gql__enum_Gender           *graphql.Enum        // enum Gender in session/v1alpha1/session.proto
	gql__type_User             *graphql.Object      // message User in session/v1alpha1/session.proto
	gql__type_GetUserResponse  *graphql.Object      // message GetUserResponse in session/v1alpha1/session.proto
	gql__type_GetUserRequest   *graphql.Object      // message GetUserRequest in session/v1alpha1/session.proto
	gql__input_User            *graphql.InputObject // message User in session/v1alpha1/session.proto
	gql__input_GetUserResponse *graphql.InputObject // message GetUserResponse in session/v1alpha1/session.proto
	gql__input_GetUserRequest  *graphql.InputObject // message GetUserRequest in session/v1alpha1/session.proto
)

func Gql__enum_Gender() *graphql.Enum {
	if gql__enum_Gender == nil {
		gql__enum_Gender = graphql.NewEnum(graphql.EnumConfig{
			Name: "Sessionv1Alpha1_Enum_Gender",
			Values: graphql.EnumValueConfigMap{
				"GENDER_UNSPECIFIED": &graphql.EnumValueConfig{
					Value: Gender(0),
				},
				"GENDER_MALE": &graphql.EnumValueConfig{
					Value: Gender(1),
				},
				"GENDER_FEMALE": &graphql.EnumValueConfig{
					Value: Gender(2),
				},
			},
		})
	}
	return gql__enum_Gender
}

func Gql__type_User() *graphql.Object {
	if gql__type_User == nil {
		gql__type_User = graphql.NewObject(graphql.ObjectConfig{
			Name: "Sessionv1Alpha1_Type_User",
			Fields: graphql.Fields{
				"user_id": &graphql.Field{
					Type: graphql.String,
				},
				"first_name": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"last_name": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"gender": &graphql.Field{
					Type:        Gql__enum_Gender(),
					Description: `google.type.Date birthday = 4;`,
				},
			},
		})
	}
	return gql__type_User
}

func Gql__type_GetUserResponse() *graphql.Object {
	if gql__type_GetUserResponse == nil {
		gql__type_GetUserResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Sessionv1Alpha1_Type_GetUserResponse",
			Fields: graphql.Fields{
				"user": &graphql.Field{
					Type: graphql.NewNonNull(Gql__type_User()),
				},
			},
		})
	}
	return gql__type_GetUserResponse
}

func Gql__type_GetUserRequest() *graphql.Object {
	if gql__type_GetUserRequest == nil {
		gql__type_GetUserRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Sessionv1Alpha1_Type_GetUserRequest",
			Fields: graphql.Fields{
				"user_id": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
		})
	}
	return gql__type_GetUserRequest
}

func Gql__input_User() *graphql.InputObject {
	if gql__input_User == nil {
		gql__input_User = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Sessionv1Alpha1_Input_User",
			Fields: graphql.InputObjectConfigFieldMap{
				"user_id": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"first_name": &graphql.InputObjectFieldConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"last_name": &graphql.InputObjectFieldConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"gender": &graphql.InputObjectFieldConfig{
					Description: `google.type.Date birthday = 4;`,
					Type:        Gql__enum_Gender(),
				},
			},
		})
	}
	return gql__input_User
}

func Gql__input_GetUserResponse() *graphql.InputObject {
	if gql__input_GetUserResponse == nil {
		gql__input_GetUserResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Sessionv1Alpha1_Input_GetUserResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"user": &graphql.InputObjectFieldConfig{
					Type: graphql.NewNonNull(Gql__input_User()),
				},
			},
		})
	}
	return gql__input_GetUserResponse
}

func Gql__input_GetUserRequest() *graphql.InputObject {
	if gql__input_GetUserRequest == nil {
		gql__input_GetUserRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Sessionv1Alpha1_Input_GetUserRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"user_id": &graphql.InputObjectFieldConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
		})
	}
	return gql__input_GetUserRequest
}

// graphql__resolver_SessionService is a struct for making query, mutation and resolve fields.
// This struct must be implemented runtime.SchemaBuilder interface.
type graphql__resolver_SessionService struct {

	// Automatic connection host
	host string

	// grpc dial options
	dialOptions []grpc.DialOption

	// grpc client connection.
	// this connection may be provided by user
	conn *grpc.ClientConn
}

// new_graphql_resolver_SessionService creates pointer of service struct
func new_graphql_resolver_SessionService(conn *grpc.ClientConn, host string, opts []grpc.DialOption) *graphql__resolver_SessionService {
	return &graphql__resolver_SessionService{
		conn:        conn,
		host:        host,
		dialOptions: opts,
	}
}

// CreateConnection() returns grpc connection which user specified or newly connected and closing function
func (x *graphql__resolver_SessionService) CreateConnection(ctx context.Context) (*grpc.ClientConn, func(), error) {
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
func (x *graphql__resolver_SessionService) GetQueries(conn *grpc.ClientConn) graphql.Fields {
	return graphql.Fields{
		"getUser": &graphql.Field{
			Type: Gql__type_GetUserResponse(),
			Args: graphql.FieldConfigArgument{
				"user_id": &graphql.ArgumentConfig{
					Type:         graphql.NewNonNull(graphql.String),
					DefaultValue: "",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req GetUserRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for getUser")
				}
				client := NewSessionServiceClient(conn)
				resp, err := client.GetUser(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC GetUser")
				}
				return resp, nil
			},
		},
	}
}

// GetMutations returns acceptable graphql.Fields for Mutation.
func (x *graphql__resolver_SessionService) GetMutations(conn *grpc.ClientConn) graphql.Fields {
	return graphql.Fields{}
}

// Register package divided graphql handler "without" *grpc.ClientConn,
// therefore gRPC connection will be opened and closed automatically.
// Occasionally you may worry about open/close performance for each handling graphql request,
// then you can call RegisterSessionServiceGraphqlHandler with *grpc.ClientConn manually.
func RegisterSessionServiceGraphql(mux *runtime.ServeMux, host string, opts ...grpc.DialOption) error {
	return RegisterSessionServiceGraphqlHandler(mux, nil, host, opts...)
}

// Register package divided graphql handler "with" *grpc.ClientConn.
// this function accepts your defined grpc connection, so that we reuse that and never close connection inside.
// You need to close it manually when application will terminate.
func RegisterSessionServiceGraphqlHandler(mux *runtime.ServeMux, conn *grpc.ClientConn, host string, opts ...grpc.DialOption) error {
	return mux.AddHandler(new_graphql_resolver_SessionService(conn, host, opts))
}
