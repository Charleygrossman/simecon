package main

import (
	"github.com/google/uuid"
	"log"
	"tradesim/market"
)

func main() {
	a := &market.Trader{
		ID:      uuid.New(),
		GraphID: uuid.New(),
		Inventory: []market.Item{
			market.Good{
				Cost: map[market.Currency]float64{market.USD: 10},
				Name: "Foo",
			}},
		Wants: []market.Item{
			market.Good{
				Cost: map[market.Currency]float64{market.USD: 10},
				Name: "Bar",
			}},
	}

	b := &market.Trader{
		ID:      uuid.New(),
		GraphID: uuid.New(),
		Inventory: []market.Item{
			market.Good{
				Cost: map[market.Currency]float64{market.USD: 10},
				Name: "Bar",
			}},
		Wants: []market.Item{
			market.Good{
				Cost: map[market.Currency]float64{market.USD: 10},
				Name: "Foo",
			}},
	}

	graph := market.NewGraph(
		[]market.Node{
			{
				Trader:  a,
				GraphID: a.GraphID,
			},
			{
				Trader:  b,
				GraphID: b.GraphID,
			},
		},
		map[uuid.UUID][]uuid.UUID{
			a.ID: {b.ID},
			b.ID: {a.ID},
		},
	)

	if err := a.SendTradeMessage(graph, b.ID); err != nil {
		log.Fatal(err)
	}
}
