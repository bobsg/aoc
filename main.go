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
	list1, list2, err := ParseInput(string(input))
	if err != nil {
		log.Fatal("Could not parse input")
	}

	result := Distance(list1, list2)
	fmt.Println(result)
}
