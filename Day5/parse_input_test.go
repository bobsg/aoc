package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestParseInput(t *testing.T) {
	input := []string{
		"47|53",
		"97|13",

		"",
		"75,47,61,53,29",
		"97,61,53,29,13",
	}

	gotRuleMap, gotUpdates := ParseInput(input)
	wantRuleMap := RuleMap{"47": []string{"53"}, "97": []string{"13"}}
	wantUpdates := [][]string{{"75", "47", "61", "53", "29"}, {"97", "61", "53", "29", "13"}}

	if !reflect.DeepEqual(gotRuleMap, wantRuleMap) {
		t.Errorf("got %#v, want %#v", gotRuleMap, wantRuleMap)
	}

	if len(gotUpdates) != len(wantUpdates) {
		t.Fatalf("got %d updates want %d", len(gotUpdates), len(wantUpdates))
	}

	for i := 0; i < len(gotUpdates); i++ {
		if !slices.Equal(gotUpdates[i], wantUpdates[i]) {
			t.Errorf("got %#v, want %#v at %d", gotUpdates[i], wantUpdates[i], i)
		}
	}
}
