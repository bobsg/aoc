package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseInput(t *testing.T) {
	input := `3267: 81 40 27
83: 17 5`

	r := strings.NewReader(input)
	got := ParseInput(r)
	want := []Equation{
		{
			TestValue: 3267,
			Numbers:   []int{81, 40, 27},
		},
		{
			TestValue: 83,
			Numbers:   []int{17, 5},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}

var longInput = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`
