package logger

import (
	"context"
	"io"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

type Logger interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(err error, msg string, fields ...Field)
	Fatal(err error, msg string, fields ...Field)

	With(fields ...Field) Logger
}

type Field struct {
	Key   string
	Value interface{}
}

func Str(key, value string) Field {
	return Field{Key: key, Value: value}
}

func Int(key string, value int) Field {
	return Field{Key: key, Value: value}
}

func Any(key string, value interface{}) Field {
	return Field{Key: key, Value: value}
}

type zerologAdapter struct {
	zlog zerolog.Logger
}

type callerHook struct{}

func (h callerHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	if level >= zerolog.ErrorLevel {
		e.Caller(3)
	}
}

func New(app string, level string, writers ...io.Writer) Logger {
	zlevel, err := zerolog.ParseLevel(strings.ToLower(level))
	if err != nil {
		zlevel = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(zlevel)

	if len(writers) == 0 {
		writers = append(writers, zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		})
	}

	multi := zerolog.MultiLevelWriter(writers...)
	zlog := zerolog.New(multi).With().
		Str("app", app).
		Timestamp().
		Logger().
		Hook(callerHook{})

	return &zerologAdapter{zlog: zlog}
}

func (a *zerologAdapter) Debug(msg string, fields ...Field) {
	logEvent(a.zlog.Debug(), msg, fields)
}

func (a *zerologAdapter) Info(msg string, fields ...Field) {
	logEvent(a.zlog.Info(), msg, fields)
}

func (a *zerologAdapter) Warn(msg string, fields ...Field) {
	logEvent(a.zlog.Warn(), msg, fields)
}

func (a *zerologAdapter) Error(err error, msg string, fields ...Field) {
	event := a.zlog.Error()
	if err != nil {
		event = event.Err(err)
	}
	logEvent(event, msg, fields)
}

func (a *zerologAdapter) Fatal(err error, msg string, fields ...Field) {
	event := a.zlog.Fatal()
	if err != nil {
		event = event.Err(err)
	}
	logEvent(event, msg, fields)
}

func (a *zerologAdapter) With(fields ...Field) Logger {
	context := a.zlog.With()
	for _, f := range fields {
		context = context.Interface(f.Key, f.Value)
	}
	newZlog := context.Logger()
	return &zerologAdapter{zlog: newZlog}
}

func logEvent(e *zerolog.Event, msg string, fields []Field) {
	if e.Enabled() {
		for _, f := range fields {
			e = e.Interface(f.Key, f.Value)
		}
		e.Msg(msg)
	}
}

type contextKey struct{}

func NewContext(ctx context.Context, log Logger) context.Context {
	if zla, ok := log.(*zerologAdapter); ok {
		return zla.zlog.WithContext(ctx)
	}
	return context.WithValue(ctx, contextKey{}, log)
}

func FromContext(ctx context.Context) Logger {
	zlog := zerolog.Ctx(ctx)

	if zlog.GetLevel() == zerolog.Disabled {
		if l, ok := ctx.Value(contextKey{}).(Logger); ok {
			return l
		}
	}

	return &zerologAdapter{zlog: *zlog}
}
