package internal

import (
	"errors"
	"fmt"
)

var ErrSim = errors.New("failed to run simulation")

func Simulate(inFilepath, outFilepath string) error {
	if err := simulate(inFilepath, outFilepath); err != nil {
		return fmt.Errorf("%w: %v", ErrSim, err)
	}
	return nil
}

func simulate(in, out string) error {
	return errors.New("not implemented")
}
