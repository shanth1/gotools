package ops

import (
	"strings"
)

// Kind defines the category of the error.
type Kind uint8

const (
	KindOther          Kind = iota // 0: Unclassified error
	KindInvalid                    // 1: Invalid operation/input (400)
	KindUnauthorized               // 2: Auth missing (401)
	KindPermission                 // 3: Auth present but forbidden (403)
	KindNotFound                   // 4: Resource not found (404)
	KindExist                      // 5: Resource conflict (409)
	KindInternal                   // 6: Internal system error (500)
	KindUnavailable                // 7: Downstream service unavailable (503)
	KindTimeout                    // 8: Operation timeout (504)
	KindNotImplemented             // 9: Not implemented (501)
)

func (k Kind) String() string {
	switch k {
	case KindOther:
		return "other_error"
	case KindInvalid:
		return "invalid_operation"
	case KindUnauthorized:
		return "unauthorized"
	case KindPermission:
		return "permission_denied"
	case KindNotFound:
		return "not_found"
	case KindExist:
		return "item_already_exists"
	case KindInternal:
		return "internal_error"
	case KindUnavailable:
		return "service_unavailable"
	case KindTimeout:
		return "operation_timeout"
	case KindNotImplemented:
		return "not_implemented"
	default:
		return "unknown_kind"
	}
}

// Error implements the standard error interface for Kind.
func (k Kind) Error() string {
	return k.String()
}

// Error is the canonical domain error.
type Error struct {
	Op      string
	Kind    Kind
	Err     error
	Message string // Safe message for the user
}

// Error implements the error interface.
func (e *Error) Error() string {
	var b strings.Builder
	b.WriteString(e.Op)

	if e.Kind != KindOther {
		b.WriteString(": ")
		b.WriteString(e.Kind.String())
	}

	if e.Err != nil {
		b.WriteString(": ")
		b.WriteString(e.Err.Error())
	}

	return b.String()
}

// Unwrap allows standard errors.Is/As to work on the underlying error.
func (e *Error) Unwrap() error {
	return e.Err
}

// Is supports checking against a specific Kind using errors.Is(err, ops.KindNotFound).
func (e *Error) Is(target error) bool {
	t, ok := target.(Kind)
	if ok {
		return e.Kind == t
	}

	return false
}

// --- Constructors ---

func New(op string, kind Kind, msg string) *Error {
	return &Error{Op: op, Kind: kind, Message: msg}
}

func Wrap(op string, kind Kind, err error) *Error {
	return &Error{Op: op, Kind: kind, Err: err}
}

func WrapMsg(op string, kind Kind, err error, msg string) *Error {
	return &Error{Op: op, Kind: kind, Err: err, Message: msg}
}
