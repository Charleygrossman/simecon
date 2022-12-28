package trade

import (
	"context"
	"math/rand"
	"time"
	"tradesim/src/prob"
	"tradesim/src/time/clock"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

// Have represents an item that a trader holds,
// along with their evaluation of its unit price
// and how many units they hold.
type Have struct {
	Item     Item
	Price    float64
	Quantity float64
}

// Want represents an item that a trader would like to hold,
// along with their evaluation of its minimum and maximum
// unit price and how many units they would like to hold.
type Want struct {
	Item     Item
	PriceMin float64
	PriceMax float64
	Quantity float64
}

// Trader represents an entity participating
// in the exchange of items with other traders.
type Trader struct {
	ID           uuid.UUID
	Haves        map[uuid.UUID]*Have
	Wants        map[uuid.UUID]*Want
	RequestSend  chan Request
	RequestRecv  chan Request
	ResponseSend chan Response
	ResponseRecv chan Responses
	Choice       chan Response
	process      *prob.Process
}

func NewTrader(haves []Have, wants []Want) *Trader {
	t := &Trader{
		ID:           uuid.New(),
		Haves:        make(map[uuid.UUID]*Have, len(haves)),
		Wants:        make(map[uuid.UUID]*Want, len(wants)),
		RequestSend:  make(chan Request, 8),
		RequestRecv:  make(chan Request, 8),
		ResponseSend: make(chan Response, 8),
		ResponseRecv: make(chan Responses, 8),
		Choice:       make(chan Response, 8),
		// TODO: configurable
		process: prob.NewProcess(prob.NewUniform(0.2), clock.NewClock(time.Second, 0)),
	}
	for _, h := range haves {
		t.Haves[h.Item.ID] = &h
	}
	for _, w := range wants {
		t.Wants[w.Item.ID] = &w
	}
	return t
}

func (t *Trader) Start(ctx context.Context) error {
	wg, c := errgroup.WithContext(ctx)
	wg.Go(func() error { return t.process.Start(c) })
	wg.Go(func() error { return t.sendRequest(c) })
	wg.Go(func() error { return t.sendResponse(c) })
	wg.Go(func() error { return t.sendChoice(c) })
	return wg.Wait()
}

func (t *Trader) sendRequest(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.process.Event:
		r, ok := t.randomRequest()
		if ok {
			t.RequestSend <- r
		}
	}
	return nil
}

func (t *Trader) randomRequest() (Request, bool) {
	if len(t.Wants) == 0 {
		return Request{}, false
	}
	ws := make([]*Want, len(t.Wants))
	i := 0
	for _, v := range t.Wants {
		ws[i] = v
		i++
	}
	w := ws[rand.Intn(len(ws))]
	return Request{
		ID:       uuid.New(),
		TraderID: t.ID,
		Item:     w.Item,
		Quantity: w.Quantity,
		Side:     SideBuy,
	}, true
}

func (t *Trader) sendResponse(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case req := <-t.RequestRecv:
		resp, ok := t.response(req)
		if ok {
			t.ResponseSend <- resp
		}
	}
	return nil
}

func (t *Trader) response(req Request) (Response, bool) {
	h, have := t.Haves[req.Item.ID]
	w, want := t.Wants[req.Item.ID]
	if !(have || want) {
		return Response{}, false
	}
	r := Response{
		ID:        uuid.New(),
		Request:   req,
		TraderID:  t.ID,
		OrderBook: OrderBook{},
	}
	if have {
		r.OrderBook.Ask.Item = h.Item
		r.OrderBook.Ask.Price = h.Price
		r.OrderBook.Ask.Quantity = h.Quantity
	}
	if want {
		r.OrderBook.Bid.Item = w.Item
		r.OrderBook.Bid.Price = w.PriceMax
		r.OrderBook.Bid.Quantity = w.Quantity
	}
	return r, true
}

func (t *Trader) sendChoice(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case resps := <-t.ResponseRecv:
		c, ok := t.randomChoice(resps)
		if ok {
			t.Choice <- c
		}
	}
	return nil
}

func (t *Trader) randomChoice(resp Responses) (Response, bool) {
	if len(resp) == 0 {
		return Response{}, false
	}
	return resp[rand.Intn(len(resp))], true
}
