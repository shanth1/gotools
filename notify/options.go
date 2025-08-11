package notify

import (
	"fmt"
	"net/http"
	"net/smtp"
	"strings"
	"time"

	"github.com/shanth1/gotools/log"
)

type option func(*service) error

func WithLogger(l log.Logger) option {
	return func(s *service) error {
		s.logger = l
		return nil
	}
}

// WithTelegram
func WithTelegram(token, chatID string) option {
	return func(s *service) error {
		if token == "" || chatID == "" {
			return fmt.Errorf("telegram token and chatID cannot be empty")
		}

		notifier := &telegramNotifier{
			apiURL: fmt.Sprintf(telegramAPIURLTemplate, token),
			chatID: chatID,
			client: &http.Client{Timeout: 10 * time.Second},
		}
		s.notifiers = append(s.notifiers, notifier)
		s.logger.Info().Msgf("Telegram notifier enabled for chat ID %s", chatID)
		return nil
	}
}

// WithEmail
func WithEmail(cfg emailConfig) option {
	return func(s *service) error {
		if cfg.Host == "" || cfg.Port == 0 || cfg.Username == "" || cfg.Password == "" || len(cfg.To) == 0 {
			return fmt.Errorf("email config is incomplete")
		}
		if cfg.From == "" {
			cfg.From = cfg.Username
		}

		auth := smtp.PlainAuth("", cfg.Username, cfg.Password, cfg.Host)
		notifier := &emailNotifier{
			cfg:  cfg,
			auth: auth,
		}
		s.notifiers = append(s.notifiers, notifier)
		s.logger.Info().Msgf("Email notifier enabled for recipients: %s", strings.Join(cfg.To, ", "))
		return nil
	}
}
