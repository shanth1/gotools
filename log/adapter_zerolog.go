package log

import (
	"io"
	"time"

	"github.com/rs/zerolog"
)

// zerologAdapter implements the Logger interface
type zerologAdapter struct {
	logger zerolog.Logger
	cfg    *config
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

	return newLoggerWithConfig(cfg)
}

func (l *zerologAdapter) With(fields ...Field) Logger {
	context := l.logger.With()
	for _, f := range fields {
		context = context.Interface(f.Key, f.Value)
	}
	return &zerologAdapter{
		logger: context.Logger(),
		cfg:    l.cfg,
	}
}

func (l *zerologAdapter) WithOptions(opts ...option) Logger {
	newCfg := l.cfg.clone()

	for _, opt := range opts {
		opt(newCfg)
	}

	return newLoggerWithConfig(newCfg)
}

// --- Logger ---

func (l *zerologAdapter) Trace() Event { return &zerologEvent{l.logger.Trace()} }
func (l *zerologAdapter) Debug() Event { return &zerologEvent{l.logger.Debug()} }
func (l *zerologAdapter) Info() Event  { return &zerologEvent{l.logger.Info()} }
func (l *zerologAdapter) Warn() Event  { return &zerologEvent{l.logger.Warn()} }
func (l *zerologAdapter) Error() Event { return &zerologEvent{l.logger.Error()} }
func (l *zerologAdapter) Fatal() Event { return &zerologEvent{l.logger.Fatal()} }
func (l *zerologAdapter) Panic() Event { return &zerologEvent{l.logger.Panic()} }

// --- Event Implementation ---

// Standard types
func (e *zerologEvent) Str(key, val string) Event             { e.event.Str(key, val); return e }
func (e *zerologEvent) Bool(key string, val bool) Event       { e.event.Bool(key, val); return e }
func (e *zerologEvent) Int(key string, val int) Event         { e.event.Int(key, val); return e }
func (e *zerologEvent) Int8(key string, val int8) Event       { e.event.Int8(key, val); return e }
func (e *zerologEvent) Int16(key string, val int16) Event     { e.event.Int16(key, val); return e }
func (e *zerologEvent) Int32(key string, val int32) Event     { e.event.Int32(key, val); return e }
func (e *zerologEvent) Int64(key string, val int64) Event     { e.event.Int64(key, val); return e }
func (e *zerologEvent) Uint(key string, val uint) Event       { e.event.Uint(key, val); return e }
func (e *zerologEvent) Uint8(key string, val uint8) Event     { e.event.Uint8(key, val); return e }
func (e *zerologEvent) Uint16(key string, val uint16) Event   { e.event.Uint16(key, val); return e }
func (e *zerologEvent) Uint32(key string, val uint32) Event   { e.event.Uint32(key, val); return e }
func (e *zerologEvent) Uint64(key string, val uint64) Event   { e.event.Uint64(key, val); return e }
func (e *zerologEvent) Float32(key string, val float32) Event { e.event.Float32(key, val); return e }
func (e *zerologEvent) Float64(key string, val float64) Event { e.event.Float64(key, val); return e }

// Time and Duration
func (e *zerologEvent) Time(key string, val time.Time) Event    { e.event.Time(key, val); return e }
func (e *zerologEvent) Dur(key string, val time.Duration) Event { e.event.Dur(key, val); return e }

// Binary and Complex
func (e *zerologEvent) Bytes(key string, val []byte) Event    { e.event.Bytes(key, val); return e }
func (e *zerologEvent) Hex(key string, val []byte) Event      { e.event.Hex(key, val); return e }
func (e *zerologEvent) RawJSON(key string, b []byte) Event    { e.event.RawJSON(key, b); return e }
func (e *zerologEvent) Err(err error) Event                   { e.event.Err(err); return e }
func (e *zerologEvent) Any(key string, val interface{}) Event { e.event.Any(key, val); return e }

// Slices
func (e *zerologEvent) Strs(key string, vals []string) Event    { e.event.Strs(key, vals); return e }
func (e *zerologEvent) Bools(key string, vals []bool) Event     { e.event.Bools(key, vals); return e }
func (e *zerologEvent) Ints(key string, vals []int) Event       { e.event.Ints(key, vals); return e }
func (e *zerologEvent) Ints64(key string, vals []int64) Event   { e.event.Ints64(key, vals); return e }
func (e *zerologEvent) Uints(key string, vals []uint) Event     { e.event.Uints(key, vals); return e }
func (e *zerologEvent) Uints64(key string, vals []uint64) Event { e.event.Uints64(key, vals); return e }
func (e *zerologEvent) Floats32(key string, vals []float32) Event {
	e.event.Floats32(key, vals)
	return e
}
func (e *zerologEvent) Floats64(key string, vals []float64) Event {
	e.event.Floats64(key, vals)
	return e
}
func (e *zerologEvent) Times(key string, vals []time.Time) Event { e.event.Times(key, vals); return e }
func (e *zerologEvent) Durs(key string, vals []time.Duration) Event {
	e.event.Durs(key, vals)
	return e
}

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

// --- Constructor ---
func newLoggerWithConfig(cfg *config) Logger {
	strLevel := levelToString(cfg.level)

	var zlevel zerolog.Level
	if cfg.level == LevelDisabled {
		zlevel = zerolog.Disabled
	} else {
		var err error
		zlevel, err = zerolog.ParseLevel(strLevel)
		if err != nil {
			zlevel = zerolog.InfoLevel
		}
	}

	var finalWriter io.Writer
	if len(cfg.writers) == 1 {
		finalWriter = cfg.writers[0]
	} else {
		finalWriter = zerolog.MultiLevelWriter(cfg.writers...)
	}

	zerologContext := zerolog.New(finalWriter).With()

	if cfg.app != "" {
		zerologContext = zerologContext.Str("app", cfg.app)
	}
	if cfg.service != "" {
		zerologContext = zerologContext.Str("service", cfg.service)
	}
	if cfg.enableCaller {
		zerologContext = zerologContext.Caller()
	}

	loggerWithContext := zerologContext.Logger().Level(zlevel)
	nanoSecondHook := zerolog.HookFunc(func(e *zerolog.Event, level zerolog.Level, message string) {
		e.Str(zerolog.TimestampFieldName, time.Now().Format(time.RFC3339Nano))
	})
	finalLogger := loggerWithContext.Hook(nanoSecondHook).Level(zlevel)

	return &zerologAdapter{
		logger: finalLogger,
		cfg:    cfg,
	}
}
