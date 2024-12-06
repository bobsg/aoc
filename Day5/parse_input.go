package main

import (
	"slices"
	"strings"
)

func ParseInput(input []string) (RuleMap, [][]string) {
	rules := RuleMap{}
	var updatesIndex int
	for i := 0; i < len(input); i++ {
		if input[i] == "" {
			updatesIndex = i + 1
			break
		}
		pair := strings.Split(input[i], "|")
		if slices.Contains(rules[pair[0]], pair[1]) {
			continue
		}
		rules[pair[0]] = append(rules[pair[0]], pair[1])
	}

	var updates [][]string
	for i := updatesIndex; i < len(input); i++ {
		parts := strings.Split(input[i], ",")
		updates = append(updates, parts)
	}

	return rules, updates
}

type RuleMap map[string][]string
