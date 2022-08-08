package main

import "fmt"

// source= https=//leetcode.com/problems/count-vowels-permutation/

func countVowelPermutation(n int) int {
	const m = 1_000_000_007
	var a, u, e, o, i = 1, 1, 1, 1, 1
	for k := 1; k < n; k++ {
		na := (e + i + u) % m
		nu := (i + o) % m
		ne := (a + i) % m
		no := i % m
		ni := (e + o) % m

		a, u, e, o, i = na, nu, ne, no, ni
	}

	return (a + u + e + o + i) % m
}

func main() {
	// Example 0
	var n0 int = 144
	fmt.Println("Expected= 18208803	Output= ", countVowelPermutation(n0))

	// Example 1
	var n1 int = 1
	fmt.Println("Expected= 5	Output= ", countVowelPermutation(n1))

	// Example 2
	var n2 int = 2
	fmt.Println("Expected= 10	Output= ", countVowelPermutation(n2))

	// Example 3
	var n3 int = 5
	fmt.Println("Expected= 68	Output= ", countVowelPermutation(n3))

}
