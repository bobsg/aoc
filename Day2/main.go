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
	parsed, err := ParseInput(string(input))
	if err != nil {
		log.Fatal("Could not parse input")
	}

	numSafe := ValidateReports(parsed)
	fmt.Printf("Number of safe reports: %d\n", numSafe)

	numSafeWithDamp := ValidateReportsWithDampener(parsed)
	fmt.Printf("Number of safe dampened reports: %d\n", numSafeWithDamp)
}
