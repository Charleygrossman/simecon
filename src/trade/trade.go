package trade

type TradeRequest struct {
	Item     Item
	Quantity float64
	Side     Side
}

type TradeResponse struct {
	OrderBooks []OrderBook
}

type OrderBook struct {
	Asks []OrderBookLevel
	Bids []OrderBookLevel
}

type OrderBookLevel struct {
	Item     Item
	Price    float64
	Quantity float64
}

type Side uint8

const (
	SideUnknown Side = iota
	SideBuy
	SideSell
)

type Transaction struct {
	Credit TransactionRecord
	Debit  TransactionRecord
}

type TransactionRecord struct {
	Trader   *Trader
	Item     Item
	Price    float64
	Quantity float64
}
