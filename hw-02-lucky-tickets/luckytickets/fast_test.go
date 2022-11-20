package luckytickets_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/hw-02-lucky-tickets/luckytickets"
)

type Solver func(N int) int

func (s Solver) Solve(input []string, output []string) error {
	if len(input) == 0 || len(output) == 0 {
		return datatesting.ErrNotEnoughArguments
	}

	n, err := strconv.Atoi(input[0])
	if err != nil {
		return fmt.Errorf("parse N: %w", err)
	}
	want, err := strconv.Atoi(output[0])
	if err != nil {
		return fmt.Errorf("parse expected result: %w", err)
	}

	return datatesting.AssertEqual(want, s(n))
}

func TestCountFast(t *testing.T) {
	runner := datatesting.NewRunner()
	runner.Run(t, Solver(luckytickets.CountFast))
}
