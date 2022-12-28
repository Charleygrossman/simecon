package exchange

import (
	"context"
	"fmt"
	"sync"
	"tradesim/src/db"
	"tradesim/src/trade"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

// Market represents a tradable item on an exchange,
// and a set of traders participating in the market.
type Market struct {
	Item       trade.Item
	TraderByID map[uuid.UUID]*trade.Trader
}

func NewMarket(item trade.Item, traders ...*trade.Trader) Market {
	m := Market{
		Item:       item,
		TraderByID: make(map[uuid.UUID]*trade.Trader, len(traders)),
	}
	for _, t := range traders {
		m.TraderByID[t.ID] = t
	}
	return m
}

// Exchange represents a centralized item exchange that facilitates
// the interactions between traders participating in a set of markets.
type Exchange struct {
	Markets map[uuid.UUID]Market
	DB      *db.Blockchain
	dbLock  sync.Mutex
}

func NewExchange(markets []Market) *Exchange {
	e := &Exchange{
		Markets: make(map[uuid.UUID]Market, len(markets)),
		DB:      db.NewBlockchain(),
	}
	for _, m := range markets {
		e.Markets[m.Item.ID] = m
	}
	return e
}

func (e *Exchange) Start(ctx context.Context) error {
	wg, c := errgroup.WithContext(ctx)
	for _, m := range e.Markets {
		for _, t := range m.TraderByID {
			_t := t
			wg.Go(func() error { return e.recvRequest(c, _t) })
			wg.Go(func() error { return e.recvResponse(c, _t) })
			wg.Go(func() error { return e.sendResponse(c, _t) })
			wg.Go(func() error { return e.recvChoice(c, _t) })
		}
	}
	return wg.Wait()
}

func (e *Exchange) recvRequest(ctx context.Context, t *trade.Trader) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case r := <-t.RequestSend:
		m, ok := e.Markets[r.Item.ID]
		if !ok {
			return fmt.Errorf("no market found for item: %+v", r.Item)
		}
		for _, t := range m.TraderByID {
			t.RequestRecv <- r
		}
	}
	return nil
}

func (e *Exchange) recvResponse(ctx context.Context, t *trade.Trader) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case resp := <-t.ResponseRecv:
		r := resp[0]
		m, ok := e.Markets[r.Request.Item.ID]
		if !ok {
			return fmt.Errorf("no market found for item: %+v", r.Request.Item.ID)
		}
		for _, t := range m.TraderByID {
			if t.ID == r.Request.TraderID {
				t.ResponseRecv <- resp
			}
		}
	}
	return nil
}

func (e *Exchange) sendResponse(ctx context.Context, t *trade.Trader) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case resp := <-t.ResponseSend:
		m, ok := e.Markets[resp.Request.Item.ID]
		if !ok {
			return fmt.Errorf("no market found for item: %+v", resp.Request.Item.ID)
		}
		r := []trade.Response{resp}
		for _, t := range m.TraderByID {
			if t.ID == resp.Request.TraderID {
				t.ResponseRecv <- r
			}
		}
	}
	return nil
}

func (e *Exchange) recvChoice(ctx context.Context, t *trade.Trader) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case c := <-t.Choice:
		if err := e.execute(c); err != nil {
			return err
		}
	}
	return nil
}

func (e *Exchange) execute(choice trade.Response) error {
	e.dbLock.Lock()
	defer e.dbLock.Unlock()

	t := trade.Transaction{
		ID: uuid.New(),
		Credit: trade.TransactionRecord{
			TraderID: choice.Request.TraderID,
			Item:     choice.Request.Item,
			Price:    choice.OrderBook.Ask.Price,
			Quantity: choice.OrderBook.Ask.Quantity,
		},
		Debit: trade.TransactionRecord{
			TraderID: choice.TraderID,
			Item:     choice.Request.Item,
			Price:    choice.OrderBook.Ask.Price,
			Quantity: choice.OrderBook.Ask.Quantity,
		},
	}
	if ok := e.DB.Append(db.NewBlock(&t)); !ok {
		return fmt.Errorf("failed to persist transaction: %+v", t)
	}
	return nil
}
