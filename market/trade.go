package market

import (
	"github.com/google/uuid"
)

type Trader struct {
	ID        uuid.UUID
	GraphID   uuid.UUID
	Inventory []Item
	Wants     []Item
}

func (t Trader) Trade(g *Graph) error {
	adjacent, err := g.Adjacent(t.ID, t.GraphID)
	if err != nil {
		return err
	}
	for _, a := range adjacent {
		msg := TradeMessage{
			FromTraderID: t.ID,
			ToTraderID:   a,
			Tradable:     t.Inventory,
			Wants:        t.Wants,
		}
		if err := g.SendTradeMessage(t.ID, t.GraphID, msg); err != nil {
			return err
		}
	}
	return nil
}

type TradeMessage struct {
	FromTraderID uuid.UUID
	ToTraderID   uuid.UUID
	Tradable     []Item
	Wants        []Item
}

// Item represents a thing traded for, including cash and goods.
type Item interface {
	// Value returns the cash quantity of the provided currency
	// of the Item. The boolean return value distinguishes
	// no value from a zero value.
	Value(ccy Currency) (float64, bool)
}

// Cash represents money in the physical
// or electronic form of currency.
type Cash struct {
	ID       uuid.UUID
	Quantity float64
	Currency Currency
}

func (c Cash) Value(ccy Currency) (float64, bool) {
	if ccy != c.Currency {
		return 0.0, false
	}
	return c.Quantity, true
}

// Commodity represents a marketable raw material.
type Commodity struct {
	Quantity float64
	Currency Currency
}

// TODO: Mapping a Commodity value to a currency.
func (c Commodity) Value(ccy Currency) (float64, bool) {
	if ccy != c.Currency {
		return 0.0, false
	}
	return c.Quantity, true
}

// Good represents an item that satisfies a want.
type Good struct {
	Cost map[Currency]float64
	Name string
}

func (g Good) Value(ccy Currency) (float64, bool) {
	cost, ok := g.Cost[ccy]
	return cost, ok
}

// Currency represents a currency code.
type Currency string

const (
	// USD represents the United States Dollar.
	USD Currency = "USD"
	// CNY represents the Chinese Yuan (Renminbi).
	CNY Currency = "CNY"
	// EUR represents the European Euro.
	EUR Currency = "EUR"
	// GBP represents the British Pound Sterling.
	GBP Currency = "GBP"
	// JPY represents the Japanese Yen.
	JPY Currency = "JPY"
)
