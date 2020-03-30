package main

import (
	"fmt"

	"github.com/ob-algdatii-20ss/ob-algdatii-20ss/graph"
)

func main() {
	var g graph.Graph = graph.NewGraphAdjLst(9,
		&graph.Edge{From: 1, To: 2},
		&graph.Edge{From: 1, To: 3},
		&graph.Edge{From: 1, To: 7},
		&graph.Edge{From: 4, To: 6},
		&graph.Edge{From: 5, To: 4},
		&graph.Edge{From: 6, To: 1},
		&graph.Edge{From: 6, To: 5},
		&graph.Edge{From: 6, To: 6},
		&graph.Edge{From: 7, To: 5},
		&graph.Edge{From: 9, To: 8})

	fmt.Println(g.ToDot().String())
	if g.IsAcyclic() {
		fmt.Println("# ist zyklenfrei")
	} else {
		fmt.Println("# ist nicht zyklenfrei")
	}
}
