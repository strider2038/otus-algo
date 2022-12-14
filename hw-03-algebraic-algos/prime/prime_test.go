package prime_test

import (
	"strconv"
	"testing"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/hw-03-algebraic-algos/prime"
)

type Solver func(n int) int

func (s Solver) Solve(t *testing.T, input, output []string) {
	if len(input) < 1 || len(output) < 1 {
		t.Fatal(datatesting.ErrNotEnoughArguments)
	}
	n, err := strconv.Atoi(input[0])
	if err != nil {
		t.Fatalf("parse N: %v", err)
	}
	want, err := strconv.Atoi(output[0])
	if err != nil {
		t.Fatalf("parse output: %v", err)
	}

	datatesting.AssertEqual(t, want, s(n))
}

func TestCountByBruteForce(t *testing.T) {
	runner := datatesting.NewRunner(datatesting.WithLimit(10))
	runner.Run(t, Solver(prime.CountByBruteForce))
}

func TestCountByBruteForceOptimized(t *testing.T) {
	runner := datatesting.NewRunner(datatesting.WithLimit(11))
	runner.Run(t, Solver(prime.CountByBruteForceOptimized))
}

func TestCountByPrimes(t *testing.T) {
	runner := datatesting.NewRunner(datatesting.WithLimit(12))
	runner.Run(t, Solver(prime.CountByPrimes))
}

func TestCountBySieveOfEratosthenes(t *testing.T) {
	runner := datatesting.NewRunner()
	runner.Run(t, Solver(prime.CountBySieveOfEratosthenes))
}

func TestCountBySieveOfEratosthenesOptimized(t *testing.T) {
	runner := datatesting.NewRunner()
	runner.Run(t, Solver(prime.CountBySieveOfEratosthenesOptimized))
}

// cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
// BenchmarkCountBySieveOfEratosthenes
// BenchmarkCountBySieveOfEratosthenes-8   	      33	  35462247 ns/op	10002494 B/op	       1 allocs/op
func BenchmarkCountBySieveOfEratosthenes(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		prime.CountBySieveOfEratosthenes(10000000)
	}
}

// cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
// BenchmarkCountBySieveOfEratosthenesOptimized
// BenchmarkCountBySieveOfEratosthenesOptimized-8   	      44	  26107945 ns/op	  630871 B/op	       1 allocs/op
func BenchmarkCountBySieveOfEratosthenesOptimized(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		prime.CountBySieveOfEratosthenesOptimized(10000000)
	}
}
