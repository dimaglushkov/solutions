package main

// source: https://leetcode.com/problems/implement-trie-prefix-tree/

type Trie struct {
	Children [26]*Trie
	Term     bool
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	for _, ch := range word {
		x := ch - 'a'
		if t.Children[x] == nil {
			t.Children[x] = new(Trie)
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

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
