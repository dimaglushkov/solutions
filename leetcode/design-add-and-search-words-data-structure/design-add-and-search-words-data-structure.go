package main

// source: https://leetcode.com/problems/design-add-and-search-words-data-structure/

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
	var dfs func(p *Trie, word []rune) bool
	dfs = func(p *Trie, word []rune) bool {
		if len(word) == 0 {
			return p.Term == true
		}
		if p == nil {
			return false
		}

		r := word[0]
		if r == '.' {
			for rr := range p.Children {
				if dfs(p.Children[rr], word[1:]) {
					return true
				}
			}
			return false
		} else if p.Children[r] == nil {
			return false
		}
		return dfs(p.Children[r], word[1:])
	}
	return dfs(t, []rune(word))
}

type WordDictionary struct {
	data Trie
}

func Constructor() WordDictionary {
	return WordDictionary{data: *NewTrie()}
}

func (this *WordDictionary) AddWord(word string) {
	this.data.Insert(word)
}

func (this *WordDictionary) Search(word string) bool {
	return this.data.Search(word)
}

func main() {
	wd := Constructor()
	wd.AddWord("a")
	wd.AddWord("a")
	println(wd.Search("."))
	println(wd.Search("a"))
	println(wd.Search("aa"))
	println(wd.Search("a"))
	println(wd.Search(".a"))
	println(wd.Search("a."))
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
