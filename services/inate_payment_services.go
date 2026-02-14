package services

import (
	"inapp_payment_service/helpers"
	"inapp_payment_service/model"
	"inapp_payment_service/repositories"

	"time"
	"github.com/google/uuid"
)

type InatePaymentService struct {
	Repository repositories.Repository // Inject Repository
}

func NewInatePaymentService(repository repositories.Repository) *InatePaymentService {
	return &InatePaymentService{Repository: repository}
}

func (ms *InatePaymentService) VerifyIapReceipt(transactionId string, userId *uuid.UUID, membershipDurationId uuid.UUID, product string, amount float64) (*model.IapTransaction, error) {
    transaction, err := helpers.VerifyAppleTransaction(transactionId, product, amount)
    if err != nil {
        return nil, err
    }

    iapInput := model.IapTransactionInput{
        UserId:               userId,
        MembershipDurationId: membershipDurationId,
        Product:              transaction.ProductId,
        TransactionId:        transaction.TransactionId,
        PurchaseDate:         time.Now(),
        Amount:               transaction.Amount,
        Status:               transaction.Status,
    }

    if err == nil && transaction.ProductId != product {
        return nil, helpers.NewValidationError("product mismatch", "product")
    }

    storedTransaction, storeErr := ms.Repository.StoreIapTransaction(iapInput)
    if storeErr != nil {
        return nil, storeErr
    }

    if err != nil {
        return storedTransaction, helpers.WrapInternal("apple verification", err)
    }

    return storedTransaction, nil
}
