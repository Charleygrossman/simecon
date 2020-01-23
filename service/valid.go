package service

import (
	"fmt"
)

type Validator interface {
	Validate() error
}

var (
	ErrInvalidTimestamp = fmt.Errorf("invalid timestamp")
)

func ValidTimestamp(t string) (bool, error) {
	return true, nil
}
