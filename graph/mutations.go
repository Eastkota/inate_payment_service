package schema

import (
	"inapp_payment_service/resolvers"

	"github.com/graphql-go/graphql"
)


func NewMutationType(resolver *resolvers.InatePaymentResolver) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"verifyIapReceipt": &graphql.Field{
				Type: GenericInatePaymentSuccessResponse,
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Type: IapReceiptInput,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					// return PublicAuthMiddleware(resolver.VerifyIapReceipt)(p), nil
					return resolver.VerifyIapReceipt(p), nil
				},
			},
		},
	})
}
