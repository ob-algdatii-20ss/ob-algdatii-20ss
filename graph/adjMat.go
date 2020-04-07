package graph

import (
	"fmt"

	"github.com/emicklei/dot"
)

type AdjMat [][]bool

func NewGraphAdjMat(nodes int, edges ...*Edge) AdjMat {
	var g AdjMat = make([][]bool, nodes+1)
	for c := range g {
		g[c] = make([]bool, nodes+1)
	}

	for _, edge := range edges {
		g.AddEdge(edge.From, edge.To)
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

// O(|V| + |E|)
func (g AdjMat) IsAcyclic() bool {
	// Eingangsgrade berechnen
	indeg := make(map[Node]int)

	// O(|V| + |E|)
	for _, elem := range g {
		for node, isEdge := range elem {
			if isEdge {
				indeg[Node(node)]++
			}
		}
	}

	// übernehme alle Knoten mit Eingangsgrad 0 in "stack"
	indegZero := make([]Node, 0)

	// O(|V|)
	for i := range g {
		if i == 0 {
			continue
		}

		n := Node(i)

		if indeg[n] == 0 {
			indegZero = append(indegZero, n)
		}
	}

	// topologische Ordnung berechnen
	seqNo := 0
	ord := make(map[Node]int)

	// O(|V| + |E|)
	// maximal durch alle Knoten
	// jeder Pfeil insgesamt über alle inneren Schleifen je einmal
	for len(indegZero) > 0 {
		// nehme den nächsten vom Stack
		lastIdx := len(indegZero) - 1
		v := indegZero[lastIdx]
		indegZero = indegZero[:lastIdx]
		seqNo++
		ord[v] = seqNo

		// dekrementiere Eingangsgrade aller Knoten die Endknoten
		// eines Pfeiles sind, von dem v der Anfangsknoten ist
		for i, isEdge := range g[v] {
			node := Node(i)

			if isEdge {
				indeg[node]--
				if indeg[node] == 0 {
					indegZero = append(indegZero, node)
				}
			}
		}
	}

	// wenn jeder Knoten eine laufende Nummer bekommen hat,
	// ist der Graph azyklisch.
	return seqNo == len(g)-1
}

func (g AdjMat) TransitiveClosure() Graph {
	closure := NewGraphAdjMat(len(g) - 1)

	for i := range g {
		copy(closure[i], g[i])
	}

	n := Node(len(closure) - 1)

	for i := Node(1); i <= n; i++ {
		closure.AddEdge(i, i)
	}

	for j := Node(1); j <= n; j++ {
		for i := Node(1); i <= n; i++ {
			if closure[i][j] {
				for k := Node(1); k <= n; k++ {
					if closure[j][k] {
						closure.AddEdge(i, k)
					}
				}
			}
		}
	}

	return closure
}
