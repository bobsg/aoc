package main

import (
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
)

func sortFunc(i, j int64) int {
	if i < j {
		return -1
	}
	return 1
}

func Distance(list1 []int64, list2 []int64) int64 {
	slices.SortFunc(list1, sortFunc)
	slices.SortFunc(list2, sortFunc)

	var distance int64
	for i := 0; i < len(list1); i++ {
		distance += int64(math.Abs(float64(list1[i]) - float64(list2[i])))
	}

	return distance
}

func ParseInput(input string) ([]int64, []int64) {
	rows := strings.Split(input, "\n")
	var list1 []int64
	var list2 []int64
	for i := 0; i < len(rows); i++ {
		cols := strings.Split(rows[i], "   ")
		if len(cols) != 2 {
			continue
		}

		num, err := strconv.ParseInt(cols[0], 10, 0)
		if err != nil {
			log.Panicf("Invalid number at row %d, number %q in column 0", i, cols[0])
		}
		list1 = append(list1, num)
		num, err = strconv.ParseInt(cols[1], 10, 0)
		if err != nil {
			log.Panicf("Invalid number at row %d, number %q in column 1", i, cols[0])
		}
		list2 = append(list2, num)
	}
	return list1, list2
}
