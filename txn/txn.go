package txn

// TxnType represents a type of transaction.
type TxnType string

const (
	TestTxnType    TxnType = "TestTxnType"
	TradeRequested TxnType = "TradeRequested"
)

// TODO: Define this interface.
type Transaction interface {
	GetHash() string
	GetTxnType() TxnType
}
