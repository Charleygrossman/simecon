package config

import (
	"testing"
)

var cfg = SimConfig{
	Items: []ItemConfig{
		{
			ID:   "1",
			Name: "a",
		},
		{
			ID:   "2",
			Name: "b",
		},
	},
	Traders: []TraderConfig{
		{
			ID: "1",
			Haves: []HaveConfig{
				{
					ItemID:   "1",
					Price:    3.14,
					Quantity: 2.718,
				},
			},
			Wants: []WantConfig{
				{
					ItemID:   "2",
					PriceMin: 4.5,
					PriceMax: 4.9,
					Quantity: 8,
				},
			},
		},
		{
			ID: "2",
			Haves: []HaveConfig{
				{
					ItemID:   "1",
					Price:    3.14,
					Quantity: 3,
				},
				{
					ItemID:   "2",
					Price:    5,
					Quantity: 10,
				},
			},
			Wants: []WantConfig{
				{
					ItemID:   "1",
					PriceMin: 3,
					PriceMax: 4,
					Quantity: 10,
				},
				{
					ItemID:   "2",
					PriceMin: 4.2,
					PriceMax: 4.4,
					Quantity: 5,
				},
			},
		},
	},
	Exchange: ExchangeConfig{
		Markets: []MarketConfig{
			{
				ItemID:    "1",
				TraderIDs: []string{"1", "2"},
			},
			{
				ItemID:    "2",
				TraderIDs: []string{"1", "2"},
			},
		},
	},
}

func TestParseItems(t *testing.T) {
	items := ParseItems(cfg.Items)
	if len(items) != 2 {
		t.Errorf("item length: expected: %d actual: %d", 2, len(items))
	}
	names := make(map[string]struct{})
	for _, i := range items {
		names[i.Name] = struct{}{}
	}
	if len(names) != 2 {
		t.Errorf("unique item names: expected: %d actual: %d", 2, len(names))
	}
	if _, ok := names[cfg.Items[0].Name]; !ok {
		t.Errorf("missing item name: %s", cfg.Items[0].Name)
	}
	if _, ok := names[cfg.Items[1].Name]; !ok {
		t.Errorf("missing item name: %s", cfg.Items[1].Name)
	}
}
