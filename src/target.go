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

func newTarget(config *Config) *Target {
	return &Target{
		ni:   config.Inputs,
		no:   config.Outputs,
		name: config.Command,
	}
}

func (self *Target) Compute(z, u []float64) error {
	inputs := make([]string, self.ni)
	for i := uint(0); i < self.ni; i++ {
		inputs[i] = fmt.Sprintf("%e", z[i])
	}
	bytes, err := exec.Command(self.name, inputs...).Output()
	if err != nil {
		return err
	}
	outputs := strings.Split(strings.TrimSpace(string(bytes)), " ")
	if uint(len(outputs)) != self.no {
		return errors.New("expected a different number of outputs")
	}
	for i := uint(0); i < self.no; i++ {
		u[i], err = strconv.ParseFloat(outputs[i], 64)
		if err != nil {
			return err
		}
	}
	return nil
}
