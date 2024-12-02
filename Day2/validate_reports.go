package main

import (
	"math"
)

func ValidateReports(input [][]int) int {
	numSafe := 0
	for _, report := range input {
		if ValidateReport(report) {
			numSafe++
		}
	}

	return numSafe
}

func ValidateReportsWithDampener(input [][]int) int {
	numSafe := 0
	for _, report := range input {
		if ValidateReport(report) {
			numSafe++
		} else if validateReportsWithDampener(report) {
			numSafe++
		}
	}

	return numSafe
}

func validateReportsWithDampener(input []int) bool {
	for i := 0; i < len(input); i++ {
		newInput := make([]int, 0)
		newInput = append(newInput, input[:i]...)
		newInput = append(newInput, input[i+1:]...)
		if ValidateReport(newInput) {
			return true
		}
	}

	return false
}

func ValidateReport(input []int) bool {
	if len(input) == 1 {
		return true
	}

	prevNum := input[0]
	increasing := true
	for i := 1; i < len(input); i++ {
		currentNum := input[i]
		if !isValidDistance(currentNum, prevNum) {
			return false
		}

		// Can compare to sorted slice?
		if i == 1 {
			increasing = currentNum > prevNum
		}

		if currentNum > prevNum && !increasing {
			return false
		}
		if prevNum > currentNum && increasing {
			return false
		}
		prevNum = currentNum
	}

	return true
}

func isValidDistance(a, b int) bool {
	distance := math.Abs(float64(a) - float64(b))
	return distance <= 3 && distance > 0
}
