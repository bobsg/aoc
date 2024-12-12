package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestCreatePairs(t *testing.T) {
	got := antennasInput.createPairs('A')
	want := []pair{
		{Antenna{'A', Position{X: 6, Y: 5}}, Antenna{'A', Position{X: 8, Y: 8}}},
		{Antenna{'A', Position{X: 6, Y: 5}}, Antenna{'A', Position{X: 9, Y: 9}}},
		{Antenna{'A', Position{X: 8, Y: 8}}, Antenna{'A', Position{X: 9, Y: 9}}},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}

func TestCalculateAllAntiNodes(t *testing.T) {
	two := ParseInput(strings.NewReader(twoAntennas))
	got := two.CalculateAllAntiNodes()
	want := []AntiNode{
		{
			Freq: 'a',
			Pos:  NewPosition(6, 7),
		},
		{
			Freq: 'a',
			Pos:  NewPosition(3, 1),
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}

func TestCreateAllPairs(t *testing.T) {
	got := antennasInput.createAllPairs()
	want := []pair{
		{Antenna{'A', Position{X: 6, Y: 5}}, Antenna{'A', Position{X: 8, Y: 8}}},
		{Antenna{'A', Position{X: 6, Y: 5}}, Antenna{'A', Position{X: 9, Y: 9}}},
		{Antenna{'A', Position{X: 8, Y: 8}}, Antenna{'A', Position{X: 9, Y: 9}}},
		{Antenna{'0', Position{X: 8, Y: 1}}, Antenna{'0', Position{X: 5, Y: 2}}},
		{Antenna{'0', Position{X: 8, Y: 1}}, Antenna{'0', Position{X: 7, Y: 3}}},
		{Antenna{'0', Position{X: 8, Y: 1}}, Antenna{'0', Position{X: 4, Y: 4}}},
		{Antenna{'0', Position{X: 5, Y: 2}}, Antenna{'0', Position{X: 7, Y: 3}}},
		{Antenna{'0', Position{X: 5, Y: 2}}, Antenna{'0', Position{X: 4, Y: 4}}},
		{Antenna{'0', Position{X: 7, Y: 3}}, Antenna{'0', Position{X: 4, Y: 4}}},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}

var antennasInput = Antennas{
	width:  12,
	height: 12,
	antennas: map[frequency][]Antenna{
		'0': {
			{'0', Position{X: 8, Y: 1}},
			{'0', Position{X: 5, Y: 2}},
			{'0', Position{X: 7, Y: 3}},
			{'0', Position{X: 4, Y: 4}},
		},
		'A': {
			{'A', Position{X: 6, Y: 5}},
			{'A', Position{X: 8, Y: 8}},
			{'A', Position{X: 9, Y: 9}},
		},
	},
}

func TestCalculateAntiNodes(t *testing.T) {
	p := pair{
		A: Antenna{
			Freq: 'A',
			Pos:  NewPosition(2, 2),
		},
		B: Antenna{
			Freq: 'A',
			Pos:  NewPosition(4, 5),
		},
	}

	t.Run("calculate antinodes p1", func(t *testing.T) {
		got := p.CalculateAntiNodes()
		want := []AntiNode{
			{
				Freq: 'A',
				Pos:  NewPosition(6, 8),
			},
			{
				Freq: 'A',
				Pos:  NewPosition(0, -1),
			},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v, want %#v", got, want)
		}
	})

	t.Run("calculate antinodes p2", func(t *testing.T) {
		p := pair{
			A: Antenna{
				Freq: '0',
				Pos:  NewPosition(0, 0),
			},
			B: Antenna{
				Freq: '0',
				Pos:  NewPosition(3, 1),
			},
		}
		got := p.CalculateAntiNodesP2(10, 10)
		want := []AntiNode{
			{
				Freq: '0',
				Pos:  NewPosition(3, 1),
			},
			{
				Freq: '0',
				Pos:  NewPosition(0, 0),
			},

			{
				Freq: '0',
				Pos:  NewPosition(6, 2),
			},
			{
				Freq: '0',
				Pos:  NewPosition(9, 3),
			},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v, want %#v", got, want)
		}
	})
}
