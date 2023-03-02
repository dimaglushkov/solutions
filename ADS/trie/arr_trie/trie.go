package arr_trie

type Trie struct {
	Children [26]*Trie
	Term     bool
}

func NewTrie() *Trie {
	return new(Trie)
}

func (t *Trie) Insert(word string) {
	for _, ch := range word {
		x := ch - 'a'
		if t.Children[x] == nil {
			t.Children[x] = NewTrie()
		}
		t = t.Children[x]
	}
	t.Term = true
}

func (t *Trie) Search(word string) bool {
	for _, ch := range word {
		x := ch - 'a'
		if t.Children[x] == nil {
			return false
		}
		t = t.Children[x]
	}
	return t.Term == true
}

func (t *Trie) StartsWith(prefix string) bool {
	for _, ch := range prefix {
		x := ch - 'a'
		if t.Children[x] == nil {
			return false
		}
		t = t.Children[x]
	}
	return true
}
