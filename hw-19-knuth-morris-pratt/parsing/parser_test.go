package parsing_test

import (
	"testing"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/hw-19-knuth-morris-pratt/parsing"
)

func TestParser_Search(t *testing.T) {
	p := parsing.NewParser("ABC", "AABAABAAABA")

	index := p.Find("AABAABAABAAABA")

	datatesting.AssertEqual(t, 3, index)
}
