package islands_test

import (
	"fmt"
	"testing"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/hw-21-dynamic-programming/islands"
)

func TestCount(t *testing.T) {
	tests := []struct {
		M    [][]int
		want int
	}{
		{
			M: [][]int{
				{1, 1, 0, 0},
				{1, 0, 0, 1},
				{1, 0, 1, 0},
				{0, 1, 1, 0},
			},
			want: 3,
		},
		{
			M: [][]int{
				{1, 1, 1, 0},
				{1, 1, 1, 1},
				{1, 0, 1, 0},
				{0, 1, 1, 0},
			},
			want: 1,
		},
		{
			M: [][]int{
				{1, 0, 1, 0},
				{0, 1, 0, 1},
				{1, 0, 1, 0},
				{0, 1, 0, 1},
			},
			want: 8,
		},
		{
			M: [][]int{
				{1, 0, 1, 0, 1},
				{0, 1, 0, 1, 1},
				{1, 0, 1, 0, 1},
				{0, 1, 0, 1, 1},
				{0, 1, 0, 1, 1},
			},
			want: 7,
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%v, count = %d", test.M, test.want), func(t *testing.T) {
			got := islands.Count(test.M)

			datatesting.AssertEqual(t, test.want, got)
		})
	}
}
