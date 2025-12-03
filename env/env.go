package env

import (
	"fmt"
	"reflect"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

// LoadIntoStruct loads data from variables and env file into structure
// Struct tags example: `env:"TEST"`
//
// Priority: System .env > File .env
// If the path to the file is not specified, only system variables are read.
func LoadIntoStruct(envPath string, cfgPtr interface{}) error {
	val := reflect.ValueOf(cfgPtr)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("expected a pointer to a struct, but got %T", cfgPtr)
	}

	if envPath != "" {
		if err := godotenv.Load(envPath); err != nil {
			return fmt.Errorf("read env file: %w", err)
		}
	}

	if err := cleanenv.ReadEnv(cfgPtr); err != nil {
		return fmt.Errorf("read environment variables: %w", err)
	}

	return nil
}
