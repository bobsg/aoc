package main

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	l := &Line{}
	l.Add(10)
	l.Add(11)
	l.Add(12)
	l.Add(13)

	if l.Len() != 4 {
		t.Errorf("got len %d, want 4", l.Len())
	}

	if l.First == nil {
		t.Fatal("stone not added")
	}

	if l.First.Value != 10 {
		t.Errorf("got value %d, want 10", l.First.Value)
	}

	if l.First.Next.Value != 11 {
		t.Errorf("got value %d, want 11", l.First.Next.Value)
	}

	if l.First.Next.Next.Value != 12 {
		t.Errorf("got value %d, want 12", l.First.Next.Next.Value)
	}

	if l.First.Next.Next.Next.Value != 13 {
		t.Errorf("got value %d, want 13", l.First.Next.Next.Next.Value)
	}

	if l.Last.Value != 13 {
		t.Errorf("got value %d, want 13", l.First.Value)

	}
}

func TestFindStone(t *testing.T) {
	l := &Line{}
	l.Add(1)
	l.Add(23)
	s := l.FindStone(l.First.Next)

	if s == nil || s.Value != 23 {
		t.Errorf("did not find stone, got %#v", s)
	}
}

func TestSplit(t *testing.T) {
	l := &Line{}
	l.Add(1)
	l.Add(23)
	l.Add(4)

	l.Split(l.First.Next)

	if l.Len() != 4 {
		t.Errorf("got %d, want length 4", l.Len())
	}

	if l.String() != "1 2 3 4" {
		t.Errorf("got '%s' want '%s'", l.String(), "1 2 3 4")
	}

	if l.StringRev() != "4 3 2 1" {
		t.Errorf("got '%s' want '%s'", l.StringRev(), "4 3 2 1")
	}

	if l.Last.Prev.Value != 3 {
		t.Errorf("got value %d, want 3", l.Last.Prev.Value)
	}

	if l.First.Next.Next != l.Last.Prev {
		t.Errorf("line not correct, first.next.next and last.prev should be the same stone")
	}
}

func TestBlink(t *testing.T) {
	l := &Line{}
	l.Add(125)
	l.Add(17)

	l.Blink()
	want := "253000 1 7"
	if l.String() != want {
		t.Errorf("got '%s', want '%s'", l.String(), want)
	}
}

func TestBlink2(t *testing.T) {
	l := []int{125, 17}

	l = Blink(l)
	want := "253000 1 7"
	if String(l) != want {
		t.Errorf("got '%s', want '%s'", String(l), want)
	}
}

func TestBlinkN(t *testing.T) {
	l := &Line{}
	l.Add(125)
	l.Add(17)

	l.BlinkN(6)
	got := l.String()
	want := "2097446912 14168 4048 2 0 2 4 40 48 2024 40 48 80 96 2 8 6 7 6 0 3 2"

	if l.Len() != 22 {
		t.Errorf("got len %d, want 22", l.Len())
	}

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

func TestBlinkN2(t *testing.T) {
	l := []int{125, 17}

	l = BlinkN(l, 6)
	//got := String(l)
	//want := "2097446912 14168 4048 2 0 2 4 40 48 2024 40 48 80 96 2 8 6 7 6 0 3 2"

	if len(l) != 22 {
		t.Errorf("got len %d, want 22", len(l))
	}

	/*if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}*/
}

func TestIterator(t *testing.T) {
	l := &Line{}
	l.Add(125)
	l.Add(17)

	l.BlinkN(6)
	want := []int{2097446912, 14168, 4048, 2, 0, 2, 4, 40, 48, 2024, 40, 48, 80, 96, 2, 8, 6, 7, 6, 0, 3, 2}
	got := []int{}
	for s := range l.Traverse() {
		got = append(got, s.Value)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}
