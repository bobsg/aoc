package main

type Operation func(a, b int) int

func Add(a, b int) int {
	return a + b
}

func Mul(a, b int) int {
	return a * b
}

func GenerateOperations[T any](a, b T, maxLength int) [][][]T {
	result := [][][]T{}
	result = append(result, [][]T{{}})
	result = append(result, [][]T{{a}, {b}})
	for i := 1; i <= maxLength-1; i++ {
		coll := [][]T{}
		for j := 0; j < len(result[i]); j++ {
			wa := make([]T, len(result[i][j]))
			wb := make([]T, len(result[i][j]))
			copy(wa, result[i][j])
			copy(wb, result[i][j])
			wa = append(wa, a)
			wb = append(wb, b)
			coll = append(coll, wa)
			coll = append(coll, wb)
		}
		result = append(result, coll)
	}

	return result
}
