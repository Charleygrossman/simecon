package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"time"
)

// DecodeBody decodes and returns an HTTP request body as a byte array.
func DecodeBody(r *http.Request) (b []byte, err error) {
	b, err = ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	return b, err
}

// JsonToMap takes a byte array it expects to be json,
// and returns a map of the json fields.
func JsonToMap(j []byte) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal(j, &m)
	return m, err
}

// Now returns current local time in the ISO 8601 format YYYY-MM-DD HH:MM:SS
func Now() string {
	t := fmt.Sprintf("%v", time.Now())
	ts := strings.Split(t, " +")
	return ts[0]
}

// StringStruct returns a one-line string representation of the provided interface.
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
