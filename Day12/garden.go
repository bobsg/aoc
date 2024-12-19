package main

import "fmt"

type Plot struct {
	PlotType string
	Region   *Region
	X        int
	Y        int
	Edges    int
}

type Region struct {
	PlotType string
	plots    map[string]*Plot
}

func NewRegion() *Region {
	return &Region{plots: make(map[string]*Plot)}
}

func NewPlot(plotType string, x, y, edges int) *Plot {
	return &Plot{
		PlotType: plotType,
		X:        x,
		Y:        y,
		Edges:    edges,
	}
}

func (r *Region) Add(p *Plot) {
	if r.PlotType != p.PlotType && r.PlotType != "" {
		return
	}
	r.PlotType = p.PlotType
	key := makeKey(p)
	r.plots[key] = p
	p.Region = r
}

func (r *Region) Len() int {
	return len(r.plots)
}

func (r *Region) Contains(p *Plot) bool {
	key := makeKey(p)
	_, ok := r.plots[key]
	return ok
}

func (r *Region) GetPlot(p *Plot) *Plot {
	rp, ok := r.plots[makeKey(p)]
	if !ok {
		return nil
	}

	return rp
}

func makeKey(p *Plot) string {
	key := fmt.Sprintf("%d,%d", p.X, p.Y)
	return key
}

func (p *Plot) Equal(other *Plot) bool {
	return p.X == other.X && p.Y == other.Y && p.PlotType == other.PlotType && p.Region == other.Region && p.Edges == other.Edges
}

func (p *Plot) HasRegion() bool {
	return p.Region != nil
}

func (r *Region) SumFencePrice() int {
	area := 0
	perimeter := 0
	for _, v := range r.plots {
		perimeter += 4 - v.Edges
		area++
	}
	return area * perimeter
}

func SumFencePrice(garden [][]*Plot) int {
	regions := MapRegions(garden)
	sum := 0
	for _, r := range regions {
		sumRegion := r.SumFencePrice()
		sum += sumRegion
	}
	return sum
}

func MapRegions(garden [][]*Plot) []*Region {
	var regions []*Region
	for y := 0; y < len(garden); y++ {
		for x := 0; x < len(garden[y]); x++ {
			plot := garden[y][x]
			if plot.Region != nil {
				continue
			}

			r := NewRegion()
			r.Add(plot)
			regions = append(regions, r)

			MapRegion(garden, r, plot)
		}
	}
	return regions
}

func MapRegion(garden [][]*Plot, region *Region, p *Plot) {
	plotType := p.PlotType
	x := p.X
	y := p.Y
	if !isInBounds(garden, p) {
		return
	}

	if !region.Contains(p) {
		region.Add(p)
	}

	if p.X > 0 && garden[y][x-1].PlotType == plotType {
		neighbour := garden[y][x-1]
		p.Edges++
		MapPlotIfNotDone(garden, region, neighbour)
	}

	if p.X < len(garden[0])-1 && garden[y][x+1].PlotType == plotType {
		neighbour := garden[y][x+1]
		p.Edges++
		MapPlotIfNotDone(garden, region, neighbour)
	}

	if p.Y > 0 && garden[y-1][x].PlotType == plotType {
		neighbour := garden[y-1][x]
		p.Edges++
		MapPlotIfNotDone(garden, region, neighbour)
	}

	if p.Y < len(garden)-1 && garden[y+1][x].PlotType == plotType {
		neighbour := garden[y+1][x]
		p.Edges++
		MapPlotIfNotDone(garden, region, neighbour)
	}

}

func MapPlotIfNotDone(garden [][]*Plot, r *Region, p *Plot) {
	if p.Region == nil {
		MapRegion(garden, r, p)
	}
}

func isInBounds(garden [][]*Plot, p *Plot) bool {
	if len(garden) == 0 {
		return false
	}
	return p.Y >= 0 && p.Y < len(garden) && p.X >= 0 && p.X < len(garden[0])
}
