package main

import (
	"bufio"
	"io"
	"strings"
)

func ParseInput(r io.Reader) [][]*Plot {
	scanner := bufio.NewScanner(r)
	result := [][]*Plot{}
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		row := []*Plot{}
		chars := strings.Split(line, "")
		for x, char := range chars {
			if char == "\n" {
				break
			}
			row = append(row, NewPlot(char, x, y, 0))
		}
		result = append(result, row)
		y++
	}
	return result
}
