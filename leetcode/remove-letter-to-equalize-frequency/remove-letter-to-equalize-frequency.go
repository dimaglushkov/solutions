package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/remove-letter-to-equalize-frequency/

func equalFrequency(word string) bool {
	if len(word) < 3 {
		return true
	}

	var m = make([]int, 26)
	for _, r := range word {
		m[r-'a']++
	}

	sort.Ints(m)
	if m[25] == 1 {
		return true
	}

	var prev = 0
	for _, frequency := range m {
		if frequency == 0 {
			continue
		}
		if prev == 0 {
			prev = frequency
		} else if prevFrequency != frequency {
			return false
		}
	}
	return false
}

func main() {
	// Example 1
	var word1 string = "abcc"
	fmt.Println("Expected: true	Output: ", equalFrequency(word1))

	// Example 2
	var word2 string = "aazz"
	fmt.Println("Expected: false	Output: ", equalFrequency(word2))

}
