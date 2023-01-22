package graph_test

import (
	"testing"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/hw-16-graph-kraskal/graph"
)

func TestEdges_FindSpanningTree(t *testing.T) {
	tests := []struct {
		name      string
		graph     graph.Edges
		want      graph.Edges
		wantError error
	}{
		{
			name: "valid case 1",
			graph: graph.Edges{
				{Vertex1: 0, Vertex2: 1, Weight: 3},
				{Vertex1: 0, Vertex2: 4, Weight: 1},
				{Vertex1: 1, Vertex2: 2, Weight: 5},
				{Vertex1: 1, Vertex2: 4, Weight: 4},
				{Vertex1: 2, Vertex2: 3, Weight: 2},
				{Vertex1: 2, Vertex2: 4, Weight: 6},
				{Vertex1: 3, Vertex2: 4, Weight: 7},
			},
			want: graph.Edges{
				{Vertex1: 0, Vertex2: 4, Weight: 1},
				{Vertex1: 2, Vertex2: 3, Weight: 2},
				{Vertex1: 0, Vertex2: 1, Weight: 3},
				{Vertex1: 1, Vertex2: 2, Weight: 5},
			},
		},
		{
			name: "valid case 2",
			graph: graph.Edges{
				{Vertex1: 0, Vertex2: 1, Weight: 9},
				{Vertex1: 0, Vertex2: 3, Weight: 10},
				{Vertex1: 0, Vertex2: 8, Weight: 3},
				{Vertex1: 1, Vertex2: 2, Weight: 4},
				{Vertex1: 1, Vertex2: 4, Weight: 8},
				{Vertex1: 1, Vertex2: 8, Weight: 16},
				{Vertex1: 2, Vertex2: 4, Weight: 14},
				{Vertex1: 2, Vertex2: 5, Weight: 1},
				{Vertex1: 3, Vertex2: 4, Weight: 7},
				{Vertex1: 3, Vertex2: 6, Weight: 13},
				{Vertex1: 3, Vertex2: 7, Weight: 5},
				{Vertex1: 3, Vertex2: 8, Weight: 11},
				{Vertex1: 4, Vertex2: 5, Weight: 12},
				{Vertex1: 4, Vertex2: 6, Weight: 15},
				{Vertex1: 5, Vertex2: 6, Weight: 2},
				{Vertex1: 6, Vertex2: 7, Weight: 6},
			},
			want: graph.Edges{
				{Vertex1: 2, Vertex2: 5, Weight: 1},
				{Vertex1: 5, Vertex2: 6, Weight: 2},
				{Vertex1: 0, Vertex2: 8, Weight: 3},
				{Vertex1: 1, Vertex2: 2, Weight: 4},
				{Vertex1: 3, Vertex2: 7, Weight: 5},
				{Vertex1: 6, Vertex2: 7, Weight: 6},
				{Vertex1: 3, Vertex2: 4, Weight: 7},
				{Vertex1: 0, Vertex2: 1, Weight: 9},
			},
		},
		{
			name: "valid case 3",
			graph: graph.Edges{
				{Vertex1: 0, Vertex2: 1, Weight: 1},
				{Vertex1: 0, Vertex2: 2, Weight: 2},
				{Vertex1: 1, Vertex2: 2, Weight: 3},
			},
			want: graph.Edges{
				{Vertex1: 0, Vertex2: 1, Weight: 1},
				{Vertex1: 0, Vertex2: 2, Weight: 2},
			},
		},
		{
			name: "non intersecting graphs",
			graph: graph.Edges{
				{Vertex1: 0, Vertex2: 1, Weight: 1},
				{Vertex1: 2, Vertex2: 3, Weight: 2},
			},
			wantError: graph.ErrNonIntersectingGraphs,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.graph.FindSpanningTree()

			if test.wantError != nil {
				if test.wantError != err {
					t.Fatalf("want error %s, got %s", test.wantError, err)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			datatesting.AssertEqualArrays(t, test.want, got)
		})
	}
}
