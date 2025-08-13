package notify

import (
	"context"
	"fmt"
	"net/smtp"
	"strings"
)

type emailConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
	To       []string
}

type emailNotifier struct {
	cfg  emailConfig
	auth smtp.Auth
}

func (e *emailNotifier) Send(ctx context.Context, message string) error {
	headers := map[string]string{
		"From":    e.cfg.From,
		"To":      strings.Join(e.cfg.To, ","),
		"Subject": "Notification", // TODO: config
	}

	var msgBuilder strings.Builder
	for k, v := range headers {
		msgBuilder.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}
	msgBuilder.WriteString("\r\n")
	msgBuilder.WriteString(message)

	addr := fmt.Sprintf("%s:%d", e.cfg.Host, e.cfg.Port)

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		err := smtp.SendMail(addr, e.auth, e.cfg.From, e.cfg.To, []byte(msgBuilder.String()))
		if err != nil {
			return fmt.Errorf("failed to send email via SMTP: %w", err)
		}
		return nil
	}
}
