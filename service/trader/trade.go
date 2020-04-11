package trader

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"tradesim/transaction"
)

// ccy represents a currency code.
type Ccy string

const (
	// United States Dollar
	USD Ccy = "USD"
	// Renminbi (Chinese Yuan)
	CNY Ccy = "CNY"
	// Euro
	EUR Ccy = "EUR"
	// Pound Sterling
	GBP Ccy = "GBP"
	// Japanese Yen
	JPY Ccy = "JPY"
)

// trade is a type of transaction
// that involves "from" and "to" traders
// as well as the thing being traded.
type trade struct {
	tradeEntity tradeEntity
	from        uint64
	to          uint64
	txnType     transaction.TxnType
	createdOn   string
}

func (t trade) GetTxnType() transaction.TxnType {
	return t.txnType
}

func (t trade) GetHash() string {
	data := strconv.FormatUint(t.from, 10) + strconv.FormatUint(t.to, 10) + string(t.txnType)
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
	value(ccy Ccy) (float64, bool)
}

type cash struct {
	qty float64
	ccy Ccy
}

func (c cash) value(ccy Ccy) (float64, bool) {
	if ccy != c.ccy {
		return 0.0, false
	}
	return c.qty, true
}

type good struct {
	cost map[Ccy]float64
	name string
}

func (g good) value(ccy Ccy) (float64, bool) {
	cost, ok := g.cost[ccy]
	return cost, ok
}
