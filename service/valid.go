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

func IsValidTimestamp(t string) (bool, error) {
	return true, nil
}
