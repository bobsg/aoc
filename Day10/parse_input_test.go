package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseInput(t *testing.T) {
	input := `0123
1234
8765
9876`

	got := ParseInput(strings.NewReader(input))
	want := TopoMap{
		{0, 1, 2, 3},
		{1, 2, 3, 4},
		{8, 7, 6, 5},
		{9, 8, 7, 6},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}
