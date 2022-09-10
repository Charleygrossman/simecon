package internal

import (
	"errors"
	"fmt"
	"io/ioutil"

	"tradesim/src/sim/config"
	"tradesim/src/trade"

	"gopkg.in/yaml.v3"
)

var (
	ErrGen = errors.New("failed to generate simulation configuration file")

	template = config.SimConfig{
		Process: config.ProcessConfig{
			Clock: config.ClockConfig{
				Frequency: 1,
			},
			Distrib: config.DistribConfig{},
		},
		Traders: []config.TraderConfig{
			{
				Haves: []trade.Have{},
				Wants: []trade.Want{},
			},
		},
	}
)

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
