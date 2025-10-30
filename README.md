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
