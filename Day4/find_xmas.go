package main

import (
	"strings"
)

var RIGHT = Direction{X: 1, Y: 0}
var LEFT = Direction{X: -1, Y: 0}
var UP = Direction{X: 0, Y: -1}
var DOWN = Direction{X: 0, Y: 1}
var UP_RIGHT = Direction{X: 1, Y: -1}
var DOWN_RIGHT = Direction{X: 1, Y: 1}
var UP_LEFT = Direction{X: -1, Y: -1}
var DOWN_LEFT = Direction{X: -1, Y: 1}

type Pos struct {
	X int
	Y int
}

type Direction struct {
	X int
	Y int
}

func CountXmases(input [][]string) int {
	var sum int
	xses := FindX(input)
	for _, xpos := range xses {
		sum += testAllDirections(input, xpos)
	}

	return sum
}

func testAllDirections(input [][]string, p Pos) int {
	var sum int
	sum += FindXmas(input, p, RIGHT)
	sum += FindXmas(input, p, LEFT)
	sum += FindXmas(input, p, UP)
	sum += FindXmas(input, p, DOWN)
	sum += FindXmas(input, p, UP_RIGHT)
	sum += FindXmas(input, p, DOWN_RIGHT)
	sum += FindXmas(input, p, UP_LEFT)
	sum += FindXmas(input, p, DOWN_LEFT)
	return sum
}

func FindX(input [][]string) []Pos {
	var ret []Pos
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == "X" {
				p := Pos{X: j, Y: i}
				ret = append(ret, p)
			}
		}
	}

	return ret
}

func FindXmas(input [][]string, p Pos, d Direction) int {
	if !isInBounds(input, p, d) {
		return 0
	}

	var matchSlice []string
	for i := 0; i < 4; i++ {
		matchSlice = append(matchSlice, input[p.Y+(i*d.Y)][p.X+(i*d.X)])
	}

	if strings.Join(matchSlice, "") == "XMAS" {
		return 1
	}

	return 0
}

func isInBounds(input [][]string, p Pos, d Direction) bool {
	endX := p.X + (d.X * 3)
	endY := p.Y + (d.Y * 3)

	if endX < 0 || endX >= len(input[p.Y]) {
		return false
	}

	if endY < 0 || endY >= len(input) {
		return false
	}

	return true
}
