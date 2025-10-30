package notify

import (
	"context"
	"fmt"
	"net/smtp"
)

type EmailConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

type EmailNotifier struct {
	cfg  EmailConfig
	auth smtp.Auth
	addr string
}

func NewEmailNotifier(cfg EmailConfig) (*EmailNotifier, error) {
	if cfg.Host == "" || cfg.Port == 0 || cfg.Username == "" || cfg.Password == "" {
		return nil, fmt.Errorf("email config is incomplete (host, port, username, password are required)")
	}
	if cfg.From == "" {
		cfg.From = cfg.Username
	}

	auth := smtp.PlainAuth("", cfg.Username, cfg.Password, cfg.Host)
	return &EmailNotifier{
		cfg:  cfg,
		auth: auth,
		addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
	}, nil
}

func (n *EmailNotifier) Send(ctx context.Context, recipientEmail string, msg Message) error {
	subject := msg.Subject
	if subject == "" {
		subject = "Notification"
	}

	emailBody := "From: " + n.cfg.From + "\r\n" +
		"To: " + recipientEmail + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		msg.Text

	errChan := make(chan error, 1)

	go func() {
		errChan <- smtp.SendMail(n.addr, n.auth, n.cfg.From, []string{recipientEmail}, []byte(emailBody))
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errChan:
		if err != nil {
			return fmt.Errorf("failed to send email via SMTP: %w", err)
		}
		return nil
	}
}
