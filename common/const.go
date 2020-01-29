package common

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

// TxnType represents a type of transaction.
type TxnType string

const (
	TradeRequested TxnType = "TradeRequested"
)
