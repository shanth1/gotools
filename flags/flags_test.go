package flags

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestFlagConfig struct {
	Host    string `flag:"host" default:"localhost" usage:"Server host"`
	Port    int    `flag:"port" default:"8080"`
	Retries int64  `flag:"retries"`
	Enabled bool   `flag:"enabled" default:"true"`
	NoFlag  string
}

func TestMain(m *testing.M) {
	if os.Getenv("GO_TEST_SUBPROCESS") == "1" {
		var cfg TestFlagConfig
		err := RegisterFromStruct(&cfg)
		if err != nil {
			fmt.Printf("Error registering from struct: %v", err)
			os.Exit(1)
		}

		flag.Parse()

		fmt.Printf("Host=%s,Port=%d,Enabled=%v", cfg.Host, cfg.Port, cfg.Enabled)
		os.Exit(0)
	}

	os.Exit(m.Run())
}

func TestRegisterFromStruct_Subprocess(t *testing.T) {
	t.Run("registers and parses flags correctly", func(t *testing.T) {
		cmd := exec.Command(os.Args[0], "-test.run", "^TestMain$")
		cmd.Env = append(os.Environ(), "GO_TEST_SUBPROCESS=1")
		cmd.Args = append(cmd.Args, "-host", "remote.server", "-port", "3000")

		output, err := cmd.CombinedOutput()
		require.NoError(t, err, "Subprocess execution failed with output: %s", string(output))

		expected := "Host=remote.server,Port=3000,Enabled=true"
		assert.Equal(t, expected, strings.TrimSpace(string(output)))
	})

	t.Run("uses default values", func(t *testing.T) {
		cmd := exec.Command(os.Args[0], "-test.run", "^TestMain$")
		cmd.Env = append(os.Environ(), "GO_TEST_SUBPROCESS=1")

		output, err := cmd.CombinedOutput()
		require.NoError(t, err, "Subprocess execution failed with output: %s", string(output))

		expected := "Host=localhost,Port=8080,Enabled=true"
		assert.Equal(t, expected, strings.TrimSpace(string(output)))
	})
}

func TestRegisterFromStruct_Unit(t *testing.T) {
	t.Run("unsupported type", func(t *testing.T) {
		type UnsupportedConfig struct {
			Value float64 `flag:"value"`
		}
		var cfg UnsupportedConfig
		err := RegisterFromStruct(&cfg)
		require.Error(t, err)
		assert.EqualError(t, err, "unsupported type for flag registration: float64")
	})

	t.Run("target is not a pointer", func(t *testing.T) {
		var cfg TestFlagConfig
		err := RegisterFromStruct(cfg)
		require.Error(t, err)
		assert.EqualError(t, err, "expected a pointer to a struct, but got flags.TestFlagConfig")
	})
}
