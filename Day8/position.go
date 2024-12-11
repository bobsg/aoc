package main

type Position struct {
	X int
	Y int
}

type Distance struct {
	X int
	Y int
}

func (p *Position) Equals(other Position) bool {
	return p.X == other.X && p.Y == other.Y
}

func NewPosition(x, y int) Position {
	return Position{X: x, Y: y}
}

func (p *Position) DistanceTo(other Position) Distance {
	d := Distance{}
	d.X = other.X - p.X
	d.Y = other.Y - p.Y
	return d
}

func NewDistance(x, y int) Distance {
	return Distance{X: x, Y: y}
}

func (p *Position) NewPositionAtDistance(d Distance) Position {
	newPosition := Position{}
	newPosition.X = p.X + d.X
	newPosition.Y = p.Y + d.Y
	return newPosition
}

func (d *Distance) Invert() Distance {
	newDistance := Distance{}
	newDistance.X = d.X * -1
	newDistance.Y = d.Y * -1
	return newDistance
}
