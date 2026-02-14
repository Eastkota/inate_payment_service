package helpers

import "fmt"

type ErrorCode string

const (
	ErrCodeValidation   ErrorCode = "VALIDATION_ERROR"
	ErrCodeNotFound     ErrorCode = "NOT_FOUND"
	ErrCodeUnauthorized ErrorCode = "UNAUTHORIZED"
	ErrCodeForbidden    ErrorCode = "FORBIDDEN"
	ErrCodeConflict     ErrorCode = "CONFLICT"
	ErrCodeInternal     ErrorCode = "INTERNAL_ERROR"
	ErrCodeRateLimited  ErrorCode = "RATE_LIMITED"
)

type AppError struct {
	Code     ErrorCode
	Message  string
	Field    string
	Internal error
}

func (e *AppError) Error() string { return e.Message }
func (e *AppError) Unwrap() error { return e.Internal }

func NewValidationError(message string, field string) *AppError {
	return &AppError{Code: ErrCodeValidation, Message: message, Field: field}
}
func NewNotFoundError(message string) *AppError {
	return &AppError{Code: ErrCodeNotFound, Message: message}
}
func NewUnauthorizedError(message string) *AppError {
	return &AppError{Code: ErrCodeUnauthorized, Message: message}
}
func NewForbiddenError(message string) *AppError {
	return &AppError{Code: ErrCodeForbidden, Message: message}
}
func NewConflictError(message string) *AppError {
	return &AppError{Code: ErrCodeConflict, Message: message}
}
func NewInternalError(message string, internal error) *AppError {
	return &AppError{Code: ErrCodeInternal, Message: message, Internal: internal}
}
func WrapInternal(context string, err error) *AppError {
	return &AppError{Code: ErrCodeInternal, Message: fmt.Sprintf("%s: an unexpected error occurred", context), Internal: err}
}
func NewRateLimitedError(message string) *AppError {
	return &AppError{Code: ErrCodeRateLimited, Message: message}
}
