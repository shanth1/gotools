package ctx

import "context"

type key int

const (
	requestIDKey key = iota
	userIDIntKey
	userIDStrKey
)

// --------------------------------------------------------------------------------

// WithRequestID adds the id to ctx
func WithRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestIDKey, id)
}

// RequestIDFrom returns request id from context
func RequestIDFrom(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(requestIDKey).(string)
	return id, ok
}

// --------------------------------------------------------------------------------

// WithUserIDInt adds the user id to context
func WithUserIDInt(ctx context.Context, id int64) context.Context {
	return context.WithValue(ctx, userIDIntKey, id)
}

// UserIDIntFrom returns user id from context
func UserIDIntFrom(ctx context.Context) (int64, bool) {
	id, ok := ctx.Value(userIDIntKey).(int64)
	return id, ok
}

// --------------------------------------------------------------------------------

// WithUserIDStr adds the user id to context
func WithUserIDStr(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, userIDStrKey, id)
}

// UserIDStrFrom returns user id from context
func UserIDStrFrom(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(userIDStrKey).(string)
	return id, ok
}
