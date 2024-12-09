package main

import (
	"fmt"
	"log"
)

type Guard struct {
	Pos     *Position
	Dir     *Direction
	History []*Entry
}

type Entry struct {
	Pos *Position
	Dir *Direction
}

func (g *Guard) Move(b Board) {
	newPos := &Position{
		X: g.Pos.X + g.Dir.X,
		Y: g.Pos.Y + g.Dir.Y,
	}

	if !b.isOutOfBounds(newPos) && b.testCollision(newPos) {
		g.Turn()
	} else if b.isOutOfBounds(newPos) {
		g.Pos = newPos
	} else {
		g.Pos = newPos
		g.History = append(g.History, &Entry{Pos: newPos, Dir: g.Dir})
	}
}

func (g *Guard) MoveWithVirtualObstruction(b Board, o *Position) {
	newPos := &Position{
		X: g.Pos.X + g.Dir.X,
		Y: g.Pos.Y + g.Dir.Y,
	}

	if !b.isOutOfBounds(newPos) && b.testCollisionWithVirtualObstruction(newPos, o) {
		g.Turn()
	} else if b.isOutOfBounds(newPos) {
		g.Pos = newPos
	} else {
		g.Pos = newPos
		g.History = append(g.History, &Entry{Pos: newPos, Dir: g.Dir})
	}
}

func (g *Guard) Turn() {
	if g.Dir.Equal(UP) {
		g.Dir = RIGHT
	} else if g.Dir.Equal(RIGHT) {
		g.Dir = DOWN
	} else if g.Dir.Equal(DOWN) {
		g.Dir = LEFT
	} else {
		g.Dir = UP
	}
}

func (g *Guard) DoRound(board Board) {
	// move while not out of bounds,
	for !board.isOutOfBounds(g.Pos) {
		g.Move(board)
	}
}

func (g *Guard) GetUniqueEntries() []*Entry {
	var entries []*Entry
	visited := map[string]bool{}
	for i := 0; i < len(g.History); i++ {
		e := g.History[i]
		key := fmt.Sprintf("%d,%d", e.Pos.X, e.Pos.Y)
		if !visited[key] {
			entries = append(entries, e)
			visited[key] = true
		}
	}
	return entries
}

func (g *Guard) IsInLoop() bool {
	visited := map[string]bool{}
	for i := 0; i < len(g.History); i++ {
		e := g.History[i]
		key := fmt.Sprintf("%d,%d,%d,%d", e.Pos.X, e.Pos.Y, e.Dir.X, e.Dir.Y)
		if visited[key] {
			return true
		}
		visited[key] = true
	}
	return false
}

func (g *Guard) DoRoundWithObstructions(board Board) int {
	found := 0
	count := 0

	g.DoRound(board)
	unique := g.GetUniqueEntries()
	resultChan := make(chan int, len(unique))

	for i := 0; i < len(unique); i++ {
		p := unique[i].Pos
		if board[p.Y][p.X] == "^" {
			continue
		}
		count++
		go DoWalk(board, resultChan, p, i)

	}
	fmt.Printf("Startat %d\n", count)

	for i := 0; i < count-1; i++ {
		found += <-resultChan
	}
	return found
}

func DoWalk(b Board, resultChan chan int, o *Position, id int) {
	found := false
	guard, err := b.FindGuard()
	if err != nil {
		log.Fatal("no guard")
	}

	for !b.isOutOfBounds(guard.Pos) {
		guard.MoveWithVirtualObstruction(b, o)
		if guard.IsInLoop() {
			found = true
			break
		}
	}
	if found {
		resultChan <- 1
	} else {
		resultChan <- 0
	}
}

func copyGuard(g *Guard) Guard {
	guard := Guard{
		Pos:     &Position{X: g.Pos.X, Y: g.Pos.Y},
		Dir:     &Direction{X: g.Dir.X, Y: g.Dir.Y},
		History: []*Entry{},
	}
	return guard
}
