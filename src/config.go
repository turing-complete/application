package main

import (
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	Inputs  uint
	Outputs uint

	Command string

	MinLevel uint
	MaxLevel uint

	AbsoluteError float64
	RelativeError float64
	ScoreError    float64
}

func newConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	config := defaultConfig()
	if err = json.NewDecoder(file).Decode(config); err != nil {
		return nil, err
	}
	if config.Inputs == 0 {
		return nil, errors.New("expected a positive number of inputs")
	}
	if config.Outputs == 0 {
		return nil, errors.New("expected a positive number of outputs")
	}
	if len(config.Command) == 0 {
		return nil, errors.New("expected a command name")
	}
	return config, nil
}

func defaultConfig() *Config {
	return &Config{
		Inputs:  1,
		Outputs: 1,

		MinLevel: 1,
		MaxLevel: 10,

		AbsoluteError: 1e-3,
		RelativeError: 1e-2,
		ScoreError:    1e-4,
	}
}
