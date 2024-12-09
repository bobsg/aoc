package main

import "strings"

func ParseInput(input string) Board {
	var result Board
	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines); i++ {
		result = append(result, strings.Split(lines[i], ""))
	}
	return result
}
