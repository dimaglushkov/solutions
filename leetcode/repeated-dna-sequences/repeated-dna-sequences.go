package main

import (
	"fmt"
	"math"
)

// source: https://leetcode.com/problems/repeated-dna-sequences/

// just store amount of every 10-digit len substrings
// return every substring that appeared more than once
func findRepeatedDnaSequences(s string) (res []string) {
	res = make([]string, 0)
	if len(s) < 10 {
		return res
	}

	seq := make(map[string]int)
	for l, r := 0, 10; r <= len(s); l, r = l+1, r+1 {
		seq[s[l:r]]++
	}
	for k, v := range seq {
		if v > 1 {
			res = append(res, k)
		}
	}
	return res
}

// same as above, but uses less memory
func findRepeatedDnaSequences_(s string) (res []string) {
	res = make([]string, 0)
	if len(s) < 10 {
		return res
	}

	seen := make(map[string]bool)
	for l, r := 0, 10; r <= len(s); l, r = l+1, r+1 {
		if used, exists := seen[s[l:r]]; exists {
			if !used {
				res = append(res, s[l:r])
				seen[s[l:r]] = true
			}
			continue
		}
		seen[s[l:r]] = false
	}
	return res
}

// below I tried to implement rolling hash algorithm
// but didn't succeed with this idea :/
func findRepeatedDnaSequences__(s string) (res []string) {
	res = make([]string, 0)
	if len(s) < 10 {
		return res
	}

	var hash int
	letters := map[rune]int{
		'A': 1,
		'B': 2,
		'C': 3,
		'D': 4,
	}
	seen := make(map[int]bool)

	runes := []rune(s)
	for i := range runes {
		hash = hash*4 + letters[runes[i]]
		if i >= 10 {
			hash -= letters[runes[i-10]] * int(math.Pow(4, 10))
		}
		if i >= 9 {
			if used, ok := seen[hash]; ok {
				if !used {
					res = append(res, s[i-9:i+1])
					seen[hash] = true
				}
			} else {
				seen[hash] = false
			}
		}
	}
	return res
}

func main() {
	// Example 2
	var s2 string = "AAAAAAAAAAAAA"
	fmt.Println("Expected: [\"AAAAAAAAAA\"]	Output: ", findRepeatedDnaSequences(s2))

	// Example 1
	var s1 string = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	fmt.Println("Expected: [\"AAAAACCCCC\",\"CCCCCAAAAA\"]	Output: ", findRepeatedDnaSequences(s1))

	// Example 3
	var s3 string = "AAAAAAAAAAA"
	fmt.Println("Expected: [\"AAAAAAAAAAA\"]	Output: ", findRepeatedDnaSequences(s3))

}
