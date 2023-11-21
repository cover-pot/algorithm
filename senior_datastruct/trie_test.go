package senior_datastruct

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrie_Add(t *testing.T) {
	context.Background()
	trie := NewTrie()
	trie.Add("apple")
	assert.Equal(t, true, trie.Contains("apple"))
	assert.Equal(t, false, trie.Contains("app"))
	assert.Equal(t, true, trie.IsPrefix("app"))
	trie.Add("app")
	assert.Equal(t, true, trie.Contains("app"))

}
