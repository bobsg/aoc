package main

type Position struct {
	X int
	Y int
}

func (p *Position) Equal(other *Position) bool {
	return p.X == other.X && p.Y == other.Y
}
