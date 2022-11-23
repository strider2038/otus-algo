package power_test

import (
	"strconv"
	"testing"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/hw-03-algebraic-algos/power"
)

type Solver func(a float64, n int) float64

func (s Solver) Solve(t *testing.T, input []string, output []string) {
	if len(input) < 2 || len(output) < 1 {
		t.Fatal(datatesting.ErrNotEnoughArguments)
	}
	a, err := strconv.ParseFloat(input[0], 64)
	if err != nil {
		t.Fatalf("parse A: %v", err)
	}
	n, err := strconv.Atoi(input[1])
	if err != nil {
		t.Fatalf("parse N: %v", err)
	}
	want, err := strconv.ParseFloat(output[0], 64)
	if err != nil {
		t.Fatalf("parse output: %v", err)
	}

	datatesting.AssertEqualFloat(t, want, s(a, n))
}

func TestIterative(t *testing.T) {
	runner := datatesting.NewRunner(datatesting.WithLimit(9))
	runner.Run(t, Solver(power.Iterative))
}

func TestLogarithmic(t *testing.T) {
	runner := datatesting.NewRunner()
	runner.Run(t, Solver(power.Logarithmic))
}

func BenchmarkLogarithmic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		power.Logarithmic(1.0000000001, 10000000000)
	}
}
