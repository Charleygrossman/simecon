package trade

import (
	"context"
	"tradesim/src/prob"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

type Trader struct {
	ID      uuid.UUID
	Haves   map[uuid.UUID]*Have
	Wants   map[uuid.UUID]*Want
	process *prob.Process
}

func NewTrader(have []Have, want []Want) *Trader {
	t := &Trader{
		ID:    uuid.New(),
		Haves: make(map[uuid.UUID]*Have, len(have)),
		Wants: make(map[uuid.UUID]*Want, len(want)),
	}
	for _, h := range have {
		t.Haves[h.Item.ID] = &h
	}
	for _, w := range want {
		t.Wants[w.Item.ID] = &w
	}
	return t
}

func (t *Trader) Start(ctx context.Context) error {
	wg, c := errgroup.WithContext(ctx)
	wg.Go(func() error { return t.process.Start(c) })
	wg.Go(func() error {
		select {
		// TODO: Send random trade request
		case <-t.process.Event:
		default:
		}
		return nil
	})
	return wg.Wait()
}

type Have struct {
	Item     Item
	Price    float64
	Quantity float64
}

type Want struct {
	Item     Item
	PriceMin float64
	PriceMax float64
	Quantity float64
}

type Item struct {
	ID   uuid.UUID
	Name string
}
