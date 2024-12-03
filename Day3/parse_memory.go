package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
)

func ParseMemory(input string) ([][]int, error) {
	r := regexp.MustCompile(`mul\((?P<X>\d{1,3}),(?P<Y>\d{1,3})\)`)
	subMatches := r.FindAllStringSubmatch(input, -1)
	var ret [][]int
	for i, match := range subMatches {
		x, err := strconv.Atoi(match[1])
		if err != nil {
			return nil, fmt.Errorf("could not parse int from match %d : %v", i, match[1])
		}
		y, err := strconv.Atoi(match[2])
		if err != nil {
			return nil, fmt.Errorf("could not parse int from match %d : %v", i, match[2])
		}
		ret = append(ret, []int{x, y})
	}

	return ret, nil
}

func ParseMemoryWithConditionals(input string) ([]Command, error) {
	r := regexp.MustCompile(`mul\((?P<X>\d{1,3}),(?P<Y>\d{1,3})\)`)
	dos := regexp.MustCompile(`do\(\)`)
	donts := regexp.MustCompile(`don't\(\)`)
	subMatches := r.FindAllStringSubmatch(input, -1)
	indices := r.FindAllStringSubmatchIndex(input, -1)

	ret := []Command{}
	for i, match := range subMatches {
		x, err := strconv.Atoi(match[1])
		if err != nil {
			return nil, fmt.Errorf("could not parse int from match %d : %v", i, match[1])
		}
		y, err := strconv.Atoi(match[2])
		if err != nil {
			return nil, fmt.Errorf("could not parse int from match %d : %v", i, match[2])
		}
		pos := indices[i][0]
		cmd := Command{
			Pos:  pos,
			Type: MUL,
			X:    x,
			Y:    y,
		}

		ret = append(ret, cmd)
	}

	dosIndices := dos.FindAllStringSubmatchIndex(input, -1)
	for _, do := range dosIndices {
		cmd := Command{
			Pos:  do[0],
			Type: DO,
		}
		ret = append(ret, cmd)
	}

	dontsIndices := donts.FindAllStringSubmatchIndex(input, -1)
	for _, dont := range dontsIndices {
		cmd := Command{
			Pos:  dont[0],
			Type: DONT,
		}
		ret = append(ret, cmd)
	}

	slices.SortFunc(ret, sortFunc)
	return ret, nil
}

func sortFunc(i, j Command) int {
	if i.Pos < j.Pos {
		return -1
	}
	return 1
}

type CmdType int

const (
	MUL CmdType = iota
	DO
	DONT
)

type Command struct {
	Pos  int
	Type CmdType
	X    int
	Y    int
}
