package conf

import (
	"fmt"
	"reflect"

	"github.com/spf13/viper"
)

// Load reads a YAML configuration file from the given path into a struct.
// The cfgPtr argument must be a pointer to the struct that will hold the configuration.
func Load(path string, cfgPtr interface{}) error {
	val := reflect.ValueOf(cfgPtr)

	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("expected a pointer to a struct, but got %T", cfgPtr)
	}

	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}
	if err := v.Unmarshal(cfgPtr); err != nil {
		return fmt.Errorf("error unmarshalling config: %w", err)
	}

	return nil
}
