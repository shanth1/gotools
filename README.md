# Go Common

A collection of common, reusable Go packages for building applications. It includes utilities for configuration, logging, context management, command-line flags, and more.

## Installation

```sh
go get github.com/shanth1/gotools
```

---

## Packages

- [**`conf`**](#conf---configuration-loader): Load YAML configuration files into Go structs.
- [**`flags`**](#flags---cli-flag-parsing): Register command-line flags from struct tags.
- [**`log`**](#log---structured-logging): A structured, leveled logging wrapper around `zerolog`.
- [**`ctx`**](#ctx---context-management): Helpers for graceful shutdown and request-scoped context values.
- [**`notify`**](#notify---notifications): Notification service with support for Telegram and Email.
- [**`errs`**](#errs---standard-errors): A set of pre-defined, common application errors.
- [**`consts`**](#consts---common-constants): Pre-defined constants for environments, statuses, etc.

---

### `conf` - Configuration Loader

The `conf` package provides a simple function to load a YAML configuration file into a struct.

**Usage**

Given a `config.yaml` file:

```yaml
# ./config.yaml
port: 8080
log_level: 'debug'
```

You can load it like this:

```go
package main

import (
	"fmt"
	"log"

	"github.com/shanth1/gotools/conf"
)

type Config struct {
	Port     int    `yaml:"port"`
	LogLevel string `yaml:"log_level"`
}

func main() {
	var cfg Config
	if err := conf.Load("config.yaml", &cfg); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	fmt.Printf("Port: %d, Log Level: %s\n", cfg.Port, cfg.LogLevel)
	// Output: Port: 8080, Log Level: debug
}
```

### `flags` - CLI Flag Parsing

This package allows you to define command-line flags using struct tags. It automatically registers flags and populates the struct instance after parsing.

**Supported tags:**

- `flag:"<name>"`: The name of the flag (e.g., `--port`).
- `default:"<value>"`: The default value for the flag.
- `usage:"<description>"`: The help text for the flag.

**Usage**

```go
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/shanth1/gotools/flags"
)

type AppConfig struct {
	Port    int    `flag:"port" default:"8080" usage:"TCP port to listen on"`
	LogLevel string `flag:"log-level" default:"info" usage:"Logging level (debug, info)"`
	DevMode  bool   `flag:"dev" default:"false" usage:"Enable development mode"`
}

func main() {
	var cfg AppConfig
	if err := flags.RegisterFromStruct(&cfg); err != nil {
		log.Fatalf("failed to register flags: %v", err)
	}
	flag.Parse()

	fmt.Printf("Starting with config: %+v\n", cfg)
}
```

**Example execution:**

```sh
# Run with custom flags
go run main.go --port=3000 --dev

# Output:
# Starting with config: {Port:3000 LogLevel:info DevMode:true}
```

### `log` - Structured Logging

A flexible logging library built on `zerolog`. It supports configuration via options or a struct, multiple writers (console, UDP), and context integration.

**Usage**

```go
package main

import (
	"errors"

	"github.com/shanth1/gotools/log"
)

func main() {
	// Simple initialization with options
	logger := log.New(
		log.WithService("my-app"),
		log.WithLevel(log.LevelDebug),
		log.WithCaller(),
	)

	logger.Info().Msg("Service starting...")
	logger.Debug().Str("user_id", "123").Msg("User logged in")
	logger.Error().Err(errors.New("database connection failed")).Msg("A critical error occurred")
}
```

You can also initialize the logger from a configuration struct, which is useful when combined with the `conf` package.

### `ctx` - Context Management

Provides helpers for handling application lifecycle signals (graceful shutdown) and managing request-scoped values in `context.Context`.

**Usage: Graceful Shutdown**

```go
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/shanth1/gotools/ctx"
)

func main() {
	// Create contexts for the app and its shutdown period.
	appCtx, shutdownCtx, cancel, shutdownCancel := ctx.WithGracefulShutdown(10 * time.Second)
	defer cancel()
	defer shutdownCancel()

	server := &http.Server{Addr: ":8080"}

	// Run the server in a goroutine
	go func() {
		fmt.Println("Server starting on :8080")
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Printf("Server error: %v\n", err)
		}
	}()

	// Wait for an OS signal (e.g., Ctrl+C)
	<-appCtx.Done()
	fmt.Println("Shutdown signal received. Shutting down gracefully...")

	// Attempt to shut down the server within the timeout
	if err := server.Shutdown(shutdownCtx); err != nil {
		fmt.Printf("Server shutdown failed: %v\n", err)
	} else {
		fmt.Println("Server stopped successfully.")
	}
}
```

**Usage: Context Values**

```go
package main

import (
	"context"
	"fmt"

	"github.com/shanth1/gotools/ctx"
)

func processRequest(c context.Context) {
	if reqID, ok := ctx.RequestIDFrom(c); ok {
		fmt.Printf("Processing request with ID: %s\n", reqID)
	} else {
		fmt.Println("No request ID found in context.")
	}
}

func main() {
	// Create a context with a request ID
	c := ctx.WithRequestID(context.Background(), "req-abc-123")
	processRequest(c) // Output: Processing request with ID: req-abc-123
}
```

### `notify` - Notifications

A service for sending notifications through different channels like Telegram or email.

**Usage**

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/shanth1/gotools/notify"
)

func main() {
	telegramToken := os.Getenv("TELEGRAM_TOKEN")
	chatID := os.Getenv("TELEGRAM_CHAT_ID")

	notifier, err := notify.New(
		notify.WithTelegram(telegramToken, chatID),
	)
	if err != nil {
		log.Fatalf("Failed to create notifier: %v", err)
	}

	err = notifier.Send(context.Background(), "Hello from the *go-common* library!")
	if err != nil {
		log.Printf("Failed to send notification: %v", err)
	}
}
```

### `errs` - Standard Errors

Defines a set of standard, exported errors for common use cases in applications. It also includes a `Wrap` function to add context to errors.

**Usage**

```go
package main

import (
	"database/sql"
	"fmt"

	"github.com/shanth1/gotools/errs"
)

func findUserByID(id int) error {
	// Simulate a database call
	err := sql.ErrNoRows
	if err == sql.ErrNoRows {
		return errs.Wrap(errs.ErrNotFound, fmt.Sprintf("user with id %d", id))
	}
	return nil
}

func main() {
	err := findUserByID(42)
	fmt.Println(err) // Output: user with id 42: resource not found
}
```

### `consts` - Common Constants

This package provides sets of string constants to ensure consistency across your application.

**Available Constants**

```go
// Environments
const (
	EnvLocal = "local"
	EnvDev   = "development"
	EnvProd  = "production"
)

// Statuses
const (
	StatusSuccess = "success"
	StatusError   = "error"
)

// Content Types
const (
	ContentTypeJSON = "application/json"
	ContentTypeText = "text/plain"
)
```
