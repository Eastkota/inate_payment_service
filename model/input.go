package model

import (
	"time"
	"github.com/google/uuid"
)

type InatePaymentResponseInput struct {
	BfsBfsTxnId          string `json:"bfs_bfsTxnId"`
	BfsDebitAuthNo       string `json:"bfs_debitAuthNo"`
	BfsRemitterName      string `json:"bfs_remitterName"`
	BfsTxnCurrency       string `json:"bfs_txnCurrency"`
	BfsBfsTxnTime        string `json:"bfs_bfsTxnTime"`
	BfsBenfId            string `json:"bfs_benfId"`
	BfsRemitterBankId    string `json:"bfs_remitterBankId"`
	BfsOrderNo           string `json:"bfs_orderNo"`
	BfsDebitAuthCode     string `json:"bfs_debitAuthCode"`
	BfsTxnAmount         string `json:"bfs_txnAmount"`
	BfsBenfTxnTime       string `json:"bfs_benfTxnTime"`
	BfsMsgType           string `json:"bfs_msgType"`
	UserId               uuid.UUID `json:"user_id"`
	MembershipDurationId uuid.UUID `json:"membership_duration_id"`
	Remarks              string `json:"remarks"`
}

type ArRequestPayload struct {
	BfsMsgType       string `json:"bfs_msgType"`
	BfsBfsTxnTime    string `json:"bfs_bfsTxnTime"`
	BfsOrderNo       string `json:"bfs_orderNo"`
	BfsBenfId        string `json:"bfs_benfId"`
	BfsBenfBankCode  string `json:"bfs_benfBankCode"`
	BfsTxnCurrency   string `json:"bfs_txnCurrency"`
	BfsTxnAmount     string `json:"bfs_txnAmount"`
	BfsRemitterEmail string `json:"bfs_remitterEmail"`
	BfsPaymentDesc   string `json:"bfs_paymentDesc"`
	BfsVersion       string `json:"bfs_version"`
	BfsCheckSum      string `json:"bfs_checkSum"`
}

type UserActivityInput struct {
    Activity string `json:"activity"`
    UserID        uuid.UUID `json:"user_id"`
    Count   int     `json:"count"`
    Month           int         `json:"month"`             
    Year            int         `json:"year"`
}

type IapTransactionInput struct {
	UserId               *uuid.UUID `json:"user_id"`
	MembershipDurationId uuid.UUID `json:"membership_duration_id"`
	Product            string    `json:"product"`
	ReceiptData          string    `json:"receipt_data"`
	TransactionId        string    `json:"transaction_id,omitempty"`
	PurchaseDate         time.Time `json:"purchase_date,omitempty"`
	ExpiresDate          string    `json:"expires_date,omitempty"`
	Amount               float64   `json:"amount"`
	Status               string    `json:"status"`
}
