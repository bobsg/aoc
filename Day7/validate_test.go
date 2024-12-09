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

	ops := GenerateOperations(Operation(Add), Operation(Mul), 3)

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
	ops := GenerateOperations(Operation(Add), Operation(Mul), 4)

	got := ValidateCalibration(input, ops)

	if got != false {
		t.Errorf("got %v, want %v", got, false)
	}
}

func TestValidateAll(t *testing.T) {
	input := ParseInput(strings.NewReader(longInput))

	got := ValidateAll(input)
	want := []int{190, 3267, 292}

	if !slices.Equal(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}

}

func TestSumValidations(t *testing.T) {
	input := ParseInput(strings.NewReader(longInput))

	got := SumValidations(input)
	want := 3749

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

}
