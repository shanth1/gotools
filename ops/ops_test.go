package ops

import (
	"errors"
	"testing"
)

func TestKind_String(t *testing.T) {
	tests := []struct {
		kind     Kind
		expected string
	}{
		{KindOther, "other_error"},
		{KindInvalid, "invalid_operation"},
		{KindUnauthorized, "unauthorized"},
		{KindPermission, "permission_denied"},
		{KindNotFound, "not_found"},
		{KindExist, "item_already_exists"},
		{KindInternal, "internal_error"},
		{KindUnavailable, "service_unavailable"},
		{KindTimeout, "operation_timeout"},
		{KindNotImplemented, "not_implemented"},
		{255, "unknown_kind"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if got := tt.kind.String(); got != tt.expected {
				t.Errorf("Kind.String() = %q, want %q", got, tt.expected)
			}
			if got := tt.kind.Error(); got != tt.expected {
				t.Errorf("Kind.Error() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestError_Error(t *testing.T) {
	baseErr := errors.New("low-level error")

	tests := []struct {
		name     string
		err      *Error
		expected string
	}{
		{
			name: "Op and Kind only",
			err: &Error{
				Op:   "user.Create",
				Kind: KindInvalid,
			},
			expected: "user.Create: invalid_operation",
		},
		{
			name: "Full Error",
			err: &Error{
				Op:   "user.Update",
				Kind: KindPermission,
				Err:  baseErr,
			},
			expected: "user.Update: permission_denied: low-level error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.expected {
				t.Errorf("Error.Error() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestConstructors(t *testing.T) {
	baseErr := errors.New("db error")
	err := Wrap("op", KindInternal, baseErr)
	if err.Kind != KindInternal {
		t.Errorf("Wrap failed kind assignment")
	}
}

func TestError_Is(t *testing.T) {
	base := errors.New("missing")
	err := Wrap("test", KindNotFound, base)

	tests := []struct {
		name   string
		target error
		want   bool
	}{
		{
			name:   "Matches Kind directly",
			target: KindNotFound,
			want:   true,
		},
		{
			name:   "Does not match different Kind",
			target: KindPermission,
			want:   false,
		},
		{
			name:   "Matches underlying error",
			target: base,
			want:   true,
		},
		{
			name:   "Does not match arbitrary error",
			target: errors.New("some other error"),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := errors.Is(err, tt.target); got != tt.want {
				t.Errorf("errors.Is() = %v, want %v", got, tt.want)
			}
		})
	}
}
