package internal

import (
	"context"
	"errors"
	"tradesim/src/sim/config"

	"golang.org/x/sync/errgroup"
)

var ErrSim = errors.New("failed to run simulation")

func Simulate(inFilepath, outFilepath string) error {
	cfg, err := config.NewSimConfig(inFilepath)
	if err != nil {
		return err
	}
	items := config.ParseItems(cfg.Items)
	traders := config.ParseTraders(cfg.Traders, items)
	exchange := config.ParseExchange(cfg.Exchange, items, traders)

	wg, ctx := errgroup.WithContext(context.Background())
	for _, t := range traders {
		_t := t
		wg.Go(func() error { return _t.Start(ctx) })
	}
	wg.Go(func() error { return exchange.Start(ctx) })
	return wg.Wait()
}
