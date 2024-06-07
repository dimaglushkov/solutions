package main

import (
	"fmt"
	"strings"
)

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

func (t *Trie) Search(word string) string {
	p := t
	for i, ch := range word {
		x := ch - 'a'
		if p.Term {
			return word[:i]
		}
		if p.Children[x] == nil {
			return ""
		}
		p = p.Children[x]
	}
	return ""
}

func replaceWords(dictionary []string, sentence string) string {
	trie := NewTrie()
	for _, word := range dictionary {
		trie.Insert(word)
	}

	words := strings.Split(sentence, " ")

	for i, w := range words {
		if x := trie.Search(w); x != "" {
			words[i] = x
		}
	}

	return strings.Join(words, " ")
}

func main() {
	testCases := []struct {
		dictionary []string
		sentence   string
		want       string
	}{
		{
			dictionary: []string{"cat", "bat", "rat"},
			sentence:   "the cattle was rattled by the battery",
			want:       "the cat was rat by the bat",
		},
		{
			dictionary: []string{"a", "b", "c"},
			sentence:   "aadsfasf absbs bbab cadsfafs",
			want:       "a a b c",
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := replaceWords(tc.dictionary, tc.sentence)
		status := "ERROR"
		if fmt.Sprint(x) == fmt.Sprint(tc.want) {
			status = "OK"
			successes++
		}
		fmt.Println(status, "	Expected: ", tc.want, " Actual: ", x)
	}
	if l := len(testCases); successes == len(testCases) {
		fmt.Printf("===\nSUCCESS: %d of %d tests ended successfully\n", successes, l)
	} else {
		fmt.Printf("===\nFAIL: %d tests failed\n", l-successes)
	}

}
