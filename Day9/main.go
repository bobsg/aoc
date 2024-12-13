package main

import (
	"fmt"
	"io"
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

	disk := ParseInput(f)
	disk.Defrag()
	checksum := disk.Checksum()
	fmt.Printf("Checksum: %d\n", checksum)

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		log.Panic("could not rewind file")
	}
	d2 := ParseInputP2(f)
	d2.Defrag()
	c2 := d2.Checksum()
	fmt.Printf("Checksum2: %d\n", c2)
}
