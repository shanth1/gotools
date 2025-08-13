package flags

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"
)

// RegisterFromStruct registers flags for the passed pointer to the structure with tags
//
// Tag description:
// - flag - flag name on command line
// - default (optional) - the default value for this flag
// - usage (optional) - the help text that will be shown when called with -h or -help
//
// Example: `flag:"level" default:"info" usage:"Logging level (debug, info, error)"`
func RegisterFromStruct(cfgPtr interface{}) error {
	val := reflect.ValueOf(cfgPtr)

	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("expected a pointer to a struct, but got %T", cfgPtr)
	}

	elem := val.Elem()
	t := elem.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldVal := elem.Field(i)

		flagName := field.Tag.Get("flag")
		if flagName == "" {
			continue
		}

		defaultValue := field.Tag.Get("default")
		usage := field.Tag.Get("usage")

		if !fieldVal.CanSet() {
			continue
		}

		switch field.Type.Kind() {
		case reflect.String:
			ptr := fieldVal.Addr().Interface().(*string)
			flag.StringVar(ptr, flagName, defaultValue, usage)
		case reflect.Int64:
			defaultValInt, _ := strconv.ParseInt(defaultValue, 10, 64)
			ptr := fieldVal.Addr().Interface().(*int64)
			flag.Int64Var(ptr, flagName, defaultValInt, usage)
		case reflect.Int:
			defaultValInt, _ := strconv.Atoi(defaultValue)
			ptr := fieldVal.Addr().Interface().(*int)
			flag.IntVar(ptr, flagName, defaultValInt, usage)
		case reflect.Bool:
			defaultValBool, _ := strconv.ParseBool(defaultValue)
			ptr := fieldVal.Addr().Interface().(*bool)
			flag.BoolVar(ptr, flagName, defaultValBool, usage)
		default:
			return fmt.Errorf("unsupported type for flag registration: %s", field.Type.Kind())
		}
	}

	return nil
}
