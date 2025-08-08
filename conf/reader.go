package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

// Load loads the yaml config by accepting a path and a pointer to the structure
func Load(path string, cfg interface{}) error {
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}
	if err := v.Unmarshal(cfg); err != nil {
		return fmt.Errorf("error unmarshalling config: %w", err)
	}

	return nil
}
