package model


type GenericInatePaymentSuccessData struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type GenericAeInatePaymentSuccessData struct {
	TxnId string `json:"txn_id"`
}

type GenericInatePaymentResponse struct {
	Data  interface{}
	Error *InatePaymentError
}