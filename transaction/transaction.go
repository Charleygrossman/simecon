package transaction

import (
	"errors"
	"simecon/currency"
	"simecon/utils"
)

type Event int

const (
	CREDIT Event = iota + 1
	DEBIT
)

type Transaction struct {
	Credit   float64
	Debit    float64
	Currency currency.Code
}

func (t *Transaction) String() string {
	return utils.StringStruct(t)
}

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
