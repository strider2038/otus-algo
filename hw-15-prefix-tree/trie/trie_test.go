package trie_test

import (
	"testing"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/hw-15-prefix-tree/trie"
)

func TestTrie(t *testing.T) {
	words := trie.New()
	words.Insert("apple")
	datatesting.AssertTrue(t, words.Search("apple"))
	datatesting.AssertFalse(t, words.Search("app"))
	datatesting.AssertTrue(t, words.StartsWith("app"))
	words.Insert("app")
	datatesting.AssertTrue(t, words.Search("app"))
}
