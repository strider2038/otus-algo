package luckytickets_test

import (
	"testing"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/hw-02-lucky-tickets/luckytickets"
)

func TestCountRecursively(t *testing.T) {
	runner := datatesting.NewRunner(datatesting.WithLimit(4))
	runner.Run(t, Solver(luckytickets.CountRecursively))
}
