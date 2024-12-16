package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestFindTrailHeads(t *testing.T) {
	input := TopoMap{
		{0, 1, 2, 3},
		{1, 2, 3, 4},
		{8, 7, 0, 5},
		{9, 8, 7, 6},
	}

	got := input.FindTrailheads()
	want := []*TrailHead{{X: 0, Y: 0}, {X: 2, Y: 2}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}

func TestSumAllTrails(t *testing.T) {
	topoMap := ParseInput(strings.NewReader(longInput))
	sum := topoMap.SumAllTrails()

	if sum != 36 {
		t.Errorf("got %d, want 36", sum)
	}
}

func TestFindTrail(t *testing.T) {
	topoMap := TopoMap{
		{0, 1, 2, 3},
		{1, 2, 3, 4},
		{8, 7, 6, 5},
		{9, 8, 7, 6},
	}

	heads := topoMap.FindTrails()

	if len(heads) != 1 {
		t.Fatalf("expected 1 trail")
	}

	if heads[0].Score != 1 {
		t.Errorf("expected score 1")
	}
}

func TestSumDistinctTrails(t *testing.T) {
	topoMap := ParseInput(strings.NewReader(longInput))
	sum := topoMap.SumDistinctTrails()

	if sum != 81 {
		t.Errorf("got %d, want 81", sum)
	}
}

func TestFindDistinctTrailsForHead(t *testing.T) {
	topoMap := ParseInput(strings.NewReader(longInput))
	head := &TrailHead{
		X:     2,
		Y:     0,
		Score: 0,
	}

	heads := topoMap.FindDistinctTrailsForHead(head)

	if len(heads) != 20 {
		t.Fatalf("expected 20 trails")
	}
}

func TestFindNextSteps(t *testing.T) {
	topoMap := TopoMap{
		{0, 1, 2, 3},
		{1, 2, 3, 4},
		{8, 7, 6, 5},
		{9, 8, 7, 6},
	}
	t.Run("one hit", func(t *testing.T) {
		step := &TrailPart{
			X:    2,
			Y:    1,
			Step: 3,
		}

		got := topoMap.FindNextSteps(step)
		want := []*TrailPart{&TrailPart{
			X:    3,
			Y:    1,
			Step: 4,
		}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v, want %#v", got, want)
		}
	})
	t.Run("one hit", func(t *testing.T) {
		step := &TrailPart{
			X:    2,
			Y:    2,
			Step: 6,
		}

		got := topoMap.FindNextSteps(step)
		want := []*TrailPart{
			{
				X:    1,
				Y:    2,
				Step: 7,
			}, {
				X:    2,
				Y:    3,
				Step: 7,
			},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v, want %#v", got, want)
		}
	})

}
