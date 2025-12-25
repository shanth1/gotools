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
		{KindOther, "other error"},
		{KindInvalid, "invalid operation"},
		{KindUnauthorized, "unauthorized"},
		{KindPermission, "permission denied"},
		{KindNotFound, "not found"},
		{KindExist, "item already exists"},
		{KindInternal, "internal error"},
		{KindUnavailable, "service unavailable"},
		{KindTimeout, "operation timeout"},
		{KindNotImplemented, "not implemented"},
		{Kind(255), "unknown error kind"},
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
			expected: "user.Create: invalid operation",
		},
		{
			name: "Op and Nested Error only",
			err: &Error{
				Op:   "user.Delete",
				Kind: KindOther,
				Err:  baseErr,
			},
			expected: "user.Delete: low-level error",
		},
		{
			name: "Full Error (Op, Kind, Err)",
			err: &Error{
				Op:   "user.Update",
				Kind: KindPermission,
				Err:  baseErr,
			},
			expected: "user.Update: permission denied: low-level error",
		},
		{
			name: "Op Only (KindOther, nil Err)",
			err: &Error{
				Op:   "process.Start",
				Kind: KindOther,
			},
			expected: "process.Start",
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

func TestE(t *testing.T) {
	baseErr := errors.New("db connection failed")

	tests := []struct {
		name     string
		op       string
		args     []interface{}
		validate func(*testing.T, error)
	}{
		{
			name: "Construct with Kind",
			op:   "auth.Login",
			args: []interface{}{KindUnauthorized},
			validate: func(t *testing.T, err error) {
				e, ok := err.(*Error)
				if !ok {
					t.Fatalf("expected *Error, got %T", err)
				}
				if e.Op != "auth.Login" {
					t.Errorf("expected Op 'auth.Login', got %q", e.Op)
				}
				if e.Kind != KindUnauthorized {
					t.Errorf("expected KindUnauthorized, got %v", e.Kind)
				}
			},
		},
		{
			name: "Construct with Error",
			op:   "file.Read",
			args: []interface{}{baseErr},
			validate: func(t *testing.T, err error) {
				e, ok := err.(*Error)
				if !ok {
					t.Fatalf("expected *Error, got %T", err)
				}
				if e.Err != baseErr {
					t.Errorf("expected wrapped error to be baseErr")
				}
				if e.Kind != KindOther {
					t.Errorf("expected KindOther, got %v", e.Kind)
				}
			},
		},
		{
			name: "Construct with Kind and Error",
			op:   "service.Call",
			args: []interface{}{KindInternal, baseErr},
			validate: func(t *testing.T, err error) {
				e, ok := err.(*Error)
				if !ok {
					t.Fatalf("expected *Error, got %T", err)
				}
				if e.Kind != KindInternal {
					t.Errorf("expected KindInternal")
				}
				if e.Err != baseErr {
					t.Errorf("expected wrapped error")
				}
			},
		},
		{
			name: "Nil return if empty",
			op:   "noop",
			args: []interface{}{},
			validate: func(t *testing.T, err error) {
				if err != nil {
					t.Errorf("expected nil error, got %v", err)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := E(tt.op, tt.args...)
			tt.validate(t, err)
		})
	}
}

func TestError_Unwrap(t *testing.T) {
	inner := errors.New("inner")
	err := &Error{Err: inner}

	if unwrapped := errors.Unwrap(err); unwrapped != inner {
		t.Errorf("Unwrap() = %v, want %v", unwrapped, inner)
	}
}

func TestError_Is(t *testing.T) {
	err := E("test", KindNotFound, errors.New("missing"))

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
