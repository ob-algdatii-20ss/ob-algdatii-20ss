package main

import (
	"fmt"

	"github.com/ob-algdatii-20ss/ob-algdatii-20ss/graph"
)

func main() {
	var g graph.Graph = graph.NewGraphAdjLst(9,
		&graph.Edge{1, 2},
		&graph.Edge{1, 3},
		&graph.Edge{1, 7},
		&graph.Edge{4, 6},
		&graph.Edge{5, 4},
		&graph.Edge{6, 1},
		&graph.Edge{6, 5},
		&graph.Edge{6, 6},
		&graph.Edge{7, 5},
		&graph.Edge{9, 8})

	fmt.Println(g.ToDot().String())
}
