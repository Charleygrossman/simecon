package mkt

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

// Currency represents an ISO 4217 alpha currency code.
type Currency = string

const (
	// Unknown represents an unknown currency.
	Unknown Currency = ""
	// AUD represents the Australian Dollar.
	AUD Currency = "AUD"
	// CAD represents the Canadian Dollar.
	CAD Currency = "CAD"
	// CNY represents the Chinese Yuan (Renminbi).
	CNY Currency = "CNY"
	// EUR represents the European Euro.
	EUR Currency = "EUR"
	// GBP represents the British Pound Sterling.
	GBP Currency = "GBP"
	// JPY represents the Japanese Yen.
	JPY Currency = "JPY"
	// USD represents the United States Dollar.
	USD Currency = "USD"
)

var Currencies = []Currency{
	AUD,
	CAD,
	CNY,
	EUR,
	GBP,
	JPY,
	USD,
}

type CurrencyError struct {
	Ccy string
}

func NewCurrencyError(ccy string) *CurrencyError {
	return &CurrencyError{Ccy: ccy}
}

func (e CurrencyError) Error() string {
	return fmt.Sprintf("unsupported currency: supported=%s got=%s", strings.Join(Currencies, ", "), e.Ccy)
}

// Instrument represents a tradable thing.
type Instrument interface {
	// Value returns the derived value
	// of the Instrument in the provided Currency.
	Value(ccy Currency) (float64, error)
}

type InstrumentSet struct {
	Cash  map[Currency]Cash
	Goods []Good
}

func (s InstrumentSet) GoodsByID() map[uuid.UUID]Good {
	if len(s.Goods) == 0 {
		return nil
	}
	m := make(map[uuid.UUID]Good, len(s.Goods))
	for _, g := range s.Goods {
		m[g.ID] = g
	}
	return m
}

func (s InstrumentSet) GoodsByName() map[string]Good {
	if len(s.Goods) == 0 {
		return nil
	}
	m := make(map[string]Good, len(s.Goods))
	for _, g := range s.Goods {
		m[g.Name] = g
	}
	return m
}

type BaseInstrument struct {
	ID        uuid.UUID
	Name      string
	Prices    map[Currency]float64
	Quantity  float64
	ValueFunc func(ccy Currency) (float64, error)
}

func (b BaseInstrument) Price(ccy Currency) (float64, error) {
	p, ok := b.Prices[ccy]
	if !ok {
		return 0, NewCurrencyError(ccy)
	}
	return p, nil
}

// Cash represents money in the form of currency.
type Cash struct {
	BaseInstrument
	Currency Currency
}

func (c Cash) Value(ccy Currency) (float64, error) {
	return c.ValueFunc(ccy)
}

// Good represents a tradable thing that satisfies a want.
type Good struct {
	BaseInstrument
}

func (g Good) Value(ccy Currency) (float64, error) {
	return g.ValueFunc(ccy)
}
