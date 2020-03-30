package graph

import "github.com/emicklei/dot"

type Node int

type Edge struct {
	From Node
	To   Node
}

type Graph interface {
	AddEdge(from Node, to Node)
	Edges() []Edge
	ToDot() dot.Graph
	IsAcyclic() bool
}
