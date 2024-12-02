package main

import "testing"

func TestValidateReports(t *testing.T) {
	input := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	got := ValidateReports(input)
	want := 2

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestValidateReportsWithDampener(t *testing.T) {
	input := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	got := ValidateReportsWithDampener(input)
	want := 4

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestValidateReport(t *testing.T) {
	input := []int{7, 6, 4, 2, 1}

	got := ValidateReport(input)

	if got != true {
		t.Errorf("got %v want %v", got, true)
	}
}

func TestValidateReportsWithDampenerInternal(t *testing.T) {
	input := []int{1, 3, 2, 4, 5}

	got := validateReportsWithDampener(input)

	if got != true {
		t.Error("got false want true")
	}
}
