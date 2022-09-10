package trade

import (
	"tradesim/src/db"

	"github.com/google/uuid"
)

type Exchange struct {
	DB      *db.Blockchain
	Markets map[uuid.UUID]Market
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

type Market struct {
	Item         Item
	Participants map[uuid.UUID]*Trader
}
