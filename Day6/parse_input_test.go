package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	want := bigBoard

	got := ParseInput(input)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("input does not match")
	}
}
