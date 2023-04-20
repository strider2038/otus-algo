package a58_test

import (
	"fmt"
	"testing"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/hw-21-dynamic-programming/a58"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{n: 0, want: 0},
		{n: 1, want: 2},
		{n: 2, want: 4},
		{n: 3, want: 6},
		{n: 4, want: 10},
		{n: 5, want: 16},
		{n: 6, want: 26},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("n = %d, f(n) = %d", test.n, test.want), func(t *testing.T) {
			got := a58.Calculate(test.n)

			datatesting.AssertEqual(t, test.want, got)
		})
	}
}
