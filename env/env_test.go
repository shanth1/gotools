package env

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestEnvConfig struct {
	DBHost   string `env:"DB_HOST"`
	DBPort   int    `env:"DB_PORT" env-default:"5432"`
	Override string `env:"OVERRIDE_ME"`
}

func createTestEnvFile(t *testing.T, content string) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, ".env")
	err := os.WriteFile(path, []byte(content), 0644)
	require.NoError(t, err)
	return path
}

func TestLoadIntoStruct(t *testing.T) {
	t.Run("success with file and system env", func(t *testing.T) {
		envContent := "DB_HOST=db.from.file\nOVERRIDE_ME=file_value"
		path := createTestEnvFile(t, envContent)

		// Set system env var which should take precedence
		require.NoError(t, os.Setenv("OVERRIDE_ME", "system_value"))
		t.Cleanup(func() {
			os.Unsetenv("OVERRIDE_ME")
		})

		var cfg TestEnvConfig
		err := LoadIntoStruct(path, &cfg)

		require.NoError(t, err)
		assert.Equal(t, "db.from.file", cfg.DBHost)
		assert.Equal(t, 5432, cfg.DBPort)             // from default tag
		assert.Equal(t, "system_value", cfg.Override) // system var wins
	})

	t.Run("success with only system env", func(t *testing.T) {
		require.NoError(t, os.Setenv("DB_HOST", "db.from.system"))
		t.Cleanup(func() {
			os.Unsetenv("DB_HOST")
		})

		var cfg TestEnvConfig
		// Empty path skips file loading
		err := LoadIntoStruct("", &cfg)

		require.NoError(t, err)
		assert.Equal(t, "db.from.system", cfg.DBHost)
		assert.Equal(t, 5432, cfg.DBPort)
	})

	t.Run("file not found", func(t *testing.T) {
		var cfg TestEnvConfig
		err := LoadIntoStruct("non-existent.env", &cfg)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "read env file")
	})

	t.Run("target is not a pointer", func(t *testing.T) {
		var cfg TestEnvConfig
		err := LoadIntoStruct("", cfg)
		require.Error(t, err)
		assert.EqualError(t, err, "expected a pointer to a struct, but got env.TestEnvConfig")
	})

	t.Run("target is not a pointer to struct", func(t *testing.T) {
		var i int
		err := LoadIntoStruct("", &i)
		require.Error(t, err)
		assert.EqualError(t, err, "expected a pointer to a struct, but got *int")
	})
}
