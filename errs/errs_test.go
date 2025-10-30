package errs

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrap(t *testing.T) {
	t.Parallel()

	t.Run("wrap non-nil error", func(t *testing.T) {
		t.Parallel()
		baseErr := errors.New("base error")
		wrappedErr := Wrap(baseErr, "additional context")

		assert.Error(t, wrappedErr)
		assert.EqualError(t, wrappedErr, "additional context: base error")
		assert.True(t, errors.Is(wrappedErr, baseErr), "wrapped error should contain the base error")
	})

	t.Run("wrap a pre-defined error", func(t *testing.T) {
		t.Parallel()
		wrappedErr := Wrap(ErrNotFound, "user not found")

		assert.Error(t, wrappedErr)
		assert.EqualError(t, wrappedErr, "user not found: resource not found")
		assert.True(t, errors.Is(wrappedErr, ErrNotFound))
	})

	t.Run("wrap nil error", func(t *testing.T) {
		t.Parallel()
		err := Wrap(nil, "this should not appear")
		assert.NoError(t, err)
	})
}
