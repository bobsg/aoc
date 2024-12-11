package main

import "fmt"

type AntiNode struct {
	Freq frequency
	Pos  Position
}

type AntiNodes struct {
	width  int
	height int
	nodes  []AntiNode
}

func NewAntiNodes(width, height int) AntiNodes {
	return AntiNodes{
		width:  width,
		height: height,
		nodes:  []AntiNode{},
	}
}

func (a *AntiNodes) AddNodes(nodes []AntiNode) {
	for _, n := range nodes {
		a.Add(n)
	}
}

func (a *AntiNodes) Add(node AntiNode) bool {
	if node.Pos.X < 0 || node.Pos.X >= a.width {
		return false
	}
	if node.Pos.Y < 0 || node.Pos.Y >= a.height {
		return false
	}
	a.nodes = append(a.nodes, node)
	return true
}

func (a *AntiNodes) UniquePositions() []Position {
	visited := map[string]bool{}
	result := []Position{}
	for _, n := range a.nodes {
		k := fmt.Sprintf("%d,%d", n.Pos.X, n.Pos.Y)
		if !visited[k] {
			result = append(result, n.Pos)
		}
		visited[k] = true
	}

	return result
}

func NewAntiNode(p Position, freq frequency) AntiNode {
	return AntiNode{
		Freq: freq,
		Pos:  p,
	}
}
