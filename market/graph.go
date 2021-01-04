package market

import (
	"errors"
	"github.com/google/uuid"
)

type Graph struct {
	nodes    map[uuid.UUID]*node
	adjacent map[uuid.UUID][]uuid.UUID
}

func (g Graph) Adjacent(traderID, graphID uuid.UUID) ([]uuid.UUID, error) {
	if err := g.authorizeTrader(traderID, graphID); err != nil {
		return nil, err
	}
	return g.adjacent[traderID], nil
}

func (g Graph) SendTradeMessage(traderID, graphID uuid.UUID, message TradeMessage) error {
	if err := g.authorizeTrader(traderID, graphID); err != nil {
		return err
	}
	if err := g.authorizeTradeMessage(traderID, message); err != nil {
		return err
	}
	n, ok := g.nodes[message.ToTraderID]
	if !ok || n == nil {
		return errors.New("receiving trader not found")
	}
	n.tradeRequests <- message
	return nil
}

func (g Graph) authorizeTrader(traderID, graphID uuid.UUID) error {
	n, ok := g.nodes[traderID]
	if !ok || n == nil || n.graphID != graphID {
		return errors.New("unauthorized trader")
	}
	return nil
}

func (g Graph) authorizeTradeMessage(traderID uuid.UUID, message TradeMessage) error {
	err := errors.New("unauthorized trade message")
	if message.FromTraderID != traderID {
		return err
	}
	// TODO: Binary search.
	for _, a := range g.adjacent[traderID] {
		if a == message.ToTraderID {
			return nil
		}
	}
	return err
}

type node struct {
	trader         *Trader
	graphID        uuid.UUID
	tradeRequests  chan TradeMessage
	tradeResponses chan TradeMessage
}

type edge struct {
	uTraderID uuid.UUID
	vTraderID uuid.UUID
	delta     float64
	closeness float64
}
