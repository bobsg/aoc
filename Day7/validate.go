package main

import "log"

func SumValidations(calibrations []Equation) int {
	valids := ValidateAll(calibrations)

	sum := 0
	for i := 0; i < len(valids); i++ {
		sum += valids[i]
	}

	return sum
}

func ValidateAll(calibrations []Equation) []int {
	maxLength := 0
	for i := 0; i < len(calibrations); i++ {
		maxLength = max(maxLength, len(calibrations[i].Numbers))
	}

	ops := GenerateOperations(Operation(Add), Operation(Mul), maxLength)
	valids := []int{}
	for i := 0; i < len(calibrations); i++ {
		if ValidateCalibration(calibrations[i], ops) {
			valids = append(valids, calibrations[i].TestValue)
		}
	}

	return valids
}

func ValidateCalibration(e Equation, o [][][]Operation) bool {
	if len(e.Numbers) == 1 {
		return e.Numbers[1] == e.TestValue
	}
	if len(e.Numbers) == 0 {
		return false
	}

	numOps := len(e.Numbers) - 1
	if numOps >= len(o) {
		log.Panic("Not enough operations for numbers")
	}
	operations := o[numOps]

	for i := 0; i < len(operations); i++ {
		sum := e.Numbers[0]
		for j := 1; j < len(e.Numbers); j++ {
			op := operations[i][j-1]
			sum = op(sum, e.Numbers[j])
		}
		if sum == e.TestValue {
			return true
		}
	}

	return false
}
