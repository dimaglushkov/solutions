package main

import (
	"fmt"
	"sort"
	"strings"
)

// source: https://leetcode.com/problems/top-k-frequent-words/

func topKFrequent(words []string, k int) []string {
	m := make(map[string]int)
	for _, w := range words {
		m[w]++
	}

	type wc struct {
		w string
		c int
	}

	wcs := make([]wc, 0)

	for w, c := range m {
		wcs = append(wcs, wc{w: w, c: c})
	}

	sort.Slice(wcs, func(i, j int) bool {
		if wcs[i].c > wcs[j].c {
			return true
		}
		if wcs[i].c == wcs[j].c {
			return strings.Compare(wcs[i].w, wcs[j].w) == -1
		}
		return false
	})

	res := make([]string, k)
	for i := range res {
		res[i] = wcs[i].w
	}
	return res
}

func main() {
	// Example 1
	var words1 = []string{"i", "love", "leetcode", "i", "love", "coding"}
	var k1 int = 2
	fmt.Println("Expected: [i, love]	Output: ", topKFrequent(words1, k1))

	// Example 2
	var words2 = []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}
	var k2 int = 4
	fmt.Println("Expected: [the, is, sunny, day]	Output: ", topKFrequent(words2, k2))

}
