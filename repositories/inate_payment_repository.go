package repositories

import (
	"inapp_payment_service/model"

	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InatePaymentRepository struct{
	DB *gorm.DB
}

func NewInatePaymentRepository(db *gorm.DB) *InatePaymentRepository {
	return &InatePaymentRepository{DB: db}
}

func (repo *InatePaymentRepository) StoreIapTransaction(iapInput model.IapTransactionInput) (*model.IapTransaction, error) {
    transaction := &model.IapTransaction{
		ID:                   uuid.New(),
		TransactionId:        iapInput.TransactionId,
		Product:            iapInput.Product,
		ReceiptData:          iapInput.ReceiptData,
		PurchaseDate:         iapInput.PurchaseDate,
		UserId:               iapInput.UserId,
		MembershipDurationId: iapInput.MembershipDurationId,
		Amount:			   iapInput.Amount,
		Status:               iapInput.Status,
	}

    result := repo.DB.Create(transaction)

    if result.Error != nil {
        return nil, fmt.Errorf("failed to save the IAP transaction: %v", result.Error)
    }

    if result.RowsAffected == 0 {
        return nil, fmt.Errorf("no rows were inserted")
    }
    
    return transaction, nil
}
