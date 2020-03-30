package graph

import (
	"fmt"

	"github.com/emicklei/dot"
)

type AdjMat [][]bool

func NewGraphAdjMat(nodes int) AdjMat {
	g := make([][]bool, nodes+1)
	for c := range g {
		g[c] = make([]bool, nodes+1)
	}

	return g
}

func (g AdjMat) AddEdge(from Node, to Node) {
	g[from][to] = true
}

func (g AdjMat) Edges() []Edge {
	es := make([]Edge, 0)

	for i, col := range g {
		if i == 0 {
			continue
		}

		for j, edge := range col {
			if j == 0 {
				continue
			}

			if edge {
				es = append(es, Edge{Node(i), Node(j)})
			}
		}
	}

	return es
}

func (g AdjMat) ToDot() dot.Graph {
	dg := dot.NewGraph(dot.Directed)

	for i, col := range g {
		if i == 0 {
			continue
		}

		f := dg.Node(fmt.Sprintf("%d", i))

		for j, edge := range col {
			if i == 0 {
				continue
			}

			if edge {
				t := dg.Node(fmt.Sprintf("%d", j))
				dg.Edge(f, t)
			}
		}
	}

	return *dg
}
