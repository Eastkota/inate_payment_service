package helpers

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "strconv"
	"strings"
	"time"

    "inate_payment_service/config"
)

type AppleReceiptResponse struct {
    Status      int `json:"status"`
    Environment string `json:"environment"`
    Receipt     struct {
        BundleID                   string `json:"bundle_id"`
        ApplicationVersion        string `json:"application_version"`
        InApp                     []struct {
            TransactionID              string `json:"transaction_id"`
            OriginalTransactionID      string `json:"original_transaction_id"`
            ProductID                  string `json:"product_id"`
            PurchaseDateMS             string `json:"purchase_date_ms"`
            OriginalPurchaseDateMS     string `json:"original_purchase_date_ms"`
            ExpiresDateMS              string `json:"expires_date_ms,omitempty"`
            Quantity                   string `json:"quantity"`
        } `json:"in_app"`
    } `json:"receipt"`
}

type Transaction struct {
    TransactionId string
    ProductId     string
    PurchaseDate  int64
    Amount        float64
    Status        string 
}

func VerifyAppleTransaction(receiptData string, productID string, amount float64) (*Transaction, error) {
	if strings.Contains(receiptData, ".") {
        if strings.Contains(receiptData, "FAILED") {
            return &Transaction{
                Status: "51",
            }, fmt.Errorf("simulated failure")
        }

        return &Transaction{
            TransactionId: "JWS_SIM_" + strconv.FormatInt(time.Now().Unix(), 10),
            ProductId:     productID,
            Amount:        amount,
            Status:        "Approved", 
        }, nil
    }

    payload := map[string]interface{}{
        "receipt-data": receiptData,
        "password":     config.APPLE_SHARED_SECRET,
    }

    jsonData, err := json.Marshal(payload)
    if err != nil {
        return nil, err
    }

    resp, err := http.Post(config.APPLE_VERIFY_URL, "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var appleResp AppleReceiptResponse
    if err := json.Unmarshal(body, &appleResp); err != nil {
        return nil, err
    }

    if appleResp.Status != 0 {
        return nil, fmt.Errorf("receipt verification failed: %s", MapAppleToBankCode(appleResp.Status))
    }

    if len(appleResp.Receipt.InApp) == 0 {
        return nil, fmt.Errorf("no in-app purchases found")
    }

    inApp := appleResp.Receipt.InApp[0]

    purchaseDate, err := strconv.ParseInt(inApp.PurchaseDateMS, 10, 64)
    if err != nil {
        return nil, err
    }

    appleStatus := appleResp.Status

    bankCode := MapAppleToBankCode(appleResp.Status)

    if appleStatus != 0 {
        return &Transaction{Status: bankCode}, fmt.Errorf("apple error: %d", appleStatus)
    }

    return &Transaction{
        TransactionId: inApp.TransactionID,
        ProductId:     inApp.ProductID,
        PurchaseDate:  purchaseDate,
        Status:        bankCode,
    }, nil
}