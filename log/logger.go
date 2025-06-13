package log

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

type callerHook struct{}

func (h callerHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	if level == zerolog.ErrorLevel {
		e.Caller()
	}
}

// GetLogger returns a custom zerolog logger
func New(app, level string, writers ...io.Writer) zerolog.Logger {
	zlevel, err := zerolog.ParseLevel(level)
	if err != nil {
		zlevel = zerolog.TraceLevel
	}
	zerolog.SetGlobalLevel(zlevel)

	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}
	writers = append(writers, output)

	multi := zerolog.MultiLevelWriter(writers...)
	logger := zerolog.New(multi).With().Str("app", app).Timestamp().Logger()

	return logger.Hook(callerHook{})
}

func FromCtx(ctx context.Context) *zerolog.Logger {
	return zerolog.Ctx(ctx)
}
