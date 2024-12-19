package main

import (
	"strings"
	"testing"
)

func TestParseInput(t *testing.T) {
	got := ParseInput(strings.NewReader(tinyExample))
	want := [][]*Plot{
		{NewPlot("A", 0, 0, 0), NewPlot("A", 1, 0, 0)},
		{NewPlot("B", 0, 1, 0), NewPlot("B", 1, 1, 0)},
	}

	assertGarden(t, got, want)
}

func assertGarden(t testing.TB, got, want [][]*Plot) {
	t.Helper()
	if len(got) != len(want) {
		t.Errorf("got height %d, want %d", len(got), len(want))
	}
	if len(got[0]) != len(want[0]) {
		t.Errorf("got width %d, want %d", len(got[0]), len(want[0]))
	}
	for y := 0; y < len(got); y++ {
		for x := 0; x < len(got[y]); x++ {
			if !got[y][x].Equal(want[y][x]) {
				t.Errorf("got %#v, want %#v", got[y][x], want[y][x])
			}
		}
	}
}
