package openmap_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/hw-12-hash-tables/openmap"
)

func TestMapBasic(t *testing.T) {
	m := openmap.Map[int]{}

	m.Put("alpha", 1)
	m.Put("beta", 2)
	m.Put("gamma", 3)
	m.Delete("gamma")
	m.Put("gamma", 3)
	m.Put("delta", 4)
	m.Put("beta", 5)
	m.Delete("delta")

	datatesting.AssertEqual(t, 3, m.Count())
	datatesting.AssertEqual(t, 1, m.Get("alpha"))
	datatesting.AssertEqual(t, 5, m.Get("beta"))
	datatesting.AssertEqual(t, 3, m.Get("gamma"))
	datatesting.AssertEqual(t, 0, m.Get("delta"))
	if _, exist := m.Find("delta"); exist {
		t.Error("delta value is found in map")
	}
}

func TestMapCollisions(t *testing.T) {
	m := openmap.Map[int]{}

	m.Put("PZD3n7 Aep", 1)
	m.Put("4TAvA5wK4f", 2)
	m.Put("tw0LcVjr v", 3)
	m.Put("VjpdKEZgoX", 4)
	m.Delete("4TAvA5wK4f")
	m.Put("FQ5rNDbYmd", 5)
	m.Delete("FQ5rNDbYmd")

	m.Get("FQ5rNDbYmd")
}

func TestMapRehash(t *testing.T) {
	m := openmap.Map[int]{}
	stdmap := map[string]int{}

	ss := datatesting.GenerateRandomStrings(10000)
	for i, s := range ss {
		m.Put(s, i)
		stdmap[s] = i
	}
	for i := 0; i < len(ss); i += 2 {
		m.Delete(ss[i])
		delete(stdmap, ss[i])
	}
	for i, s := range ss {
		m.Put(s, len(ss)+i)
		stdmap[s] = len(ss) + i
	}

	if len(stdmap) != m.Count() {
		t.Errorf("unexpected map count: want %d, got %d", len(stdmap), m.Count())
	}
	for i, s := range ss {
		datatesting.AssertEqual(t, len(ss)+i, m.Get(s))
	}
}

func TestMap_Put(t *testing.T) {
	counts := []int{1000, 10_000, 100_000, 1_000_000}

	for _, count := range counts {
		t.Run(fmt.Sprintf("%d", count), func(t *testing.T) {
			ss := datatesting.GenerateRandomStrings(count)

			start := time.Now()
			m := openmap.Map[int]{}
			for i, s := range ss {
				m.Put(s, i)
			}
			t.Log("elapsed time:", time.Since(start).String())
		})
	}
}
