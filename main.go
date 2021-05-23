package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	help := flag.Bool("help", false, "print description and available command options")

	flag.Parse()

	if *help {
		fmt.Printf("tradesim runs simulations of trades within a market system\n\n")
		fmt.Println("options")
		flag.PrintDefaults()
		os.Exit(1)
	}
}
