package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

var (
	configPath = flag.String("c", "", "a configuration file (required)")
)

func main() {
	flag.Usage = usage
	flag.Parse()
	if len(*configPath) == 0 {
		abort(errors.New("expected a configuration file"))
	}
	config, err := newConfig(*configPath)
	if err != nil {
		abort(err)
	}
	target := newTarget(config)
	algorithm := newAlgorithm(config)
	surrogate := algorithm.Compute(target)
	fmt.Printf("Surrogate: %s\n", surrogate)
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
