package log

import (
	"context"
)

type contextKey struct{}

// NewContext returns a new context derived from ctx that embeds the provided logger.
func NewContext(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, contextKey{}, logger)
}

// FromContext retrieves the logger from the provided context.
// If no logger is found in the context, it returns a new logger with default configuration.
func FromContext(ctx context.Context) Logger {
	if l, ok := ctx.Value(contextKey{}).(Logger); ok {
		return l
	}
	return New()
}

// FromContextOr retrieves the logger from the provided context.
// If no logger is found in the context, it returns the provided fallback logger.
// This is recommended when you want to ensure the logger retains specific configurations
// (e.g., from a service struct) even if the context is missing a logger.
func FromContextOr(ctx context.Context, fallback Logger) Logger {
	if l, ok := ctx.Value(contextKey{}).(Logger); ok {
		return l
	}
	return fallback
}
