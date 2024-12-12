package main

type frequency rune

type Antennas struct {
	width    int
	height   int
	antennas map[frequency][]Antenna
}

type Antenna struct {
	Freq frequency
	Pos  Position
}

type pair struct {
	A Antenna
	B Antenna
}

type Pairings map[frequency][]pair

func (a *Antennas) createAllPairs() []pair {
	result := []pair{}
	for freq := range a.antennas {
		p := a.createPairs(freq)
		result = append(result, p...)
	}
	return result
}

func (a *Antennas) CalculateAllAntiNodes() AntiNodes {
	pairs := a.createAllPairs()
	nodes := NewAntiNodes(a.width, a.height)
	for _, p := range pairs {
		nodes.AddNodes(p.CalculateAntiNodes())
	}
	return nodes
}

func (a *Antennas) CalculateAllAntiNodesP2() AntiNodes {
	pairs := a.createAllPairs()
	nodes := NewAntiNodes(a.width, a.height)
	for _, p := range pairs {
		nodes.AddNodes(p.CalculateAntiNodesP2(a.width, a.height))
	}
	return nodes
}

// createPairs creates all possible pairs for the given frequency.
func (a *Antennas) createPairs(f frequency) []pair {
	antennas, ok := a.antennas[f]
	if !ok || len(antennas) == 0 || len(antennas) == 1 {
		return []pair{}
	}
	if len(antennas) == 2 {
		return []pair{
			{A: antennas[0], B: antennas[1]},
		}
	}

	pairs := []pair{}
	for i := 1; i < len(antennas); i++ {
		for j := i; j < len(antennas); j++ {
			p := pair{
				A: antennas[i-1],
				B: antennas[j],
			}
			pairs = append(pairs, p)
		}
	}
	return pairs
}

func (p *pair) CalculateAntiNodes() []AntiNode {
	distance := p.A.Pos.DistanceTo(p.B.Pos)
	inverted := distance.Invert()

	p1 := p.B.Pos.NewPositionAtDistance(distance)
	p2 := p.A.Pos.NewPositionAtDistance(inverted)

	a1 := NewAntiNode(p1, p.A.Freq)
	a2 := NewAntiNode(p2, p.A.Freq)

	return []AntiNode{a1, a2}
}

func (p *pair) CalculateAntiNodesP2(width, height int) []AntiNode {
	result := []AntiNode{}
	freq := p.A.Freq
	distance := p.A.Pos.DistanceTo(p.B.Pos)
	inverted := distance.Invert()

	foundNode := true
	for x := 0; foundNode; x++ {
		foundNode = false
		d := distance.Mul(x)
		i := inverted.Mul(x)

		p1 := p.B.Pos.NewPositionAtDistance(d)
		p2 := p.A.Pos.NewPositionAtDistance(i)

		if isInBounds(width, height, p1) {
			result = append(result, NewAntiNode(p1, freq))
			foundNode = true
		}
		if isInBounds(width, height, p2) {
			result = append(result, NewAntiNode(p2, freq))
			foundNode = true
		}
	}

	return result
}

func isInBounds(width, height int, p Position) bool {
	return p.X >= 0 && p.X < width && p.Y >= 0 && p.Y < height
}
