package graph

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
