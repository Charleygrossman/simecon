package internal

import (
	"context"
	"errors"
	"fmt"
	"tradesim/src/exchange"
	"tradesim/src/sim/config"
)

var ErrSim = errors.New("failed to run simulation")

func Simulate(inFilepath, outFilepath string) error {
	exchange, err := parseExchange(inFilepath)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrSim, err)
	}
	return exchange.Start(context.Background())
}

func parseExchange(inFilepath string) (*exchange.Exchange, error) {
	c, err := config.NewSimConfig(inFilepath)
	if err != nil {
		return nil, err
	}
	items := config.ParseItems(c.Items)
	traders := config.ParseTraders(c.Traders, items)
	exchange := config.ParseExchange(c.Exchange, items, traders)
	return exchange, nil
}
