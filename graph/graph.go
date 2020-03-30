package graph

import "github.com/emicklei/dot"

type Node int

type Edge struct {
	from Node
	to   Node
}

type Graph interface {
	AddEdge(from Node, to Node)
	Edges() []Edge
	ToDot() dot.Graph
}
