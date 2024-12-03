package main

import (
	"slices"
	"testing"
)

func TestParseMemory(t *testing.T) {
	input := `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`

	got, err := ParseMemory(input)
	if err != nil {
		t.Fatal("Got error but did not expect one", err)
	}
	want := [][]int{{2, 4}, {5, 5}, {11, 8}, {8, 5}}

	if len(got) != len(want) {
		t.Errorf("got %d matches, want %d matches", len(got), len(want))
	}

	for i := 0; i < len(got); i++ {
		if !slices.Equal(got[i], want[i]) {
			t.Errorf("got %v, want %v for match %d", got[i], want[i], i)
		}
	}
}

func TestParseMemoryWithConditionals(t *testing.T) {
	input := `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

	got, err := ParseMemoryWithConditionals(input)
	if err != nil {
		t.Error("Got error but did not expect one", err)
	}

	want := []Command{{
		Pos:  1,
		Type: MUL,
		X:    2,
		Y:    4,
	}, {
		Pos:  20,
		Type: DONT,
		X:    0,
		Y:    0,
	},
		{
			Pos:  28,
			Type: MUL,
			X:    5,
			Y:    5,
		}, {
			Pos:  48,
			Type: MUL,
			X:    11,
			Y:    8,
		}, {
			Pos:  59,
			Type: DO,
			X:    0,
			Y:    0,
		}, {
			Pos:  64,
			Type: MUL,
			X:    8,
			Y:    5,
		},
	}

	if !slices.Equal(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}
