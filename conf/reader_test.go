package conf

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestConfig struct {
	Service string `yaml:"service"`
	Port    int    `yaml:"port"`
	Enabled bool   `yaml:"enabled"`
}

func createTestYAML(t *testing.T, content string) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "config.yaml")
	err := os.WriteFile(path, []byte(content), 0644)
	require.NoError(t, err)
	return path
}

func TestLoad(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		yamlContent := `
service: my-app
port: 8080
enabled: true
`
		path := createTestYAML(t, yamlContent)

		var cfg TestConfig
		err := Load(path, &cfg)

		require.NoError(t, err)
		assert.Equal(t, "my-app", cfg.Service)
		assert.Equal(t, 8080, cfg.Port)
		assert.True(t, cfg.Enabled)
	})

	t.Run("file not found", func(t *testing.T) {
		t.Parallel()
		var cfg TestConfig
		err := Load("non-existent-file.yaml", &cfg)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "error reading config file")
	})

	t.Run("invalid yaml format", func(t *testing.T) {
		t.Parallel()
		path := createTestYAML(t, "service: my-app\n  port: 8080") // bad indentation
		var cfg TestConfig
		err := Load(path, &cfg)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "error reading config file")
	})

	t.Run("target is not a pointer", func(t *testing.T) {
		t.Parallel()
		path := createTestYAML(t, "port: 8080")
		var cfg TestConfig
		err := Load(path, cfg)
		require.Error(t, err)
		assert.EqualError(t, err, "expected a pointer to a struct, but got conf.TestConfig")
	})

	t.Run("target is not a pointer to struct", func(t *testing.T) {
		t.Parallel()
		path := createTestYAML(t, "port: 8080")
		var i int
		err := Load(path, &i)
		require.Error(t, err)
		assert.EqualError(t, err, "expected a pointer to a struct, but got *int")
	})
}
