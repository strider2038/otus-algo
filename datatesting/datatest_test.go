package datatesting_test

import (
	"strconv"
	"testing"
	"unicode/utf8"

	"github.com/strider2038/otus-algo/datatesting"
)

func TestRun(t *testing.T) {
	datatesting.Run(t, datatesting.SolverFunc(func(input, output []string) error {
		if len(input) == 0 || len(output) == 0 {
			return datatesting.ErrNotEnoughArguments
		}

		want := output[0]
		length := utf8.RuneCountInString(input[0])
		got := strconv.Itoa(length)

		return datatesting.AssertEqual(want, got)
	}))
}
