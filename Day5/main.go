package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input")
	if err != nil {
		log.Fatal("Could not read input")
	}
	rules, updates := ParseInput(strings.Split(string(input), "\n"))

	sum := GetSum(rules, updates)
	fmt.Printf("Sum for valid updates: %d\n", sum)

	rules2, updates2 := ParseInput(strings.Split(string(input), "\n"))
	invalidsSum := GetSumFixedInvalids(rules2, updates2)
	fmt.Printf("Sum for fixed invalid updates: %d\n", invalidsSum)

}
