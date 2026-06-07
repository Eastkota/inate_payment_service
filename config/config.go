package config

import (
	"os"
)

func PostgresUser() string     { return os.Getenv("POSTGRES_USER") }
func PostgresPassword() string { return os.Getenv("POSTGRES_PASSWORD") }
func PostgresHost() string     { return os.Getenv("POSTGRES_HOST") }
func PostgresPort() string     { return os.Getenv("POSTGRES_PORT") }
func PostgresDB() string       { return os.Getenv("POSTGRES_DB") }

const (
	PaymentResponseCollection = "payment_responses"
	BENF_ID                   = "BE10000169"
	BENF_BANK_CODE            = "01"
	CURRENCY                  = "BTN"
	PAYMENT_PRIVATE_KEY       = "keys/inate-payment.key"
	APPLE_SHARED_SECRET       = "your_shared_secret_from_app_store_connect"
	APPLE_VERIFY_URL          = "https://sandbox.itunes.apple.com/verifyReceipt"
	// APPLE_VERIFY_URL          = "https://buy.itunes.apple.com/verifyReceipt"
	APPLE_BUNDLE_ID           = "com.yourcompany.yourapp"
	APPLE_PRIVATE_KEY         = "your_private_key_content_here" // Replace with actual .p8 content
	APPLE_KEY_ID              = "your_key_id_here"              // e.g., "2X9R4HXF34"
	APPLE_ISSUER_ID           = "your_issuer_id_here"           // Found in App Store Connect
)

func MembershipApi() string  { return os.Getenv("MEMBERSHIP_API") }
func AuthServiceApi() string { return os.Getenv("AUTH_SERVICE_API") }

func CouponPostgresUser() string     { return os.Getenv("COUPON_POSTGRES_USER") }
func CouponPostgresPassword() string { return os.Getenv("COUPON_POSTGRES_PASSWORD") }
func CouponPostgresHost() string     { return os.Getenv("COUPON_POSTGRES_HOST") }
func CouponPostgresPort() string     { return os.Getenv("COUPON_POSTGRES_PORT") }
func CouponPostgresDB() string       { return os.Getenv("COUPON_POSTGRES_DB") }

func CouponDiscountEndDate() string { return os.Getenv("COUPON_DISCOUNT_END_DATE") }
func CouponDiscountPercent() string { return os.Getenv("COUPON_DISCOUNT_PERCENT") }

