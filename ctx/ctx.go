package ctx

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// GetAppCtx creates a context that is terminated by operating system signals
//
// accepts a set of os signals for cancel (if not passed syscall.SIGINT and syscall.SIGTERM will be used)
func GetAppCtx(sigs ...os.Signal) (context.Context, context.CancelFunc) {
	return withSignals(context.Background(), sigs...)
}

// WithGracefulShutdown creates two contexts to manage the application lifecycle
//
// accepts a timeout for shutdown and a set of os signals for cancel (if not passed syscall.SIGINT and syscall.SIGTERM will be used)
func WithGracefulShutdown(shutdownTimeout time.Duration, sigs ...os.Signal) (ctx, shutdownCtx context.Context, cancel, shutdownCancel context.CancelFunc) {
	ctx, cancel = withSignals(context.Background(), sigs...)
	shutdownCtx, shutdownCancel = context.WithCancel(context.Background())

	go func() {
		<-ctx.Done()
		shutdownTimer := time.AfterFunc(shutdownTimeout, func() {
			shutdownCancel()
		})
		defer shutdownTimer.Stop()
	}()

	return ctx, shutdownCtx, cancel, shutdownCancel
}

// withSignals creates a new context with the given signals
//
// if signals are not specified, syscall.SIGINT and syscall.SIGTERM are used
func withSignals(parent context.Context, sigs ...os.Signal) (context.Context, context.CancelFunc) {
	if len(sigs) == 0 {
		sigs = []os.Signal{syscall.SIGINT, syscall.SIGTERM}
	}

	return signal.NotifyContext(parent, sigs...)
}
