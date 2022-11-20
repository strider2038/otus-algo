package power_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/hw-03-algebraic-algos/power"
)

type Solver func(a float64, n int) float64

func (s Solver) Solve(input []string, output []string) error {
	if len(input) < 2 || len(output) < 1 {
		return datatesting.ErrNotEnoughArguments
	}
	a, err := strconv.ParseFloat(input[0], 64)
	if err != nil {
		return fmt.Errorf("parse A: %w", err)
	}
	n, err := strconv.Atoi(input[1])
	if err != nil {
		return fmt.Errorf("parse N: %w", err)
	}
	want, err := strconv.ParseFloat(output[0], 64)
	if err != nil {
		return fmt.Errorf("parse output: %w", err)
	}

	return datatesting.AssertEqualFloat(want, s(a, n))
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
