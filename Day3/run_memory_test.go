package main

import "testing"

func TestRunMemory(t *testing.T) {
	input := [][]int{{2, 4}, {5, 5}, {11, 8}, {8, 5}}

	got := RunMemory(input)
	want := 161

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestRunMemoryWithConditionals(t *testing.T) {
	input := `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`
	parsedInput, err := ParseMemoryWithConditionals(input)
	if err != nil {
		t.Error("got error but did not expect one", err)
	}

	got := RunMemoryWithConditionals(parsedInput)
	want := 48

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
