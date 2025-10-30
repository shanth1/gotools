package ctx

import (
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetAppCtx(t *testing.T) {
	t.Run("cancels on default signal", func(t *testing.T) {
		ctx, cancel := GetAppCtx()
		defer cancel()

		go func() {
			// Allow time for context setup
			time.Sleep(10 * time.Millisecond)
			syscall.Kill(syscall.Getpid(), syscall.SIGINT)
		}()

		select {
		case <-ctx.Done():
			// Success
		case <-time.After(1 * time.Second):
			t.Fatal("context was not canceled by the signal")
		}
	})

	t.Run("cancels on custom signal", func(t *testing.T) {
		ctx, cancel := GetAppCtx(syscall.SIGHUP)
		defer cancel()

		go func() {
			time.Sleep(10 * time.Millisecond)
			syscall.Kill(syscall.Getpid(), syscall.SIGHUP)
		}()

		select {
		case <-ctx.Done():
			// Success
		case <-time.After(1 * time.Second):
			t.Fatal("context was not canceled by the custom signal")
		}
	})
}

func TestWithGracefulShutdown(t *testing.T) {
	const shutdownTimeout = 50 * time.Millisecond

	appCtx, shutdownCtx, cancel, shutdownCancel := WithGracefulShutdown(shutdownTimeout, syscall.SIGHUP)
	defer cancel()
	defer shutdownCancel()

	// Ensure contexts are not done initially
	require.NoError(t, appCtx.Err())
	require.NoError(t, shutdownCtx.Err())

	// Send signal to trigger shutdown
	go func() {
		time.Sleep(10 * time.Millisecond)
		syscall.Kill(syscall.Getpid(), syscall.SIGHUP)
	}()

	// Wait for app context to be canceled
	select {
	case <-appCtx.Done():
		// App context is done, as expected
	case <-time.After(1 * time.Second):
		t.Fatal("app context was not canceled by the signal")
	}

	// At this point, shutdown context should still be active
	require.NoError(t, shutdownCtx.Err(), "shutdown context should not be canceled immediately")

	// Wait for the shutdown timeout to pass
	select {
	case <-shutdownCtx.Done():
		// Shutdown context is now done, as expected
	case <-time.After(shutdownTimeout + 20*time.Millisecond):
		t.Fatal("shutdown context was not canceled after the timeout")
	}
}
