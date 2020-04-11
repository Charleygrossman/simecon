package transaction

type Transaction interface {
	GetHash() string
	GetTxnType() TxnType
}

// TxnType represents a type of transaction.
type TxnType string

const (
	TradeRequested TxnType = "TradeRequested"
)
