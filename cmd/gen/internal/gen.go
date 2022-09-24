package internal

import (
	"errors"
	"fmt"
	"io/ioutil"

	"tradesim/src/sim/config"

	"gopkg.in/yaml.v3"
)

var (
	ErrGen = errors.New("failed to generate simulation configuration file")

	template = config.SimConfig{
		Items: []config.ItemConfig{{}, {}},
		Traders: []config.TraderConfig{
			{
				Haves: []config.HaveConfig{{}, {}},
				Wants: []config.WantConfig{{}, {}},
			},
			{
				Haves: []config.HaveConfig{{}, {}},
				Wants: []config.WantConfig{{}, {}},
			},
		},
		Exchange: config.ExchangeConfig{
			Markets: []config.MarketConfig{{}, {}},
		},
	}
)

func Generate(filepath string) error {
	content, err := yaml.Marshal(template)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath, content, 0644); err != nil {
		return fmt.Errorf("%w: %v", ErrGen, err)
	}
	return nil
}
