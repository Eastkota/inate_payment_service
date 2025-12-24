package services

import (
	"inate_payment_service/helpers"
	"inate_payment_service/model"
	"inate_payment_service/repositories"

	"fmt"
	"time"
	"github.com/google/uuid"
)

type InatePaymentService struct {
	Repository repositories.Repository // Inject Repository
}

func NewInatePaymentService(repository repositories.Repository) *InatePaymentService {
	return &InatePaymentService{Repository: repository}
}

func (ms *InatePaymentService) VerifyIapReceipt(transactionId string, userId, membershipDurationId uuid.UUID, product string, amount float64) (*model.IapTransaction, error) {
    transaction, err := helpers.VerifyAppleTransaction(transactionId, product, amount)
    

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
        return nil, fmt.Errorf("product mismatch")
    }

    storedTransaction, storeErr := ms.Repository.StoreIapTransaction(iapInput)
    if storeErr != nil {
        return nil, storeErr
    }

    if err != nil {
        return storedTransaction, fmt.Errorf("apple verification failed: %v", err)
    }

    return storedTransaction, nil
}
