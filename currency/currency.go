package currency

type Code int

const (
	USD Code = iota + 1
	CNY
	EUR
	GBP
	JPY
)
