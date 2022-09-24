package config

import (
	"strings"
	"time"
	"tradesim/src/exchange"
	"tradesim/src/prob"
	"tradesim/src/time/clock"
	"tradesim/src/trade"
)

func ParseExchange(config ExchangeConfig, items map[string]trade.Item, traders map[string]*trade.Trader) *exchange.Exchange {
	markets := make([]exchange.Market, 0, len(config.Markets))
	for _, c := range config.Markets {
		i, ok := items[c.ItemID]
		if !ok {
			continue
		}
		ts := make([]*trade.Trader, 0, len(c.TraderIDs))
		for _, id := range c.TraderIDs {
			t, ok := traders[id]
			if ok {
				ts = append(ts, t)
			}
		}
		m := exchange.NewMarket(i, ts...)
		markets = append(markets, m)
	}
	return exchange.NewExchange(markets)
}

func ParseItems(config []ItemConfig) map[string]trade.Item {
	result := make(map[string]trade.Item, len(config))
	for _, v := range config {
		result[v.ID] = trade.NewItem(v.Name)
	}
	return result
}

func ParseTraders(config []TraderConfig, items map[string]trade.Item) map[string]*trade.Trader {
	result := make(map[string]*trade.Trader, len(config))
	for _, v := range config {
		result[v.ID] = parseTrader(v, items)
	}
	return result
}

func parseTrader(config TraderConfig, items map[string]trade.Item) *trade.Trader {
	haves := make([]trade.Have, 0, len(config.Haves))
	for _, c := range config.Haves {
		i, ok := items[c.ItemID]
		if ok {
			h := parseHave(c, i)
			haves = append(haves, h)
		}
	}
	wants := make([]trade.Want, 0, len(config.Wants))
	for _, c := range config.Wants {
		i, ok := items[c.ItemID]
		if ok {
			w := parseWant(c, i)
			wants = append(wants, w)
		}
	}
	return trade.NewTrader(haves, wants)
}

func parseHave(config HaveConfig, item trade.Item) trade.Have {
	return trade.Have{
		Item:     item,
		Price:    config.Price,
		Quantity: config.Quantity,
	}
}

func parseWant(config WantConfig, item trade.Item) trade.Want {
	return trade.Want{
		Item:     item,
		PriceMin: config.PriceMin,
		PriceMax: config.PriceMax,
		Quantity: config.Quantity,
	}
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
