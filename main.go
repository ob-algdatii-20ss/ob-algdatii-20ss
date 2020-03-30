package main

import (
	"fmt"

	"github.com/ob-algdatii-20ss/ob-algdatii-20ss/graph"
)

func main() {
	var g graph.Graph = graph.NewGraphAdjMat(9)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 7)
	g.AddEdge(4, 6)
	g.AddEdge(5, 4)
	g.AddEdge(6, 1)
	g.AddEdge(6, 5)
	g.AddEdge(6, 6)
	g.AddEdge(7, 5)
	g.AddEdge(9, 8)

	fmt.Println(g.ToDot().String())
}
