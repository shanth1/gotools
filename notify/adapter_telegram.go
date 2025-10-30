package notify

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const telegramAPIURLTemplate = "https://api.telegram.org/bot%s/sendMessage"

var ErrMarkdownFallback = errors.New("message sent successfully, but markdown formatting was stripped")

type TelegramNotifier struct {
	client *http.Client
	token  string
}

func NewTelegramNotifier(token string) (*TelegramNotifier, error) {
	if token == "" {
		return nil, fmt.Errorf("telegram token cannot be empty")
	}
	return &TelegramNotifier{
		client: &http.Client{Timeout: 10 * time.Second},
		token:  token,
	}, nil
}

func (n *TelegramNotifier) Send(ctx context.Context, chatID string, msg Message) error {
	err := n.trySend(ctx, chatID, msg.Text, "MarkdownV2")
	if err == nil {
		return nil
	}

	if strings.Contains(err.Error(), "can't parse entities") {
		fallbackErr := n.trySend(ctx, chatID, msg.Text, "")
		if fallbackErr != nil {
			return fallbackErr
		}
		return ErrMarkdownFallback
	}

	return err
}

func (n *TelegramNotifier) trySend(ctx context.Context, chatID, text, parseMode string) error {
	apiURL := fmt.Sprintf(telegramAPIURLTemplate, n.token)

	payload := map[string]string{
		"chat_id": chatID,
		"text":    text,
	}
	if parseMode != "" {
		payload["parse_mode"] = parseMode
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal telegram payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, apiURL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("failed to create telegram request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := n.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request to telegram: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("telegram API error: status %d, response: %s", resp.StatusCode, string(body))
	}

	return nil
}
