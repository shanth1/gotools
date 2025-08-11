package log

import (
	"io"

	"github.com/rs/zerolog"
)

// zerologAdapter implements the Logger interface
type zerologAdapter struct {
	logger zerolog.Logger
}

// zerologEvent implements the Event interface
type zerologEvent struct {
	event *zerolog.Event
}

func newZerologLogger(opts ...option) Logger {
	cfg := &config{
		level:   LevelInfo,
		writers: []io.Writer{},
	}

	for _, opt := range opts {
		opt(cfg)
	}

	if len(cfg.writers) == 0 {
		WithConsoleWriter()(cfg)
	}

	zlevel, err := zerolog.ParseLevel(levelToString(cfg.level))
	if err != nil {
		zlevel = zerolog.InfoLevel
	}

	multi := zerolog.MultiLevelWriter(cfg.writers...)

	context := zerolog.New(multi).With().Timestamp()

	if cfg.service != "" {
		context = context.Str("service", cfg.service)
	}
	if cfg.enableCaller {
		context = context.Caller()
	}

	logger := context.Logger().Level(zlevel)

	return &zerologAdapter{logger: logger}
}

func (l *zerologAdapter) With(fields ...Field) Logger {
	context := l.logger.With()
	for _, f := range fields {
		context = context.Interface(f.Key, f.Value)
	}
	return &zerologAdapter{logger: context.Logger()}
}

// --- Logger ---

func (l *zerologAdapter) Trace() Event { return &zerologEvent{l.logger.Trace()} }
func (l *zerologAdapter) Debug() Event { return &zerologEvent{l.logger.Debug()} }
func (l *zerologAdapter) Info() Event  { return &zerologEvent{l.logger.Info()} }
func (l *zerologAdapter) Warn() Event  { return &zerologEvent{l.logger.Warn()} }
func (l *zerologAdapter) Error() Event { return &zerologEvent{l.logger.Error()} }
func (l *zerologAdapter) Fatal() Event { return &zerologEvent{l.logger.Fatal()} }

// --- Event ---

func (e *zerologEvent) Str(key, val string) Event             { e.event.Str(key, val); return e }
func (e *zerologEvent) Strs(key string, vals []string) Event  { e.event.Strs(key, vals); return e }
func (e *zerologEvent) Int(key string, val int) Event         { e.event.Int(key, val); return e }
func (e *zerologEvent) Bool(key string, val bool) Event       { e.event.Bool(key, val); return e }
func (e *zerologEvent) Any(key string, val interface{}) Event { e.event.Any(key, val); return e }
func (e *zerologEvent) Err(err error) Event                   { e.event.Err(err); return e }

func (e *zerologEvent) Fields(fields ...Field) Event {
	for _, f := range fields {
		e.event.Interface(f.Key, f.Value)
	}
	return e
}

func (e *zerologEvent) Msg(msg string) {
	e.event.Msg(msg)
}

func (e *zerologEvent) Msgf(format string, v ...interface{}) {
	e.event.Msgf(format, v...)
}
