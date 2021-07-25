package internal

import (
	"errors"
	"fmt"
	"tradesim/src/sim"
)

var ErrSim = errors.New("failed to run simulation")

func Simulate(inFilepath, outFilepath string) error {
	if err := simulate(inFilepath, outFilepath); err != nil {
		return fmt.Errorf("%w: %v", ErrSim, err)
	}
	return nil
}

// TODO: initialize and run market graph
func simulate(in, out string) error {
	config, err := sim.NewSimConfig(in)
	if err != nil {
		return err
	}

	sim.ParseClock(config.Clock)
	sim.ParseTraders(config.Traders)
	return nil
}
