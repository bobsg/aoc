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
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	topoMap := ParseInput(f)
	sum := topoMap.SumAllTrails()
	fmt.Printf("Sum of scores: %d\n", sum)
}
