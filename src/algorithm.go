package main

import (
	"log"

	"github.com/ready-steady/adapt/algorithm"
	"github.com/ready-steady/adapt/algorithm/hybrid"
	"github.com/ready-steady/adapt/basis/polynomial"
	"github.com/ready-steady/adapt/grid/equidistant"
)

type Algorithm struct {
	hybrid.Algorithm
	strategy func() *hybrid.Strategy
}

type Surrogate struct {
	algorithm.Surrogate
}

func newAlgorithm(config *Config) *Algorithm {
	ni, no := config.Inputs, config.Outputs
	grid := equidistant.NewOpen(ni)
	basis := polynomial.NewOpen(ni, 1)
	strategy := func() *hybrid.Strategy {
		return hybrid.NewStrategy(ni, no, grid, config.MinLevel, config.MaxLevel,
			config.AbsoluteError, config.RelativeError, config.ScoreError)
	}
	return &Algorithm{
		Algorithm: *hybrid.New(ni, no, grid, basis),
		strategy:  strategy,
	}
}

func (self *Algorithm) Compute(target *Target) *Surrogate {
	log.Printf("[Algorithm] Start constructing with %d input(s) and output %d (s).\n",
		target.ni, target.no)
	compute := func(z, u []float64) {
		log.Printf("[Target] Start computing at node %v.\n", z)
		if err := target.Compute(z, u); err != nil {
			abort(err)
		}
		log.Printf("[Target] Finish computing at node %v.\n", z)
	}
	surrogate := self.Algorithm.Compute(compute, self.strategy())
	log.Printf("[Algorithm] Finish constructing with %d node(s).\n", surrogate.Nodes)
	return &Surrogate{*surrogate}
}

func (self *Algorithm) Evaluate(surrogate *Surrogate, points []float64) []float64 {
	log.Printf("[Algorithm] Start evaluating at %d point(s).\n",
		uint(len(points))/surrogate.Inputs)
	values := self.Algorithm.Evaluate(&surrogate.Surrogate, points)
	log.Printf("[Algorithm] Finish evaluating.\n")
	return values
}
