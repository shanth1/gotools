package log

import (
	"fmt"
	"sort"
	"strings"
)

var levelToStringMap = map[level]string{
	LevelTrace:    "trace",
	LevelDebug:    "debug",
	LevelInfo:     "info",
	LevelWarn:     "warn",
	LevelError:    "error",
	LevelFatal:    "fatal",
	LevelPanic:    "panic",
	LevelDisabled: "disabled",
}

func levelToString(level level) string {
	if s, ok := levelToStringMap[level]; ok {
		return s
	}
	return "info"
}

var stringToLevelMap = map[string]level{
	"trace":    LevelTrace,
	"debug":    LevelDebug,
	"info":     LevelInfo,
	"warn":     LevelWarn,
	"error":    LevelError,
	"fatal":    LevelFatal,
	"panic":    LevelPanic,
	"disabled": LevelDisabled,
	"off":      LevelDisabled,
	"none":     LevelDisabled,
}

// ParseLevel converts a string to a Level.
// It returns an error if the string is not a valid log level.
func ParseLevel(lvl string) (level, error) {
	if l, ok := stringToLevelMap[strings.ToLower(lvl)]; ok {
		return l, nil
	}
	return LevelInfo, fmt.Errorf("unknown log level: %q. Valid levels: %s", lvl, validLevelsString())
}

// validLevelsString returns a sorted, comma-separated string of valid levels for error messages.
func validLevelsString() string {
	var levels []string
	for k := range stringToLevelMap {
		levels = append(levels, k)
	}
	sort.Strings(levels)
	return strings.Join(levels, ", ")
}
