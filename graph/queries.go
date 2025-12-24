package schema

import (
	"inate_payment_service/helpers"
	"inate_payment_service/model"
	"inate_payment_service/resolvers"

	"github.com/graphql-go/graphql"
)

// var query = (&queries.Query{Resolver: resolver})
func NewQueryType(resolver *resolvers.InatePaymentResolver) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"service": &graphql.Field{
				Type: graphql.NewNonNull(Service),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					schema, err := GetSchema()
					if err != nil {
						return nil, err
					}

					serviceInfo := model.Service{
						Name:    "InatePaymentService",
						Version: "1.0.0",
						Schema:  helpers.ConvertSchemaToString(schema),
					}
					return serviceInfo, nil
				},
			},
		},
	})
}
