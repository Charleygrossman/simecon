package internal

import (
	"context"
	"errors"
	"fmt"
	"tradesim/src/mkt"
	"tradesim/src/sim"
)

var ErrSim = errors.New("failed to run simulation")

func Simulate(inFilepath, outFilepath string) error {
	if err := simulate(inFilepath, outFilepath); err != nil {
		return fmt.Errorf("%w: %v", ErrSim, err)
	}
	return nil
}

func simulate(in, out string) error {
	config, err := sim.NewSimConfig(in)
	if err != nil {
		return err
	}

	clock := sim.ParseClock(config.Clock)
	traders := sim.ParseTraders(config.Traders)
	if len(traders) == 0 {
		return nil
	}

	// TODO: configurable
	edges := make([]mkt.Edge, 0, len(traders)-1)
	for i := 0; i < len(traders)-1; i++ {
		edges = append(edges, mkt.Edge{
			UTraderID: traders[i].ID,
			VTraderID: traders[i+1].ID,
			Delta:     1,
		})
	}

	graph := mkt.NewGraph(traders, edges, clock)
	return graph.Run(context.Background())
}
