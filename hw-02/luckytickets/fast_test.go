package luckytickets_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/hw-02/luckytickets"
)

type Solver func(N int) int

func (s Solver) Solve(input []string) (string, error) {
	if len(input) == 0 {
		return "", datatesting.ErrEmptyInput
	}

	n, err := strconv.Atoi(input[0])
	if err != nil {
		return "", fmt.Errorf("parse N: %w", err)
	}

	count := s(n)

	return strconv.Itoa(count), nil
}

func TestCountFast(t *testing.T) {
	runner := datatesting.NewRunner()
	runner.Run(t, Solver(luckytickets.CountFast))
}
