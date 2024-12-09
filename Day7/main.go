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
	ops := []Operation{Add, Mul}
	sum := SumValidations(input, ops)
	fmt.Printf("Sum: %d\n", sum)

	opsPart2 := []Operation{Add, Mul, Conc}
	sumPart2 := SumValidations(input, opsPart2)
	fmt.Printf("Sum part2: %d\n", sumPart2)
}
