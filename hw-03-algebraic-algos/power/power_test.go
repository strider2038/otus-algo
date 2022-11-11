package power_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/strider2038/otus-algo/datatesting"
)

func TestIterative(t *testing.T) {
	runner := datatesting.NewRunner()
	runner.Run(t, datatesting.SolverFunc(func(input []string) (string, error) {
		if len(input) < 2 {
			return "", datatesting.ErrEmptyInput
		}
		a, err := strconv.ParseFloat(input[0], 64)
		if err != nil {
			return "", fmt.Errorf("parse A: %w", err)
		}
		n, err := strconv.Atoi(input[1])
		if err != nil {
			return "", fmt.Errorf("parse N: %w", err)
		}

	}))
}
