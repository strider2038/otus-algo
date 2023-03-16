package kmp_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/hw-18-boyer-moore-algo/pattern"
	"github.com/strider2038/otus-algo/hw-19-knuth-morris-pratt/kmp"
	"github.com/strider2038/otus-algo/hw-19-knuth-morris-pratt/parsing"
)

var algos = []struct {
	name string
	find func(text string, pattern string) int
}{
	{name: "full", find: pattern.FindFull},
	{name: "pattern prefix", find: pattern.FindByPatternPrefix},
	{name: "text suffix", find: pattern.FindByTextSuffix},
	{
		name: "state machine parser",
		find: func(text string, pattern string) int {
			parser := parsing.NewParser("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ ", pattern)

			return parser.Find(text)
		},
	},
	{
		name: "KMP slow",
		find: func(text string, pattern string) int {
			return kmp.NewSlowParser(pattern).Find(text)
		},
	},
	{
		name: "KMP fast",
		find: func(text string, pattern string) int {
			return kmp.NewParser(pattern).Find(text)
		},
	},
	{
		name: "strings.Index",
		find: strings.Index,
	},
}

var cases = []struct {
	text    string
	pattern string
	want    int
}{
	{text: "AABAABAABAAABA", pattern: "AABAABAAABA", want: 3},
	{text: "ababababababc", pattern: "abc", want: 10},
	{text: "abababababababababababababababababababababc", pattern: "abc", want: 40},
	{text: "the quick brown fox jumps over the lazy dog", pattern: "jump", want: 20},
	{text: "the quick brown fox jumps over the lazy dog", pattern: "dog", want: 40},
	{text: "the quick brown fox jumps over the lazy dog", pattern: "the", want: 0},
	{text: "the quick brown fox jumps over the lazy dog", pattern: "dot", want: -1},
	{
		text:    "the quick doggy brown dog named dancing doggy do the jump over the lazy dancing doggy dog and doggy dog do the best",
		pattern: "doggy dog do",
		want:    94,
	},
}

var benchmarkCases = []struct {
	name    string
	text    string
	pattern string
}{
	{
		name:    "aab",
		text:    "ABCAABAABBAAABAABAABAAABABAABBAAABAAAB",
		pattern: "AABAABAAABA",
	},
	{
		name:    "dog",
		text:    "the quick doggy brown dog named dancing doggy do the jump over the lazy dancing doggy dog and doggy dog do the best",
		pattern: "doggy dog do",
	},
	{
		name: "big string",
		text: strings.Repeat(" the quick brown fox jumps over the lazy dog ", 100) +
			"brown fox jumps over the dog jumps over the dog" +
			strings.Repeat(" the quick brown fox jumps over the lazy dog ", 100),
		pattern: "brown fox jumps over the dog jumps over the dog",
	},
}

func TestParser_Find(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%s (%s)", test.text, test.pattern), func(t *testing.T) {
			p := kmp.NewParser(test.pattern)

			got := p.Find(test.text)

			datatesting.AssertEqual(t, test.want, got)
		})
	}
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
	for _, test := range benchmarkCases {
		for _, algo := range algos {
			b.Run(test.name+": "+algo.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					algo.find(test.text, test.pattern)
				}
			})
		}

		b.Run(test.name+": KMP prebuilt", func(b *testing.B) {
			p := kmp.NewParser(test.pattern)
			for i := 0; i < b.N; i++ {
				p.Find(test.text)
			}
		})
	}
}
