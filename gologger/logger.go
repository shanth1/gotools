package logger

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
func GetLogger(app string, level zerolog.Level, writers ...io.Writer) zerolog.Logger {
	zerolog.SetGlobalLevel(level)

	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}
	writers = append(writers, output)

	multi := zerolog.MultiLevelWriter(writers...)
	logger := zerolog.New(multi).With().Str("app", app).Timestamp().Logger()

	return logger.Hook(callerHook{})
}

func GetLoggerFromCtx(ctx context.Context) *zerolog.Logger {
	return zerolog.Ctx(ctx)
}
