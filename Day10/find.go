package main

import (
	"fmt"
	"maps"
	"slices"
)

type TrailHead struct {
	X     int
	Y     int
	Score int
}

type TrailPart struct {
	X    int
	Y    int
	Step int // 0 - 9
}

type TopoMap [][]int

func (t TopoMap) FindNextSteps(tp *TrailPart) []*TrailPart {
	found := []*TrailPart{}
	height := len(t)
	width := len(t[0])
	nextStep := tp.Step + 1
	x := tp.X
	y := tp.Y

	if x != 0 && t[y][x-1] == nextStep {
		found = append(found, &TrailPart{X: x - 1, Y: y, Step: nextStep})
	}
	if x != width-1 && t[y][x+1] == nextStep {
		found = append(found, &TrailPart{X: x + 1, Y: y, Step: nextStep})
	}
	if y != 0 && t[y-1][x] == nextStep {
		found = append(found, &TrailPart{X: x, Y: y - 1, Step: nextStep})
	}
	if y != height-1 && t[y+1][x] == nextStep {
		found = append(found, &TrailPart{X: x, Y: y + 1, Step: nextStep})
	}
	return found
}

func (t TopoMap) SumAllTrails() int {
	heads := t.FindTrails()
	sum := 0
	for _, h := range heads {
		sum += h.Score
	}
	return sum
}

func (t TopoMap) FindTrailheads() []*TrailHead {
	var result []*TrailHead
	for y := 0; y < len(t); y++ {
		for x := 0; x < len(t[y]); x++ {
			if t[y][x] == 0 {
				result = append(result, &TrailHead{X: x, Y: y})
			}
		}
	}

	return result
}

func (t TopoMap) FindTrails() []*TrailHead {
	heads := t.FindTrailheads()
	for _, h := range heads {
		ends := t.FindTrailsForHead(h)
		h.Score = len(ends)
	}

	return heads
}

func (t TopoMap) FindTrailsForHead(head *TrailHead) []*TrailPart {
	found := map[string]*TrailPart{}
	start := &TrailPart{
		X:    head.X,
		Y:    head.Y,
		Step: 0,
	}
	found[makeKey(start)] = start
	for !allAtEnd(found) {
		prev := found
		found = map[string]*TrailPart{}
		for _, p := range prev {
			steps := t.FindNextSteps(p)
			addSteps(found, steps)
		}
	}

	return slices.Collect(maps.Values(found))
}

func (t TopoMap) FindDistinctTrailsForHead(head *TrailHead) []*TrailPart {
	found := []*TrailPart{}
	start := &TrailPart{
		X:    head.X,
		Y:    head.Y,
		Step: 0,
	}
	found = append(found, start)
	for !allAtEndDistinct(found) {
		prev := found
		found = []*TrailPart{}
		for _, p := range prev {
			steps := t.FindNextSteps(p)
			found = append(found, steps...)
		}
	}

	return found
}

func (t TopoMap) SumDistinctTrails() int {
	heads := t.FindTrailheads()
	sum := 0
	for _, h := range heads {
		ends := t.FindDistinctTrailsForHead(h)
		h.Score = len(ends)
		sum += h.Score
	}

	return sum
}

func addSteps(dst map[string]*TrailPart, steps []*TrailPart) {
	for _, s := range steps {
		dst[makeKey(s)] = s
	}
}

func makeKey(t *TrailPart) string {
	return fmt.Sprintf("%d,%d", t.X, t.Y)
}

func allAtEnd(parts map[string]*TrailPart) bool {
	for _, v := range parts {
		if v.Step != 9 {
			return false
		}
	}
	return true
}

func allAtEndDistinct(parts []*TrailPart) bool {
	for _, v := range parts {
		if v.Step != 9 {
			return false
		}
	}
	return true
}
