package txn

type Transaction interface {
	GetHash() string
	GetTxnType() TxnType
}

// TxnType represents a type of txn.
type TxnType string

const (
	TradeRequested TxnType = "TradeRequested"
)
