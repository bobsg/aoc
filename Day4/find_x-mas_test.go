package main

import "testing"

func TestFindA(t *testing.T) {
	input := `.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........`
	parsed := ParseInput(input)
	as := FindA(parsed)

	if len(as) != 9 {
		t.Errorf("expected 9 matches")
	}
}

func TestIsInBounds2(t *testing.T) {
	var input [][]string = make([][]string, 5)
	for i := 0; i < len(input); i++ {
		input[i] = make([]string, 5)
	}

	tests := []struct {
		Expected bool
		P        Pos
	}{
		{false, Pos{0, 0}},
		{false, Pos{4, 4}},
		{false, Pos{0, 1}},
		{false, Pos{1, 0}},
		{false, Pos{2, 4}},
		{false, Pos{4, 2}},
		{true, Pos{2, 2}},
		{true, Pos{3, 1}},
	}

	for _, test := range tests {
		got := isInBounds2(input, test.P)
		if got != test.Expected {
			t.Errorf("got %v expected %v for x: %d y: %d", got, test.Expected, test.P.X, test.P.Y)
		}
	}
}

func TestTestPosition(t *testing.T) {
	// a at x 2 y 1
	input := [][]string{
		{".", "M", ".", "S", ".", ".", ".", ".", ".", "."},
		{".", ">", "A", "<", ".", "M", "S", "M", "S", "."},
		{".", "M", ".", "S", ".", "M", "A", "A", ".", "."},
		{".", ".", "A", ".", "A", "S", "M", "S", "M", "."},
		{".", "M", ".", "S", ".", "M", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{"S", ".", "S", ".", "S", ".", "S", ".", "S", "."},
		{".", "A", ".", "A", ".", "A", ".", "A", ".", "."},
		{"M", ".", "M", ".", "M", ".", "M", ".", "M", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
	}

	got := testPosition(input, Pos{X: 2, Y: 1})

	if got != 1 {
		t.Errorf("got %d want 1", got)
	}
}

func TestCountCrossXmases(t *testing.T) {
	input := [][]string{
		{".", "M", ".", "S", ".", ".", ".", ".", ".", "."},
		{".", ">", "A", "<", ".", "M", "S", "M", "S", "."},
		{".", "M", ".", "S", ".", "M", "A", "A", ".", "."},
		{".", ".", "A", ".", "A", "S", "M", "S", "M", "."},
		{".", "M", ".", "S", ".", "M", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{"S", ".", "S", ".", "S", ".", "S", ".", "S", "."},
		{".", "A", ".", "A", ".", "A", ".", "A", ".", "."},
		{"M", ".", "M", ".", "M", ".", "M", ".", "M", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
	}

	count := CountCrossXmases(input)

	if count != 9 {
		t.Errorf("got %d expected 9", count)
	}

}
