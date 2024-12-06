package main

import (
	"log"
	"math"
	"strconv"
)

func GetSum(rules RuleMap, updates [][]string) int {
	validUpdates := CheckUpdates(rules, updates)
	var sum int
	for i := 0; i < len(validUpdates); i++ {
		index := int(math.Floor(float64(len(validUpdates[i])) / 2))
		num, err := strconv.Atoi(validUpdates[i][index])
		if err != nil {
			log.Panicf("could not parse %s in valid update #%d > %#v", validUpdates[i][index], i, validUpdates[i])
		}

		sum += num
	}

	return sum
}

func CheckUpdates(rules RuleMap, updates [][]string) [][]string {
	var validUpdates [][]string
	for i := 0; i < len(updates); i++ {
		if CheckUpdate(rules, updates[i]) {
			validUpdates = append(validUpdates, updates[i])
		}
	}
	return validUpdates
}

func CheckUpdate(rules RuleMap, update []string) bool {
	visited := map[string]bool{}

	for i := 0; i < len(update); i++ {
		curr := update[i]
		visited[curr] = true
		mustBeAfter, ok := rules[curr]
		if !ok {
			// no rule for this page
			continue
		}

		for _, page := range mustBeAfter {
			if visited[page] {
				// mustBeAfter was visited before current but must be after
				return false
			}
		}
	}

	return true
}
