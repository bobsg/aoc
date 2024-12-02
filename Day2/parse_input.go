package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseInput(input string) ([][]int, error) {
	rows := strings.Split(input, "\n")
	result := make([][]int, len(rows), len(rows))
	for i := 0; i < len(rows); i++ {
		cols := strings.Split(rows[i], " ")
		colsParsed := make([]int, len(cols), len(cols))
		for j := 0; j < len(cols); j++ {
			if len(cols) == 1 {
				continue
			}
			n, err := strconv.Atoi(cols[j])
			if err != nil {
				return nil, fmt.Errorf("cannot parse input at row %d, col %d, val %s", i, j, cols[j])
			}
			colsParsed[j] = n
		}
		result[i] = colsParsed
	}
	return result, nil
}
