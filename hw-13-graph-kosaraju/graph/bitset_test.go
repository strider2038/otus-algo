package graph_test

import (
	"testing"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/hw-13-graph-kosaraju/graph"
)

func TestBitSet_IsSet(t *testing.T) {
	bits := graph.NewBitSet(100)

	bits.Set(0)
	bits.Set(2)
	bits.Set(63)
	bits.Set(64)
	bits.Set(99)

	for i := 0; i <= 100; i++ {
		if i == 0 || i == 2 || i == 63 || i == 64 || i == 99 {
			datatesting.AssertTrue(t, bits.IsSet(i))
		} else {
			datatesting.AssertFalse(t, bits.IsSet(i))
		}
	}
}

func TestBitSet_OnesCount(t *testing.T) {
	bits := graph.NewBitSet(100)
	bits.Set(0)
	bits.Set(2)
	bits.Set(63)
	bits.Set(64)
	bits.Set(99)
	bits.Unset(2)

	count := bits.OnesCount()

	datatesting.AssertEqual(t, 4, count)
}
