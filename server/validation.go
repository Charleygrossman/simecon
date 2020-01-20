package server

import (
	"errors"
)

type Validator interface {
	Validate() error
}

var (
	ErrInvalidTimestamp = errors.New("invalid timestamp")
)

// NOTE: Booleans for these validators not only indicates a match, but also a failure to validate.

func IsValidTimestamp(t string) (bool, error) {
	return true, nil
}
