package graph

import (
	"fmt"
	"testing"
)

func TestAdjLst(t *testing.T) {
	g := NewGraphAdjLst(9)
	testGraph(t, g)
}

func testGraph(t *testing.T, g Graph) {
	wantedEdges := []Edge{
		{1, 2},
		{1, 3},
		{1, 7},
		{4, 6},
		{5, 4},
		{6, 1},
		{6, 5},
		{6, 6},
		{7, 5},
		{9, 8},
	}

	for _, edge := range wantedEdges {
		g.AddEdge(edge.From, edge.To)
	}

	gotEdges := g.Edges()

outerWanted:
	for _, wanted := range wantedEdges {
		for _, got := range gotEdges {
			if wanted.From == got.From && wanted.To == got.To {
				continue outerWanted
			}
		}
		t.Errorf("Edge %v wanted, but not found", wanted)
	}

outerGot:
	for _, got := range gotEdges {
		for _, wanted := range wantedEdges {
			if wanted.From == got.From && wanted.To == got.To {
				continue outerGot
			}
		}
		t.Errorf("Edge %v found, but not wanted", got)
	}
}

func TestIsAcyclicLst(t *testing.T) {
	cases := []struct {
		graph     AdjLst
		isAcyclic bool
	}{
		{NewGraphAdjLst(1), true},
		{NewGraphAdjLst(1, &Edge{1, 1}), false},
		{NewGraphAdjLst(2, &Edge{1, 2}), true},
		{NewGraphAdjLst(2, &Edge{1, 2}, &Edge{2, 1}), false},
	}

	for _, c := range cases {
		got := c.graph.IsAcyclic()
		if got != c.isAcyclic {
			t.Errorf("%v IsAcyclic() == %v, want %v", c.graph, got, c.isAcyclic)
		}
	}
}

func ExampleAdjLst() {
	g := NewGraphAdjLst(9)
	mkExampleGraph(&g)
	fmt.Printf("%v\n", g.Edges())
	// Output:
	// [{1 7} {1 3} {1 2} {4 6} {5 4} {6 6} {6 5} {6 1} {7 5} {9 8}]
}

func BenchmarkAdjLstEdges(b *testing.B) {
	g := NewGraphAdjLst(9)
	benchmarkHelper(b, g)
}
