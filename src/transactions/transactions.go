// A transaction represents a single Credit or Debit from one entity to another.
// For any transaction, exactly one of Credit or Debit must have a nonzero,
// positive value.

package transactions

import (
    "fmt"
    "strings"
    "errors"
    "currency"
)

// An atomic, monetary value represented in a Transaction
type Credit float64
type Debit float64

type Transaction struct {
    Credit Credit
    Debit Debit
}

// Transaction.ToString: Returns a self-representational string
//
// TODO: Reflect on all properties,convert each to a string,
// concatenate them together and return
func (t *Transaction) String() string {
    var rep []string
    c := fmt.Sprintf("Credit: %f ", t.Credit)
    d := fmt.Sprintf("Debit: %f ", t.Debit)
    rep = append(rep, c)
    rep = append(rep, d)
    return fmt.Sprintf(strings.Join(rep, ""))
}

// TODO: amount and t handled differently depending on curr
func (t *Transaction) SetCredit(amt Credit, curr currency.Currency) error {
    if amt > 0 {
        t.Debit = 0
        t.Credit = amt
    } else {
        return errors.New("amount must be nonzero positive")
    }
    // TODO: How to have procedure with optional error return?
    return nil
}

// TODO: amount and t handled differently depending on curr
func (t *Transaction) SetDebit(amt Debit, curr currency.Currency) error {
    if amt > 0 {
        t.Credit = 0
        t.Debit = amt
    } else {
        return errors.New("amount must be nonzero positive")
    }
    // TODO: How to have procedure with optional error return?
    return nil
}
