package ctx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContextValues(t *testing.T) {
	t.Parallel()

	t.Run("RequestID", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()
		id, ok := RequestIDFrom(ctx)
		assert.False(t, ok)
		assert.Empty(t, id)

		ctxWithID := WithRequestID(ctx, "req-123")
		id, ok = RequestIDFrom(ctxWithID)
		assert.True(t, ok)
		assert.Equal(t, "req-123", id)
	})

	t.Run("UserIDInt", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()
		id, ok := UserIDIntFrom(ctx)
		assert.False(t, ok)
		assert.Zero(t, id)

		ctxWithID := WithUserIDInt(ctx, 42)
		id, ok = UserIDIntFrom(ctxWithID)
		assert.True(t, ok)
		assert.Equal(t, int64(42), id)
	})

	t.Run("UserIDStr", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()
		id, ok := UserIDStrFrom(ctx)
		assert.False(t, ok)
		assert.Empty(t, id)

		ctxWithID := WithUserIDStr(ctx, "user-abc")
		id, ok = UserIDStrFrom(ctxWithID)
		assert.True(t, ok)
		assert.Equal(t, "user-abc", id)
	})
}
