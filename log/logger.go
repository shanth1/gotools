package log

type level int8

const (
	LevelTrace level = iota - 1
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

type Field struct {
	Key   string
	Value interface{}
}

type Logger interface {
	Trace() Event
	Debug() Event
	Info() Event
	Warn() Event
	Error() Event
	Fatal() Event
	With(fields ...Field) Logger
	WithOptions(opts ...option) Logger
}

// New creates a logger with the given options.
//
// Use With\* functions from this package for configuration.
// - Default log level: info
// - Default writer: formatted console writer
func New(opts ...option) Logger {
	return newZerologLogger(opts...)
}

// NewFromConfig creates a logger from a configuration struct.
// This is convenient for initializing the logger from a YAML file.
func NewFromConfig(cfg Config) Logger {
	return New(WithConfig(cfg))
}

type Event interface {
	Str(key, val string) Event
	Strs(key string, vals []string) Event
	Int(key string, val int) Event
	Bool(key string, val bool) Event
	Err(err error) Event
	Any(key string, val interface{}) Event
	Fields(fields ...Field) Event
	Msg(msg string)
	Msgf(format string, v ...interface{})
}

func Str(key, value string) Field             { return Field{key, value} }
func Int(key string, value int) Field         { return Field{key, value} }
func Err(err error) Field                     { return Field{"error", err} }
func Any(key string, value interface{}) Field { return Field{key, value} }
