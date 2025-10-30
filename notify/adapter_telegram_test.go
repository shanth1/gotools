package notify

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTelegramNotifier(t *testing.T) {
	t.Parallel()
	_, err := NewTelegramNotifier("some-token")
	assert.NoError(t, err)

	_, err = NewTelegramNotifier("")
	assert.Error(t, err)
	assert.EqualError(t, err, "telegram token cannot be empty")
}

func TestTelegramNotifier_Send(t *testing.T) {
	t.Parallel()

	testMsg := Message{Text: "Test *message*"}
	expectedChatID := "12345"

	t.Run("success with markdown", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			body, _ := io.ReadAll(r.Body)
			var payload map[string]string
			json.Unmarshal(body, &payload)

			assert.Equal(t, expectedChatID, payload["chat_id"])
			assert.Equal(t, testMsg.Text, payload["text"])
			assert.Equal(t, "MarkdownV2", payload["parse_mode"])

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"ok":true}`))
		}))
		defer server.Close()

		originalURL := telegramAPIURLTemplate
		defer func() {
			telegramAPIURLTemplate = originalURL
		}()
		telegramAPIURLTemplate = server.URL + "/bot%s/sendMessage"

		notifier := &TelegramNotifier{client: server.Client(), token: "test-token"}

		err := notifier.Send(context.Background(), expectedChatID, testMsg)
		assert.NoError(t, err)
	})

	t.Run("fallback to plain text", func(t *testing.T) {
		var requestCount int32
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			count := atomic.AddInt32(&requestCount, 1)
			body, _ := io.ReadAll(r.Body)
			var payload map[string]string
			json.Unmarshal(body, &payload)

			if count == 1 { // First request with markdown fails
				assert.Equal(t, "MarkdownV2", payload["parse_mode"])
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(`{"description":"Bad Request: can't parse entities"}`))
			} else { // Second request without markdown succeeds
				assert.NotContains(t, payload, "parse_mode")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"ok":true}`))
			}
		}))
		defer server.Close()

		originalURL := telegramAPIURLTemplate
		defer func() {
			telegramAPIURLTemplate = originalURL
		}()
		telegramAPIURLTemplate = server.URL + "/bot%s/sendMessage"

		notifier := &TelegramNotifier{client: server.Client(), token: "test-token"}

		err := notifier.Send(context.Background(), expectedChatID, testMsg)
		require.Error(t, err)
		assert.True(t, errors.Is(err, ErrMarkdownFallback))
		assert.Equal(t, int32(2), atomic.LoadInt32(&requestCount))
	})

	t.Run("api error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"description":"Internal server error"}`))
		}))
		defer server.Close()

		originalURL := telegramAPIURLTemplate
		defer func() {
			telegramAPIURLTemplate = originalURL
		}()
		telegramAPIURLTemplate = server.URL + "/bot%s/sendMessage"

		notifier := &TelegramNotifier{client: server.Client(), token: "test-token"}

		err := notifier.Send(context.Background(), expectedChatID, testMsg)
		require.Error(t, err)
		assert.False(t, errors.Is(err, ErrMarkdownFallback))
		assert.True(t, strings.Contains(err.Error(), "telegram API error: status 500"))
	})

	t.Run("context cancelled", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(50 * time.Millisecond) // Ensure context is cancelled before response
			w.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		originalURL := telegramAPIURLTemplate
		defer func() {
			telegramAPIURLTemplate = originalURL
		}()
		telegramAPIURLTemplate = server.URL + "/bot%s/sendMessage"

		notifier := &TelegramNotifier{client: server.Client(), token: "test-token"}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
		defer cancel()

		err := notifier.Send(ctx, expectedChatID, testMsg)
		require.Error(t, err)
		assert.ErrorIs(t, err, context.DeadlineExceeded)
	})
}
