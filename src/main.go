package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

var (
	configFile = flag.String("c", "", "a configuration file (required)")
)

func main() {
	flag.Usage = usage
	flag.Parse()

	if len(*configFile) == 0 {
		fail(errors.New("expected a configuration file"))
	}
	config, err := newConfig(*configFile)
	if err != nil {
		fail(err)
	}
	if len(config.Target) == 0 {
		fail(errors.New("expected a target command"))
	}
	fmt.Printf("Target: %s\n", config.Target)
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
