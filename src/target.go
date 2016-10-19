package main

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type Target struct {
	ni   uint
	no   uint
	name string
}

func newTarget(config *Config) (*Target, error) {
	if config.Inputs == 0 {
		return nil, errors.New("expected a positive number of inputs")
	}
	if config.Outputs == 0 {
		return nil, errors.New("expected a positive number of outputs")
	}
	if len(config.Command) == 0 {
		return nil, errors.New("expected a command")
	}
	target := &Target{
		ni:   config.Inputs,
		no:   config.Outputs,
		name: config.Command,
	}
	return target, nil
}

func (self *Target) evaluate(z []float64) ([]float64, error) {
	inputs := make([]string, self.ni)
	for i := uint(0); i < self.ni; i++ {
		inputs[i] = fmt.Sprintf("%.15e", z[i])
	}
	bytes, err := exec.Command(self.name, inputs...).Output()
	if err != nil {
		return nil, err
	}
	outputs := strings.Split(strings.TrimSpace(string(bytes)), " ")
	if uint(len(outputs)) != self.no {
		return nil, errors.New("expected a different number of outputs")
	}
	u := make([]float64, self.no)
	for i := uint(0); i < self.no; i++ {
		u[i], err = strconv.ParseFloat(outputs[i], 64)
		if err != nil {
			return nil, err
		}
	}
	return u, nil
}
