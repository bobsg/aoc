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
