package log

import (
	"context"
	"log/slog"

	"github.com/rs/zerolog"
)

// ToSlog converts the current Logger into a standard library *slog.Logger.
// This allows using this library in code that expects slog.
func ToSlog(l Logger) *slog.Logger {
	// Try to get the underlying zerolog logger for better performance
	if adapter, ok := l.(*zerologAdapter); ok {
		return slog.New(&zeroSlogHandler{
			logger: adapter.logger,
		})
	}

	// Fallback if implementation changes (unlikely)
	// In a real scenario, we might want to panic or handle this gracefully,
	// but currently only zerologAdapter exists.
	return slog.Default()
}

// zeroSlogHandler implements slog.Handler
type zeroSlogHandler struct {
	logger zerolog.Logger
	group  string
}

func (h *zeroSlogHandler) Enabled(_ context.Context, level slog.Level) bool {
	zLevel := slogLevelToZerolog(level)
	return h.logger.GetLevel() <= zLevel
}

func (h *zeroSlogHandler) Handle(ctx context.Context, record slog.Record) error {
	zLevel := slogLevelToZerolog(record.Level)

	// Create a zerolog event
	event := h.logger.WithLevel(zLevel)

	// If logging is disabled for this level, WithLevel returns nil (in zerolog internals)
	// but the public API usually returns a safe object.
	// However, zerolog.Logger.WithLevel() returns *Event.
	if !event.Enabled() {
		return nil
	}

	// Add fields
	if h.group != "" {
		event = event.Dict(h.group, zerolog.Dict().Func(func(e *zerolog.Event) {
			record.Attrs(func(attr slog.Attr) bool {
				addSlogAttr(e, attr)
				return true
			})
		}))
	} else {
		record.Attrs(func(attr slog.Attr) bool {
			addSlogAttr(event, attr)
			return true
		})
	}

	// Add caller if enabled in config (handled by zerolog internally if configured),
	// but slog record has its own PC. We might want to use slog's PC if zerolog's caller is disabled,
	// but usually, we trust the underlying logger config.

	event.Msg(record.Message)
	return nil
}

func (h *zeroSlogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	// Zerolog Context pattern
	zCtx := h.logger.With()
	for _, attr := range attrs {
		// Since zerolog.Context doesn't expose full typing like Event,
		// we mainly rely on Interface for complex types or known basic types.
		// For simplicity/performance in WithAttrs, Interface is safest.
		zCtx = zCtx.Interface(attr.Key, attr.Value.Any())
	}

	return &zeroSlogHandler{
		logger: zCtx.Logger(),
		group:  h.group,
	}
}

func (h *zeroSlogHandler) WithGroup(name string) slog.Handler {
	// Zerolog doesn't support stateful groups on the Logger object easily without wrappers.
	// We handle this by storing the group name or nesting loggers.
	// For deep nesting, this simple string approach is limited,
	// but sufficient for basic usage.
	// A proper implementation would need a stack of groups.

	// Simplified implementation:
	return &zeroSlogHandler{
		logger: h.logger,
		group:  name, // Limitation: only supports one level of grouping in this simple adapter
	}
}

// slogLevelToZerolog converts slog levels to zerolog levels.
func slogLevelToZerolog(l slog.Level) zerolog.Level {
	switch {
	case l >= slog.LevelError:
		return zerolog.ErrorLevel
	case l >= slog.LevelWarn:
		return zerolog.WarnLevel
	case l >= slog.LevelInfo:
		return zerolog.InfoLevel
	case l >= slog.LevelDebug:
		return zerolog.DebugLevel
	default:
		return zerolog.TraceLevel
	}
}

func addSlogAttr(e *zerolog.Event, attr slog.Attr) {
	// Resolve lazy values
	attr.Value = attr.Value.Resolve()

	switch attr.Value.Kind() {
	case slog.KindString:
		e.Str(attr.Key, attr.Value.String())
	case slog.KindInt64:
		e.Int64(attr.Key, attr.Value.Int64())
	case slog.KindUint64:
		e.Uint64(attr.Key, attr.Value.Uint64())
	case slog.KindFloat64:
		e.Float64(attr.Key, attr.Value.Float64())
	case slog.KindBool:
		e.Bool(attr.Key, attr.Value.Bool())
	case slog.KindDuration:
		e.Dur(attr.Key, attr.Value.Duration())
	case slog.KindTime:
		e.Time(attr.Key, attr.Value.Time())
	case slog.KindGroup:
		e.Dict(attr.Key, zerolog.Dict().Func(func(dictEvent *zerolog.Event) {
			for _, groupAttr := range attr.Value.Group() {
				addSlogAttr(dictEvent, groupAttr)
			}
		}))
	case slog.KindAny:
		fallthrough
	default:
		if err, ok := attr.Value.Any().(error); ok {
			e.Err(err)
		} else {
			e.Any(attr.Key, attr.Value.Any())
		}
	}
}
