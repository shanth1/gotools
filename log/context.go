package log

import (
	"context"
)

type contextKey struct{}

// NewContext puts the logger into context
func NewContext(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, contextKey{}, logger)
}

// FromContext extracts the logger from the context
func FromContext(ctx context.Context) Logger {
	if l, ok := ctx.Value(contextKey{}).(Logger); ok {
		return l
	}
	return New()
}
