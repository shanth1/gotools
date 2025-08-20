package log

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"

	"github.com/rs/zerolog"
)

type config struct {
	level        level
	writers      []io.Writer
	service      string
	enableCaller bool
}

// option defines a function for configuring the logger.
type option func(*config)

// WithLevel sets the logging level (debug, info, warn, error).
func WithLevel(level level) option {
	return func(c *config) {
		c.level = level
	}
}

// WithWriter adds a custom writer for outputting logs.
func WithWriter(w io.Writer) option {
	return func(c *config) {
		c.writers = append(c.writers, w)
	}
}

// WithConsoleWriter adds a user-friendly console writer.
func WithConsoleWriter() option {
	return func(c *config) {
		c.writers = append(c.writers, zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		})
	}
}

// WithService adds the service name to all log entries
func WithService(service string) option {
	return func(c *config) {
		c.service = service
	}
}

// WithCaller add information about the file and line where the call came from
func WithCaller() option {
	return func(c *config) {
		c.enableCaller = true
	}
}

// WithUDPWriter adds a writer that sends JSON logs over UDP to the specified address.
func WithUDPWriter(addr string) option {
	return func(c *config) {
		udpAddr, err := net.ResolveUDPAddr("udp", addr)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Failed to resolve UDP address %s: %v\n", addr, err)
			return
		}

		conn, err := net.DialUDP("udp", nil, udpAddr)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Failed to connect to UDP address %s: %v\n", addr, err)
			return
		}

		c.writers = append(c.writers, conn)
	}
}
