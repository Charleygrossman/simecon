package internal

import (
	"context"
	"errors"

	"time"
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

	ctx := context.Background()
	if cfg.Duration > 0 {
		c, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Duration)*time.Second)
		ctx = c
		defer cancel()
	}
	wg, c := errgroup.WithContext(ctx)
	for _, t := range traders {
		_t := t
		wg.Go(func() error { return _t.Start(c) })
	}
	wg.Go(func() error { return exchange.Start(c) })
	if err := wg.Wait(); err != nil && !errors.Is(err, context.DeadlineExceeded) {
		return err
	}
	return exchange.DB.Write(outFilepath)
}
