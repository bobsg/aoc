package main

import (
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	a := AntiNodes{
		width:  12,
		height: 12,
		nodes:  []AntiNode{},
	}

	nodes := []AntiNode{
		{
			Freq: 'a',
			Pos:  NewPosition(2, 2),
		},
		{
			Freq: 'b',
			Pos:  NewPosition(0, 0),
		},
		{
			Freq: 'c',
			Pos:  NewPosition(11, 11),
		},
		{
			Freq: 'd',
			Pos:  NewPosition(-1, 1),
		},
		{
			Freq: 'e',
			Pos:  NewPosition(12, 1),
		},
		{
			Freq: 'f',
			Pos:  NewPosition(1, -1),
		},
		{
			Freq: 'g',
			Pos:  NewPosition(1, 12),
		},
	}

	a.AddNodes(nodes)

	if len(a.nodes) != 3 {
		t.Errorf("got %d but expected 3", len(a.nodes))
	}
}

func TestUniquePositions(t *testing.T) {
	a := AntiNodes{
		width:  12,
		height: 12,
		nodes:  []AntiNode{},
	}

	nodes := []AntiNode{
		{
			Freq: 'a',
			Pos:  NewPosition(2, 2), // 1
		},
		{
			Freq: 'b',
			Pos:  NewPosition(0, 0), // 2
		},
		{
			Freq: 'c',
			Pos:  NewPosition(2, 1), // 3
		},
		{
			Freq: 'd',
			Pos:  NewPosition(0, 0), // -
		},
		{
			Freq: 'e',
			Pos:  NewPosition(1, 3), // 4
		},
		{
			Freq: 'f',
			Pos:  NewPosition(2, 2), // -
		},
		{
			Freq: 'g',
			Pos:  NewPosition(2, 1), // -
		},
	}

	a.AddNodes(nodes)
	u := a.UniquePositions()

	if len(u) != 4 {
		t.Errorf("got %d want 4", len(u))
	}
}

func TestUniqueExampleInput(t *testing.T) {
	antennas := ParseInput(strings.NewReader(longInput))

	t.Run("Part1", func(t *testing.T) {
		antiNodes := antennas.CalculateAllAntiNodes()
		u := antiNodes.UniquePositions()

		if len(u) != 14 {
			t.Errorf("got %d want 14", len(u))
		}
	})
	t.Run("Part2", func(t *testing.T) {
		antiNodes := antennas.CalculateAllAntiNodesP2()
		u := antiNodes.UniquePositions()

		if len(u) != 34 {
			t.Errorf("got %d want 14", len(u))
		}
	})

}
