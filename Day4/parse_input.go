package main

import (
	"strings"
)

func ParseInput(input string) [][]string {
	var ret [][]string
	rows := strings.Split(input, "\n")
	for i := 0; i < len(rows); i++ {
		chars := strings.Split(rows[i], "")
		ret = append(ret, chars)
	}
	return ret
}
