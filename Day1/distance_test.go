package main

import (
	"slices"
	"testing"
)

func TestDistance(t *testing.T) {
	list1 := []int64{3, 4, 2, 1, 3, 3}
	list2 := []int64{4, 3, 5, 3, 9, 3}

	got := Distance(list1, list2)
	want := int64(11)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestParseInput(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3

`
	gotList1, gotList2, err := ParseInput(input)
	if err != nil {
		t.Error("Got error but did not expect one")
	}
	wantList1 := []int64{3, 4, 2, 1, 3, 3}
	wantList2 := []int64{4, 3, 5, 3, 9, 3}

	if !slices.Equal(gotList1, wantList1) {
		t.Errorf("got %v, want %v for list1", gotList1, wantList1)
	}

	if !slices.Equal(gotList2, wantList2) {
		t.Errorf("got %v, want %v for list2", gotList1, wantList1)
	}
}
