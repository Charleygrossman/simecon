package config

import (
	"strings"
	"time"
	"tradesim/src/prob"
	"tradesim/src/time/clock"
	"tradesim/src/trade"
)

func ParseTraders(config []TraderConfig) []*trade.Trader {
	if len(config) == 0 {
		return nil
	}
	traders := make([]*trade.Trader, len(config))
	for i, v := range config {
		traders[i] = parseTrader(v)
	}
	return traders
}

func parseTrader(config TraderConfig) *trade.Trader {
	return trade.NewTrader(config.Haves, config.Wants)
}

func ParseProcess(config ProcessConfig) *prob.Process {
	return prob.NewProcess(
		parseDistribution(config.Distrib),
		parseClock(config.Clock),
	)
}

func parseClock(config ClockConfig) clock.Clock {
	return clock.NewClock(
		time.Second*time.Duration(config.Frequency),
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
