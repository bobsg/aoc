package main

import (
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("can add one block", func(t *testing.T) {
		d := NewDisk()
		id := 10
		size := 5
		d.AddBlock(id, size)

		if d.Len() != 1 {
			t.Errorf("got len %d, want 1", d.Len())
		}

		if d.First == nil {
			t.Fatal("block not added")
		}

		if d.First.Id != 10 {
			t.Errorf("got id %d, want 10", d.First.Id)
		}

		if d.Last.Id != 10 {
			t.Errorf("got id %d, want 10", d.First.Id)

		}
	})
	t.Run("can add two blocks", func(t *testing.T) {
		d := NewDisk()
		d.AddBlock(10, 5)
		d.AddBlock(5, 10)

		if d.Len() != 2 {
			t.Errorf("got len %d, want 2", d.Len())
		}

		if d.First.Id != 10 {
			t.Errorf("got id %d, want 10", d.First.Id)
		}

		if d.Last != d.First.Next {
			t.Error("last and first.next should point to same block")
		}

		if d.Last.Id != 5 {
			t.Errorf("got id %d, want 5", d.Last.Id)
		}

		if d.Last.Previous != d.First {
			t.Errorf("expexted last.prev to equal d.first")
		}
	})
}

func TestInsertAfter(t *testing.T) {
	t.Run("insert in the middle", func(t *testing.T) {
		r := strings.NewReader("12345")
		disk := ParseInput(r)
		b := disk.FindBlockById(1)

		if b == nil {
			t.Fatal("expected to find block with id 1")
		}

		disk.InsertAfter(3, 2, b)
		got := disk.String()
		want := "0..11133....22222"
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})
	t.Run("insert free at end", func(t *testing.T) {
		r := strings.NewReader("12345")
		disk := ParseInput(r)
		b := disk.FindBlockById(2)

		if b == nil {
			t.Fatal("expected to find block with id 2")
		}

		disk.InsertFreeAfter(1, b)
		got := disk.String()
		want := "0..111....22222."
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})
}

func TestInsertBefore(t *testing.T) {
	t.Run("insert in the middle", func(t *testing.T) {
		r := strings.NewReader("12345")
		disk := ParseInput(r)
		b := disk.FirstFree()

		if b == nil {
			t.Fatal("expected to find free block")
		}

		disk.InsertBefore(3, 2, b)
		got := disk.String()
		want := "033..111....22222"
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})
}

func TestDelete(t *testing.T) {
	r := strings.NewReader("12345")
	disk := ParseInput(r)
	b := disk.FindBlockById(1)
	disk.Delete(b)
	got := disk.String()
	want := "0......22222"
	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}

	if disk.Len() != 3 {
		t.Errorf("expected length 3")
	}
}

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

		if disk.Len() != 14 {
			t.Errorf("expected 14 blocks")
		}
	})
}
