package main

import (
	"slices"
	"strings"
	"testing"
)

func TestValidateSuccess(t *testing.T) {
	input := Equation{
		TestValue: 3267,
		Numbers:   []int{81, 40, 27},
	}
	o := []Operation{Add, Mul}
	ops := GenerateOperations(3, o)

	got := ValidateCalibration(input, ops)

	if got != true {
		t.Errorf("got %v, want %v", got, true)
	}
}

func TestValidateFailure(t *testing.T) {
	input := Equation{
		TestValue: 21037,
		Numbers:   []int{9, 7, 18, 13},
	}
	o := []Operation{Add, Mul}
	ops := GenerateOperations(4, o)

	got := ValidateCalibration(input, ops)

	if got != false {
		t.Errorf("got %v, want %v", got, false)
	}
}

func TestValidateAll(t *testing.T) {
	input := ParseInput(strings.NewReader(longInput))
	ops := []Operation{Add, Mul}
	got := ValidateAll(input, ops)
	want := []int{190, 3267, 292}

	if !slices.Equal(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}

}

func TestSumValidations(t *testing.T) {
	input := ParseInput(strings.NewReader(longInput))

	ops := []Operation{Add, Mul}
	got := SumValidations(input, ops)
	want := 3749

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

}

func TestSumValidationsPart2(t *testing.T) {
	input := ParseInput(strings.NewReader(longInput))

	ops := []Operation{Add, Mul, Conc}
	got := SumValidations(input, ops)
	want := 11387

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

}
