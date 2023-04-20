package goroh_test

import (
	"fmt"
	"testing"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/hw-21-dynamic-programming/goroh"
)

func TestGCD(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{a: 0, b: 1, want: 1},
		{a: 12, b: 0, want: 12},
		{a: 52, b: 48, want: 4},
		{a: 62, b: 45, want: 1},
		{a: 65, b: 45, want: 5},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf(`GCD(%d, %d)=%d`, test.a, test.b, test.want), func(t *testing.T) {
			got := goroh.GCD(test.a, test.b)

			datatesting.AssertEqual(t, test.want, got)
		})
	}
}
