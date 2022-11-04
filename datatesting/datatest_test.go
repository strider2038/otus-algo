package datatesting_test

import (
	"strconv"
	"testing"
	"unicode/utf8"

	"github.com/strider2038/otus-algo/datatesting"
)

func TestRun(t *testing.T) {
	datatesting.Run(t, datatesting.SolverFunc(func(input []string) (string, error) {
		if len(input) == 0 {
			return "", datatesting.ErrEmptyInput
		}

		length := utf8.RuneCountInString(input[0])

		return strconv.Itoa(length), nil
	}))
}
