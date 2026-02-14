package helpers

import "inapp_payment_service/model"

func FormatError(err error) *model.GenericInatePaymentResponse {
	return &model.GenericInatePaymentResponse{
		Data: nil,
		Error: &model.InatePaymentError{
			Message: err.Error(),
		},
	}
}
