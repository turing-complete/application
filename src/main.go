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
		fail(errors.New("expected a configuration file"))
	}
	config, err := newConfig(*configPath)
	if err != nil {
		fail(err)
	}
	target, err := newTarget(config)
	if err != nil {
		fail(err)
	}
	target.evaluate(make([]float64, target.ni))
}

func fail(err error) {
	fmt.Printf("Error: %s.\n", err)
	os.Exit(1)
}

func usage() {
	fmt.Printf("Usage: %s [options]\n\n", os.Args[0])
	fmt.Printf("Options:\n")
	flag.PrintDefaults()
}
