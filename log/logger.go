package log

import (
	"time"
)

type level int8

const (
	LevelTrace level = iota - 1
	// LevelDebug logs useful for debugging.
	LevelDebug
	// LevelInfo logs standard informational messages.
	LevelInfo
	// LevelWarn logs warnings that doesn't stop the flow.
	LevelWarn
	// LevelError logs errors that should be investigated.
	LevelError
	// LevelFatal logs fatal errors and CALLS os.Exit(1).
	// CAUTION: This will terminate the application immediately.
	LevelFatal
	// LevelPanic logs the error and CALLS panic().
	LevelPanic
	// LevelDisabled disables logging completely.
	LevelDisabled
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
	Fatal() Event // Fatal logs a message at fatal level and then calls os.Exit(1).
	Panic() Event // Panic logs a message at panic level and then calls panic().
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
	// Standard types
	Str(key, val string) Event
	Bool(key string, val bool) Event
	Int(key string, val int) Event
	Int8(key string, val int8) Event
	Int16(key string, val int16) Event
	Int32(key string, val int32) Event
	Int64(key string, val int64) Event
	Uint(key string, val uint) Event
	Uint8(key string, val uint8) Event
	Uint16(key string, val uint16) Event
	Uint32(key string, val uint32) Event
	Uint64(key string, val uint64) Event
	Float32(key string, val float32) Event
	Float64(key string, val float64) Event

	// Time and Duration
	Time(key string, val time.Time) Event
	Dur(key string, val time.Duration) Event

	// Binary and Complex
	Bytes(key string, val []byte) Event
	Hex(key string, val []byte) Event
	RawJSON(key string, b []byte) Event
	Err(err error) Event
	Any(key string, val interface{}) Event

	// Slices
	Strs(key string, vals []string) Event
	Bools(key string, vals []bool) Event
	Ints(key string, vals []int) Event
	Ints64(key string, vals []int64) Event
	Uints(key string, vals []uint) Event
	Uints64(key string, vals []uint64) Event
	Floats32(key string, vals []float32) Event
	Floats64(key string, vals []float64) Event
	Times(key string, vals []time.Time) Event
	Durs(key string, vals []time.Duration) Event

	// Common
	Fields(fields ...Field) Event
	Msg(msg string)
	Msgf(format string, v ...interface{})
}

// --- Field Helpers for With() ---

func Str(key, value string) Field             { return Field{key, value} }
func Bool(key string, value bool) Field       { return Field{key, value} }
func Int(key string, value int) Field         { return Field{key, value} }
func Int8(key string, value int8) Field       { return Field{key, value} }
func Int16(key string, value int16) Field     { return Field{key, value} }
func Int32(key string, value int32) Field     { return Field{key, value} }
func Int64(key string, value int64) Field     { return Field{key, value} }
func Uint(key string, value uint) Field       { return Field{key, value} }
func Uint8(key string, value uint8) Field     { return Field{key, value} }
func Uint16(key string, value uint16) Field   { return Field{key, value} }
func Uint32(key string, value uint32) Field   { return Field{key, value} }
func Uint64(key string, value uint64) Field   { return Field{key, value} }
func Float32(key string, value float32) Field { return Field{key, value} }
func Float64(key string, value float64) Field { return Field{key, value} }

func Time(key string, value time.Time) Field    { return Field{key, value} }
func Dur(key string, value time.Duration) Field { return Field{key, value} }
func Err(err error) Field                       { return Field{"error", err} }
func Any(key string, value interface{}) Field   { return Field{key, value} }
func Bytes(key string, value []byte) Field      { return Field{key, value} }
func Hex(key string, value []byte) Field        { return Field{key, value} } // Note: With() uses reflection, might log as b64 unless adapter is smart
func RawJSON(key string, value []byte) Field    { return Field{key, value} }
