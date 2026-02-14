package helpers

import (
	"errors"
	"inapp_payment_service/model"
	"log"
)

func FormatError(err error) *model.GenericInatePaymentResponse {
	var appErr *AppError
	if errors.As(err, &appErr) {
		if appErr.Internal != nil {
			log.Printf("[ERROR] code=%s message=%q internal=%q", appErr.Code, appErr.Message, appErr.Internal.Error())
		}
		return &model.GenericInatePaymentResponse{
			Data: nil,
			Error: &model.InatePaymentError{
				Message: appErr.Message,
				Code:    string(appErr.Code),
				Field:   appErr.Field,
			},
		}
	}
	log.Printf("[ERROR] code=INTERNAL_ERROR untyped_error=%q", err.Error())
	return &model.GenericInatePaymentResponse{
		Data: nil,
		Error: &model.InatePaymentError{
			Message: err.Error(),
			Code:    string(ErrCodeInternal),
		},
	}
}
