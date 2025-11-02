package notify

import (
	"context"
	"fmt"
	"net/smtp"
)

// EmailConfig holds the connection parameters for an SMTP server.
type EmailConfig struct {
	Host     string `mapstructure:"host" yaml:"host" json:"host" toml:"host" env:"EMAIL_HOST"`
	Port     int    `mapstructure:"port" yaml:"port" json:"port" toml:"port" env:"EMAIL_PORT"`
	Username string `mapstructure:"username" yaml:"username" json:"username" toml:"username" env:"EMAIL_USERNAME"`
	Password string `mapstructure:"password" yaml:"password" json:"password" toml:"password" env:"EMAIL_PASSWORD"`
	From     string `mapstructure:"from" yaml:"from" json:"from" toml:"from" env:"EMAIL_FROM"`
}

// EmailNotifier implements a notifier for sending emails.
type EmailNotifier struct {
	cfg  EmailConfig
	auth smtp.Auth
	addr string
}

// NewEmailNotifier creates and configures a new notifier for sending emails via SMTP.
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

// Send sends an email to the specified recipient.
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
