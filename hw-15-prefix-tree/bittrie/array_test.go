package bittrie_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/datatesting/randomize"
	"github.com/strider2038/otus-algo/hw-15-prefix-tree/bittrie"
)

var testChars = []rune("abcdefghijklmnopqrstuvwxyz")

func TestArray64_Basic(t *testing.T) {
	items := bittrie.NewArray64[int]("abcdefghijklmnopqrstuvwxyz")

	items.Put("alpha", 1)
	items.Put("beta", 2)
	items.Put("gamma", 3)
	items.Put("delta", 4)
	items.Delete("beta")
	items.Put("beta", 5)
	items.Put("cap", 6)
	items.Put("cat", 7)
	items.Put("car", 8)
	items.Delete("delta")
	items.Delete("delta")
	items.Delete("unknown")

	datatesting.AssertEqual(t, 6, items.Count())
	datatesting.AssertEqual(t, 1, items.Get("alpha"))
	datatesting.AssertEqual(t, 5, items.Get("beta"))
	datatesting.AssertEqual(t, 3, items.Get("gamma"))
	datatesting.AssertEqual(t, 6, items.Get("cap"))
	datatesting.AssertEqual(t, 7, items.Get("cat"))
	datatesting.AssertEqual(t, 8, items.Get("car"))
	datatesting.AssertEqual(t, 0, items.Get("delta"))
	if _, exist := items.Find("delta"); exist {
		t.Error("delta value is found in map")
	}
}

func TestArray64_RealData(t *testing.T) {
	countries := bittrie.NewArray64[int](`abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ "-&`)
	m := map[string]int{}

	for i, country := range randomize.Countries {
		countries.Put(country, i+1)
		m[country] = i + 1
	}

	countries.Iterate(func(key string, value int) error {
		datatesting.AssertEqual(t, m[key], value)

		return nil
	})
}

func TestArray64_MarshalJSON(t *testing.T) {
	items := bittrie.NewArray64[int]("abcdefghijklmnopqrstuvwxyz")
	items.Put("alpha", 1)
	items.Put("beta", 2)
	items.Put("gamma", 3)
	items.Put("delta", 4)

	data, err := json.Marshal(items)
	if err != nil {
		t.Fatal(err)
	}

	datatesting.AssertEqual(t, `{"alpha":1,"beta":2,"delta":4,"gamma":3}`, string(data))
}

func TestArray64_Put(t *testing.T) {
	counts := []int{100, 1000, 10_000, 100_000, 1_000_000}

	for _, count := range counts {
		t.Run(fmt.Sprintf("%d", count), func(t *testing.T) {
			ss := datatesting.GenerateRandomStrings(count, testChars...)

			start := time.Now()
			m := bittrie.NewArray64[int](string(testChars))
			for i, s := range ss {
				m.Put(s, i)
			}
			t.Log("elapsed time:", time.Since(start).String())
			t.Log("mem size (KB):", m.MemSize()/1024)
			t.Log("mem size (MB):", m.MemSize()/1024/1024)
		})
	}
}

func TestArray64_Get(t *testing.T) {
	counts := []int{100, 1000, 10_000, 100_000, 1_000_000}

	for _, count := range counts {
		t.Run(fmt.Sprintf("%d", count), func(t *testing.T) {
			ss := datatesting.GenerateRandomStrings(count, testChars...)
			m := bittrie.NewArray64[int](string(testChars))
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
