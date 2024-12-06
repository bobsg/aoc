package main

import (
	"slices"
	"strings"
	"testing"
)

func TestCheckUpdatePass(t *testing.T) {
	update := []string{"75", "47", "61", "53", "29"}
	ruleMap := RuleMap{"47": []string{"53"}, "61": []string{"29"}} // 47 before 53, 61 before 29

	got := CheckUpdate(ruleMap, update)

	if !got {
		t.Errorf("expected success but did not reach it :( ")
	}
}

func TestCheckUpdateFail(t *testing.T) {
	update := []string{"75", "47", "29", "53", "61"}
	ruleMap := RuleMap{"47": []string{"53"}, "61": []string{"29"}} // 47 before 53, 61 before 29

	got := CheckUpdate(ruleMap, update)

	if got {
		t.Errorf("expected failure but did not reach it :( ")
	}
}

func TestCheckUpdates(t *testing.T) {
	ruleMap, updates := ParseInput(strings.Split(longInput, "\n"))

	got := CheckUpdates(ruleMap, updates)
	want := [][]string{
		{"75", "47", "61", "53", "29"},
		{"97", "61", "53", "29", "13"},
		{"75", "29", "13"}}

	if len(got) != len(want) {
		t.Fatalf("got %d valid updates, want %d", len(got), len(want))
	}

	for i := 0; i < len(got); i++ {
		if !slices.Equal(got[i], want[i]) {
			t.Errorf("got %#v, want %#v at %d", got[i], want[i], i)
		}
	}
}

func TestGetSumForValidUpdates(t *testing.T) {
	ruleMap, updates := ParseInput(strings.Split(longInput, "\n"))

	got := GetSum(ruleMap, updates)

	if got != 143 {
		t.Errorf("got %d want 143", got)
	}
}

var longInput = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`
