package log

import "strings"

var levelToStringMap = map[level]string{
	LevelTrace: "trace",
	LevelDebug: "debug",
	LevelInfo:  "info",
	LevelWarn:  "warn",
	LevelError: "error",
	LevelFatal: "fatal",
	LevelPanic: "panic",
}

func levelToString(level level) string {
	if s, ok := levelToStringMap[level]; ok {
		return s
	}
	return "info"
}

var stringToLevelMap = map[string]level{
	"trace": LevelTrace,
	"debug": LevelDebug,
	"info":  LevelInfo,
	"warn":  LevelWarn,
	"error": LevelError,
	"fatal": LevelFatal,
	"panic": LevelPanic,
}

func stringToLevel(level string) level {
	if l, ok := stringToLevelMap[strings.ToLower(level)]; ok {
		return l
	}
	return LevelInfo
}
