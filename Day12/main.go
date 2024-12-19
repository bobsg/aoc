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

	garden := ParseInput(f)
	sum := SumFencePrice(garden)
	fmt.Printf("Price for fence: %d\n", sum)
}
