package ops

import (
	"strings"
)

type Kind uint8

const (
	KindOther Kind = iota
	KindPermission
	KindNotFound
	KindInvalid
)

func (k Kind) String() string {
	switch k {
	case KindPermission:
		return "permission denied"
	case KindNotFound:
		return "not found"
	case KindInvalid:
		return "invalid"
	default:
		return "unknown error"
	}
}

func (k Kind) Error() string {
	return k.String()
}

type Error struct {
	Op   string
	Kind Kind
	Err  error
}

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

func (e *Error) Unwrap() error {
	return e.Err
}

func (e *Error) Is(target error) bool {
	t, ok := target.(Kind)
	if ok {
		return e.Kind == t
	}
	return false
}

func E(op string, args ...interface{}) error {
	e := &Error{Op: op}
	for _, arg := range args {
		switch arg := arg.(type) {
		case Kind:
			e.Kind = arg
		case error:
			e.Err = arg
		}
	}

	if e.Err == nil && e.Kind == KindOther {
		return nil
	}

	return e
}
