package ctx

import "context"

type key int

const (
	requestIDKey key = iota
	userIDIntKey
	userIDStrKey
)

// --------------------------------------------------------------------------------

// WithRequestID adds a request ID to the context.
func WithRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestIDKey, id)
}

// RequestIDFrom extracts a request ID from the context.
func RequestIDFrom(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(requestIDKey).(string)
	return id, ok
}

// --------------------------------------------------------------------------------

// WithUserIDInt adds an integer user ID to the context.
func WithUserIDInt(ctx context.Context, id int64) context.Context {
	return context.WithValue(ctx, userIDIntKey, id)
}

// UserIDIntFrom extracts an integer user ID from the context.
func UserIDIntFrom(ctx context.Context) (int64, bool) {
	id, ok := ctx.Value(userIDIntKey).(int64)
	return id, ok
}

// --------------------------------------------------------------------------------

// WithUserIDStr adds a string user ID to the context.
func WithUserIDStr(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, userIDStrKey, id)
}

// UserIDStrFrom extracts a string user ID from the context.
func UserIDStrFrom(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(userIDStrKey).(string)
	return id, ok
}
