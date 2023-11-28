package senior_datastruct

// TrieNode node of Trie
type TrieNode struct {
	IsWord bool
	Next   map[string]*TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		Root: &TrieNode{
			Next: make(map[string]*TrieNode),
		},
		Size: 0,
	}
}

// Trie struct
type Trie struct {
	Root *TrieNode
	Size int
}

// Add put a word to Trie
func (t *Trie) Add(word string) {
	cur := t.Root
	for i := 0; i < len(word); i++ {
		c := word[i : i+1]
		if _, ok := cur.Next[c]; !ok {
			cur.Next[c] = &TrieNode{Next: make(map[string]*TrieNode)}
		}
		cur = cur.Next[c]
	}
	// marker last node true
	if !cur.IsWord {
		cur.IsWord = true
		t.Size++
	}
}

// Contains trie is contained the word
func (t *Trie) Contains(word string) bool {
	cur := t.Root
	for i := 0; i < len(word); i++ {
		c := word[i : i+1]
		if _, ok := cur.Next[c]; !ok {
			return false
		}
		cur = cur.Next[c]
	}
	return cur.IsWord
}

func (t *Trie) IsPrefix(prefix string) bool {
	cur := t.Root
	for i := 0; i < len(prefix); i++ {
		c := prefix[i : i+1]
		if _, ok := cur.Next[c]; !ok {
			return false
		}
		cur = cur.Next[c]
	}

	return true
}
