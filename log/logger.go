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
}

// New creates the logger
//
// use With* funcs from package to add custom options
// - default log level: info
// - default writer: formatted console writer
func New(opts ...option) Logger {
	return newZerologLogger(opts...)
}

func NewFromConfig(cfg Config) Logger {
	opts := []option{
		WithService(cfg.Service),
	}

	if cfg.Level != "" {
		opts = append(opts, WithLevel(stringToLevel(cfg.Level)))
	}

	if cfg.EnableCaller {
		opts = append(opts, WithCaller())
	}

	if cfg.UDPAddress != "" {
		opts = append(opts, WithUDPWriter(cfg.UDPAddress))
	}

	if cfg.UDPAddress == "" || cfg.Console {
		opts = append(opts, WithConsoleWriter())
	}

	return newZerologLogger(opts...)
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
