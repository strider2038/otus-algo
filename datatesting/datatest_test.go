package datatesting_test

import (
	"strconv"
	"testing"
	"unicode/utf8"

	"github.com/strider2038/otus-algo/datatesting"
)

func TestRun(t *testing.T) {
	datatesting.Run(t, datatesting.SolverFunc(func(t *testing.T, input, output []string) {
		if len(input) == 0 || len(output) == 0 {
			t.Fatal(datatesting.ErrNotEnoughArguments)
		}

		want := output[0]
		length := utf8.RuneCountInString(input[0])
		got := strconv.Itoa(length)

		datatesting.AssertEqual(t, want, got)
	}))
}
