package repositories

import "inapp_payment_service/model"

type Repository interface {
	StoreIapTransaction(iapInput model.IapTransactionInput) (*model.IapTransaction, error)
}
