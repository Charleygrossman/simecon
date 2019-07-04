// Package utils provides utility functions for general use.
package utils

import (
	"fmt"
	"reflect"
	"strings"
)

// StringStruct returns a oneline string representation of the provided interface.
func StringStruct(t interface{}) string {
	s := reflect.ValueOf(t).Elem()
	typeOfT := s.Type()
	rep := []string{}
	for i := 0; i < s.NumField(); i++ {
		f := fmt.Sprintf("%s: %v", typeOfT.Field(i).Name, s.Field(i).Interface())
		rep = append(rep, f)
	}
	return fmt.Sprintf(strings.Join(rep, ", "))
}
