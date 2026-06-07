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

	// Apple verifyReceipt endpoints. Production is the safer default; sandbox
	// must be opted into via APPLE_VERIFY_URL env var for test builds.
	AppleVerifyURLProduction = "https://buy.itunes.apple.com/verifyReceipt"
	AppleVerifyURLSandbox    = "https://sandbox.itunes.apple.com/verifyReceipt"
)

// Apple IAP credentials. All driven by env so prod / staging / dev can use
// different App Store Connect keys without rebuild.
func AppleSharedSecret() string { return os.Getenv("APPLE_SHARED_SECRET") }
func AppleBundleID() string     { return os.Getenv("APPLE_BUNDLE_ID") }

// AppleVerifyURL returns the configured verifyReceipt endpoint, defaulting to
// the production URL when unset (safer than silently using sandbox in prod).
func AppleVerifyURL() string {
	if v := os.Getenv("APPLE_VERIFY_URL"); v != "" {
		return v
	}
	return AppleVerifyURLProduction
}

// App Store Server API JWT auth (unused today; kept here so the env names are
// stable once the new flow is wired in).
func ApplePrivateKey() string { return os.Getenv("APPLE_PRIVATE_KEY") }
func AppleKeyID() string      { return os.Getenv("APPLE_KEY_ID") }
func AppleIssuerID() string   { return os.Getenv("APPLE_ISSUER_ID") }

// IapSimulatedMode gates the testing-only short-circuit in VerifyAppleTransaction
// that accepts any receipt containing "." (i.e. any JWS-looking string) as
// "Approved" without contacting Apple. Must be "true" to enable.
func IapSimulatedMode() bool { return os.Getenv("IAP_SIMULATED_MODE") == "true" }

func MembershipApi() string  { return os.Getenv("MEMBERSHIP_API") }
func AuthServiceApi() string { return os.Getenv("AUTH_SERVICE_API") }

func CouponPostgresUser() string     { return os.Getenv("COUPON_POSTGRES_USER") }
func CouponPostgresPassword() string { return os.Getenv("COUPON_POSTGRES_PASSWORD") }
func CouponPostgresHost() string     { return os.Getenv("COUPON_POSTGRES_HOST") }
func CouponPostgresPort() string     { return os.Getenv("COUPON_POSTGRES_PORT") }
func CouponPostgresDB() string       { return os.Getenv("COUPON_POSTGRES_DB") }

func CouponDiscountEndDate() string { return os.Getenv("COUPON_DISCOUNT_END_DATE") }
func CouponDiscountPercent() string { return os.Getenv("COUPON_DISCOUNT_PERCENT") }

