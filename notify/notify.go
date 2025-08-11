package notify

import (
	"context"
	"fmt"
	"strings"

	"github.com/shanth1/gotools/log"
)

type notifier interface {
	Send(ctx context.Context, message string) error
}

type service struct {
	notifiers []notifier
	logger    log.Logger
}

func New(opts ...option) (*service, error) {
	s := &service{
		logger: log.New(),
	}

	for _, opt := range opts {
		if err := opt(s); err != nil {
			return nil, fmt.Errorf("apply option: %w", err)
		}
	}

	return s, nil
}

func (s *service) Send(ctx context.Context, message string) error {
	var errs []error

	for _, n := range s.notifiers {
		if err := n.Send(ctx, message); err != nil {
			errs = append(errs, err)
			s.logger.Error().Err(err).Msg("send notification")
		}
	}

	// TODO: refactor
	if len(errs) > 0 {
		var strErrs []string
		for _, err := range errs {
			strErrs = append(strErrs, err.Error())
		}
		return fmt.Errorf("send notification(s): %s", strings.Join(strErrs, ", "))
	}

	return nil
}
