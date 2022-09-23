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

func NewTrader(have []Have, want []Want) *Trader {
	t := &Trader{
		ID:           uuid.New(),
		Haves:        make(map[uuid.UUID]*Have, len(have)),
		Wants:        make(map[uuid.UUID]*Want, len(want)),
		RequestSend:  make(chan Request),
		RequestRecv:  make(chan Request),
		ResponseSend: make(chan Response),
		ResponseRecv: make(chan Responses),
		Choice:       make(chan Response),
		process:      prob.NewProcess(prob.NewUniform(0.5), clock.NewClock(time.Second, 0)),
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
	wg.Go(t.sendRequest)
	wg.Go(t.sendResponse)
	wg.Go(t.choose)
	return wg.Wait()
}

func (t *Trader) sendRequest() error {
	select {
	case <-t.process.Event:
		r, ok := t.randomRequest()
		if ok {
			t.RequestSend <- r
		}
	default:
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
		Item:     w.Item,
		Quantity: w.Quantity,
		Side:     SideBuy,
	}, true
}

func (t *Trader) sendResponse() error {
	select {
	case req := <-t.RequestRecv:
		resp, ok := t.response(req)
		if ok {
			t.ResponseSend <- resp
		}
	default:
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

func (t *Trader) choose() error {
	select {
	case resps := <-t.ResponseRecv:
		c, ok := t.randomChoice(resps)
		if ok {
			t.Choice <- c
		}
	default:
	}
	return nil
}

func (t *Trader) randomChoice(resp Responses) (Response, bool) {
	if len(resp) == 0 {
		return Response{}, false
	}
	return resp[rand.Intn(len(resp))], true
}
