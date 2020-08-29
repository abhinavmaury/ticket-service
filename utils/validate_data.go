package utils

import (
	"errors"
	"reflect"
)

func ValidateMandatoryData(args map[string]interface{}) error {
	for x, y := range args {
		if reflect.Indirect(reflect.ValueOf(y)) == reflect.ValueOf(nil) {
			return errors.New("No Details Given For " + x)
		} else if reflect.Indirect(reflect.ValueOf(y)) == reflect.ValueOf(0) {
			return errors.New("No Details Given For " + x)
		} else if y == "" {
			return errors.New("No Details Given For " + x)
		}
	}
	return nil
}
