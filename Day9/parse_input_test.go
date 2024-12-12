package main

import (
	"strings"
	"testing"
)

func TestParseInput(t *testing.T) {
	t.Run("short input", func(t *testing.T) {
		r := strings.NewReader("12345")

		disk := ParseInput(r)
		got := disk.String()
		want := "0..111....22222"

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("long input with zero length free block", func(t *testing.T) {
		r := strings.NewReader("2333133121414131402")

		disk := ParseInput(r)
		got := disk.String()
		want := "00...111...2...333.44.5555.6666.777.888899"
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})
}
