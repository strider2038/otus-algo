package graph_test

import (
	"fmt"
	"testing"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/hw-13-graph-kosaraju/graph"
)

func TestAdjacencyMatrix_FindStronglyConnectedComponents(t *testing.T) {
	tests := []struct {
		name           string
		graph          graph.AdjacencyMatrix
		wantComponents [][]int
	}{
		{
			name: "case 1",
			graph: graph.AdjacencyMatrix{
				0: {1},
				1: {2, 4, 5},
				2: {3, 6},
				3: {2, 7},
				4: {0, 5},
				5: {6},
				6: {5},
				7: {3, 6},
			},
			wantComponents: [][]int{
				{5, 6},
				{2, 7, 3},
				{1, 0, 4},
			},
		},
		{
			name: "case 2",
			graph: graph.AdjacencyMatrix{
				0: {1, 4},
				1: {2},
				2: {3},
				3: {1},
				4: {3},
			},
			wantComponents: [][]int{
				{2, 1, 3},
				{4},
				{0},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.graph.FindStronglyConnectedComponents()

			datatesting.AssertEqualMatrix(t, test.wantComponents, got)
		})
	}
}

func TestAdjacencyMatrix_Inverse(t *testing.T) {
	g := graph.AdjacencyMatrix{
		0: {1},
		1: {2, 4, 5},
		2: {3, 6},
		3: {2, 7},
		4: {0, 5},
		5: {6},
		6: {5, 7},
		7: {3},
	}

	got := g.Inverse()

	want := graph.AdjacencyMatrix{
		0: {4},
		1: {0},
		2: {1, 3},
		3: {2, 7},
		4: {1},
		5: {1, 4, 6},
		6: {2, 5},
		7: {3, 6},
	}
	if len(want) != len(got) {
		t.Errorf("different matrix rows count: want %d, got %d", len(want), len(got))
		return
	}
	for i := 0; i < len(want); i++ {
		datatesting.AssertEqualArrays(t, want[i], got[i])
	}
}

func TestAdjacencyMatrix_PathExistsDFS(t *testing.T) {
	tests := []struct {
		from int
		to   int
		want bool
	}{
		{from: 0, to: 4, want: true},
		{from: 0, to: 2, want: true},
		{from: 5, to: 6, want: true},
		{from: 1, to: 6, want: false},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("from vertex %d to %d", test.from, test.to), func(t *testing.T) {
			g := graph.AdjacencyMatrix{
				0: {1, 3},
				1: {0, 4},
				2: {0, 3, 4},
				3: {0, 2},
				4: {1, 2},
				5: {6},
				6: {5},
			}

			got := g.PathExistsDFS(test.from, test.to)

			datatesting.AssertEqual(t, test.want, got)
		})
	}
}

func TestAdjacencyMatrix_WalkDFS(t *testing.T) {
	tests := []struct {
		vertex   int
		wantPath []int
	}{
		{vertex: 0, wantPath: []int{0, 1, 4, 2, 3}},
		{vertex: 2, wantPath: []int{2, 0, 1, 4, 3}},
		{vertex: 4, wantPath: []int{4, 1, 0, 2, 3}},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("from vertex %d", test.vertex), func(t *testing.T) {
			g := graph.AdjacencyMatrix{
				0: {1, 2, 3},
				1: {0, 4},
				2: {0, 3, 4},
				3: {0, 2},
				4: {1, 2},
			}

			var path []int
			g.WalkDFS(test.vertex, func(vertex int) {
				path = append(path, vertex)
			})

			datatesting.AssertEqualArrays(t, test.wantPath, path)
		})
	}
}
