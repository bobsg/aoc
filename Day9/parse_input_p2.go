package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
)

func ParseInputP2(r io.Reader) *DiskP2 {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)
	disk := NewDiskP2()
	id := 0
	free := false
	for scanner.Scan() {
		s := scanner.Text()
		size, err := strconv.Atoi(s)
		if err != nil {
			log.Panicf("could not parse %s as int", s)
		}
		if !free {
			disk.AddBlock(id, size)
			id++
		} else {
			disk.AddFreeBlock(size)
		}
		free = !free
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
		return nil
	}

	return disk
}
