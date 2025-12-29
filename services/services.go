package services

import (
	"inate_payment_service/model"
	"github.com/google/uuid"
)

type Services interface {
	VerifyIapReceipt(receiptData string, userId *uuid.UUID, membershipDurationId uuid.UUID, productId string, amount float64) (*model.IapTransaction, error)
}