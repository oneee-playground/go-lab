package trie_test

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/onee-only/go-data-structures/data-structure/non-linear/tree/trie"
	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	tr := trie.NewTrie()

	tr.Insert("this")
	tr.Insert("thinner")

	assert.True(t, tr.Find("this"))
}

func TestTrieDelete(t *testing.T) {
	tr := trie.NewTrie()

	tr.Insert("this")
	tr.Insert("thinner")

	tr.Delete("thinner")

	assert.True(t, tr.Find("this"))
	assert.False(t, tr.Find("thinner"))
}

func BenchmarkTrie(b *testing.B) {
	tr := trie.NewTrie()

	for i := 0; i < b.N; i++ {
		tr.Insert(strconv.Itoa(rand.Int()))
	}
}
