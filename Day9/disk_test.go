package main

import (
	"strings"
	"testing"
)

func TestDefrag(t *testing.T) {
	t.Run("defrag1", func(t *testing.T) {
		r := strings.NewReader("12345")
		disk := ParseInput(r)
		disk.Defrag()
		layout := disk.String()
		want := "022111222......"

		if layout != want {
			t.Errorf("got '%s' want '%s", layout, want)
		}
	})
	t.Run("long input", func(t *testing.T) {
		r := strings.NewReader("2333133121414131402")
		disk := ParseInput(r)
		disk.Defrag()
		layout := disk.String()
		want := "0099811188827773336446555566.............."

		if layout != want {
			t.Errorf("got '%s' want '%s", layout, want)
		}
	})
}

func TestChecksum(t *testing.T) {
	r := strings.NewReader("2333133121414131402")
	disk := ParseInput(r)
	disk.Defrag()
	checksum := disk.Checksum()

	if checksum != 1928 {
		t.Errorf("got %d, want 1928", checksum)
	}
}
