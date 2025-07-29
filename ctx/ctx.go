package ctx

import (
	"context"
	"os/signal"
	"syscall"
)

func GetAppCtx() (context.Context, context.CancelFunc) {
	return signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
}
