package log

import "strings"

func levelToString(level level) string {
	switch level {
	case LevelTrace:
		return "trace"
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	default:
		return "info"
	}
}

func stringToLevel(level string) level {
	level = strings.ToLower(level)
	switch level {
	case "trace":
		return LevelTrace
	case "debug":
		return LevelDebug
	case "info":
		return LevelInfo
	case "warn":
		return LevelWarn
	case "error":
		return LevelError
	case "fatal":
		return LevelFatal
	case "panic":
		return LevelPanic
	default:
		return LevelInfo
	}
}
