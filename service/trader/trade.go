package trader

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"tradesim/common"
)

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

// TODO: Full hash.
func (t trade) getHash() string {
	data := strconv.FormatUint(t.from, 10) + strconv.FormatUint(t.to, 10) + string(t.txnType) + t.createdOn
	return fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
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
