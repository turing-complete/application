package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	help = flag.Bool("h", false, "show this message")
)

func main() {
	flag.Usage = usage
	flag.Parse()
	if *help {
		usage()
		return
	}
}

func usage() {
	fmt.Printf("Usage: %s [options]\n\n", os.Args[0])
	fmt.Printf("Options:\n")
	flag.PrintDefaults()
}
