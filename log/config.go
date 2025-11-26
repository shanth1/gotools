package log

type Config struct {
	// Level sets the logging level.
	// Valid values: trace, debug, info, warn, error, fatal, panic, disabled, off, none.
	Level string `mapstructure:"level" yaml:"level" json:"level" toml:"level"`

	// App name to be included in all logs.
	App string `mapstructure:"app" yaml:"app" json:"app" toml:"app"`

	// Service name to be included in all logs.
	Service string `mapstructure:"service" yaml:"service" json:"service" toml:"service"`

	// EnableCaller adds file and line number to logs.
	EnableCaller bool `mapstructure:"enable_caller" yaml:"enable_caller" json:"enable_caller" toml:"enable_caller"`

	// UDPAddress is the address to send JSON logs to (e.g. "127.0.0.1:1234").
	UDPAddress string `mapstructure:"udp_address" yaml:"udp_address" json:"udp_address" toml:"udp_address"`

	// Console enables pretty printing to stdout/stderr instead of JSON.
	Console bool `mapstructure:"console" yaml:"console" json:"console" toml:"console"`

	// JSONOutput enforces JSON output even if a terminal is detected.
	JSONOutput bool `mapstructure:"json_output" yaml:"json_output" json:"json_output" toml:"json_output"`
}
