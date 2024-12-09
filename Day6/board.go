package main

import (
	"errors"
	"slices"
)

type Board [][]string

func (b Board) FindGuard() (*Guard, error) {
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			if b[i][j] == "^" {
				return &Guard{
					Pos: &Position{X: j, Y: i},
					Dir: UP,
					History: []*Entry{
						{Pos: &Position{X: j, Y: i}, Dir: UP},
					},
				}, nil
			}
		}
	}

	return nil, errors.New("no guard found")
}

func (b Board) testCollision(p *Position) bool {
	return b[p.Y][p.X] == "#"
}

func (b Board) testCollisionWithVirtualObstruction(p *Position, o *Position) bool {
	return p.Equal(o) || b[p.Y][p.X] == "#"
}

func (b Board) isOutOfBounds(p *Position) bool {
	return p.Y < 0 || p.Y >= len(b) || p.X < 0 || p.X >= len(b[p.Y])
}

func copyBoard(board Board) Board {
	boardCopy := Board{}
	for i := 0; i < len(board); i++ {
		boardCopy = append(boardCopy, slices.Clone(board[i]))
	}
	return boardCopy
}
