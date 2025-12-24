package schema

import "github.com/graphql-go/graphql"

var InatePaymentError = graphql.NewObject(graphql.ObjectConfig{
	Name: "InatePaymentError",
	Fields: graphql.Fields{
		"message": &graphql.Field{Type: graphql.String},
		"code":    &graphql.Field{Type: graphql.String},
		"field":   &graphql.Field{Type: graphql.String},
	},
})
