package main

import "fmt"

// source: https://leetcode.com/problems/longest-palindromic-substring/

// first find the shortest palindromes possible and then try to extend
func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}

	var ml, mr int
	runes := []rune(s)
	size := len(runes)

	extendPalindrome := func(l, r int) (int, int) {
		for l-1 >= 0 && r+1 < size && runes[l-1] == runes[r+1] {
			l--
			r++
		}
		return l, r
	}

	for i := 0; i < size-2; i++ {
		if runes[i] == runes[i+1] {
			l, r := extendPalindrome(i, i+1)
			if r-l > mr-ml {
				ml, mr = l, r
			}
		}
		if runes[i] == runes[i+2] {
			l, r := extendPalindrome(i, i+2)
			if r-l > mr-ml {
				ml, mr = l, r
			}
		}
	}
	if runes[size-2] == runes[size-1] {
		l, r := extendPalindrome(size-2, size-1)
		if r-l > mr-ml {
			ml, mr = l, r
		}
	}

	return s[ml : mr+1]
}

func main() {
	// Example 7
	var s7 string = "aaaa"
	fmt.Println("Expected: \"aaaa\"	Output: ", longestPalindrome(s7))

	// Example 6
	var s6 string = "ccc"
	fmt.Println("Expected: \"ccc\"	Output: ", longestPalindrome(s6))

	// Example 4
	var s5 string = "abc"
	fmt.Println("Expected: \"a\"	Output: ", longestPalindrome(s5))

	// Example 3
	var s4 string = "b"
	fmt.Println("Expected: \"b\"	Output: ", longestPalindrome(s4))

	// Example 3
	var s3 string = "bb"
	fmt.Println("Expected: \"bb\"	Output: ", longestPalindrome(s3))

	// Example 2
	var s2 string = "cbbd"
	fmt.Println("Expected: \"bb\"	Output: ", longestPalindrome(s2))

	// Example 1
	var s1 string = "babad"
	fmt.Println("Expected: \"bab\"	Output: ", longestPalindrome(s1))

}
