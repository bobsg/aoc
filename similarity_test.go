package main

import "testing"

func TestSimilarity(t *testing.T) {
	list1 := []int64{3, 4, 2, 1, 3, 3}
	list2 := []int64{4, 3, 5, 3, 9, 3}

	got := Similarity(list1, list2)
	want := int64(31)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
