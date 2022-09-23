package exchange

import (
	"context"
	"fmt"
	"tradesim/src/db"
	"tradesim/src/trade"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

type Exchange struct {
	DB      *db.Blockchain
	Markets map[uuid.UUID]Market
}

type Market struct {
	Item       trade.Item
	TraderByID map[uuid.UUID]*trade.Trader
}

func NewExchange(markets []Market) *Exchange {
	e := &Exchange{
		DB:      db.NewBlockchain(),
		Markets: make(map[uuid.UUID]Market, len(markets)),
	}
	for _, m := range markets {
		e.Markets[m.Item.ID] = m
	}
	return e
}

func (e *Exchange) Start(ctx context.Context) error {
	wg, _ := errgroup.WithContext(ctx)
	for _, m := range e.Markets {
		for _, t := range m.TraderByID {
			_t := t
			wg.Go(func() error { return e.recvRequest(_t) })
			wg.Go(func() error { return e.recvResponse(_t) })
			wg.Go(func() error { return e.recvChoice(_t) })
		}
	}
	return wg.Wait()
}

func (e *Exchange) recvRequest(t *trade.Trader) error {
	select {
	case r := <-t.RequestSend:
		m, ok := e.Markets[r.Item.ID]
		if !ok {
			return fmt.Errorf("no market found for item: %+v", r.Item)
		}
		for _, t := range m.TraderByID {
			t.RequestRecv <- r
		}
	default:
	}
	return nil
}

func (e *Exchange) recvResponse(t *trade.Trader) error {
	return nil
}

func (e *Exchange) recvChoice(t *trade.Trader) error {
	select {
	case c := <-t.Choice:
		e.execute(c)
	default:
	}
	return nil
}

func (e *Exchange) execute(choice trade.Response) error {
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
