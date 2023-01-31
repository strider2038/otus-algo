package pattern_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/hw-18-boyer-moore-algo/pattern"
)

var algos = []struct {
	name string
	find func(text string, pattern string) int
}{
	{name: "full", find: pattern.FindFull},
	{name: "pattern prefix", find: pattern.FindByPatternPrefix},
	{name: "text suffix", find: pattern.FindByTextSuffix},
	{name: "Boyer Moore Horspool", find: pattern.FindBMH},
	{name: "strings.Index", find: strings.Index},
}

var cases = []struct {
	text    string
	pattern string
	want    int
}{
	{text: "ababababababc", pattern: "abc", want: 10},
	{text: "abababababababababababababababababababababc", pattern: "abc", want: 40},
	{text: "the quick brown fox jumps over the lazy dog", pattern: "jump", want: 20},
	{text: "the quick brown fox jumps over the lazy dog", pattern: "dog", want: 40},
	{text: "the quick brown fox jumps over the lazy dog", pattern: "the", want: 0},
	{text: "the quick brown fox jumps over the lazy dog", pattern: "dot", want: -1},
}

func TestFind(t *testing.T) {
	for _, algo := range algos {
		for _, test := range cases {
			t.Run(fmt.Sprintf("%s: %s (%s)", algo.name, test.text, test.pattern), func(t *testing.T) {
				got := algo.find(test.text, test.pattern)

				datatesting.AssertEqual(t, test.want, got)
			})
		}
	}
}

func BenchmarkFind(b *testing.B) {
	for _, algo := range algos {
		b.Run(algo.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				algo.find("the quick brown fox jumps over the lazy dog", "dog")
			}
		})
	}
}

func BenchmarkFind_BigString(b *testing.B) {
	text := strings.Repeat(" the quick brown fox jumps over the lazy ", 10) +
		"dog" + strings.Repeat(" the quick brown fox jumps over the lazy ", 10)
	for _, algo := range algos {
		b.Run(algo.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				algo.find(text, "dog")
			}
		})
	}
}

func BenchmarkFindBMHWithTable(b *testing.B) {
	text := strings.Repeat(" the quick brown fox jumps over the lazy ", 10) +
		"dog" + strings.Repeat(" the quick brown fox jumps over the lazy ", 10)
	shifts := pattern.CreateShiftTable("dog")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pattern.FindBMHWithTable(text, "dog", shifts)
	}
}
