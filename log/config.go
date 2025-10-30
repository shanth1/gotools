package log

type Config struct {
	Level        string `yaml:"level"`
	Service      string `yaml:"service"`
	EnableCaller bool   `yaml:"enable_caller"`
	UDPAddress   string `yaml:"udp_address"`
	Console      bool   `yaml:"console"`
}
