package trader

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"tradesim/txn"
)

// Ccy represents a currency code.
type Ccy string

const (
	// USD represents the United States Dollar.
	USD Ccy = "USD"
	// CNY represents the Chinese Yuan (Renminbi).
	CNY Ccy = "CNY"
	// EUR represents the European Euro.
	EUR Ccy = "EUR"
	// GBP represents the British Pound Sterling.
	GBP Ccy = "GBP"
	// JPY represents the Japanese Yen.
	JPY Ccy = "JPY"
)

// trade represents a type of transaction
// that involves "from" and "to" traders,
// as well as the thing being traded.
type trade struct {
	tradeEntity tradeEntity
	from        uint64
	to          uint64
	txnType     txn.TxnType
	createdOn   string
}

func (t trade) GetTxnType() txn.TxnType {
	return t.txnType
}

func (t trade) GetHash() string {
	data := strconv.FormatUint(t.from, 10) + strconv.FormatUint(t.to, 10) + string(t.txnType)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
}

// tradeEntity represents a thing traded for, including cash and goods.
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
