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
	board := ParseInput(string(input))

	guard, err := board.FindGuard()
	if err != nil {
		log.Fatal("did not find guard")
	}

	guard.DoRound(board)
	unique := guard.GetUniqueEntries()
	fmt.Printf("Unique positions in round history: %d\n", len(unique))

	newGuard, err := board.FindGuard()
	if err != nil {
		log.Fatal("did not find guard")
	}
	possible := newGuard.DoRoundWithObstructions(board)
	fmt.Printf("Possible locations for obstructions: %d\n", possible)
}
