package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./input")
	if err != nil {
		log.Fatal("Could not open input file")
	}

	antennas := ParseInput(f)
	antiNodes := antennas.CalculateAllAntiNodes()
	u := antiNodes.UniquePositions()

	fmt.Printf("Unique locations %d\n", len(u))

}
