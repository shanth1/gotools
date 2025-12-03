package errs

import (
	"errors"
	"fmt"
)

// --- General Standard Errors ---
var (
	ErrNotFound       = errors.New("resource not found")      // 404
	ErrUnauthorized   = errors.New("unauthorized access")     // 401
	ErrForbidden      = errors.New("forbidden action")        // 403
	ErrInternal       = errors.New("internal server error")   // 500
	ErrNotImplemented = errors.New("feature not implemented") // 501
	ErrUnknown        = errors.New("unknown error occurred")  // Fallback
)

// --- Input & Validation ---
var (
	ErrInvalidInput = errors.New("invalid input provided")    // 400 Generic
	ErrMissingField = errors.New("required field is missing") // Mandatory field missing
	ErrInvalidJSON  = errors.New("failed to decode json")     // Malformed JSON
	ErrValidation   = errors.New("validation failed")         // Struct validation error
	ErrBadRequest   = errors.New("bad request")               // Generic client error
)

// --- Database & Storage ---
var (
	ErrAlreadyExists = errors.New("resource already exists") // 409 Conflict (e.g., unique constraint)
	ErrConflict      = errors.New("state conflict")          // 409 Generic conflict
	ErrDBQuery       = errors.New("database query failed")   // SQL/NoSQL execution error
	ErrTransaction   = errors.New("transaction failed")      // DB transaction error
	ErrNoRows        = errors.New("no rows in result set")   // Specific DB empty result
)

// --- Authentication & Identity ---
var (
	ErrTokenExpired = errors.New("auth token expired")   // JWT/Session expired
	ErrTokenInvalid = errors.New("auth token invalid")   // Signature mismatch or malformed
	ErrUserNotFound = errors.New("user not found")       // Specific user lookup fail
	ErrUserBlocked  = errors.New("user account blocked") // Account suspended
	ErrPassMismatch = errors.New("password mismatch")    // Wrong password
)

// --- System & Network ---
var (
	ErrTimeout      = errors.New("operation timed out")   // 408 / 504
	ErrConnection   = errors.New("connection failed")     // Network layer fail
	ErrNotAvailable = errors.New("service not available") // 503
	ErrRateLimit    = errors.New("rate limit exceeded")   // 429
	ErrContext      = errors.New("context canceled")      // Context cancellation
	ErrConfig       = errors.New("configuration error")   // Startup/Config issue
)

// --- File & I/O ---
var (
	ErrFileOpen  = errors.New("failed to open file")
	ErrFileRead  = errors.New("failed to read file")
	ErrFileWrite = errors.New("failed to write file")
	ErrEOF       = errors.New("unexpected end of file")
)

// -------------------------------------------------------------------------
// Helpers
// -------------------------------------------------------------------------

// Wrap adds a context message to an existing error.
// Returns nil if err is nil.
// Example: errs.Wrap(err, "failed to query db")
func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	// %w allows using errors.Is() and errors.As() on the returned error
	return fmt.Errorf("%s: %w", message, err)
}

// Wrapf adds a formatted context message to an existing error.
// Returns nil if err is nil.
// Example: errs.Wrapf(err, "failed to load user %d", userID)
func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	msg := fmt.Sprintf(format, args...)
	return fmt.Errorf("%s: %w", msg, err)
}

// Is reports whether any error in err's chain matches target.
// Alias for standard errors.Is to avoid importing "errors" package alongside "errs".
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// As finds the first error in err's chain that matches target, and if so, sets
// target to that error value and returns true.
// Alias for standard errors.As.
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

// New is a simple alias for errors.New, useful to keep imports clean.
func New(text string) error {
	return errors.New(text)
}
