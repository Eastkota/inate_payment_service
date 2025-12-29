package resolvers

import (
	"inate_payment_service/helpers"
	"inate_payment_service/model"
	"inate_payment_service/services"
	
	"github.com/graphql-go/graphql"
	"github.com/google/uuid"
)

type InatePaymentResolver struct {
	Services services.Services
}

func NewInatePaymentResolver(service services.Services) *InatePaymentResolver {

	return &InatePaymentResolver{Services: service}
}

func (r *InatePaymentResolver) VerifyIapReceipt(p graphql.ResolveParams) *model.GenericInatePaymentResponse {
	inputData := p.Args["input"].(map[string]interface{})
	receiptData := inputData["receipt_data"].(string)
	var userId *uuid.UUID
    if userData, ok := inputData["user_id"].(string); ok && userData != "" {
        parsedID, err := uuid.Parse(userData)
        if err == nil {
            userId = &parsedID
        }
    }

	membershipDurationId := inputData["membership_duration_id"].(uuid.UUID)
	product := inputData["product"].(string)
	amount := inputData["amount"].(float64)


	result, err := r.Services.VerifyIapReceipt(receiptData, userId, membershipDurationId, product, amount)
	if err != nil {
		return helpers.FormatError(err)
	}

	return &model.GenericInatePaymentResponse{
		Data: result,
		Error: nil,
	}
}
