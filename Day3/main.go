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
	parsed, err := ParseMemory(string(input))
	if err != nil {
		log.Fatal("Could not parse input")
	}

	result := RunMemory(parsed)
	fmt.Printf("Result: %d\n", result)

	parsedCond, err := ParseMemoryWithConditionals(string(input))
	if err != nil {
		log.Fatal("Could not parse input with conditionals", err)
	}
	resultCond := RunMemoryWithConditionals(parsedCond)
	fmt.Printf("Result with conditionals: %d", resultCond)
}
