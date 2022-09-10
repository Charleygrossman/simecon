package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"tradesim/src/prob"
	"tradesim/src/trade"
	"tradesim/src/util"

	"gopkg.in/yaml.v3"
)

const (
	minClockFrequency = 1
	minDistribProb    = 0.0
	maxDistribProb    = 1.0
	minDistribMean    = 0.0
	minDistribStdDev  = 0.0
	minDistribLambda  = 0.0
)

var (
	ErrParse      = errors.New("failed to parse simulation configuration from file")
	ErrInvalid    = errors.New("invalid simulation configuration")
	ErrOutOfRange = errors.New("value out of range")
)

type ProcessConfig struct {
	Clock   ClockConfig   `yaml:"clock"`
	Distrib DistribConfig `yaml:"distribution"`
}

type ClockConfig struct {
	// Frequency represents the time between each clock tick in seconds.
	Frequency uint64 `yaml:"frequency"`
	// Limit represents the maximum number of ticks the clock can reach before stopping.
	Limit uint64 `yaml:"limit"`
}

type DistribConfig struct {
	Type   string  `yaml:"type"`
	Prob   float64 `yaml:"probability_measure"`
	Mean   float64 `yaml:"mean"`
	StdDev float64 `yaml:"standard_deviation"`
	Lambda float64 `yaml:"lambda"`
}

type TraderConfig struct {
	Haves []trade.Have `yaml:"haves"`
	Wants []trade.Want `yaml:"wants"`
}

type SimConfig struct {
	Process ProcessConfig  `yaml:"process"`
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
	return validateProcessConfig(config.Process)
}

func validateProcessConfig(config ProcessConfig) error {
	if err := validateClockConfig(config.Clock); err != nil {
		return err
	}
	return validateDistribConfig(config.Distrib)
}

func validateClockConfig(config ClockConfig) error {
	if config.Frequency < minClockFrequency {
		return fmt.Errorf("%w: min=%d got=%d", ErrOutOfRange, minClockFrequency, config.Frequency)
	}
	return nil
}

func validateDistribConfig(config DistribConfig) error {
	distribType := strings.ToLower(strings.TrimSpace(config.Type))
	if !util.ContainsString(prob.DistribTypes, distribType) {
		return prob.NewDistribTypeError(config.Type)
	}
	if config.Prob < minDistribProb || config.Prob > maxDistribProb {
		return fmt.Errorf("%w: name=probability_measure min=%f max=%f got=%f",
			ErrOutOfRange, minDistribProb, maxDistribProb, config.Prob)
	}
	switch distribType {
	case prob.DistribExp:
		if config.Lambda < minDistribLambda {
			return fmt.Errorf("%w: name=lambda min=%f got=%f",
				ErrOutOfRange, minDistribLambda, config.Lambda)
		}
	case prob.DistribNorm:
		if config.StdDev < minDistribStdDev {
			return fmt.Errorf("%w: name=standard_deviation min=%f got=%f",
				ErrOutOfRange, minDistribStdDev, config.StdDev)
		}
	}
	return nil
}
