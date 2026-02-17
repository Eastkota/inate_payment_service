package helpers

import (
	"regexp"

	"github.com/google/uuid"
)

func isValidEmail(email string) bool {
	// Regular expression for validating an email address
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regex
	re := regexp.MustCompile(emailRegex)

	// Return whether the email matches the regex
	return re.MatchString(email)
}

func ValidateArRequest(amount float64, product, remitterEmail string) error {
	if !isValidEmail(remitterEmail) {
		return NewValidationError("Please provide a valid email address", "remitter_email")
	}
	if amount <= 0 {
		return NewValidationError("Please provide a valid amount", "amount")
	}
	if product == "" {
		return NewValidationError("Product is required", "product")
	}
	return nil
}

func ValidateAeRequest(remitterAccNo, remitterBankId, txnId string) error {
	if remitterAccNo == "" {
		return NewValidationError("Account number is required", "remitter_acc_no")
	}
	if txnId == "" {
		return NewValidationError("Transaction ID is required", "txn_id")
	}
	if remitterBankId == "" {
		return NewValidationError("Please choose a bank", "remitter_bank_id")
	}
	return nil
}

func ValidateDrRequest(txnId, remitterOtp string, userId, membershipDurationId uuid.UUID) error {
	if remitterOtp == "" {
		return NewValidationError("Verification code is required", "remitter_otp")
	}
	if txnId == "" {
		return NewValidationError("Transaction is required", "txn_id")
	}
	if membershipDurationId == uuid.Nil {
		return NewValidationError("Membership duration is required", "membership_duration_id")
	}
	return nil
}
