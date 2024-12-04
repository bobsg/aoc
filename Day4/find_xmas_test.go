package main

import "testing"

func TestFindXmas(t *testing.T) {
	input := [][]string{
		{".", ".", "X", ".", ".", "."},
		{".", "S", "A", "M", "X", "."},
		{".", "A", ".", ".", "A", "."},
		{"X", "M", "A", "S", ".", "S"},
		{".", "X", ".", ".", ".", "."},
	}

	got := FindXmas(input, Pos{X: 0, Y: 3}, RIGHT)

	if got != 1 {
		t.Errorf("expected match")
	}
}

func TestFindXmases(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	parsed := ParseInput(input)
	count := CountXmases(parsed)

	if count != 18 {
		t.Errorf("got %d but expected 18", count)
	}
}
