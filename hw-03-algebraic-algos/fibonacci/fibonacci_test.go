package fibonacci_test

import (
	"fmt"
	"math/big"
	"strconv"
	"testing"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/hw-03-algebraic-algos/fibonacci"
)

type BigIntSolver func(n int) *big.Int

func (s BigIntSolver) Solve(input []string, output string) error {
	if len(input) < 1 {
		return datatesting.ErrNotEnoughArguments
	}
	n, err := strconv.Atoi(input[0])
	if err != nil {
		return fmt.Errorf("parse N: %w", err)
	}
	want := new(big.Int)
	want, ok := want.SetString(output, 10)
	if !ok {
		return fmt.Errorf("parse output")
	}

	got := s(n)
	if want.Cmp(got) != 0 {
		return fmt.Errorf(`test failed: want %s, got %s`, want.String(), got.String())
	}

	return nil
}

func TestRecursive(t *testing.T) {
	runner := datatesting.NewRunner(datatesting.WithLimit(7))
	runner.Run(t, BigIntSolver(fibonacci.Recursive))
}

func TestIterative(t *testing.T) {
	runner := datatesting.NewRunner(datatesting.WithLimit(12))
	runner.Run(t, BigIntSolver(fibonacci.Iterative))
}

func TestByMatrix(t *testing.T) {
	runner := datatesting.NewRunner()
	runner.Run(t, BigIntSolver(fibonacci.ByMatrix))
}

func TestMatrix_Mul(t *testing.T) {
	m1 := fibonacci.Matrix{
		{big.NewInt(1), big.NewInt(2)},
		{big.NewInt(3), big.NewInt(4)},
	}
	m2 := fibonacci.Matrix{
		{big.NewInt(5), big.NewInt(6)},
		{big.NewInt(7), big.NewInt(8)},
	}

	m := m1.Mul(m2)

	err := datatesting.AssertEqual("((19, 22), (43, 50))", m.String())
	if err != nil {
		t.Error(err)
	}
}

func TestMatrix_Pow(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{n: 2, want: "((7, 10), (15, 22))"},
		{n: 3, want: "((37, 54), (81, 118))"},
		{n: 5, want: "((1069, 1558), (2337, 3406))"},
	}
	for _, test := range tests {
		t.Run(strconv.Itoa(test.n), func(t *testing.T) {
			m := fibonacci.Matrix{
				{big.NewInt(1), big.NewInt(2)},
				{big.NewInt(3), big.NewInt(4)},
			}

			got := m.Pow(test.n)

			err := datatesting.AssertEqual(test.want, got.String())
			if err != nil {
				t.Error(err)
			}
		})
	}
}
