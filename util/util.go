package util

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"
)

// MappedJson returns a map of all fields in the provided JSON-encoded data.
func MappedJson(data []byte) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal(data, &m)
	return m, err
}

// Now returns the current local time in the ISO 8601 format YYYY-MM-DD HH:MM:SS.
func Now() string {
	return strings.TrimSpace(strings.Split(time.Now().String(), "+")[0])
}

// ReversedStringSlice returns the provided string slice, reversed in-place.
func ReversedStringSlice(s []string) []string {
	for i := len(s)/2 - 1; i >= 0; i-- {
		j := len(s) - 1 - i
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// StringStruct returns a one-line string representation of the provided interface.
func StringStruct(t interface{}) string {
	v := reflect.ValueOf(t).Elem()
	rep := []string{}
	for i := 0; i < v.NumField(); i++ {
		rep = append(rep, fmt.Sprintf("%s:%v", v.Type().Field(i).Name, v.Field(i).Interface()))
	}
	return fmt.Sprint(strings.Join(rep, ","))
}
