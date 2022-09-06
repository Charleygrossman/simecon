package trade

import (
	"tradesim/src/db"

	"github.com/google/uuid"
)

type Exchange struct {
	DB      *db.Blockchain
	Markets map[uuid.UUID]*Market
}

type Market struct {
	Item         Item
	Participants map[uuid.UUID]*Trader
}
