package main

import (
	"flag"
	"fmt"
	"tradesim/cmd/sim/internal"
)

var (
	help    bool
	in, out string
)

func init() {
	flag.BoolVar(&help, "h", false, "")
	flag.BoolVar(&help, "help", false, "print description and available command options")
	flag.StringVar(&in, "i", "", "path to simulation configuration file")
	flag.StringVar(&out, "o", "", "path to simulation output file")
}

func main() {
	flag.Parse()

	if help {
		fmt.Printf("run trade simulations from input configuration\n\noptions\n")
		flag.PrintDefaults()
		return
	}

	if err := internal.Simulate(in, out); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
