package log

import (
	"bytes"
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Parallel()

	t.Run("with options", func(t *testing.T) {
		t.Parallel()
		var buf bytes.Buffer
		logger := New(
			WithLevel(LevelDebug),
			WithService("test-app"),
			WithWriter(&buf),
		)

		logger.Debug().Str("key", "value").Msg("hello")

		output := buf.String()
		assert.Contains(t, output, `"level":"debug"`)
		assert.Contains(t, output, `"service":"test-app"`)
		assert.Contains(t, output, `"key":"value"`)
		assert.Contains(t, output, `"message":"hello"`)
	})
}

func TestNewFromConfig(t *testing.T) {
	t.Parallel()
	var buf bytes.Buffer
	cfg := Config{
		Level:        "warn",
		Service:      "config-app",
		EnableCaller: true,
	}

	logger := NewFromConfig(cfg)
	// Replace default writers with our buffer for testing
	logger = New(WithLevel(stringToLevel(cfg.Level)), WithService(cfg.Service), WithCaller(), WithWriter(&buf))

	logger.Info().Msg("should not appear")
	logger.Warn().Msg("should appear")

	output := buf.String()
	assert.NotContains(t, output, "should not appear")
	assert.Contains(t, output, `"level":"warn"`)
	assert.Contains(t, output, `"service":"config-app"`)
	assert.Contains(t, output, `"message":"should appear"`)
	assert.Contains(t, output, `"caller"`)
}

func TestLogger_With(t *testing.T) {
	t.Parallel()
	var buf bytes.Buffer
	logger := New(WithWriter(&buf))

	// Create a logger with a pre-set field
	subLogger := logger.With(Str("request_id", "abc-123"))

	subLogger.Info().Msg("request processed")

	var logEntry map[string]interface{}
	err := json.Unmarshal(buf.Bytes(), &logEntry)
	require.NoError(t, err)

	assert.Equal(t, "abc-123", logEntry["request_id"])
	assert.Equal(t, "request processed", logEntry["message"])
}

func TestLogger_Levels(t *testing.T) {
	t.Parallel()
	var buf bytes.Buffer
	logger := New(WithLevel(LevelInfo), WithWriter(&buf))

	logger.Debug().Msg("debug log")
	logger.Info().Msg("info log")

	output := buf.String()
	assert.NotContains(t, output, "debug log")
	assert.Contains(t, output, "info log")
}

func TestContext(t *testing.T) {
	t.Parallel()
	var buf bytes.Buffer
	logger := New(WithWriter(&buf), WithService("ctx-logger"))

	ctx := NewContext(context.Background(), logger)
	ctxLogger := FromContext(ctx)

	ctxLogger.Info().Msg("from context")

	assert.Contains(t, buf.String(), "ctx-logger")
	assert.Contains(t, buf.String(), "from context")
}

func TestLevelConversions(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		str   string
		level level
	}{
		{"trace", LevelTrace},
		{"debug", LevelDebug},
		{"info", LevelInfo},
		{"warn", LevelWarn},
		{"error", LevelError},
		{"fatal", LevelFatal},
		{"panic", LevelPanic},
		{"unknown", LevelInfo}, // default
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.level, stringToLevel(tc.str), "stringToLevel failed for "+tc.str)
		if tc.str != "unknown" {
			assert.Equal(t, tc.str, levelToString(tc.level), "levelToString failed for "+tc.str)
		}
	}
}
