package market

import "github.com/google/uuid"

type Trader struct {
	ID        uuid.UUID
	GraphID   uuid.UUID
	Inventory []Item
	Wants     []Item
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
			Tradable:     t.Inventory,
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
				Tradable:     t.Inventory,
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

func (t Trader) hasRequestWant(request TradeMessage) bool {
	for _, item := range t.Inventory {
		for _, v := range request.Wants {
			if item == v {
				return true
			}
		}
	}
	return false
}

func (t Trader) wantsRequestTradable(request TradeMessage) bool {
	for _, item := range t.Wants {
		for _, v := range request.Tradable {
			if item == v {
				return true
			}
		}
	}
	return false
}
