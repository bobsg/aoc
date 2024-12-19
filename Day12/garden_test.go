package main

import (
	"strings"
	"testing"
)

func TestMapRegion(t *testing.T) {
	input := ParseInput(strings.NewReader(shortExample))
	r := NewRegion()
	// The first 'C' in reading order
	start := input[1][2]
	r.Add(start)
	MapRegion(input, r, start)

	if r.Len() != 4 {
		t.Errorf("got %d, want 4 plots in region", r.Len())
	}

	plots := []*Plot{
		NewPlot("C", 2, 1, 1),
		NewPlot("C", 2, 2, 2),
		NewPlot("C", 3, 2, 2),
		NewPlot("C", 3, 3, 1)}
	for _, p := range plots {
		p.Region = r
		assertContainsPlot(t, r, p)
		assertPlotsEquals(t, r, p)
	}

}

func TestMapRegions(t *testing.T) {
	input := ParseInput(strings.NewReader(shortExample))

	got := MapRegions(input)

	if len(got) != 5 {
		t.Errorf("got %d, want %d regions", len(got), 5)
	}
}

func assertContainsPlot(t testing.TB, r *Region, p *Plot) {
	t.Helper()
	if !r.Contains(p) {
		t.Errorf("expected region to contain plot %#v", p)
	}
}

func assertPlotsEquals(t testing.TB, r *Region, p *Plot) {
	t.Helper()
	rp := r.GetPlot(p)
	if !p.Equal(rp) {
		t.Errorf("got plot %#v, want %#v", rp, p)
	}
}

func TestSumFencePriceRegion(t *testing.T) {
	input := ParseInput(strings.NewReader(shortExample))
	r := NewRegion()
	start := input[1][2]
	/* 10 fences, 4 area, price of fence = 10 * 4 = 40
	+-+
	|C|
	+ +-+
	|C C|
	+-+ +
	  |C|
	  +-+
	*/

	r.Add(start)
	MapRegion(input, r, start)
	got := r.SumFencePrice()
	want := 40

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSumFencePrice(t *testing.T) {
	input := ParseInput(strings.NewReader(largerExample))

	got := SumFencePrice(input)
	want := 1930

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestRegion(t *testing.T) {
	t.Run("can add a plot to a region", func(t *testing.T) {
		r := NewRegion()
		p := &Plot{
			X: 0,
			Y: 0,
		}

		r.Add(p)

		if p.Region != r {
			t.Errorf("Expected plot to be added to region")
		}

		if !r.Contains(p) {
			t.Errorf("Expected region to contain plot")
		}

	})
}
