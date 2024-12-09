package main

import (
	"fmt"
	"log"
	"strconv"
)

type Operation func(a, b int) int

func Add(a, b int) int {
	return a + b
}

func Mul(a, b int) int {
	return a * b
}

func Conc(a, b int) int {
	ab := fmt.Sprintf("%d%d", a, b)
	res, err := strconv.Atoi(ab)
	if err != nil {
		log.Fatalf("could not concatenate %d and %d %s", a, b, ab)
	}
	return res
}

func GenerateOperations[T any](maxLength int, inputs []T) [][][]T {
	result := [][][]T{}
	result = append(result, [][]T{{}})
	first := [][]T{}
	for i := 0; i < len(inputs); i++ {
		first = append(first, []T{inputs[i]})
	}
	result = append(result, first)
	for i := 1; i <= maxLength-1; i++ {
		coll := [][]T{}
		for j := 0; j < len(result[i]); j++ {
			for k := 0; k < len(inputs); k++ {
				o := make([]T, len(result[i][j]))
				copy(o, result[i][j])
				o = append(o, inputs[k])
				coll = append(coll, o)
			}
		}
		result = append(result, coll)
	}

	return result
}
