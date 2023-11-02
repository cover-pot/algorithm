package senior_datastruct

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrie_Add(t *testing.T) {
	trie := NewTrie()
	trie.Add("apple")
	assert.Equal(t, true, trie.Contains("apple"))
	assert.Equal(t, false, trie.Contains("app"))
	assert.Equal(t, true, trie.IsPrefix("app"))
	trie.Add("app")
	assert.Equal(t, true, trie.Contains("app"))

}
