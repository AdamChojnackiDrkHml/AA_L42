package mis

import (
	"math/rand"
	"time"
)

type Color int

const (
	Red Color = iota
	Yellow
	White
	Black
)

type Node struct {
	Id         uint
	Color      Color
	Neighbours map[uint]struct{}
}

func Node_NewNode(id uint) *Node {
	return &Node{
		Id:         id,
		Color:      Yellow,
		Neighbours: make(map[uint]struct{}),
	}
}

func (node *Node) UpdateColor(independent map[uint]struct{}) {
	_, isIndependet := independent[node.Id]
	independentNeighbour := node.anyIndependentNeighbour(independent)

	if isIndependet {
		if independentNeighbour {
			node.Color = Red
			return
		}

		node.Color = Black
		return
	}

	if independentNeighbour {
		node.Color = White
		return
	}

	node.Color = Yellow

}

func (node *Node) anyIndependentNeighbour(independent map[uint]struct{}) bool {
	for neigh := range node.Neighbours {
		if _, in := independent[neigh]; in {
			return true
		}
	}

	return false
}

type Graph struct {
	N     uint
	Nodes []Node
}

func Graph_NewGraph(n uint) *Graph {
	g := &Graph{
		N:     n,
		Nodes: make([]Node, n),
	}

	for i := range g.Nodes {
		g.Nodes[i] = *Node_NewNode(uint(i))
	}

	return g
}

func (g *Graph) AddEdge(u, v uint) {
	g.Nodes[u].Neighbours[v] = struct{}{}
	g.Nodes[v].Neighbours[u] = struct{}{}
}

func Graph_NewRandGraph(n uint, p float64) *Graph {
	g := Graph_NewGraph(n)
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	for u := uint(0); u < n; u++ {
		for v := u + 1; v < n; v++ {
			if r.Float64() < p {
				g.AddEdge(u, v)
			}
		}
	}

	independent := make(map[uint]struct{})
	for i := range g.Nodes {
		g.Nodes[i].UpdateColor(independent)
	}

	return g
}
