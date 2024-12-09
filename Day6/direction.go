package main

type Direction struct {
	X int
	Y int
}

func (p *Direction) Equal(other *Direction) bool {
	return p.X == other.X && p.Y == other.Y
}

var UP = &Direction{X: 0, Y: -1}
var RIGHT = &Direction{X: 1, Y: 0}
var DOWN = &Direction{X: 0, Y: 1}
var LEFT = &Direction{X: -1, Y: 0}
