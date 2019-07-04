// Package currency provides primitives for representing world currencies.
package currency

// Code represents a currency code as an int.
type Code int

const (
	USD Code = iota + 1
	CNY
	EUR
	GBP
	JPY
)
