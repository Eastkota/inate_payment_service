package helpers

import (
	"fmt"
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
		return fmt.Errorf("the provided email, was not a valid email: %v", remitterEmail)
	}
	if amount <= 0 {
		return fmt.Errorf("provide valid amount")
	}
	if product == "" {
		return fmt.Errorf("product is required")
	}
	return nil
}

func ValidateAeRequest(remitterAccNo, remitterBankId, txnId string) error {
	if remitterAccNo == "" {
		return fmt.Errorf("account number is required")
	}
	if txnId == "" {
		return fmt.Errorf("transaction id is required")
	}
	if remitterBankId == "" {
		return fmt.Errorf("please choose the bank")
	}
	return nil
}

func ValidateDrRequest(txnId, remitterOtp string, userId, membershipDurationId uuid.UUID) error {
	if remitterOtp == "" {
		return fmt.Errorf("otp is required")
	}
	if txnId == "" {
		return fmt.Errorf("transaction is required")
	}
	if membershipDurationId == uuid.Nil {
		return fmt.Errorf("membership duration is required")
	}
	return nil
}
