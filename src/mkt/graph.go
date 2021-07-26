package mkt

import (
	"container/list"
	"context"
	"errors"
	"fmt"
	"tradesim/src/time"

	"github.com/google/uuid"
)

var (
	ErrNotFound     = errors.New("trader not found")
	ErrNotFoundSend = errors.New("sending trader not found")
	ErrNotFoundRecv = errors.New("receiving trader not found")
)

type node struct {
	trader         *Trader
	tradeRequests  *list.List
	tradeResponses *list.List
}

func newNode(trader *Trader) *node {
	return &node{
		trader:         trader,
		tradeRequests:  list.New(),
		tradeResponses: list.New(),
	}
}

type edgeID struct {
	uTraderID uuid.UUID
	vTraderID uuid.UUID
}

type Edge struct {
	UTraderID uuid.UUID
	VTraderID uuid.UUID
	Delta     float64
}

type Graph struct {
	// nodeByID maps the trader ID of a node within the graph to the node itself.
	nodeByID map[uuid.UUID]*node
	// edgeByID maps the trader IDs of two nodes within the graph to their edge.
	// Every edge corresponds to an adjacency.
	edgeByID map[edgeID]Edge
	// adjacentByID maps the trader ID of a node within the graph to the trader IDs
	// of its adjacent nodes within the graph. Every adjacency corresponds to an edge.
	adjacentByID map[uuid.UUID][]uuid.UUID
	// clock drives the global timing and events of the graph.
	clock time.Clock
}

func NewGraph(traders []*Trader, edges []Edge, clock time.Clock) *Graph {
	if len(traders) == 0 {
		return nil
	}

	nodeByID := make(map[uuid.UUID]*node)
	for _, v := range traders {
		if v == nil || v.ID == uuid.Nil {
			continue
		}
		nodeByID[v.ID] = newNode(v)
	}

	edgeByID := make(map[edgeID]Edge)
	adjacentByID := make(map[uuid.UUID][]uuid.UUID)
	for _, v := range edges {
		if _, ok := nodeByID[v.UTraderID]; !ok {
			continue
		}
		if _, ok := nodeByID[v.VTraderID]; !ok {
			continue
		}
		edgeByID[edgeID{v.UTraderID, v.VTraderID}] = v
		adjacentByID[v.UTraderID] = append(adjacentByID[v.UTraderID], v.VTraderID)
		adjacentByID[v.VTraderID] = append(adjacentByID[v.VTraderID], v.UTraderID)
	}

	return &Graph{
		nodeByID:     nodeByID,
		edgeByID:     edgeByID,
		adjacentByID: adjacentByID,
		clock:        clock,
	}
}

func (g Graph) Run(ctx context.Context) error {
	go g.clock.Start()
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-g.clock.Done:
			return nil
		// TODO
		case <-g.clock.Tick:
		}
	}
}

func (g Graph) Adjacent(traderID uuid.UUID) ([]uuid.UUID, error) {
	v, ok := g.adjacentByID[traderID]
	if !ok {
		return nil, fmt.Errorf("%w: id=%s", ErrNotFound, traderID)
	}
	return v, nil
}

func (g Graph) SendTradeRequest(message TradeMessage) error {
	if _, ok := g.nodeByID[message.FromTraderID]; !ok {
		return fmt.Errorf("%w: id=%s", ErrNotFoundSend, message.FromTraderID)
	}
	n, ok := g.nodeByID[message.ToTraderID]
	if !ok {
		return fmt.Errorf("%w: id=%s", ErrNotFoundRecv, message.ToTraderID)
	}
	n.tradeRequests.PushBack(message)
	return nil
}

func (g Graph) SendTradeResponse(message TradeMessage) error {
	if _, ok := g.nodeByID[message.FromTraderID]; !ok {
		return fmt.Errorf("%w: id=%s", ErrNotFoundSend, message.FromTraderID)
	}
	n, ok := g.nodeByID[message.ToTraderID]
	if !ok {
		return fmt.Errorf("%w: id=%s", ErrNotFoundRecv, message.ToTraderID)
	}
	n.tradeResponses.PushBack(message)
	return nil
}

func (g Graph) TradeRequests(traderID uuid.UUID) ([]TradeMessage, error) {
	n, ok := g.nodeByID[traderID]
	if !ok {
		return nil, fmt.Errorf("%w: id=%s", ErrNotFound, traderID)
	}
	requests := make([]TradeMessage, 0, n.tradeRequests.Len())
	for e := n.tradeRequests.Front(); e != nil; e = e.Next() {
		msg := e.Value.(TradeMessage)
		requests = append(requests, msg)
	}
	return requests, nil
}

func (g Graph) TradeResponses(traderID uuid.UUID) ([]TradeMessage, error) {
	n, ok := g.nodeByID[traderID]
	if !ok {
		return nil, fmt.Errorf("%w: id=%s", ErrNotFound, traderID)
	}
	responses := make([]TradeMessage, 0, n.tradeResponses.Len())
	for e := n.tradeResponses.Front(); e != nil; e = e.Next() {
		msg := e.Value.(TradeMessage)
		responses = append(responses, msg)
	}
	return responses, nil
}
