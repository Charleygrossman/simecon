package trader

import "tradesim/common"

// trade is a type of transaction
// that involves "from" and "to" traders
// as well as the thing being traded.
type trade struct {
	tradeEntity tradeEntity
	from        uint64
	to          uint64
	txnType     common.TxnType
	createdOn   string
}

func (t trade) getTxnType() common.TxnType {
	return t.txnType
}

func (t trade) getCreatedOn() string {
	return t.createdOn
}

// tradeEntity is a thing traded for, including cash and goods.
type tradeEntity interface {
	// value returns the cash quantity of the provided currency
	// of the tradeEntity. The boolean return value distinguishes
	// no value from a zero value.
	value(ccy common.Ccy) (float64, bool)
}

type cash struct {
	qty float64
	ccy common.Ccy
}

func (c cash) value(ccy common.Ccy) (float64, bool) {
	if ccy != c.ccy {
		return 0.0, false
	}
	return c.qty, true
}

type good struct {
	cost map[common.Ccy]float64
	name string
}

func (g good) value(ccy common.Ccy) (float64, bool) {
	cost, ok := g.cost[ccy]
	return cost, ok
}
