package notify

import "context"

type Message struct {
	Subject string
	Text    string
}

// Notifier defines the interface for a notification service.
type Notifier interface {
	// Send dispatches a message `msg` to a recipient `to`.
	// The `to` string can be a chat ID for Telegram, an email address for SMTP, etc.
	Send(ctx context.Context, to string, msg Message) error
}
