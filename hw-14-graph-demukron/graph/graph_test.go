package graph_test

import (
	"testing"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/hw-14-graph-demukron/graph"
)

func TestAdjacencyArray_SortByDemukron(t *testing.T) {
	tests := []struct {
		name      string
		graph     graph.AdjacencyArray
		want      [][]int
		wantError error
	}{
		{
			name: "case 1",
			graph: graph.AdjacencyArray{
				0: {1},
				1: {4},
				2: {3},
				3: {0, 1, 4, 5},
				4: {6},
				5: {4, 7},
				6: {7},
				7: {},
			},
			want: [][]int{
				{2},
				{3},
				{0, 5},
				{1},
				{4},
				{6},
				{7},
			},
		},
		{
			name: "case 2",
			graph: graph.AdjacencyArray{
				0: {3},
				1: {2, 7},
				2: {4},
				3: {2},
				4: {5},
				5: {},
				6: {2},
				7: {4},
			},
			want: [][]int{
				{0, 1, 6},
				{3, 7},
				{2},
				{4},
				{5},
			},
		},
		{
			name: "cyclic graph",
			graph: graph.AdjacencyArray{
				0: {1},
				1: {4},
				2: {3},
				3: {0, 1, 4, 5},
				4: {6},
				5: {4, 7},
				6: {7},
				7: {1},
			},
			wantError: graph.ErrNotSortable,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.graph.SortByDemukron()

			if test.wantError != nil {
				if test.wantError != err {
					t.Fatalf("want error %s, got %s", test.wantError, err)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			datatesting.AssertEqualMatrix(t, test.want, got)
		})
	}
}
