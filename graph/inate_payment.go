package schema

import "github.com/graphql-go/graphql"

var GenericInatePaymentSuccessData = graphql.NewObject(graphql.ObjectConfig{
	Name: "GenericInatePaymentSuccessData",
	Fields: graphql.Fields{
		"code":    &graphql.Field{Type: graphql.String},
		"Message": &graphql.Field{Type: graphql.String},
	},
})


var GenericInatePaymentSuccessDataResult = graphql.NewObject(graphql.ObjectConfig{
	Name: "GenericInatePaymentSuccessDataResult",
	Fields: graphql.Fields{
		"generic_success_response": &graphql.Field{Type: GenericInatePaymentSuccessData},
	},
})

var GenericAeInatePaymentSuccessData = graphql.NewObject(graphql.ObjectConfig{
	Name: "GenericAeInatePaymentSuccessData",
	Fields: graphql.Fields{
		"txn_id":    &graphql.Field{Type: graphql.String},
	},
})

var IapTransactionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "IapTransaction",
	Fields: graphql.Fields{
		"transaction_id":        &graphql.Field{Type: graphql.String},
		"product":             &graphql.Field{Type: graphql.String},
		"receipt_data":          &graphql.Field{Type: graphql.String},
		"purchase_date":         &graphql.Field{Type: graphql.DateTime},
		"user_id":               &graphql.Field{Type: graphql.String}, // UUID as string
		"membership_duration_id": &graphql.Field{Type: graphql.String}, // UUID as string
		"amount":               &graphql.Field{Type: graphql.Float},
		"status":               &graphql.Field{Type: graphql.String},
		"created_at":            &graphql.Field{Type: graphql.DateTime},
		"updated_at":            &graphql.Field{Type: graphql.DateTime},
	},
})
