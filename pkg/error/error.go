package error

import (
	"fmt"
)

const (
	StatusOK       = 200
	StatusCreated  = 201
	StatusAccepted = 202

	StatusBadRequest   = 400
	StatusUnauthorized = 401
	StatusNotFound     = 404

	StatusInternalServerError = 500

	errCodeValidation        = 400
	errCodeAuthentication    = 401
	errCodeNotFound          = 404
	errCodeInternalServer    = 500
	errCodeDatabaseOperation = 503
)

var (
	ErrInvalidCartID      = newError(StatusInternalServerError, "invalid cart id")
	ErrInvalidQuantity    = newError(StatusInternalServerError, "invalid quantity")
	ErrCartNotFound       = newError(errCodeNotFound, "cart not found")
	ErrCartItemNotFound   = newError(errCodeNotFound, "cart item not found")
	ErrFailedClearCart    = newError(StatusInternalServerError, "failed to clear cart")
	ErrProductNotFound    = newError(errCodeNotFound, "product not found")
	ErrInvalidInput       = newError(errCodeValidation, "invalid input data")
	ErrInvalidCredentials = newError(errCodeAuthentication, "invalid credentials")
	ErrDatabaseConnection = newError(errCodeDatabaseOperation, "database connection error")
	ErrTimeout            = newError(errCodeInternalServer, "request timeout")

	ErrOpenTransaction   = newError(errCodeInternalServer, "failed to open transaction")
	ErrCommitTransaction = newError(errCodeInternalServer, "failed to commit transaction")
)

type CustomError struct {
	Code    int
	Message string
}

func newError(code int, message string) *CustomError {
	return &CustomError{Code: code, Message: message}
}

func NewErrorWithDetails(code int, message string, details interface{}) *CustomError {
	return &CustomError{
		Code:    code,
		Message: fmt.Sprintf("%s: %v", message, details),
	}
}

func ErrOpenTransactionWithDetails(details interface{}) *CustomError {
	return NewErrorWithDetails(ErrOpenTransaction.Code, ErrOpenTransaction.Message, details)
}

func ErrCommitTransactionWithDetails(details interface{}) *CustomError {
	return NewErrorWithDetails(ErrCommitTransaction.Code, ErrCommitTransaction.Message, details)
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func (e *CustomError) Is(target error) bool {
	t, ok := target.(*CustomError)
	if !ok {
		return false
	}
	return e.Code == t.Code
}

func GetErrorCode(err error) int {
	if customErr, ok := err.(*CustomError); ok {
		return customErr.Code
	}
	return StatusInternalServerError // default error code
}

func GetErrorMessage(err error) string {
	if customErr, ok := err.(*CustomError); ok {
		return customErr.Message
	}
	return "internal server error" // default error message
}

// func GetErrorDetail(err error) interface{} {
// 	if customErr, ok := err.(*CustomError); ok {
// 		return customErr.Detail
// 	}
// 	return nil // default error message
// }
