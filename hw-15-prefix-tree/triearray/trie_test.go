package triearray_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/hw-12-hash-tables/chainmap"
	"github.com/strider2038/otus-algo/hw-15-prefix-tree/triearray"
)

var testChars = []rune("abcdefghijklmnopqrstuvwxyz")

func TestTrieArray(t *testing.T) {
	items := triearray.New[int]()
	items.Put("apple", 1)
	datatesting.AssertEqual(t, 1, items.Get("apple"))
	datatesting.AssertEqual(t, 0, items.Get("app"))
	items.Put("app", 2)
	datatesting.AssertEqual(t, 2, items.Get("app"))
	items.Put("abcdefghijklmnopqrstuvwxyz", 3)
}

func TestTrieArrayBasic(t *testing.T) {
	items := triearray.New[int]()

	items.Put("alpha", 1)
	items.Put("beta", 2)
	items.Put("gamma", 3)
	items.Put("delta", 4)
	items.Put("beta", 5)
	items.Delete("delta")
	items.Delete("delta")

	datatesting.AssertEqual(t, 3, items.Count())
	datatesting.AssertEqual(t, 1, items.Get("alpha"))
	datatesting.AssertEqual(t, 5, items.Get("beta"))
	datatesting.AssertEqual(t, 3, items.Get("gamma"))
	datatesting.AssertEqual(t, 0, items.Get("delta"))
	if _, exist := items.Find("delta"); exist {
		t.Error("delta value is found in map")
	}
}

func TestTrieArray_Put(t *testing.T) {
	counts := []int{100, 1000, 10_000, 100_000, 1_000_000}

	for _, count := range counts {
		t.Run(fmt.Sprintf("%d", count), func(t *testing.T) {
			ss := datatesting.GenerateRandomStrings(count, testChars...)

			start := time.Now()
			m := triearray.Trie[int]{}
			for i, s := range ss {
				m.Put(s, i)
			}
			t.Log("elapsed time:", time.Since(start).String())
			t.Log("mem size (KB):", m.MemSize()/1024)
			t.Log("mem size (MB):", m.MemSize()/1024/1024)
		})
	}
}

func TestTrieArray_Get(t *testing.T) {
	counts := []int{100, 1000, 10_000, 100_000, 1_000_000}

	for _, count := range counts {
		t.Run(fmt.Sprintf("%d", count), func(t *testing.T) {
			ss := datatesting.GenerateRandomStrings(count, testChars...)
			m := triearray.Trie[int]{}
			for i, s := range ss {
				m.Put(s, i)
			}

			var sum time.Duration
			for _, s := range ss {
				start := time.Now()
				m.Get(s)
				sum += time.Since(start)
			}
			t.Log("average time:", (sum / time.Duration(len(ss))).String())
		})
	}
}

func TestMap_Put(t *testing.T) {
	counts := []int{100, 1000, 10_000, 100_000, 1_000_000}

	for _, count := range counts {
		t.Run(fmt.Sprintf("%d", count), func(t *testing.T) {
			ss := datatesting.GenerateRandomStrings(count, testChars...)

			start := time.Now()
			m := chainmap.Map[int]{}
			for i, s := range ss {
				m.Put(s, i)
			}
			t.Log("elapsed time:", time.Since(start).String())
			t.Log("mem size (KB):", m.MemSize()/1024)
			t.Log("mem size (MB):", m.MemSize()/1024/1024)
		})
	}
}

func TestMap_Get(t *testing.T) {
	counts := []int{100, 1000, 10_000, 100_000, 1_000_000}

	for _, count := range counts {
		t.Run(fmt.Sprintf("%d", count), func(t *testing.T) {
			ss := datatesting.GenerateRandomStrings(count, testChars...)
			m := chainmap.Map[int]{}
			for i, s := range ss {
				m.Put(s, i)
			}

			var sum time.Duration
			for _, s := range ss {
				start := time.Now()
				m.Get(s)
				sum += time.Since(start)
			}
			t.Log("average time:", (sum / time.Duration(len(ss))).String())
		})
	}
}

func TestStandardMap_Put(t *testing.T) {
	counts := []int{100, 1000, 10_000, 100_000, 1_000_000}

	for _, count := range counts {
		t.Run(fmt.Sprintf("%d", count), func(t *testing.T) {
			ss := datatesting.GenerateRandomStrings(count)

			start := time.Now()
			m := map[string]int{}
			for i, s := range ss {
				m[s] = i
			}
			t.Log("elapsed time:", time.Since(start).String())
		})
	}
}

func TestStandardMap_Get(t *testing.T) {
	counts := []int{100, 1000, 10_000, 100_000, 1_000_000}

	for _, count := range counts {
		t.Run(fmt.Sprintf("%d", count), func(t *testing.T) {
			ss := datatesting.GenerateRandomStrings(count)
			m := map[string]int{}
			for i, s := range ss {
				m[s] = i
			}

			var sum time.Duration
			for _, s := range ss {
				start := time.Now()
				_ = m[s]
				sum += time.Since(start)
			}
			t.Log("average time:", (sum / time.Duration(len(ss))).String())
		})
	}
}
