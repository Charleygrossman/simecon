package sim

import (
	"errors"
	"fmt"
	"io/ioutil"
	"tradesim/src/mkt"
	"tradesim/src/util"

	"gopkg.in/yaml.v3"
)

var (
	ErrParse   = errors.New("failed to parse simulation configuration from file")
	ErrInvalid = errors.New("invalid simulation configuration")
)

type ClockConfig struct {
	// Frequency represents the time between each clock tick in seconds.
	Frequency uint64 `yaml:"frequency"`
	// Limit represents the maximum number of ticks the clock can reach before stopping.
	Limit uint64 `yaml:"limit"`
}

type TraderConfig struct {
	Inventory InstrumentSetConfig `yaml:"inventory"`
	Wants     InstrumentSetConfig `yaml:"wants"`
}

type InstrumentSetConfig struct {
	Cash  []CashConfig `yaml:"cash"`
	Goods []GoodConfig `yaml:"goods"`
}

type CashConfig struct {
	Currency string  `yaml:"currency"`
	Quantity float64 `yaml:"quantity"`
}

type GoodConfig struct {
	Name     string  `yaml:"name"`
	Quantity float64 `yaml:"quantity"`
}

type SimConfig struct {
	Clock   ClockConfig    `yaml:"clock"`
	Traders []TraderConfig `yaml:"traders"`
}

func NewSimConfig(filepath string) (SimConfig, error) {
	config, err := parseSimConfig(filepath)
	if err != nil {
		return SimConfig{}, fmt.Errorf("%w: %v", ErrParse, err)
	}
	if err := validateSimConfig(config); err != nil {
		return SimConfig{}, fmt.Errorf("%w: %v", ErrInvalid, err)
	}
	return config, nil
}

func parseSimConfig(filepath string) (SimConfig, error) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return SimConfig{}, err
	}

	var config SimConfig
	if err := yaml.Unmarshal(content, &config); err != nil {
		return SimConfig{}, err
	}
	return config, nil
}

func validateSimConfig(config SimConfig) error {
	for _, t := range config.Traders {
		for _, c := range t.Inventory.Cash {
			if !util.ContainsString(mkt.Currencies, c.Currency) {
				return mkt.NewCurrencyError(c.Currency)
			}
		}
	}
	return nil
}
