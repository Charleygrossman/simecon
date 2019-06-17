// Currency type with enumerated codes

package currency

type Currency int

const (
    USD Currency = iota
    CNY
    EUR
    GBP
    JPY
)
