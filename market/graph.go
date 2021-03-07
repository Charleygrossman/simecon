package market

import (
	"container/list"
	"errors"
	"github.com/google/uuid"
	"math"
	"time"
)

type Graph struct {
	// node maps the trader ID of a node within the graph to the node itself.
	node map[uuid.UUID]*Node
	// adjacent maps the trader ID of a node within the graph to the trader IDs
	// of its adjacent nodes within the graph.
	adjacent map[uuid.UUID][]uuid.UUID
}

func NewGraph(nodes []Node, adjacencies map[uuid.UUID][]uuid.UUID) *Graph {
	g := &Graph{
		adjacent: adjacencies,
	}
	m := make(map[uuid.UUID]*Node, len(nodes))
	for _, n := range nodes {
		m[n.Trader.ID] = &Node{
			Trader:         n.Trader,
			GraphID:        n.GraphID,
			TradeRequests:  list.New(),
			TradeResponses: list.New(),
		}
	}
	g.node = m
	return g
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

// Clock represents a monotonic clock with a frequency and duration.
type Clock struct {
	// Ticker is the underlying time.Ticker of Clock.
	Ticker *time.Ticker
	// Frequency represents the time interval between ticks.
	Frequency time.Duration
	// Limit represents the maximum value Tick can reach;
	// there is no such value if Limit is nil.
	Limit *uint64
	// Tick is a count of the number of ticks Clock has seen.
	Tick uint64
	// Done receives true when the clock has reached its limit,
	// or is stopped or reset.
	Done chan bool
}

func NewClock(frequency time.Duration, limit *uint64) *Clock {
	return &Clock{
		Ticker:    nil,
		Frequency: frequency,
		Limit:     limit,
		Tick:      0,
		Done:      make(chan bool),
	}
}

func (c *Clock) Start() {
	// Creating a new ticker will start it.
	c.Ticker = time.NewTicker(c.Frequency)
	go func() {
		// Run the clock until Limit is reached or Tick
		// reaches system limits (maximum integer value).
		for (c.Limit == nil || c.Tick < *c.Limit) && c.Tick < math.MaxUint64 {
			time.Sleep(c.Frequency)
			c.Tick++
		}
		c.Done <- true
	}()
}

func (c *Clock) Stop() {
	c.Ticker.Stop()
}

func (c *Clock) Reset() {
	c.Ticker.Stop()
	c.Ticker = nil
	c.Tick = 0
	c.Done = make(chan bool)
}

type Node struct {
	Clock          Clock
	GraphID        uuid.UUID
	Trader         *Trader
	TradeRequests  *list.List
	TradeResponses *list.List
}

type Edge struct {
	UTraderID uuid.UUID
	VTraderID uuid.UUID
	Delta     float64
	Closeness float64
}
