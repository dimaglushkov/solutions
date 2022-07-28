package main

import "fmt"

// source: https://leetcode.com/problems/valid-anagram/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var l [27]int
	rtoi := func(r rune) int {
		return int(r - 'a')
	}
	rs, rt := []rune(s), []rune(t)

	for i := 0; i < len(s); i++ {
		l[rtoi(rs[i])]++
		l[rtoi(rt[i])]--
	}
	for _, x := range l {
		if x != 0 {
			return false
		}
	}

	return true
}

func main() {
	// Example 1
	var s1 string = "anagram"
	var t1 string = "nagaram"
	fmt.Println("Expected: true	Output: ", isAnagram(s1, t1))

	// Example 2
	var s2 string = "rat"
	var t2 string = "car"
	fmt.Println("Expected: false	Output: ", isAnagram(s2, t2))

}
