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

	input := ParseInput(f)
	sum := SumValidations(input)
	fmt.Printf("Sum: %d\n", sum)
}
