package trade

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

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
	TraderID uuid.UUID
	Item     Item
	Quantity float64
	Side     Side
}

type Responses []Response

type Response struct {
	ID        uuid.UUID
	Request   Request
	TraderID  uuid.UUID
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

type Choice struct {
	Request  Request
	Response Response
}

type Transaction struct {
	ID     uuid.UUID
	Credit TransactionRecord
	Debit  TransactionRecord
}

type TransactionRecord struct {
	TraderID uuid.UUID
	Item     Item
	Price    float64
	Quantity float64
}

func (t *Transaction) Hash() string {
	var s strings.Builder

	s.WriteString(t.ID.String())

	s.WriteString(t.Credit.TraderID.String())
	s.WriteString(t.Credit.Item.ID.String())
	s.WriteString(t.Credit.Item.Name)
	s.WriteString(strconv.FormatFloat(t.Credit.Price, 'f', -1, 64))
	s.WriteString(strconv.FormatFloat(t.Credit.Quantity, 'f', -1, 64))

	s.WriteString(t.Debit.TraderID.String())
	s.WriteString(t.Debit.Item.ID.String())
	s.WriteString(t.Debit.Item.Name)
	s.WriteString(strconv.FormatFloat(t.Debit.Price, 'f', -1, 64))
	s.WriteString(strconv.FormatFloat(t.Debit.Quantity, 'f', -1, 64))

	return fmt.Sprintf("%x", sha256.Sum256([]byte(s.String())))
}
