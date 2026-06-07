package helpers

import (
	"context"
	"fmt"
	"inapp_payment_service/config"
	"log"
	"math"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	defaultDiscountPercent  = 30.0
	defaultDiscountEndDate  = "2027-07-31"
	discountDateLayout      = "2006-01-02"
	discountAmountTolerance = 0.01
)

var (
	couponDB   *gorm.DB
	couponDBMu sync.Mutex
)

// getCouponDB lazily opens a pooled connection to the coupon DB. On failure
// (DB unreachable at boot, bad config, etc.) it returns the error and leaves
// the cache empty, so the next call will retry.
func getCouponDB() (*gorm.DB, error) {
	couponDBMu.Lock()
	defer couponDBMu.Unlock()
	if couponDB != nil {
		return couponDB, nil
	}
	host := config.CouponPostgresHost()
	if host == "" {
		return nil, fmt.Errorf("COUPON_POSTGRES_HOST not configured")
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host,
		config.CouponPostgresUser(),
		config.CouponPostgresPassword(),
		config.CouponPostgresDB(),
		config.CouponPostgresPort(),
	)
	gormLogger := logger.New(
		log.New(os.Stdout, "[coupon-db] ", log.LstdFlags),
		logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
		},
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: gormLogger})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(5)
	sqlDB.SetMaxIdleConns(2)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	couponDB = db
	return couponDB, nil
}

func discountPercent() float64 {
	raw := config.CouponDiscountPercent()
	if raw == "" {
		return defaultDiscountPercent
	}
	pct, err := strconv.ParseFloat(raw, 64)
	if err != nil || pct < 0 || pct > 100 {
		return defaultDiscountPercent
	}
	return pct
}

func discountEndDate() time.Time {
	raw := config.CouponDiscountEndDate()
	if raw == "" {
		raw = defaultDiscountEndDate
	}
	t, err := time.ParseInLocation(discountDateLayout, raw, time.Local)
	if err != nil {
		t, _ = time.ParseInLocation(discountDateLayout, defaultDiscountEndDate, time.Local)
	}
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
}

func IsDiscountWindowOpen() bool {
	return !time.Now().After(discountEndDate())
}

func DiscountedAmount(original float64) float64 {
	if original <= 0 {
		return original
	}
	discounted := original * (1.0 - discountPercent()/100.0)
	return math.Round(discounted*100) / 100
}

func hasUserRedeemedCoupon(ctx context.Context, userId uuid.UUID) (bool, error) {
	if userId == uuid.Nil {
		return false, nil
	}
	db, err := getCouponDB()
	if err != nil {
		return false, err
	}
	var count int64
	if err := db.WithContext(ctx).
		Table("coupon.coupons").
		Where("user_id = ? AND redeemed = ?", userId, true).
		Limit(1).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func IsUserEligibleForCouponDiscount(ctx context.Context, userId uuid.UUID) bool {
	if userId == uuid.Nil || !IsDiscountWindowOpen() {
		return false
	}
	ok, err := hasUserRedeemedCoupon(ctx, userId)
	if err != nil {
		log.Printf("[DISCOUNT] eligibility lookup failed user=%s: %v", userId, err)
		return false
	}
	return ok
}

// fetchDurationUSDPrices returns the full USD price plus the optional
// pre-registered discounted IAP price for a membership duration. discountedIap
// is 0 when not configured.
func fetchDurationUSDPrices(ctx context.Context, membershipDurationId uuid.UUID) (price float64, discountedIap float64, err error) {
	db, err := GetGormDB()
	if err != nil {
		return 0, 0, err
	}
	var (
		fullPrice    float32
		discountedPx float32
	)
	row := db.WithContext(ctx).
		Table("subscription.membership_durations").
		Where("id = ?", membershipDurationId).
		Select("price_usd, COALESCE(discounted_usd_iap, 0)").
		Row()
	if err := row.Scan(&fullPrice, &discountedPx); err != nil {
		return 0, 0, err
	}
	return float64(fullPrice), float64(discountedPx), nil
}

// WarnOnAmountMismatch logs (but never blocks) when paidAmount diverges from
// the expected USD price after applying the coupon discount if eligible.
// Bounded by a short timeout so a stalled coupon DB never delays a payment
// response.
//
// For eligible users we prefer the exact pre-registered IAP discounted price
// (discounted_usd_iap) over the mathematically computed 30%-off price, because
// Apple price tiers may not land exactly on full*0.70 and the discrepancy
// would otherwise show up as a constant [DISCOUNT-MISMATCH] log noise.
func WarnOnAmountMismatch(ctx context.Context, userId, membershipDurationId uuid.UUID, paidAmount float64) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	fullPrice, discountedIap, err := fetchDurationUSDPrices(ctx, membershipDurationId)
	if err != nil {
		log.Printf("[DISCOUNT-LOOKUP] price lookup failed user=%s duration=%s: %v", userId, membershipDurationId, err)
		return
	}
	if fullPrice <= 0 {
		return
	}
	eligible := IsUserEligibleForCouponDiscount(ctx, userId)
	expected := fullPrice
	if eligible {
		if discountedIap > 0 {
			expected = discountedIap
		} else {
			expected = DiscountedAmount(fullPrice)
		}
	}
	if math.Abs(paidAmount-expected) > discountAmountTolerance {
		log.Printf("[DISCOUNT-MISMATCH] user=%s duration=%s currency=USD paid=%.2f expected=%.2f eligible=%v",
			userId, membershipDurationId, paidAmount, expected, eligible)
	}
}
