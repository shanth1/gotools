package notify

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const telegramAPIURLTemplate = "https://api.telegram.org/bot%s/sendMessage"

type telegramNotifier struct {
	apiURL string
	chatID string
	client *http.Client
}

func (t *telegramNotifier) Send(ctx context.Context, message string) error {
	payload := map[string]string{
		"chat_id":    t.chatID,
		"text":       message,
		"parse_mode": "Markdown",
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal telegram payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, t.apiURL, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create telegram request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := t.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request to telegram: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var result struct {
			Description string `json:"description"`
		}
		if json.NewDecoder(resp.Body).Decode(&result) == nil {
			return fmt.Errorf("telegram API error: %s (status code: %d)", result.Description, resp.StatusCode)
		}
		return fmt.Errorf("telegram API returned non-200 status code: %d", resp.StatusCode)
	}

	return nil
}
