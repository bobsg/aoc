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
	list1, list2 := ParseInput(string(input))

	distance := Distance(list1, list2)
	fmt.Printf("Distance: %d\n", distance)
	similarity := Similarity(list1, list2)
	fmt.Printf("Similarity: %d", similarity)
}
