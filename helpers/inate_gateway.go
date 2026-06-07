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

    "inapp_payment_service/config"
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
    // Simulation backdoor — only when IAP_SIMULATED_MODE=true. Real receipts
    // from the App Store contain a "." (they're JWS tokens), so without this
    // gate the entire verification step would be skipped in production.
    if config.IapSimulatedMode() && strings.Contains(receiptData, ".") {
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

    if config.AppleSharedSecret() == "" {
        return nil, fmt.Errorf("APPLE_SHARED_SECRET not configured")
    }

    payload := map[string]interface{}{
        "receipt-data": receiptData,
        "password":     config.AppleSharedSecret(),
    }

    jsonData, err := json.Marshal(payload)
    if err != nil {
        return nil, err
    }

    client := &http.Client{Timeout: 30 * time.Second}
    resp, err := client.Post(config.AppleVerifyURL(), "application/json", bytes.NewBuffer(jsonData))
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

    // Apple returns 21007 when a sandbox receipt is posted to the production
    // endpoint, and 21008 in the opposite case. Retry once on the other
    // endpoint per Apple's recommendation.
    if appleResp.Status == 21007 || appleResp.Status == 21008 {
        retryURL := config.AppleVerifyURLSandbox
        if appleResp.Status == 21008 {
            retryURL = config.AppleVerifyURLProduction
        }
        resp2, err := client.Post(retryURL, "application/json", bytes.NewBuffer(jsonData))
        if err != nil {
            return nil, err
        }
        defer resp2.Body.Close()
        body2, err := io.ReadAll(resp2.Body)
        if err != nil {
            return nil, err
        }
        if err := json.Unmarshal(body2, &appleResp); err != nil {
            return nil, err
        }
    }

    if appleResp.Status != 0 {
        return nil, fmt.Errorf("receipt verification failed: %s", MapAppleToBankCode(appleResp.Status))
    }

    // Bundle ID check — reject receipts that didn't originate from our app.
    // Skipped when APPLE_BUNDLE_ID is unset so non-prod installs don't break.
    if want := config.AppleBundleID(); want != "" && appleResp.Receipt.BundleID != want {
        return nil, fmt.Errorf("bundle id mismatch: got %q want %q", appleResp.Receipt.BundleID, want)
    }

    if len(appleResp.Receipt.InApp) == 0 {
        return nil, fmt.Errorf("no in-app purchases found")
    }

    inApp := appleResp.Receipt.InApp[0]

    purchaseDate, err := strconv.ParseInt(inApp.PurchaseDateMS, 10, 64)
    if err != nil {
        return nil, err
    }

    bankCode := MapAppleToBankCode(appleResp.Status)

    return &Transaction{
        TransactionId: inApp.TransactionID,
        ProductId:     inApp.ProductID,
        PurchaseDate:  purchaseDate,
        Status:        bankCode,
    }, nil
}