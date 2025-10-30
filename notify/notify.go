package notify

import "context"

type Message struct {
	Subject string
	Text    string
}

type Notifier interface {
	Send(ctx context.Context, to string, msg Message) error
}
