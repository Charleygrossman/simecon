// Package transaction provides primitives
// needed to work with financial transactions
package transaction

import (
	"errors"
	"simecon/currency"
	"simecon/utils"
)

// Event represents a transactional event
type Event int

const (
	CREDIT Event = iota + 1
	DEBIT
)

// Transaction represents a financial transaction
type Transaction struct {
	// The credit amount of Transaction.
	Credit float64
	// The Debit amount of transaction.
	Debit float64
	// The currency code of Transaction.
	Currency currency.Code
}

// String returns the string representation of Transaction
func (t *Transaction) String() string {
	return utils.StringStruct(t)
}

// NewTransaction instantiates and returns a new Transaction
// of the given amount, Event and currency Code.
func NewTransaction(amount float64, event Event, currency currency.Code) (*Transaction, error) {
	if amount <= 0 {
		return nil, errors.New("amount must be nonzero positive")
	}
	if event == CREDIT {
		t := &Transaction{
			Credit:   amount,
			Debit:    0.0,
			Currency: currency,
		}
		return t, nil
	} else if event == DEBIT {
		t := &Transaction{
			Credit:   0.0,
			Debit:    amount,
			Currency: currency,
		}
		return t, nil
	} else {
		return nil, errors.New("Invalid Event type")
	}
}
