package log

import (
	"context"

	"github.com/rs/zerolog"
)

// NewContext puts the logger into context
func NewContext(ctx context.Context, logger Logger) context.Context {
	if z, ok := logger.(*zerologAdapter); ok {
		return z.logger.WithContext(ctx)
	}
	return context.WithValue(ctx, contextKey{}, logger)
}

// FromContext extracts the logger from the context
func FromContext(ctx context.Context) Logger {
	l := zerolog.Ctx(ctx)
	return &zerologAdapter{logger: *l}
}

type contextKey struct{}
