package luckytickets_test

import (
	"strconv"
	"testing"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/hw-02-lucky-tickets/luckytickets"
)

type Solver func(N int) int

func (s Solver) Solve(t *testing.T, input []string, output []string) {
	if len(input) == 0 || len(output) == 0 {
		t.Fatal(datatesting.ErrNotEnoughArguments)
	}

	n, err := strconv.Atoi(input[0])
	if err != nil {
		t.Fatalf("parse N: %v", err)
	}
	want, err := strconv.Atoi(output[0])
	if err != nil {
		t.Fatalf("parse expected result: %v", err)
	}

	datatesting.AssertEqual(t, want, s(n))
}

func TestCountFast(t *testing.T) {
	runner := datatesting.NewRunner()
	runner.Run(t, Solver(luckytickets.CountFast))
}
