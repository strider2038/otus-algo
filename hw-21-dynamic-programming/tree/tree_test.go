package tree_test

import (
	"fmt"
	"testing"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/hw-21-dynamic-programming/tree"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		A    [][]int
		want int
	}{
		{
			A: [][]int{
				{1},
				{2, 3},
				{4, 5, 6},
				{9, 8, 0, 3},
			},
			want: 17,
		},
		{
			A: [][]int{
				{1},
				{4, 3},
				{1, 5, 9},
				{6, 2, 1, 3},
			},
			want: 16,
		},
		{
			A: [][]int{
				{1},
				{4, 3},
				{3, 5, 9},
				{8, 2, 1, 2},
			},
			want: 16,
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%v -> %d", test.A, test.want), func(t *testing.T) {
			got := tree.Calculate(test.A)

			datatesting.AssertEqual(t, test.want, got)
		})
	}
}
