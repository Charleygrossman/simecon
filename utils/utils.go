package utils

import (
	"fmt"
	"reflect"
	"strings"
)

// Returns a oneline string representation of a given struct
func StringStruct(t interface{}) string {
	s := reflect.ValueOf(t).Elem()
	typeOfT := s.Type()
	rep := []string{}
	for i := 0; i < s.NumField(); i++ {
		r := fmt.Sprintf("%s: %v", typeOfT.Field(i).Name, s.Field(i).Interface())
		rep = append(rep, r)
	}
	return fmt.Sprintf(strings.Join(rep, ", "))
}
