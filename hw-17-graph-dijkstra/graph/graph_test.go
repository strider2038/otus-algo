package graph_test

import (
	"testing"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/hw-17-graph-dijkstra/graph"
)

func TestAdjacencyArray_FindShortestPath(t *testing.T) {
	tests := []struct {
		name      string
		graph     graph.AdjacencyArray
		from      int
		to        int
		wantEdges []graph.Edge
		wantError error
	}{
		{
			name: "positive case 1",
			graph: graph.AdjacencyArray{
				0: {{Vertex: 1, Distance: 7}, {Vertex: 2, Distance: 9}, {Vertex: 5, Distance: 14}},
				1: {{Vertex: 0, Distance: 7}, {Vertex: 2, Distance: 10}, {Vertex: 3, Distance: 15}},
				2: {{Vertex: 0, Distance: 9}, {Vertex: 1, Distance: 10}, {Vertex: 3, Distance: 11}, {Vertex: 5, Distance: 2}},
				3: {{Vertex: 1, Distance: 15}, {Vertex: 2, Distance: 11}, {Vertex: 4, Distance: 6}},
				4: {{Vertex: 3, Distance: 6}, {Vertex: 5, Distance: 9}},
				5: {{Vertex: 0, Distance: 14}, {Vertex: 2, Distance: 2}, {Vertex: 4, Distance: 9}},
			},
			from: 0,
			to:   4,
			wantEdges: []graph.Edge{
				{Vertex1: 0, Vertex2: 2},
				{Vertex1: 2, Vertex2: 5},
				{Vertex1: 5, Vertex2: 4},
			},
		},
		{
			name: "positive case 2",
			graph: graph.AdjacencyArray{
				0: {{Vertex: 1, Distance: 7}, {Vertex: 2, Distance: 9}, {Vertex: 5, Distance: 14}},
				1: {{Vertex: 0, Distance: 7}, {Vertex: 2, Distance: 10}, {Vertex: 3, Distance: 15}},
				2: {{Vertex: 0, Distance: 9}, {Vertex: 1, Distance: 10}, {Vertex: 3, Distance: 11}, {Vertex: 5, Distance: 2}},
				3: {{Vertex: 1, Distance: 15}, {Vertex: 2, Distance: 11}, {Vertex: 4, Distance: 7}},
				4: {{Vertex: 3, Distance: 7}, {Vertex: 5, Distance: 9}},
				5: {{Vertex: 0, Distance: 14}, {Vertex: 2, Distance: 2}, {Vertex: 4, Distance: 9}},
			},
			from: 4,
			to:   1,
			wantEdges: []graph.Edge{
				{Vertex1: 4, Vertex2: 5},
				{Vertex1: 5, Vertex2: 2},
				{Vertex1: 2, Vertex2: 1},
			},
		},
		{
			name: "loop",
			graph: graph.AdjacencyArray{
				0: {{Vertex: 1, Distance: 1}},
			},
			from:      0,
			to:        0,
			wantEdges: []graph.Edge{},
		},
		{
			name: "no path",
			graph: graph.AdjacencyArray{
				0: {},
				1: {},
			},
			from:      0,
			to:        1,
			wantError: graph.ErrPathNotFound,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			path, err := test.graph.FindShortestPath(test.from, test.to)

			if test.wantError != nil {
				if test.wantError != err {
					t.Fatalf("want error %s, got %s", test.wantError, err)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			datatesting.AssertEqualArrays(t, test.wantEdges, path)
		})
	}
}
