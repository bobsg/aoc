package main

import (
	"bufio"
	"io"
)

func ParseInput(r io.Reader) TopoMap {
	scanner := bufio.NewScanner(r)
	result := TopoMap{}
	for scanner.Scan() {
		line := scanner.Text()
		row := []int{}
		for _, char := range line {
			if char == '\n' {
				break
			}
			n := 0
			if char == '.' {
				n = -1
			} else {
				n = int(char - '0')
			}
			row = append(row, n)
		}
		result = append(result, row)
	}
	return result
}
