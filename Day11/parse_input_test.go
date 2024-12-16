package main

import "testing"

func TestParseInput(t *testing.T) {
	in := "0 1 10 99 999"

	got := ParseInput(in)

	if got.String() != in {
		t.Errorf("got '%s' want '%s'", got.String(), in)
	}

	if got.StringRev() != "999 99 10 1 0" {
		t.Errorf("got '%s' want '%s'", got.StringRev(), "999 99 10 1 0")
	}
}
