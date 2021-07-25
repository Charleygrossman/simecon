package mkt

import (
	"github.com/google/uuid"
)

type TradeMessage struct {
	FromTraderID uuid.UUID
	ToTraderID   uuid.UUID
	Available    InstrumentSet
	Wants        InstrumentSet
}

type Trader struct {
	ID        uuid.UUID
	GraphID   uuid.UUID
	Inventory InstrumentSet
	Wants     InstrumentSet
}

func NewTrader(inventory, wants InstrumentSet) *Trader {
	return &Trader{
		ID:        uuid.New(),
		GraphID:   uuid.Nil,
		Inventory: inventory,
		Wants:     wants,
	}
}

func (t Trader) SendTradeRequests(g *Graph) error {
	adjacent, err := g.Adjacent(t.GraphID, t.ID)
	if err != nil {
		return err
	}
	for _, adjTraderID := range adjacent {
		msg := TradeMessage{
			FromTraderID: t.ID,
			ToTraderID:   adjTraderID,
			Available:    t.Inventory,
			Wants:        t.Wants,
		}
		if err := g.SendTradeRequest(t.GraphID, msg); err != nil {
			return err
		}
	}
	return nil
}

func (t Trader) SendTradeResponses(g *Graph) error {
	requests, err := g.TradeRequests(t.GraphID, t.ID)
	if err != nil {
		return err
	}
	for _, r := range requests {
		if t.acceptRequest(r) {
			msg := TradeMessage{
				FromTraderID: t.ID,
				ToTraderID:   r.FromTraderID,
				Available:    t.Inventory,
				Wants:        t.Wants,
			}
			if err := g.SendTradeResponse(t.GraphID, msg); err != nil {
				return err
			}
		}
	}
	return nil
}

func (t Trader) acceptRequest(request TradeMessage) bool {
	return t.hasRequestWant(request) && t.wantsRequestTradable(request)
}

// TODO
func (t Trader) hasRequestWant(request TradeMessage) bool {
	return false
}

// TODO
func (t Trader) wantsRequestTradable(request TradeMessage) bool {
	return false
}
