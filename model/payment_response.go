package model

import (
	"time"
	
	"github.com/google/uuid"
)

type IapTransaction struct {
	ID                   uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	TransactionId        string    `json:"transaction_id" gorm:"type:string"`
	Product            string    `json:"product" gorm:"type:string"`
	ReceiptData          string    `json:"receipt_data" gorm:"type:text"`
	PurchaseDate         time.Time `json:"purchase_date" gorm:"type:timestamptz"`
	UserId               *uuid.UUID `json:"user_id" gorm:"type:uuid"`
	MembershipDurationId uuid.UUID `json:"membership_duration_id" gorm:"type:uuid"`
	Amount			   float64   `json:"amount" gorm:"type:numeric"`
	Status               string    `json:"status" gorm:"type:string"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

type AuthUserMembership struct {
    ID                   uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    MembershipJoinDate   time.Time `gorm:"type:timestamptz;not null" json:"membership_join_date"`
    MembershipEndDate    time.Time `gorm:"type:timestamptz;not null" json:"membership_end_date"`
    UserId               uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
    MembershipDurationId uuid.UUID `gorm:"type:uuid;not null" json:"membership_duration_id"`
    CreatedAt            time.Time `json:"created_at"`
    UpdatedAt            time.Time `json:"updated_at"`
}

func (IapTransaction) TableName() string {
    return "payment.iap_transactions"
}

type PaymentResult struct {
	Message string `json:"message"`
}

type GenericInateResponse struct {
	Data  string `json:"data"`
	Error *InatePaymentError
}
type Bank struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type ArResponseData struct {
	TxnId  string `json:"txnId"`
	Status string `json:"status"`
	Banks  []Bank `json:"banks"`
}
