// For now a transaction is a trade between two traders.
// Transaction is a concept useful to the Accountant and database.

package trader

import (
	"tradesim/utils"
)

type TxnType uint

const (
	TRADE TxnType = iota + 1
)

// TradeEvent is an executed trade between two traders.
// The trade was matched by a Broker and accepted by both traders,
// therefore it's executed and qualifies as a transaction event.
type TradeTxn struct {
	CreatedOn string
	TxnType   TxnType
	TraderA   Trader
	TraderB   Trader
	// The subset of TraderA's inventory it's trading to B.
	ATrades Inventory
	// The subset of TraderB's inventory it's trading to A.
	BTrades Inventory
}

func (t TradeTxn) String() string {
	return utils.StringStruct(t)
}

func NewTradeTxn(a Trader, b Trader, ai Inventory, bi Inventory) *TradeTxn {
	t := &TradeTxn{
		TxnType: TRADE,
		TraderA: a,
		TraderB: b,
		ATrades: ai,
		BTrades: bi,
	}
	t.CreatedOn = utils.Now()
	return t
}
