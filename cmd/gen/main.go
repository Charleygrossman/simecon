package main

import (
	"flag"
	"fmt"
	"tradesim/cmd/gen/internal"
)

var (
	help bool
	out  string
)

func init() {
	flag.BoolVar(&help, "h", false, "")
	flag.BoolVar(&help, "help", false, "print description and available command options")
	flag.StringVar(&out, "o", "", "path to output file")
}

func main() {
	flag.Parse()

	if help {
		fmt.Printf("generate simulation configuration files\n\noptions\n")
		flag.PrintDefaults()
		return
	}

	if err := internal.Generate(out); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
