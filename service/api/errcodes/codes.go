package errcodes

import (
	"github.com/Confialink/wallet-kyc/internal/internal_errors"
	"net/http"
)

const (
	CodeForbidden       = "FORBIDDEN"
	InternalServerError = "INTERNAL_SERVER_ERROR"
	SerializeError      = "SERIALIZE_ERROR"
)

var statusCodes = map[string]int{
	CodeForbidden:       http.StatusForbidden,
	InternalServerError: http.StatusInternalServerError,

	internal_errors.RequestAlreadyExists: http.StatusBadRequest,
	internal_errors.TierNotFound:         http.StatusBadRequest,
	internal_errors.RequestNotFound:      http.StatusBadRequest,
	internal_errors.RequestNotPending:    http.StatusBadRequest,

	internal_errors.LimitingExceededMaxBalance:        http.StatusBadRequest,
	internal_errors.LimitingExceededMaxPerDay:         http.StatusBadRequest,
	internal_errors.LimitingExceededMaxPerMonth:       http.StatusBadRequest,
	internal_errors.LimitingExceededMaxPerTransaction: http.StatusBadRequest,
	internal_errors.LimitingExceededMaxSingleTransfer: http.StatusBadRequest,
}
