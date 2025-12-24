package schema

import "github.com/graphql-go/graphql"

var GenericInatePaymentSuccessResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "GenericInatePaymentSuccessResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: IapTransactionType},
		"error": &graphql.Field{Type: InatePaymentError},
	},
})
var GenericAeInatePaymentSuccessResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "GenericAeInatePaymentSuccessResponse",
	Fields: graphql.Fields{
		"data": &graphql.Field{Type: GenericAeInatePaymentSuccessData},
		"error": &graphql.Field{Type: InatePaymentError},
	},
})

