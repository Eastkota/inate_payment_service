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
		return NewValidationError("the provided email was not a valid email", "remitter_email")
	}
	if amount <= 0 {
		return NewValidationError("provide valid amount", "amount")
	}
	if product == "" {
		return NewValidationError("product is required", "product")
	}
	return nil
}

func ValidateAeRequest(remitterAccNo, remitterBankId, txnId string) error {
	if remitterAccNo == "" {
		return NewValidationError("account number is required", "remitter_acc_no")
	}
	if txnId == "" {
		return NewValidationError("transaction id is required", "txn_id")
	}
	if remitterBankId == "" {
		return NewValidationError("please choose the bank", "remitter_bank_id")
	}
	return nil
}

func ValidateDrRequest(txnId, remitterOtp string, userId, membershipDurationId uuid.UUID) error {
	if remitterOtp == "" {
		return NewValidationError("otp is required", "remitter_otp")
	}
	if txnId == "" {
		return NewValidationError("transaction is required", "txn_id")
	}
	if membershipDurationId == uuid.Nil {
		return NewValidationError("membership duration is required", "membership_duration_id")
	}
	return nil
}
