package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Inputs  uint
	Outputs uint
	Command string
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
	return config, nil
}

func defaultConfig() *Config {
	return &Config{
		Inputs:  1,
		Outputs: 1,
	}
}
