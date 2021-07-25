package internal

import (
	"errors"
	"fmt"
	"io/ioutil"

	"tradesim/src/sim"

	"gopkg.in/yaml.v3"
)

var (
	ErrGen   = errors.New("failed to generate simulation configuration file")
	template = sim.SimConfig{
		Traders: []sim.TraderConfig{
			{
				Inventory: sim.InstrumentSetConfig{
					Cash:  []sim.CashConfig{{}},
					Goods: []sim.GoodConfig{{}},
				},
				Wants: sim.InstrumentSetConfig{
					Cash:  []sim.CashConfig{{}},
					Goods: []sim.GoodConfig{{}},
				},
			},
		},
	}
)

// TODO: generate with comments
func Generate(filepath string) error {
	if err := generate(filepath); err != nil {
		return fmt.Errorf("%w: %v", ErrGen, err)
	}
	return nil
}

func generate(filepath string) error {
	content, err := yaml.Marshal(template)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath, content, 0644)
}
