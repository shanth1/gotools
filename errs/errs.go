package errs

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound       = errors.New("resource not found")
	ErrUnauthorized   = errors.New("unauthorized access")
	ErrForbidden      = errors.New("forbidden action")
	ErrInvalidInput   = errors.New("invalid input provided")
	ErrInternal       = errors.New("internal server error")
	ErrTimeout        = errors.New("operation timed out")
	ErrNotImplemented = errors.New("feature not implemented")
	ErrConnection     = errors.New("connection failed")
	ErrNotAvailable   = errors.New("service not available")
)

// Wrap adds context to an existing error
func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", message, err)
}
