package graph

import (
	"fmt"

	"github.com/emicklei/dot"
)

type Elem struct {
	value Node
	next  *Elem
}

type AdjLst []*Elem

func NewGraphAdjLst(nodes int, edges ...*Edge) AdjLst {
	var g AdjLst = make([]*Elem, nodes+1)

	for _, edge := range edges {
		g.AddEdge(edge.From, edge.To)
	}

	return g
}

func (g AdjLst) AddEdge(from Node, to Node) {
	g[from] = &Elem{to, g[from]}
}

func (g AdjLst) Edges() []Edge {
	es := make([]Edge, 0)

	for i, elem := range g {
		for elem != nil {
			es = append(es, Edge{Node(i), elem.value})
			elem = elem.next
		}
	}

	return es
}

func (g AdjLst) ToDot() dot.Graph {
	dg := dot.NewGraph(dot.Directed)

	for i, elem := range g {
		if i == 0 {
			continue
		}

		f := dg.Node(fmt.Sprintf("%d", i))

		for elem != nil {
			t := dg.Node(fmt.Sprintf("%d", elem.value))
			dg.Edge(f, t)

			elem = elem.next
		}
	}

	return *dg
}

// O(|V| + |E|)
func (g AdjLst) IsAcyclic() bool {
	// Eingangsgrade berechnen
	indeg := make(map[Node]int)

	// O(|V| + |E|)
	for _, elem := range g {
		for current := elem; current != nil; current = current.next {
			indeg[current.value]++
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

	// topologische Sortierung berechnen
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
		for current := g[v]; current != nil; current = current.next {
			indeg[current.value]--

			if indeg[current.value] == 0 {
				indegZero = append(indegZero, current.value)
			}
		}
	}

	// wenn jeder Knoten eine laufende Nummer bekommen hat,
	// ist der Graph azyklisch.
	return seqNo == len(g)-1
}

func (g AdjLst) TransitiveClosure() Graph {
	panic("not implemented")
}
