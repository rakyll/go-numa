package main

import (
	"log"

	"github.com/rakyll/go-numa"
)

func main() {
	if !numa.IsAvailable() {
		log.Fatalln("NUMA is not available on this machine.")
	}

	// Runs the current goroutine in node 0.
	numa.SetPreferred(0)
	// Do work...
}
