package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/ready-steady/hdf5"
	"github.com/ready-steady/sequence"
)

var (
	configPath = flag.String("c", "", "a configuration file (required)")
	outputPath = flag.String("o", "", "an output file (required)")
)

func main() {
	flag.Usage = usage
	flag.Parse()
	if len(*configPath) == 0 {
		abort(errors.New("expected a configuration file"))
	}
	if len(*outputPath) == 0 {
		abort(errors.New("expected an output file"))
	}
	config, err := newConfig(*configPath)
	if err != nil {
		abort(err)
	}
	target := newTarget(config)
	algorithm := newAlgorithm(config)
	surrogate := algorithm.Compute(target)
	points := sequence.NewSobol(target.ni, config.Seed).Next(config.Samples)
	values := algorithm.Evaluate(surrogate, points)
	output, err := hdf5.Create(*outputPath)
	if err != nil {
		abort(err)
	}
	defer output.Close()
	if err := output.Put("surrogate", *surrogate); err != nil {
		abort(err)
	}
	if err := output.Put("points", points); err != nil {
		abort(err)
	}
	if err := output.Put("values", values); err != nil {
		abort(err)
	}
}

func abort(err error) {
	fmt.Printf("Error: %s.\n", err)
	os.Exit(1)
}

func usage() {
	fmt.Printf("Usage: %s [flags]\n\n", os.Args[0])
	fmt.Printf("Flags:\n")
	flag.PrintDefaults()
}
