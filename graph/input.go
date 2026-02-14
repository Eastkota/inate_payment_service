package schema

import (
	"inapp_payment_service/graph/scalar"

	"github.com/graphql-go/graphql"
)

var IapReceiptInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "IapReceiptInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"receipt_data": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"user_id": &graphql.InputObjectFieldConfig{
				Type: scalar.UUID,
			},
			"membership_duration_id": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(scalar.UUID),
			},
			"product": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"amount": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Float),
			},
		},
	},
)
