package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Open("./input")
	if err != nil {
		log.Fatal("Could not open input file")
	}

	disk := ParseInput(f)
	disk.Defrag()
}
