# Go Common

A collection of common, reusable Go packages for building applications. It includes utilities for configuration, logging, context management, command-line flags, and more.

Of course. I will update all documentation and comments to English, restructure the documentation into separate `README.md` files for each package, and ensure the formatting is correct.

Here is the complete set of updated files for your project.

---

### **Root Project Files**

This is the main `README.md` with an overview and links to each package's documentation.

================================================
FILE: README.md
================================================

# Go Common

A collection of common, reusable Go packages for building applications. It includes utilities for configuration, logging, context management, command-line flags, and more.

## Installation

```sh
go get github.com/shanth1/gotools
```

## Packages

- [**`conf`**](./conf/.md): Load configuration from YAML files into Go structs.
- [**`flags`**](./flags/.md): Register command-line flags from struct tags.
- [**`log`**](./log/.md): A structured, leveled logging wrapper around `zerolog`.
- [**`ctx`**](./ctx/.md): Helpers for graceful shutdown and request-scoped context values.
- [**`env`**](./env/.md): Load environment variables from the system and `.env` files.
- [**`notify`**](./notify/.md): Notification services with support for Telegram and Email.
- [**`errs`**](./errs/.md): A set of pre-defined, common application errors.
- [**`consts`**](./consts/.md): Pre-defined constants for environments, statuses, etc.

You are absolutely right, my apologies for missing that package. I will create the documentation and add comments for the `env` package now.

I will also update the main `README.md` file to include a link to this new package's documentation.

---

### **Package `env`**

Here are the new documentation file and the updated source file for the `env` package.

================================================
FILE: env/README.md
================================================

================================================
FILE: env/env.go
================================================
// Package env provides a utility for loading environment variables
// from the system and .env files into a Go struct.
package env

import (
"fmt"
"reflect"

    "github.com/ilyakaznacheev/cleanenv"
    "github.com/joho/godotenv"

)

// LoadIntoStruct loads configuration from system environment variables and,
// optionally, from a .env file into a target struct.
//
// The cfgPtr argument must be a pointer to the struct. The envPath is the
// path to the .env file; if it's an empty string, the file is skipped.
//
// Loading Priority: System Environment Variables > .env file.
// This means existing environment variables will not be overwritten by the .env file.
//
// Fields in the struct must be tagged with `env`, e.g., `env:"DB_HOST"`.
func LoadIntoStruct(envPath string, cfgPtr interface{}) error {
val := reflect.ValueOf(cfgPtr)
if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
return fmt.Errorf("expected a pointer to a struct, but got %T", cfgPtr)
}

    if envPath != "" {
    	// godotenv.Load does not override existing environment variables.
    	if err := godotenv.Load(envPath); err != nil {
    		return fmt.Errorf("read env file: %w", err)
    	}
    }

    if err := cleanenv.ReadEnv(cfgPtr); err != nil {
    	return fmt.Errorf("read environment variables: %w", err)
    }

    return nil

}
