package main

import (
	"fmt"
	"math"
	"time"

	"github.com/strider2038/otus-algo/hw-03-algebraic-algos/fibonacci"
	"github.com/strider2038/otus-algo/hw-03-algebraic-algos/power"
	"github.com/strider2038/otus-algo/hw-03-algebraic-algos/prime"
)

func main() {
	testPowerAlgorithms()
	testFibonacciAlgorithms()
	testPrimeAlgorithms()
}

func testPowerAlgorithms() {
	testAlgorithmPow10("power.Iterative", 10, func(n int) {
		power.Iterative(1.0001, n)
	})
	testAlgorithmPow10("power.Logarithmic", 15, func(n int) {
		power.Logarithmic(1.0001, n)
	})
	testAlgorithmPow10("math.Pow", 15, func(n int) {
		math.Pow(1.0001, float64(n))
	})
}

func testFibonacciAlgorithms() {
	testAlgorithm(
		"fibonacci.Recursive",
		func(n int) { fibonacci.Recursive(n) },
		10, 20, 30, 40,
	)
	testAlgorithm(
		"fibonacci.Iterative",
		func(n int) { fibonacci.Iterative(n) },
		10, 20, 30, 40, 50, 100, 1_000, 10_000, 100_000, 1_000_000,
	)
	testAlgorithm(
		"fibonacci.ByMatrix",
		func(n int) { fibonacci.ByMatrix(n) },
		10, 20, 30, 40, 50, 100, 1_000, 10_000, 100_000, 1_000_000, 10_000_000,
	)
}

func testPrimeAlgorithms() {
	testAlgorithmPow10("prime.CountByBruteForce", 5, func(n int) {
		prime.CountByBruteForce(n)
	})
	testAlgorithmPow10("prime.CountByBruteForceOptimized", 7, func(n int) {
		prime.CountByBruteForceOptimized(n)
	})
	testAlgorithmPow10("prime.CountByPrimes", 7, func(n int) {
		prime.CountByPrimes(n)
	})
	testAlgorithmPow10("prime.CountBySieveOfEratosthenes", 9, func(n int) {
		prime.CountBySieveOfEratosthenes(n)
	})
	testAlgorithmPow10("prime.CountBySieveOfEratosthenesOptimized", 9, func(n int) {
		prime.CountBySieveOfEratosthenesOptimized(n)
	})
}

func testAlgorithmPow10(name string, maxN int, run func(n int)) {
	n := 10
	for i := 1; i <= maxN; i++ {
		start := time.Now()
		run(n)
		elapsed := time.Since(start)
		fmt.Printf("%s: n=%d, elapsed=%s\n", name, n, elapsed.String())
		n = n * 10
	}
}

func testAlgorithm(name string, run func(n int), n ...int) {
	for i := 0; i < len(n); i++ {
		start := time.Now()
		run(n[i])
		elapsed := time.Since(start)
		fmt.Printf("%s: n=%d, elapsed=%s\n", name, n[i], elapsed.String())
	}
}
