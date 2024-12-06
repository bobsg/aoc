package main

import (
	"log"
	"math"
	"slices"
	"strconv"
)

func GetSumFixedInvalids(rules RuleMap, updates [][]string) int {
	invalidUpdates := GetInvalidUpdates(rules, updates)
	nowValidUpdates := SortAllInvalid(rules, invalidUpdates)

	var sum int
	for i := 0; i < len(nowValidUpdates); i++ {
		index := int(math.Floor(float64(len(nowValidUpdates[i])) / 2))
		num, err := strconv.Atoi(nowValidUpdates[i][index])
		if err != nil {
			log.Panicf("could not parse %s in valid update #%d > %#v", nowValidUpdates[i][index], i, nowValidUpdates[i])
		}

		sum += num
	}

	return sum
}

func SortAllInvalid(rules RuleMap, updates [][]string) [][]string {
	var result [][]string
	for i := 0; i < len(updates); i++ {
		sorted := SortInvalidBuiltIn(rules, updates[i])
		result = append(result, sorted)
	}
	return result
}

func GetInvalidUpdates(rules RuleMap, updates [][]string) [][]string {
	var invalidUpdates [][]string
	for i := 0; i < len(updates); i++ {
		if !CheckUpdate(rules, updates[i]) {
			invalidUpdates = append(invalidUpdates, updates[i])
		}
	}
	return invalidUpdates
}

func SortInvalid(rules RuleMap, update []string) []string {
	visited := map[string]int{}
	var result = slices.Clone(update)
	for i := 0; i < len(result); i++ {
		curr := result[i]
		visited[curr] = i
		mustBeAfter, ok := rules[curr]
		if !ok {
			// no rule for this page
			continue
		}

		lowestIndex := -1
		for _, page := range mustBeAfter {
			index, ok := visited[page]
			if ok {
				if lowestIndex == -1 {
					lowestIndex = index
				}
				// mustBeAfter was visited before current but must be after
				lowestIndex = min(lowestIndex, index)
			}
		}

		if lowestIndex != -1 {
			newResult := []string{}
			newResult = append(newResult, result[:lowestIndex]...)
			newResult = append(newResult, curr)
			newResult = append(newResult, result[lowestIndex:i]...) // don't add current again
			newResult = append(newResult, result[i+1:]...)
			result = newResult
			// reset
			i = 0
			visited = map[string]int{}
		}

	}

	return result
}

func (r RuleMap) SortFunc(a, b string) int {
	if r.ShouldBeBefore(a, b) {
		return -1
	} else if r.ShouldBeBefore(b, a) {
		return 1
	}

	return 0
}

func (r RuleMap) ShouldBeBefore(a, b string) bool {
	vals, ok := r[a]
	if ok {
		return slices.Contains(vals, b)
	}
	return false
}

func SortInvalidBuiltIn(rules RuleMap, update []string) []string {
	newUpdate := slices.Clone(update)
	slices.SortFunc(newUpdate, rules.SortFunc)
	return newUpdate
}
