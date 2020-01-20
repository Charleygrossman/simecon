package trader

import (
	"tradesim/utils"
)

type Code uint

const (
	USD Code = iota + 1
	CNY
	EUR
	GBP
	JPY
)

type TxnType uint

const (
	TRADE TxnType = iota + 1
)

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
