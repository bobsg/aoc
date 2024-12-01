package aoc1

import "testing"

func TestDistance(t *testing.T) {
	list1 := []int{3, 4, 2, 1, 3, 3}
	list2 := []int{4, 3, 5, 3, 9, 3}

	got := Distance(list1, list2)
	want := 11

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
