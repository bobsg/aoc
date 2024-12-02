package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestParseInput(t *testing.T) {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
	got, err := ParseInput(input)
	want := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	if err != nil {
		fmt.Println(err.Error())
		t.Fatalf("Got error but did not expect one")
	}

	if len(got) != len(want) {
		t.Fatalf("got length %d, want length %d for outer array", len(got), len(want))
	}

	for i := 0; i < len(got); i++ {
		if len(got[i]) != len(want[i]) {
			t.Fatalf("got length %d, want length %d for inner array i = %d", len(got[i]), len(want[i]), i)
		}

		if !slices.Equal(got[i], want[i]) {
			t.Errorf("got %v, want %v", got[i], want[i])
		}
	}
}
