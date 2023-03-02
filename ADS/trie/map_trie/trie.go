package map_trie

type Trie struct {
	Children map[rune]*Trie
	Term     bool
}

func NewTrie() *Trie {
	return &Trie{Children: make(map[rune]*Trie)}
}

func (t *Trie) Insert(word string) {
	for _, r := range word {
		if t.Children[r] == nil {
			t.Children[r] = NewTrie()
		}
		t = t.Children[r]
	}
	t.Term = true
}

func (t *Trie) Search(word string) bool {
	for _, r := range word {
		if t.Children[r] == nil {
			return false
		}
		t = t.Children[r]
	}
	return t.Term == true
}

func (t *Trie) StartsWith(prefix string) bool {
	for _, r := range prefix {
		if t.Children[r] == nil {
			return false
		}
		t = t.Children[r]
	}
	return true
}
