package log

type Config struct {
	Level        string `mapstructure:"level" yaml:"level" json:"level" toml:"level"`
	App          string `mapstructure:"app" yaml:"app" json:"app" toml:"app"`
	Service      string `mapstructure:"service" yaml:"service" json:"service" toml:"service"`
	EnableCaller bool   `mapstructure:"enable_caller" yaml:"enable_caller" json:"enable_caller" toml:"enable_caller"`
	UDPAddress   string `mapstructure:"udp_address" yaml:"udp_address" json:"udp_address" toml:"udp_address"`
	Console      bool   `mapstructure:"console" yaml:"console" json:"console" toml:"console"`
}
