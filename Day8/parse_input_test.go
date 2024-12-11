package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseInput(t *testing.T) {
	input := longInput

	got := ParseInput(strings.NewReader(input))
	want := Antennas{
		width:  12, // 0 - 11
		height: 12, // 0 - 11
		antennas: map[frequency][]Antenna{
			'0': {
				{'0', Position{X: 8, Y: 1}},
				{'0', Position{X: 5, Y: 2}},
				{'0', Position{X: 7, Y: 3}},
				{'0', Position{X: 4, Y: 4}},
			},
			'A': {
				{'A', Position{X: 6, Y: 5}},
				{'A', Position{X: 8, Y: 8}},
				{'A', Position{X: 9, Y: 9}},
			},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}
