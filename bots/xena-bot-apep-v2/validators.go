package main

import (
	"errors"
	"reflect"
)

func validString(data string, exception string) string {
	if reflect.TypeOf(data).Kind() != reflect.String {
		errors.New(exception)
	}

	return data
}
