package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Target string
}

func newConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	config := &Config{}
	if err = json.NewDecoder(file).Decode(config); err != nil {
		return nil, err
	}
	return config, nil
}
