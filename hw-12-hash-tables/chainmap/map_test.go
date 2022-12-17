package chainmap_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/hw-12-hash-tables/chainmap"
)

func TestMapBasic(t *testing.T) {
	m := chainmap.Map[int]{}

	m.Put("alpha", 1)
	m.Put("beta", 2)
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

func TestMapRehash(t *testing.T) {
	m := chainmap.Map[int]{}
	stdmap := map[string]int{}

	ss := GenerateRandomStrings(10000)
	for i, s := range ss {
		m.Put(s, i)
		stdmap[s] = i
	}

	if len(stdmap) != m.Count() {
		t.Errorf("unexpected map count: want %d, got %d", len(stdmap), m.Count())
	}
	for i, s := range ss {
		datatesting.AssertEqual(t, i, m.Get(s))
	}
}

func TestMap_Put(t *testing.T) {
	counts := []int{1000, 10_000, 100_000, 1_000_000}

	for _, count := range counts {
		t.Run(fmt.Sprintf("%d", count), func(t *testing.T) {
			ss := GenerateRandomStrings(count)

			start := time.Now()
			m := chainmap.Map[int]{}
			for i, s := range ss {
				m.Put(s, i)
			}
			t.Log("elapsed time:", time.Since(start).String())
		})
	}
}

func TestStandardMap_Put(t *testing.T) {
	counts := []int{1000, 10_000, 100_000, 1_000_000}

	for _, count := range counts {
		t.Run(fmt.Sprintf("%d", count), func(t *testing.T) {
			ss := GenerateRandomStrings(count)

			start := time.Now()
			m := map[string]int{}
			for i, s := range ss {
				m[s] = i
			}
			t.Log("elapsed time:", time.Since(start).String())
		})
	}
}

var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ 0123456789")

func GenerateRandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

func GenerateRandomStrings(n int) []string {
	rand.Seed(time.Now().UnixNano())

	ss := make([]string, n)
	for i := 0; i < n; i++ {
		ss[i] = GenerateRandomString(15)
	}

	return ss
}
