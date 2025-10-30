package notify

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewEmailNotifier(t *testing.T) {
	t.Parallel()

	validConfig := EmailConfig{
		Host:     "smtp.example.com",
		Port:     587,
		Username: "user",
		Password: "password",
	}

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		notifier, err := NewEmailNotifier(validConfig)
		require.NoError(t, err)
		assert.NotNil(t, notifier)
		assert.Equal(t, "user", notifier.cfg.From, "From should default to Username")
	})

	t.Run("success with from field", func(t *testing.T) {
		t.Parallel()
		cfg := validConfig
		cfg.From = "sender@example.com"
		notifier, err := NewEmailNotifier(cfg)
		require.NoError(t, err)
		assert.Equal(t, "sender@example.com", notifier.cfg.From)
	})

	t.Run("incomplete config", func(t *testing.T) {
		t.Parallel()
		testCases := []struct {
			name string
			cfg  EmailConfig
		}{
			{"missing host", EmailConfig{Port: 587, Username: "u", Password: "p"}},
			{"missing port", EmailConfig{Host: "h", Username: "u", Password: "p"}},
			{"missing username", EmailConfig{Host: "h", Port: 587, Password: "p"}},
			{"missing password", EmailConfig{Host: "h", Port: 587, Username: "u"}},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				_, err := NewEmailNotifier(tc.cfg)
				require.Error(t, err)
				assert.EqualError(t, err, "email config is incomplete (host, port, username, password are required)")
			})
		}
	})
}

func TestEmailNotifier_Send_ContextCancelled(t *testing.T) {
	// This test relies on the network call to a non-existent host hanging
	// long enough for the context to be cancelled.
	cfg := EmailConfig{
		Host:     "10.255.255.1", // Non-routable address to ensure a timeout
		Port:     12345,
		Username: "user",
		Password: "password",
	}

	notifier, err := NewEmailNotifier(cfg)
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()

	msg := Message{Subject: "Test", Text: "Hello"}
	sendErr := notifier.Send(ctx, "recipient@example.com", msg)

	require.Error(t, sendErr)
	assert.ErrorIs(t, sendErr, context.DeadlineExceeded)
}
