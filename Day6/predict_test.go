package main

import (
	"reflect"
	"testing"
)

func TestFindGuard(t *testing.T) {
	got, err := bigBoard.FindGuard()
	if err != nil {
		t.Fatal("got error but did not expect one")
	}

	want := &Guard{
		Pos: &Position{X: 4, Y: 6},
		Dir: &Direction{X: 0, Y: -1},
	}

	if !got.Pos.Equal(want.Pos) {
		t.Errorf("got %#v want %#v", got, want)
	}
}

func TestMoveGuardCanMove(t *testing.T) {
	board := Board{{"."}, {"^"}}
	guard, err := board.FindGuard()
	if err != nil {
		t.Fatal("got error")
	}
	guard.Move(board)

	if guard.Pos.Y != 0 && guard.Pos.X != 0 {
		t.Errorf("got %d, want 0", guard.Pos.Y)
	}
	if len(guard.History) != 2 {
		t.Errorf("guard did not record move")
	}
	if !reflect.DeepEqual(guard.History[0], &Position{X: 0, Y: 1}) {
		t.Errorf("got first move as %#v want %#v", guard.History[0], Position{X: 0, Y: 1})
	}

	if !reflect.DeepEqual(guard.History[1], &Position{X: 0, Y: 0}) {
		t.Errorf("got second move as %#v want %#v", guard.History[1], Position{X: 0, Y: 0})
	}
}

func TestMoveGuardObstructed(t *testing.T) {
	board := Board{{"#"}, {"^"}}
	guard, err := board.FindGuard()
	if err != nil {
		t.Fatal("got error")
	}
	guard.Move(board)

	if guard.Pos.Y != 1 && guard.Pos.X != 0 {
		t.Errorf("got %d, want 0", guard.Pos.Y)
	}
	if !reflect.DeepEqual(guard.Dir, RIGHT) {
		t.Errorf("got %#v, want %#v", guard.Dir, RIGHT)
	}
	if len(guard.History) != 1 {
		t.Errorf("guard recorded move")
	}
}

func TestGuardDoRound(t *testing.T) {
	guard, err := bigBoard.FindGuard()
	if err != nil {
		t.Fatal("did not find guard")
	}

	guard.DoRound(bigBoard)

	if len(guard.GetUniqueEntries()) != 41 {
		t.Errorf("got %d, want %d", len(guard.GetUniqueEntries()), 41)
	}
}

func TestSimulateNewObstruction(t *testing.T) {
	guard, err := bigBoard.FindGuard()
	if err != nil {
		t.Fatal("did not find guard")
	}

	got := guard.DoRoundWithObstructions(bigBoard)

	if got != 6 {
		t.Errorf("got %d, want 6", got)
	}
}

func TestDoWalk(t *testing.T) {
	resultChan := make(chan int)

	obstruction := &Position{X: 3, Y: 6}
	go DoWalk(bigBoard, resultChan, obstruction, 0)
	result := <-resultChan

	if result != 1 {
		t.Errorf("got %d want 1", result)
	}
}
