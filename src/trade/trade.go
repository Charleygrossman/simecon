package trade

import "github.com/google/uuid"

type Item struct {
	ID   uuid.UUID
	Name string
}

type Side uint8

const (
	SideUnknown Side = iota
	SideBuy
	SideSell
)

type Request struct {
	ID       uuid.UUID
	Item     Item
	Quantity float64
	Side     Side
}

type Responses []Response

type Response struct {
	ID        uuid.UUID
	OrderBook OrderBook
}

type OrderBook struct {
	Ask struct {
		Item     Item
		Price    float64
		Quantity float64
	}
	Bid struct {
		Item     Item
		Price    float64
		Quantity float64
	}
}

type Transaction struct {
	ID     uuid.UUID
	Credit TransactionRecord
	Debit  TransactionRecord
}

type TransactionRecord struct {
	Trader   *Trader
	Item     Item
	Price    float64
	Quantity float64
}
