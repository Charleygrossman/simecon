package market

import (
	"container/list"
	"errors"
	"github.com/google/uuid"
)

type Node struct {
	GraphID        uuid.UUID
	Trader         *Trader
	TradeRequests  *list.List
	TradeResponses *list.List
	Clock          Clock
}

type Edge struct {
	UTraderID uuid.UUID
	VTraderID uuid.UUID
	Delta     float64
}

type Graph struct {
	// node maps the trader ID of a node within the graph to the node itself.
	node map[uuid.UUID]*Node
	// edge maps the trader IDs of two nodes within the graph to their edge.
	// Every edge corresponds to an adjacency.
	edge map[struct{ uTraderID, vTraderID uuid.UUID }]*Edge
	// adjacent maps the trader ID of a node within the graph to the trader IDs
	// of its adjacent nodes within the graph. Every adjacency corresponds to an edge.
	adjacent map[uuid.UUID][]uuid.UUID
}

func NewGraph(nodes []Node, edges []Edge) *Graph {
	node := make(map[uuid.UUID]*Node)
	for _, n := range nodes {
		if n.Trader == nil {
			continue
		}
		node[n.Trader.ID] = &Node{
			Trader:         n.Trader,
			GraphID:        n.GraphID,
			TradeRequests:  list.New(),
			TradeResponses: list.New(),
		}
	}

	edge := make(map[struct{ uTraderID, vTraderID uuid.UUID }]*Edge)
	adjacent := make(map[uuid.UUID][]uuid.UUID)
	for _, e := range edges {
		if _, ok := node[e.UTraderID]; !ok {
			continue
		}
		if _, ok := node[e.VTraderID]; !ok {
			continue
		}

		edge[struct {
			uTraderID, vTraderID uuid.UUID
		}{e.UTraderID, e.VTraderID}] = &Edge{
			UTraderID: e.UTraderID,
			VTraderID: e.VTraderID,
			Delta:     e.Delta,
		}

		adjacent[e.UTraderID] = append(adjacent[e.UTraderID], e.VTraderID)
		adjacent[e.VTraderID] = append(adjacent[e.VTraderID], e.UTraderID)
	}

	return &Graph{
		node:     node,
		edge:     edge,
		adjacent: adjacent,
	}
}

func (g Graph) Adjacent(graphID, traderID uuid.UUID) ([]uuid.UUID, error) {
	if err := g.authorizeTrader(graphID, traderID); err != nil {
		return nil, err
	}
	return g.adjacent[traderID], nil
}

func (g Graph) SendTradeRequest(graphID uuid.UUID, message TradeMessage) error {
	if err := g.authorizeTradeMessage(graphID, message); err != nil {
		return err
	}
	n, ok := g.node[message.ToTraderID]
	if !ok || n == nil || n.Trader.ID != message.ToTraderID {
		return errors.New("receiving trader not found")
	}
	n.TradeRequests.PushBack(message)
	return nil
}

func (g Graph) SendTradeResponse(graphID uuid.UUID, message TradeMessage) error {
	if err := g.authorizeTradeMessage(graphID, message); err != nil {
		return err
	}
	n, ok := g.node[message.ToTraderID]
	if !ok || n == nil || n.Trader.ID != message.ToTraderID {
		return errors.New("receiving trader not found")
	}
	n.TradeResponses.PushBack(message)
	return nil
}

func (g Graph) TradeRequests(graphID uuid.UUID, traderID uuid.UUID) ([]TradeMessage, error) {
	if err := g.authorizeTrader(graphID, traderID); err != nil {
		return nil, err
	}
	n, ok := g.node[traderID]
	if !ok || n == nil || n.Trader.ID != traderID {
		return nil, errors.New("trader not found")
	}
	requests := make([]TradeMessage, 0, n.TradeRequests.Len())
	for e := n.TradeRequests.Front(); e != nil; e = e.Next() {
		msg, ok := e.Value.(TradeMessage)
		if !ok {
			return nil, errors.New("invalid trade request type")
		}
		requests = append(requests, msg)
	}
	return requests, nil
}

func (g Graph) TradeResponses(graphID uuid.UUID, traderID uuid.UUID) ([]TradeMessage, error) {
	if err := g.authorizeTrader(graphID, traderID); err != nil {
		return nil, err
	}
	n, ok := g.node[traderID]
	if !ok || n == nil || n.Trader.ID != traderID {
		return nil, errors.New("trader not found")
	}
	responses := make([]TradeMessage, 0, n.TradeResponses.Len())
	for e := n.TradeResponses.Front(); e != nil; e = e.Next() {
		msg, ok := e.Value.(TradeMessage)
		if !ok {
			return nil, errors.New("invalid trade response type")
		}
		responses = append(responses, msg)
	}
	return responses, nil
}

func (g Graph) authorizeTrader(graphID, traderID uuid.UUID) error {
	n, ok := g.node[traderID]
	if !ok || n == nil || n.GraphID != graphID || n.Trader.ID != traderID {
		return errors.New("unauthorized trader")
	}
	return nil
}

func (g Graph) authorizeTradeMessage(graphID uuid.UUID, message TradeMessage) error {
	if err := g.authorizeTrader(graphID, message.FromTraderID); err != nil {
		return err
	}
	for _, adjTraderID := range g.adjacent[message.FromTraderID] {
		if adjTraderID == message.ToTraderID &&
			g.node[adjTraderID].Trader.ID == message.ToTraderID {
			return nil
		}
	}
	return errors.New("unauthorized trade message")
}
