package graph

type Node int

type Graph interface {
	AddEdge(from Node, to Node)
}