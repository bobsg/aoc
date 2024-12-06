package main

import (
	"slices"
	"strings"
	"testing"
)

func TestGetInvalidUpdates(t *testing.T) {
	ruleMap, updates := ParseInput(strings.Split(longInput, "\n"))

	got := GetInvalidUpdates(ruleMap, updates)
	want := [][]string{
		{"75", "97", "47", "61", "53"},
		{"61", "13", "29"},
		{"97", "13", "75", "29", "47"}}

	if len(got) != len(want) {
		t.Fatalf("got %d valid updates, want %d", len(got), len(want))
	}

	for i := 0; i < len(got); i++ {
		if !slices.Equal(got[i], want[i]) {
			t.Errorf("got %#v, want %#v at %d", got[i], want[i], i)
		}
	}

}

func TestSortInvalid(t *testing.T) {
	ruleMap, _ := ParseInput(strings.Split(longInput, "\n"))
	updates := []struct {
		Update   []string
		Expected []string
	}{
		{[]string{"75", "97", "47", "61", "53"}, []string{"97", "75", "47", "61", "53"}},
		{[]string{"61", "13", "29"}, []string{"61", "29", "13"}},
		{[]string{"97", "13", "75", "29", "47"}, []string{"97", "75", "47", "29", "13"}},
	}

	for _, test := range updates {
		got := SortInvalid(ruleMap, test.Update)
		if !slices.Equal(got, test.Expected) {
			t.Errorf("got %#v want %#v", got, test.Expected)
		}
	}
}

func TestSortInvalidBuiltIn(t *testing.T) {
	ruleMap, _ := ParseInput(strings.Split(longInput, "\n"))
	updates := []struct {
		Update   []string
		Expected []string
	}{
		{[]string{"75", "97", "47", "61", "53"}, []string{"97", "75", "47", "61", "53"}},
		{[]string{"61", "13", "29"}, []string{"61", "29", "13"}},
		{[]string{"97", "13", "75", "29", "47"}, []string{"97", "75", "47", "29", "13"}},
	}

	for _, test := range updates {
		got := SortInvalidBuiltIn(ruleMap, test.Update)
		if !slices.Equal(got, test.Expected) {
			t.Errorf("got %#v want %#v", got, test.Expected)
		}
	}
}

func TestGetSumFixedInvalids(t *testing.T) {
	ruleMap, updates := ParseInput(strings.Split(longInput, "\n"))

	got := GetSumFixedInvalids(ruleMap, updates)

	if got != 123 {
		t.Errorf("got %d want 123", got)
	}
}
