package market

import (
	"errors"
	"github.com/google/uuid"
)

type Graph struct {
	// node maps the trader ID of a node within the graph to the node itself.
	node map[uuid.UUID]*node
	// adjacent maps the trader ID of a node within the graph to the trader IDs
	// of its adjacent nodes within the graph.
	adjacent map[uuid.UUID][]uuid.UUID
}

func (g Graph) Adjacent(graphID, traderID uuid.UUID) ([]uuid.UUID, error) {
	if err := g.authorizeTrader(graphID, traderID); err != nil {
		return nil, err
	}
	return g.adjacent[traderID], nil
}

func (g Graph) SendTradeMessage(graphID uuid.UUID, message TradeMessage) error {
	if err := g.authorizeTradeMessage(graphID, message); err != nil {
		return err
	}
	n, ok := g.node[message.ToTraderID]
	if !ok || n == nil || n.trader.ID != message.ToTraderID {
		return errors.New("receiving trader not found")
	}
	n.tradeRequests <- message
	return nil
}

func (g Graph) authorizeTrader(graphID, traderID uuid.UUID) error {
	n, ok := g.node[traderID]
	if !ok || n == nil || n.graphID != graphID || n.trader.ID != traderID {
		return errors.New("unauthorized trader")
	}
	return nil
}

func (g Graph) authorizeTradeMessage(graphID uuid.UUID, message TradeMessage) error {
	if err := g.authorizeTrader(graphID, message.FromTraderID); err != nil {
		return err
	}
	// TODO: Binary search.
	for _, adjTraderID := range g.adjacent[message.FromTraderID] {
		if adjTraderID == message.ToTraderID && g.node[adjTraderID].trader.ID == message.ToTraderID {
			return nil
		}
	}
	return errors.New("unauthorized trade message")
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
