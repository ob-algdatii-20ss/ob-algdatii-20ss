package graph

import (
	"fmt"
	"testing"
)

func TestAdjMat(t *testing.T) {
	// Init
	g := NewGraphAdjMat(9)

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

	// Method under test
	gotEdges := g.Edges()

	// check
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

func ExampleAdjMat() {
	g := NewGraphAdjMat(9)
	mkExampleGraph(&g)
	fmt.Printf("%v\n", g.Edges())
	// Output:
	// [{1 2} {1 3} {1 7} {4 6} {5 4} {6 1} {6 5} {6 6} {7 5} {9 8}]
}

func mkExampleGraph(g Graph) {
	edges := []Edge{
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
	for _, edge := range edges {
		g.AddEdge(edge.From, edge.To)
	}
}

func BenchmarkAdjMatEdges(b *testing.B) {
	g := NewGraphAdjMat(9)
	mkExampleGraph(&g)
	benchmarkHelper(b, g)
}

func benchmarkHelper(b *testing.B, g Graph) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		g.Edges()
	}
}
