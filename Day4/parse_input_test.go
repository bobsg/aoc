package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestParseInput(t *testing.T) {
	input := `..X...
.SAMX.
.A..A.
XMAS.S
.X....`

	got := ParseInput(input)

	want := [][]string{
		{".", ".", "X", ".", ".", "."},
		{".", "S", "A", "M", "X", "."},
		{".", "A", ".", ".", "A", "."},
		{"X", "M", "A", "S", ".", "S"},
		{".", "X", ".", ".", ".", "."},
	}

	if len(got) != len(want) {
		t.Fatal("result differs in length")
	}

	for i := 0; i < len(got); i++ {
		if !slices.Equal(got[i], want[i]) {
			t.Errorf("got %#v\nwant %#v\n", got[i], want[i])
		}
	}
}

func TestFindX(t *testing.T) {
	input := [][]string{
		{".", ".", "X", ".", ".", "."},
		{".", "S", "A", "M", "X", "."},
		{".", "A", ".", ".", "A", "."},
		{"X", "M", "A", "S", ".", "S"},
		{".", "X", ".", ".", ".", "."},
	}

	got := FindX(input)
	want := []Pos{
		{X: 2, Y: 0}, {X: 4, Y: 1}, {X: 0, Y: 3}, {X: 1, Y: 4},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}
