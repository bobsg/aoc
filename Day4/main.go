package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.ReadFile("./input")
	if err != nil {
		log.Fatal("Could not read input")
	}
	parsed := ParseInput(string(input))

	count := CountXmases(parsed)
	fmt.Printf("Result: %d\n", count)

	xcount := CountCrossXmases(parsed)
	fmt.Printf("Result x-mas: %d\n", xcount)
}
