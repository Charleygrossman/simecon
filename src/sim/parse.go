package sim

import (
	"strings"
	libTime "time"
	"tradesim/src/mkt"
	"tradesim/src/prob"
	"tradesim/src/time"

	"github.com/google/uuid"
)

func ParseProcess(config ProcessConfig) prob.Process {
	return prob.NewProcess(
		parseDistribution(config.Distrib),
		parseClock(config.Clock),
	)
}

func parseClock(config ClockConfig) time.Clock {
	return time.NewClock(
		libTime.Second*libTime.Duration(config.Frequency),
		config.Limit,
	)
}

func parseDistribution(config DistribConfig) prob.Distribution {
	switch strings.ToLower(strings.TrimSpace(config.Type)) {
	case prob.DistribExp:
		return prob.NewExponential(config.Prob, config.Lambda)
	case prob.DistribNorm:
		return prob.NewNormal(config.Prob, config.Mean, config.StdDev)
	case prob.DistribUni:
		return prob.NewUniform(config.Prob)
	default:
		return nil
	}
}

func ParseTraders(config []TraderConfig) []*mkt.Trader {
	if len(config) == 0 {
		return nil
	}
	traders := make([]*mkt.Trader, len(config))
	for i, v := range config {
		traders[i] = ParseTrader(v)
	}
	return traders
}

func ParseTrader(config TraderConfig) *mkt.Trader {
	return mkt.NewTrader(
		parseInstrumentSet(config.Inventory),
		parseInstrumentSet(config.Wants),
	)
}

func parseInstrumentSet(c InstrumentSetConfig) mkt.InstrumentSet {
	cash := make(map[mkt.Currency]mkt.Cash, len(c.Cash))
	for _, v := range c.Cash {
		cash[v.Currency] = parseCash(v)
	}
	goods := make([]mkt.Good, len(c.Goods))
	for i, v := range c.Goods {
		goods[i] = parseGood(v)
	}
	return mkt.InstrumentSet{
		Cash:  cash,
		Goods: goods,
	}
}

func parseCash(c CashConfig) mkt.Cash {
	return mkt.Cash{
		BaseInstrument: mkt.BaseInstrument{
			ID:       uuid.New(),
			Prices:   make(map[mkt.Currency]float64),
			Quantity: c.Quantity,
		},
		Currency: c.Currency,
	}
}

func parseGood(c GoodConfig) mkt.Good {
	return mkt.Good{
		BaseInstrument: mkt.BaseInstrument{
			ID:       uuid.New(),
			Name:     c.Name,
			Prices:   make(map[mkt.Currency]float64),
			Quantity: c.Quantity,
		},
	}
}
