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
